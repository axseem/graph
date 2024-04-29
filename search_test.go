package graph_test

import (
	"reflect"
	"testing"

	"github.com/axseem/graph"
)

func TestDFS(t *testing.T) {
	testCases := []struct {
		desc   string
		order  uint
		edges  [][2]uint
		entry  uint
		output []uint
	}{
		{
			desc:   "empty graph",
			order:  0,
			edges:  nil,
			entry:  0,
			output: nil,
		},
		{
			desc:   "trivial graph",
			order:  1,
			edges:  nil,
			entry:  0,
			output: []uint{0},
		},
		{
			desc:   "null graph",
			order:  2,
			edges:  nil,
			entry:  0,
			output: []uint{0},
		},
		{
			desc:   "graph:0→1",
			order:  2,
			edges:  [][2]uint{{0, 1}},
			entry:  0,
			output: []uint{0, 1},
		},
		{
			desc:   "graph:0←1",
			order:  2,
			edges:  [][2]uint{{1, 0}},
			entry:  0,
			output: []uint{0},
		},
		{
			desc:   "2 vertices cycled undirected graph",
			order:  2,
			edges:  [][2]uint{{0, 1}, {1, 0}},
			entry:  0,
			output: []uint{0, 1},
		},
		{
			desc:   "3 vertices cycled directed graph",
			order:  3,
			edges:  [][2]uint{{0, 1}, {1, 2}, {2, 0}},
			entry:  0,
			output: []uint{0, 1, 2},
		},
		{
			desc:   "graph:2←1←0→3→4",
			order:  7,
			edges:  [][2]uint{{1, 2}, {0, 1}, {0, 3}, {3, 4}},
			entry:  0,
			output: []uint{0, 1, 2, 3, 4},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewIndexed[uint]()

			if err := g.AddVertices(tC.order); err != nil {
				panic(err)
			}

			if err := g.AddEdges(tC.edges...); err != nil {
				panic(err)
			}

			output := graph.DFS(g, tC.entry)
			if !reflect.DeepEqual(output, tC.output) {
				t.Errorf("expected: %v, got: %v", tC.output, output)
			}
		})
	}
}
