package asm

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
}

func allocateRegisters(curr *Block, colors []*Register, ignored map[*Register]bool) {
	ranges := map[*Block]*rangeSets{}

	// Create range sets
	createGenRemove(curr, ranges, ignored)

	// Create liveouts
	createLiveout(ranges)

	// Create interference graph from liveouts
	graph := createInterfGraph(curr, colors, ranges, ignored)

	// Create register->color (color being a fancy word for physical register) mapping
	coloring, _ := colorGraph(graph, colors)

	// TODO: Make this better?
	for old, nw := range coloring {
		old.Name = nw.Name
	}
}

func createGenRemove(curr *Block, ranges map[*Block]*rangeSets, ignored map[*Register]bool) {
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
	ranges[curr] = &rangeSets{
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

func createLiveout(ranges map[*Block]*rangeSets) {
	// Create working block set
	workSet := make(map[*Block]struct{}, len(ranges))

	// Initially add all blocks to the working set
	for v := range ranges {
		workSet[v] = struct{}{}
	}

	// Process the work set until it's empty
	for len(workSet) > 0 {
		// Pop a random block from the work set
		var curr *Block
		for curr = range workSet {
			break
		}
		delete(workSet, curr)

		// Compute new liveout
		newLiveout := map[*Register]struct{}{}
		oldLiveout := ranges[curr].liveout

		var changed bool
		for _, next := range curr.Next {
			changed = changed || unionLiveoutNext(ranges[next], newLiveout, oldLiveout)
		}

		// Check if length of liveout has changed (and mark as changed if so)
		if len(newLiveout) != len(oldLiveout) {
			changed = true
		}

		// Set new liveout
		ranges[curr].liveout = newLiveout

		// Add predecessor blocks to the work set if the current block's liveout changed
		if changed {
			for _, b := range curr.Prev {
				workSet[b] = struct{}{}
			}
		}
	}
}

func unionLiveoutNext(rng *rangeSets, newLiveout map[*Register]struct{},
	oldLiveout map[*Register]struct{}) bool {

	// Main goal: get genSet U (liveout - removeSet) for some block (set arithmetic)
	// Then union that expression into the previous block's liveout
	// So liveout = U{for all successors} (genSet U (liveout - removeSet))
	liveRem := map[*Register]struct{}{}

	// First, calculate (liveout - removeSet)
	for reg := range rng.liveout {
		if _, present := rng.removeSet[reg]; !present {
			liveRem[reg] = struct{}{}
		}
	}

	// Then, union (liveout - removeSet) into previous liveout
	var changed bool
	for reg := range liveRem {
		if _, present := oldLiveout[reg]; !present {
			changed = true
		}

		newLiveout[reg] = struct{}{}
	}

	// Now, union genSet into previous liveout
	for reg := range rng.genSet {
		if _, present := oldLiveout[reg]; !present {
			changed = true
		}

		newLiveout[reg] = struct{}{}
	}

	return changed
}

func createInterfGraph(curr *Block, colors []*Register,
	ranges map[*Block]*rangeSets, ignored map[*Register]bool) map[*Register]*node {

	// Create graph map
	graph := map[*Register]*node{}

	// Process interference within each block
	for b, sets := range ranges {
		processBlockInterf(b, sets.liveout, graph, colors, ignored)
	}

	return graph
}

func processBlockInterf(curr *Block, livenow map[*Register]struct{},
	graph map[*Register]*node, colors []*Register, ignored map[*Register]bool) {

	// Add interference for each instruction (bottom to top)
	for i := len(curr.Terminals) - 1; i >= 0; i-- {
		processInstrInterf(curr.Terminals[i], livenow, graph, colors, ignored)
	}

	for i := len(curr.PhiOuts) - 1; i >= 0; i-- {
		processInstrInterf(curr.PhiOuts[i], livenow, graph, colors, ignored)
	}

	for i := len(curr.Instrs) - 1; i >= 0; i-- {
		processInstrInterf(curr.Instrs[i], livenow, graph, colors, ignored)
	}

	// Add remaining livenow registers to graph (required to capture some disconnected nodes)
	for reg := range livenow {
		findNode(reg, graph)
	}
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

		// Get its node in the graph
		n := findNode(dst, graph)

		// Add interference edge with all livenow registers
		for reg := range livenow {
			n.addNeighbor(reg, findNode(reg, graph))
		}

		// TODO: Needed? Why is this here?
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

			livenow[reg] = struct{}{}
		}
	}
}

func colorGraph(graph map[*Register]*node,
	colors []*Register) (ret map[*Register]*Register, err error) {

	// Make a working copy of the interference graph (nodes will be removed from the working copy)
	workGraph := maps.Clone(graph)

	for _, n := range graph {
		// Make a working copy of each node's neighbors (again, working neighbors will be removed)
		n.workNeighbors = map[*node]struct{}{}

		for _, neigh := range n.neighbors {
			n.workNeighbors[neigh] = struct{}{}
		}

		// Clear neighbor colors for each node
		n.neighborColors = map[*Register]bool{}
	}

	// Push virtual nodes into a stack
	stack := pushNodes(workGraph, len(colors))

	// Pre-color physical nodes
	colorPhysical(graph, colors)

	// Pop and color virtual nodes
	ret, err = colorVirtual(graph, stack, colors)
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

			// Push (virtual) nodes to the stack and remove from working graph
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

		// TODO: Move this up the call chain
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
