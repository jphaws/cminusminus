package ir

import (
	"fmt"

	dot "github.com/awalterschulze/gographviz"
)

const fontName = "Noto Sans Mono"

var visitedGraph = map[*Block]bool{}

func (p ProgramIr) ToDot() string {
	graph := dot.NewEscape()
	graph.SetName("G")
	graph.SetDir(true)

	for k, v := range p.Functions {
		clusterName := "cluster_" + k

		graph.AddSubGraph("G", clusterName, map[string]string{"label": k})
		processBlock(v.Cfg, graph, k)
	}

	return graph.String()
}

func processBlock(block *Block, graph *dot.Escape, fn string) {
	visitedGraph[block] = true
	label := block.Label()

	// Generate block label (for dot)
	body := "\"{" + label + "|"
	for _, v := range block.Instrs {
		body += fmt.Sprintf("%v\\l", v)
	}
	body += "|{<next>next|<else>else}}\""

	// Set default node attributes
	attrs := map[string]string{
		"shape":    "record",
		"fontsize": "5",
		"fontname": fontName,
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
		"fontname": fontName,
		"label":    "next",
	}

	// Add next edge
	if block.Next != nil {
		if !visitedGraph[block.Next] {
			processBlock(block.Next, graph, fn)
		}

		graph.AddPortEdge(label, "next", block.Next.Label(), "", true, edgeAttrs)
	}

	// Add else edge
	if block.Els != nil {
		if !visitedGraph[block.Els] {
			processBlock(block.Els, graph, fn)
		}

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
