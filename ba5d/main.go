// Package provides solution to the
// Find the Longest Path in a DAG
// <http://rosalind.info/problems/ba5d/>
package ba5d

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"github.com/inkuzmin/ba5d/graph"
)

const Inf = int(0x7FF0000000000000)

func main() {
	file := flag.String("f", "", "Input file")
	flag.Parse()

	var source int
	var sink int

	g := parseGraph(*file, &source, &sink)

	//fmt.Printf("original graph: %v\n", g)

	lp, bt := longestPath(g, source, sink)

	fmt.Printf("%v\n%v->%v->%v\n", lp, source, bt, sink)

}

func topologicalOrdering(g graph.Graph, sink int) []graph.Node {
	size := len(g.Nodes)

	orig_g := graph.Graph{
		Nodes: append([]graph.Node(nil), g.Nodes...),
	}
	for i, node := range orig_g.Nodes {
		p := append([]int(nil), node.Predecessors...)
		w := append([]int(nil), node.Weights...)

		orig_g.Nodes[i].Predecessors = p
		orig_g.Nodes[i].Weights = w
	}

	list := make([]graph.Node, size)
	candidates := make([]graph.Node, 0)

	for _, node := range g.Nodes {
		if len(node.Predecessors) == 0 {
			candidates = append(candidates, node)
		}
	}

	i := 0
	l := len(candidates)
	for l > 0 {
		if len(candidates) > 0 {
			n := candidates[0]
			candidates = candidates[1:]

			for _, node := range orig_g.Nodes {
				if node.Label == n.Label {
					list[i] = node
				}
			}

			for i := range g.Nodes {
				g.Nodes[i].RemovePredecessor(n.Label)
				node := g.Nodes[i]

				if len(node.Predecessors) == 0 {

					to_add := true
					for _, candidate := range candidates {
						if candidate.Label == node.Label {
							to_add = false
						}
					}
					for _, ordered_node := range list {
						if ordered_node.Label == node.Label {
							to_add = false
						}
					}
					if to_add {
						candidates = append(candidates, node)
						l++
					}
				}
			}
			g.RemoveNode(n.Label)
		}

		i++
		l--
	}

	final_list := make([]graph.Node, 0)
	var sink_node graph.Node
	for _, n := range list {
		if n.Label == sink {
			sink_node = n
		} else {
			final_list = append(final_list, n)
		}
	}
	final_list = append(final_list, sink_node)

	final_final_list := make([]graph.Node, 0)
	l = len(final_list) - 1
	for l >= 0 {
		node := final_list[l]

		if len(final_final_list) == 0 {
			final_final_list = append(final_final_list, node)
		} else {
			for _, n := range final_final_list {
				for _, p := range n.Predecessors {
					if p == node.Label {
						to_add := true
						for _, e := range final_final_list {
							if e.Label == node.Label {
								to_add = false
							}
						}
						if to_add {
							final_final_list = append(final_final_list, node)
						}
					}
				}
			}
		}

		l--
	}

	final_final_final_list := make([]graph.Node, len(final_final_list))
	l = len(final_final_list) - 1
	for i := 0; i < len(final_final_list); i++ {
		final_final_final_list[i] = final_final_list[l]

		l--
	}

	// TODO: refactor
	return final_final_final_list
}

func parseGraph(file string, source *int, sink *int) graph.Graph {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		panic(err)
	}

	data := string(bytes)
	var g graph.Graph

	for i, ln := range strings.Split(data, "\n") {

		if i == 0 {
			*source, err = strconv.Atoi(ln)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
		}

		if i == 1 {
			*sink, err = strconv.Atoi(ln)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			g = graph.Graph{
				Nodes: make([]graph.Node, *sink + 100), // TODO: make with append
			}
			for i := range g.Nodes {
				g.Nodes[i].AddLabel(i)
			}

		}

		if i > 1 {
			s := strings.Split(ln, ":")
			r := strings.Split(s[0], "->")

			weight, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			from, err := strconv.Atoi(r[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			to, err := strconv.Atoi(r[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			g.Nodes[to].AddPredecessor(from, weight)
		}
	}

	return g
}

func longestPath(graph graph.Graph, source int, sink int) (int, []int) {
	size := len(graph.Nodes)
	s := make([]int, size)
	backtrack := make([]int, 0)

	for _, b := range graph.Nodes {
		s[b.Label] = -Inf
	}
	s[source] = 0

	ordered_graph := topologicalOrdering(graph, sink)

	possible_predecessors := make([]int, 1) // TODO: refactor
	possible_predecessors[0] = source

	for _, b := range ordered_graph {
		s[b.Label] = maxAll(s, b, &possible_predecessors)

		for i := range b.Predecessors {
			p := b.Predecessors[i]
			w := b.Weights[i]

			if s[b.Label] == s[p] + w {
				if p != 0 {
					backtrack = append(backtrack, p)
				}
			}
		}


	}

	return s[sink], backtrack
}

func maxAll(s []int, b graph.Node, possible_predecessors *[]int) int {
	max := 0
	is_possible := false

	for i := range b.Predecessors {
		p := b.Predecessors[i]
		w := b.Weights[i]

		is_possible = false
		for _, j := range *possible_predecessors {
			if p == j {
				is_possible = true
			}
		}
		if is_possible {
			*possible_predecessors = append(*possible_predecessors, b.Label)

			t := s[p] + w
			if t > max {
				max = t
			}
		}
	}
	return max
}
