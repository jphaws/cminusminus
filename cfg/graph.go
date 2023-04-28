package cfg

import (
	dot "github.com/awalterschulze/gographviz"
)

var visited = map[*Block]struct{}{}

func CreateGraph(block *Block) (s string, err error) {
	graph := dot.NewGraph()
	graph.SetName(block.function)
	graph.SetDir(true)

	processBlock(block, graph) // TODO err

	s = graph.String()
	return
}

func processBlock(block *Block, graph *dot.Graph) (lab string, err error) {
	if _, present := visited[block]; present {
		return
	}
	visited[block] = struct{}{}

	function := block.function
	lab = block.Label()
	// fmt.Println(function)
	// fmt.Println(lab)
	// fmt.Println(graph.String())

	graph.AddNode(function, lab, nil)

	if block.Next != nil {
		processBlock(block.Next, graph)                   // TODO err
		graph.AddEdge(lab, block.Next.Label(), true, nil) // TODO err
	}

	if block.Els != nil {
		processBlock(block.Els, graph)                   // TODO err
		graph.AddEdge(lab, block.Els.Label(), true, nil) // TODO err
	}

	for _, v := range block.Prev {
		graph.AddEdge(lab, v.Label(), true, nil)
	}

	return
}
