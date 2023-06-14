package asm

// TODO: Remove debug statements

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

type rangeSets struct {
	genSet    map[*Register]struct{}
	removeSet map[*Register]struct{}
	liveout   map[*Register]struct{}
}

type node struct {
	reg            *Register
	neighbors      map[*Register]*node
	workNeighbors  map[*Register]*node
	neighborColors map[*Register]bool
}

func (n *node) addNeighbor(reg *Register, neighbor *node) {
	// Don't add a node as it's own neighbor
	if n != neighbor {
		// Add neighbor to this node
		n.neighbors[reg] = neighbor

		// Add this node to neighbor's neighbor list
		neighbor.neighbors[n.reg] = n
	}
}

func (n *node) String() string {
	neighStrs := make([]string, 0, len(n.neighbors))

	for n := range n.neighbors {
		neighStrs = append(neighStrs, n.Name)
	}

	ret := fmt.Sprintf("%v:\n", n.reg)
	ret += "    neighbors: " + strings.Join(neighStrs, ", ") + "\n"
	return ret
}

type graphInfo struct {
	graph     map[*Register]*node
	workGraph map[*Register]*node
	colors    map[*Register]*Register
}

func allocateRegisters(curr *Block, colors []*Register) {
	ranges := map[*Block]rangeSets{}

	createGenRemove(curr, ranges)

	fmt.Printf("=== Range sets ===\n")
	for k, v := range ranges {
		fmt.Printf("%v:\n", k.Label)
		fmt.Printf("    genSet: %v\n", v.genSet)
		fmt.Printf("    removeSet: %v\n", v.removeSet)
		fmt.Printf("    liveout: %v\n", v.liveout)
	}
	fmt.Println()

	fmt.Printf("=== Bulding interference... ===\n")
	graph := buildInterfGraph(curr, colors, ranges)
	fmt.Println()

	fmt.Printf("=== Final interference graph ===\n")
	for _, v := range graph {
		fmt.Print(v)
	}
	fmt.Println()
}

func createGenRemove(curr *Block, ranges map[*Block]rangeSets) {
	fmt.Printf("curr: %v\n", curr.Label)
	for _, v := range curr.Next {
		fmt.Printf("    next: %v\n", v.Label)
	}

	// Create range sets for the current block
	genSet := map[*Register]struct{}{}
	removeSet := map[*Register]struct{}{}
	liveout := map[*Register]struct{}{}

	// Create gen and remove sets for all instructions (top to bottom)
	for _, v := range curr.Instrs {
		processInstrGenRemove(v, genSet, removeSet)
	}

	for _, v := range curr.PhiOuts {
		processInstrGenRemove(v, genSet, removeSet)
	}

	for _, v := range curr.Terminals {
		processInstrGenRemove(v, genSet, removeSet)
	}

	// Insert range sets (also marks this block as visited)
	ranges[curr] = rangeSets{
		genSet:    genSet,
		removeSet: removeSet,
		liveout:   liveout,
	}

	// Loop through successor blocks
	for _, next := range curr.Next {
		// Create gen and remove sets for unvisited successors
		if _, ok := ranges[next]; !ok {
			createGenRemove(next, ranges)
		}

		// Goal: get genSet U (liveout - removeSet) for next block (set arithmetic)
		// Then union that expression from the next block into the current liveout
		nextLive := map[*Register]struct{}{}
		nr := ranges[next]

		// Calculate (liveout - removeSet)
		for reg := range nr.liveout {
			if _, present := nr.removeSet[reg]; !present {
				nextLive[reg] = struct{}{}
			}
		}

		// Union (liveout - removeSet) into current liveout
		maps.Copy(liveout, nextLive)

		// Union genSet into current liveout
		maps.Copy(liveout, nr.genSet)
	}
}

func processInstrGenRemove(instr Instr, genSet map[*Register]struct{}, removeSet map[*Register]struct{}) {
	// Loop through instruction sources
	for _, src := range instr.getSrcs() {
		var reg *Register
		var ok bool

		// Ignore immediate sources
		if reg, ok = src.(*Register); !ok {
			continue
		}

		// Ignore registers that are already in the remove set
		if _, present := removeSet[reg]; present {
			continue
		}

		// Add source register to the generate set
		genSet[reg] = struct{}{}
	}

	// Loop through instruction destinations
	for _, dst := range instr.getDsts() {
		// Add destination register to the remove set
		removeSet[dst] = struct{}{}
	}
}

func buildInterfGraph(curr *Block, colors []*Register,
	ranges map[*Block]rangeSets) map[*Register]*node {

	// Create graph map
	graph := map[*Register]*node{}

	// Pre-populate graph with nodes for physical registers
	for _, reg := range colors {
		graph[reg] = &node{
			reg:       reg,
			neighbors: map[*Register]*node{},
		}
	}

	// Process interference within each block
	for b, sets := range ranges {
		processBlockInterf(b, sets.liveout, graph)
	}

	return graph
}

func processBlockInterf(curr *Block, livenow map[*Register]struct{}, graph map[*Register]*node) {
	fmt.Printf("Processing interference for block: %v\n", curr.Label)
	// Copy liveout as livenow
	fmt.Printf("Initial livenow: %v\n", livenow)

	// Add interference for each instruction (bottom to top)
	for i := len(curr.Terminals) - 1; i >= 0; i-- {
		fmt.Printf("Instr: %v\n", curr.Terminals[i])
		processInstrInterf(curr.Terminals[i], livenow, graph)
	}

	for i := len(curr.PhiOuts) - 1; i >= 0; i-- {
		fmt.Printf("Instr: %v\n", curr.PhiOuts[i])
		processInstrInterf(curr.PhiOuts[i], livenow, graph)
	}

	for i := len(curr.Instrs) - 1; i >= 0; i-- {
		fmt.Printf("Instr: %v\n", curr.Instrs[i])
		processInstrInterf(curr.Instrs[i], livenow, graph)
	}

	fmt.Println()
}

func processInstrInterf(instr Instr, livenow map[*Register]struct{}, graph map[*Register]*node) {
	// For each instruction target...
	for _, dst := range instr.getDsts() {
		fmt.Printf("    Target: %v\n", dst)
		nod := findNode(dst, graph)

		// Add interference edge with all livenow registers
		for reg := range livenow {
			fmt.Printf("        Adding interf w/: %v\n", reg)
			nod.addNeighbor(reg, findNode(reg, graph))
		}

		graph[dst] = nod

		// Remove target from livenow
		delete(livenow, dst)
	}

	// Add instruction sources to livenow
	for _, src := range instr.getSrcs() {
		if reg, ok := src.(*Register); ok {
			fmt.Printf("    Src: %v\n", reg)
			livenow[reg] = struct{}{}
		}
	}

	fmt.Printf("Livenow: %v\n", livenow)
	fmt.Println()
}

func findNode(reg *Register, graph map[*Register]*node) *node {
	// Check for existing node and return it
	if n, ok := graph[reg]; ok {
		return n
	}

	// Otherwise, create one and return it
	n := &node{
		reg:       reg,
		neighbors: map[*Register]*node{},
	}
	graph[reg] = n

	return n
}
