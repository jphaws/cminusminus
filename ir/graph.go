package ir

import (
	"fmt"

	dot "github.com/awalterschulze/gographviz"
)

var visited = map[*Block]struct{}{}

func CreateGraph(blocks []*Block) string {
	graph := dot.NewGraph()
	graph.SetName("G")
	graph.SetDir(true)

	for _, v := range blocks {
		fn := "cluster_" + v.function
		graph.AddSubGraph("G", fn, map[string]string{"label": v.function})
		processBlock(v, graph, fn)
	}

	return graph.String()
}

func processBlock(block *Block, graph *dot.Graph, fn string) {
	// Check if this block has been visited before
	if _, present := visited[block]; present {
		return
	}
	visited[block] = struct{}{}

	label := block.Label()

	// Generate block label (for dot)
	body := "\"{" + label + "|"
	for _, v := range block.Instrs {
		body += fmt.Sprintf("%v\\n", v)
	}
	body += "|{<next>next|<else>else}}\""

	// Set default node attributes
	attrs := map[string]string{
		"shape":    "record",
		"fontsize": "5",
		"penwidth": "0.5",
		"label":    body,
	}

	// Add color to entry and exit blocks
	if len(block.Prev) == 0 || block.Next == nil {
		attrs["color"] = "red"
	}

	// Create node for this block
	graph.AddNode(fn, label, attrs)

	// Set default edge attributes
	edgeAttrs := map[string]string{
		"penwidth": "0.5",
		"fontsize": "4",
		"label":    "next",
	}

	// Add next edge
	if block.Next != nil {
		processBlock(block.Next, graph, fn)
		graph.AddPortEdge(label, "next", block.Next.Label(), "", true, edgeAttrs)
	}

	// Add else edge
	if block.Els != nil {
		processBlock(block.Els, graph, fn)
		edgeAttrs["label"] = "else"
		graph.AddPortEdge(label, "else", block.Els.Label(), "", true, edgeAttrs)
	}

	/*
		for _, v := range block.Prev {
			edgeAttrs["label"] = "prev"
			graph.AddEdge(label, v.Label(), true, edgeAttrs)
		}
	*/

	return
}
