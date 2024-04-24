package graph_test

import (
	"reflect"
	"testing"

	"github.com/axseem/graph"
)

func TestGridAdjacency(t *testing.T) {
	testCases := []struct {
		desc           string
		deleteVertices [][2]int
		addVertices    [][2]int
		deleteEdges    [][2][2]int
		addEdges       [][2][2]int
		entry          [2]int
		result         [][2]int
	}{
		{
			desc:   "x:0;y:0",
			entry:  [2]int{0, 0},
			result: [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
		},
		{
			desc:           "delete y+1 neighbor",
			entry:          [2]int{0, 0},
			deleteVertices: [][2]int{{0, 1}},
			result:         [][2]int{{1, 0}, {0, -1}, {-1, 0}},
		},
		{
			desc:           "delete x+1 neighbor",
			entry:          [2]int{0, 0},
			deleteVertices: [][2]int{{1, 0}},
			result:         [][2]int{{0, 1}, {0, -1}, {-1, 0}},
		},
		{
			desc:           "delete y-1 neighbor",
			entry:          [2]int{0, 0},
			deleteVertices: [][2]int{{0, -1}},
			result:         [][2]int{{0, 1}, {1, 0}, {-1, 0}},
		},
		{
			desc:           "delete x-1 neighbor",
			entry:          [2]int{0, 0},
			deleteVertices: [][2]int{{-1, 0}},
			result:         [][2]int{{0, 1}, {1, 0}, {0, -1}},
		},
		{
			desc:     "add edge",
			entry:    [2]int{0, 0},
			addEdges: [][2][2]int{{{0, 0}, {1, 1}}},
			result:   [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}},
		},
		{
			desc:           "nil vertex",
			entry:          [2]int{0, 0},
			deleteVertices: [][2]int{{0, 0}},
			result:         nil,
		},
		{
			desc:           "zero neighbors",
			entry:          [2]int{0, 0},
			deleteVertices: [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
			result:         [][2]int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewGrid()

			for _, v := range tC.deleteVertices {
				g.DeleteVertex(v)
			}
			for _, v := range tC.addVertices {
				g.AddVertex(v)
			}
			for _, e := range tC.addEdges {
				g.AddEdge(e[0], e[1])
			}
			for _, e := range tC.deleteEdges {
				g.DeleteEdge(e[0], e[1])
			}

			n := g.Adjacency(tC.entry)
			if !reflect.DeepEqual(n, tC.result) {
				t.Errorf("expected: %v, got: %v", tC.result, n)
			}
		})
	}
}

func BenchmarkDeleteVertex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := graph.NewGrid()

		g.DeleteVertex([2]int{0, 0})
	}
}
