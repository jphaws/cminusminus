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
	workNeighbors  map[*node]struct{}
	neighborColors map[*Register]bool
}

func (n *node) addNeighbor(reg *Register, neigh *node) {
	// Don't add a node as its own neighbor
	if n != neigh {
		// Add neighbor to this node
		n.neighbors[reg] = neigh

		// Add this node to neighbor's neighbor list
		neigh.neighbors[n.reg] = n
	}
}

func (n *node) color(color *Register) {
	// Add color to each neighbor's color map
	for _, neigh := range n.neighbors {
		neigh.neighborColors[color] = true
	}
}

func (n *node) removeFromWorkNeighbors() {
	// Remove this node from all its neighbors' workNeighbors maps
	for neigh := range n.workNeighbors {
		delete(neigh.workNeighbors, n)
	}
}

func (n *node) String() string {
	neighStrs := make([]string, 0, len(n.neighbors))
	colorStrs := make([]string, 0, len(n.neighborColors))

	for neigh := range n.neighbors {
		neighStrs = append(neighStrs, neigh.Name)
	}

	for color := range n.neighborColors {
		colorStrs = append(colorStrs, color.Name)
	}

	ret := fmt.Sprintf("%v:\n", n.reg)
	ret += "    neighbors: " + strings.Join(neighStrs, ", ") + "\n"
	ret += "    neighbor colors: " + strings.Join(colorStrs, ", ") + "\n"
	return ret
}

type graphInfo struct {
	graph     map[*Register]*node
	workGraph map[*Register]*node
	colors    map[*Register]*Register // TODO: Necessary?
}

func allocateRegisters(curr *Block, colors []*Register, ignored map[*Register]bool) {
	ranges := map[*Block]rangeSets{}

	// Create range sets
	createGenRemove(curr, ranges, ignored)

	fmt.Printf("=== Range sets ===\n")
	for k, v := range ranges {
		fmt.Printf("%v:\n", k.Label)
		fmt.Printf("    genSet: %v\n", v.genSet)
		fmt.Printf("    removeSet: %v\n", v.removeSet)
		fmt.Printf("    liveout: %v\n", v.liveout)
	}
	fmt.Println()

	fmt.Printf("=== Bulding interference... ===\n")
	graph := createInterfGraph(curr, colors, ranges, ignored)
	fmt.Println()

	fmt.Printf("=== Final interference graph ===\n")
	for _, v := range graph {
		fmt.Print(v)
	}
	fmt.Println()

	coloring, _ := colorGraph(graph, colors)

	// TODO: Make this better?
	for old, nw := range coloring {
		old.Name = nw.Name
	}
}

func createGenRemove(curr *Block, ranges map[*Block]rangeSets, ignored map[*Register]bool) {
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
		processInstrGenRemove(v, genSet, removeSet, ignored)
	}

	for _, v := range curr.PhiOuts {
		processInstrGenRemove(v, genSet, removeSet, ignored)
	}

	for _, v := range curr.Terminals {
		processInstrGenRemove(v, genSet, removeSet, ignored)
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
			createGenRemove(next, ranges, ignored)
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

func processInstrGenRemove(instr Instr, genSet map[*Register]struct{},
	removeSet map[*Register]struct{}, ignored map[*Register]bool) {

	// Loop through instruction sources
	for _, src := range instr.getSrcs() {
		var reg *Register
		var ok bool

		// Ignore immediate sources
		if reg, ok = src.(*Register); !ok {
			continue
		}

		// Ignore special registers
		if ignored[reg] {
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
		// Ignore special registers
		if ignored[dst] {
			continue
		}

		// Add destination register to the remove set
		removeSet[dst] = struct{}{}
	}
}

func createInterfGraph(curr *Block, colors []*Register,
	ranges map[*Block]rangeSets, ignored map[*Register]bool) map[*Register]*node {

	// Create graph map
	graph := map[*Register]*node{}

	// Pre-populate graph with nodes for physical registers
	/* TODO: Remove? Unnecessary?
	for _, reg := range colors {
		graph[reg] = &node{
			reg:       reg,
			neighbors: map[*Register]*node{},
		}
	}
	*/

	// Process interference within each block
	for b, sets := range ranges {
		processBlockInterf(b, sets.liveout, graph, colors, ignored)
	}

	return graph
}

func processBlockInterf(curr *Block, livenow map[*Register]struct{},
	graph map[*Register]*node, colors []*Register, ignored map[*Register]bool) {

	fmt.Printf("Processing interference for block: %v\n", curr.Label)
	// Copy liveout as livenow
	fmt.Printf("Initial livenow: %v\n", livenow)

	// Add interference for each instruction (bottom to top)
	for i := len(curr.Terminals) - 1; i >= 0; i-- {
		fmt.Printf("Instr: %v\n", curr.Terminals[i])
		processInstrInterf(curr.Terminals[i], livenow, graph, colors, ignored)
	}

	for i := len(curr.PhiOuts) - 1; i >= 0; i-- {
		fmt.Printf("Instr: %v\n", curr.PhiOuts[i])
		processInstrInterf(curr.PhiOuts[i], livenow, graph, colors, ignored)
	}

	for i := len(curr.Instrs) - 1; i >= 0; i-- {
		fmt.Printf("Instr: %v\n", curr.Instrs[i])
		processInstrInterf(curr.Instrs[i], livenow, graph, colors, ignored)
	}

	fmt.Println()
}

func processInstrInterf(instr Instr, livenow map[*Register]struct{},
	graph map[*Register]*node, colors []*Register, ignored map[*Register]bool) {

	// Determine instruction targets
	var targets []*Register
	if _, ok := instr.(*BranchLinkInstr); ok {
		// For branch and link instructions, use all caller saved registers (x0-x15)
		// These registers may be overwritten during the subroutine call
		targets = colors[genRegsCallerStart:]
	} else {
		// For all others, just use their own targets
		targets = instr.getDsts()
	}

	// For each instruction target...
	for _, dst := range targets {
		// Ignore special registers
		if ignored[dst] {
			continue
		}

		fmt.Printf("    Target: %v\n", dst)
		n := findNode(dst, graph)

		// Add interference edge with all livenow registers
		for reg := range livenow {
			fmt.Printf("        Adding interf w/: %v\n", reg)
			n.addNeighbor(reg, findNode(reg, graph))
		}

		graph[dst] = n
	}

	// Remove instruction targets from livenow
	for _, dst := range targets {
		delete(livenow, dst)
	}

	// Add instruction sources to livenow
	for _, src := range instr.getSrcs() {
		if reg, ok := src.(*Register); ok {
			// Ignore special registers
			if ignored[reg] {
				continue
			}
			fmt.Printf("    Src: %v\n", reg)
			livenow[reg] = struct{}{}
		}
	}

	fmt.Printf("Livenow: %v\n", livenow)
	fmt.Println()
}

func colorGraph(graph map[*Register]*node,
	colors []*Register) (ret map[*Register]*Register, err error) {

	// Make a working copy of the interference graph
	workGraph := maps.Clone(graph)

	for _, n := range graph {
		// Make a working copy of each node's neighbors
		n.workNeighbors = map[*node]struct{}{}

		for _, neigh := range n.neighbors {
			n.workNeighbors[neigh] = struct{}{}
		}

		// Clear neighbor colors for each node
		n.neighborColors = map[*Register]bool{}
	}

	// Push virtual nodes into a stack
	stack := pushNodes(workGraph, len(colors))
	fmt.Printf("Stack: %v\n", stack)

	// Pre-color physical nodes
	colorPhysical(graph, colors)

	// Pop and color virtual nodes
	ret, err = colorVirtual(graph, stack, colors)

	fmt.Printf("=== Colored interference graph ===\n")
	for _, v := range graph {
		fmt.Print(v)
	}
	fmt.Println()

	fmt.Printf("=== Color maps ===\n")
	fmt.Println(ret)

	return
}

func pushNodes(wg map[*Register]*node, maxDegree int) []*node {
	// Create node stack
	stack := []*node{}

	// Loop until all nodes are removed from the working graph
	for len(wg) > 0 {
		i := 0

		// Check if each node is constrained (degree >= #colors) or unconstrained (degree < #colors)
		for k, node := range wg {
			i++

			// Skip constrained nodes, unless all nodes are constrained
			if len(node.workNeighbors) >= maxDegree && i != len(wg) {
				continue
			}

			// Remove node from its neighbors' maps
			node.removeFromWorkNeighbors()

			// Push node to the stack and remove from working graph
			// fmt.Printf("Remove: %v", node) // TODO: Remove me
			if node.reg.Virtual {
				stack = append(stack, node)
			}
			delete(wg, k)
			break
		}
	}

	return stack
}

func colorPhysical(graph map[*Register]*node, colors []*Register) {
	// For all physical registers...
	for _, phys := range colors {
		n := graph[phys]

		// That are also in the graph...
		if n == nil {
			continue
		}

		// Color register as itself
		n.color(phys)
	}
}

func colorVirtual(graph map[*Register]*node, stack []*node,
	colors []*Register) (ret map[*Register]*Register, err error) {

	ret = make(map[*Register]*Register, len(stack))

	// "Pop" nodes from stack
	for i := len(stack) - 1; i >= 0; i-- {
		n := stack[i]

		// Attempt to color node
		c, err := colorNode(n, colors)
		if err != nil {
			panic(err)
		}

		// Save coloring in final map
		ret[n.reg] = c
	}

	return
}

func colorNode(nod *node, colors []*Register) (reg *Register, err error) {
	// Loop through colors until no neighbors have that color
	for _, c := range colors {
		if nod.neighborColors[c] {
			continue
		}

		// Color this node with the first available color
		nod.color(c)

		// Return chosen color
		reg = c
		return
	}

	// If no colors available, return an error
	err = fmt.Errorf("Node %v triggered spill", nod.reg)
	return
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
