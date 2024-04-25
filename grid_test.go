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
			desc:        "delete edge",
			entry:       [2]int{0, 0},
			deleteEdges: [][2][2]int{{{0, 0}, {-1, 0}}},
			result:      [][2]int{{0, 1}, {1, 0}, {0, -1}},
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
			var g = graph.NewGrid()

			g.DeleteVertices(tC.deleteVertices...)
			if err := g.AddVertices(tC.addVertices...); err != nil {
				t.Error(err)
			}
			if err := g.AddEdges(tC.addEdges...); err != nil {
				t.Error(err)
			}
			g.DeleteEdges(tC.deleteEdges...)

			n := g.Adjacency(tC.entry)
			if !reflect.DeepEqual(n, tC.result) {
				t.Errorf("expected: %v, got: %v", tC.result, n)
			}
		})
	}
}

func TestGridVertices(t *testing.T) {
	testCases := []struct {
		desc   string
		change func(*graph.Grid) *graph.Grid
		result [][2]int
	}{
		{
			desc: "delete predictable",
			change: func(g *graph.Grid) *graph.Grid {
				g.DeleteVertices([2]int{0, 0})
				return g
			},
			result: [][2]int{{0, 1}, {-1, 0}, {0, 0}, {1, 0}, {0, -1}},
		},
		{
			desc: "delete predictable and create new",
			change: func(g *graph.Grid) *graph.Grid {
				g.DeleteVertices([2]int{0, 0})
				if err := g.AddVertices([2]int{0, 0}); err != nil {
					t.Error(err)
				}
				return g
			},
			result: [][2]int{{0, 1}, {-1, 0}, {0, 0}, {1, 0}, {0, -1}},
		},
		{
			desc: "add edge",
			change: func(g *graph.Grid) *graph.Grid {
				g.DeleteVertices([2]int{0, 0})
				if err := g.AddVertices([2]int{0, 0}); err != nil {
					t.Error(err)
				}
				edges := [][2][2]int{
					{{0, 0}, {0, 1}},
					{{0, 0}, {-1, 0}},
					{{0, 0}, {1, 0}},
					{{0, 0}, {0, -1}},
					{{0, 1}, {0, 0}},
					{{-1, 0}, {0, 0}},
					{{1, 0}, {0, 0}},
					{{0, -1}, {0, 0}},
				}
				if err := g.AddEdges(edges...); err != nil {
					t.Error(err)
				}
				return g
			},
			result: [][2]int{},
		},
		{
			desc: "delete, create new and it make predictable",
			change: func(g *graph.Grid) *graph.Grid {
				g.DeleteVertices([2]int{0, 0})
				if err := g.AddVertices([2]int{0, 0}); err != nil {
					t.Error(err)
				}
				edges := [][2][2]int{
					{{0, 0}, {0, 1}},
					{{0, 0}, {-1, 0}},
					{{0, 0}, {1, 0}},
					{{0, 0}, {0, -1}},
					{{0, 1}, {0, 0}},
					{{-1, 0}, {0, 0}},
					{{1, 0}, {0, 0}},
					{{0, -1}, {0, 0}},
				}
				if err := g.AddEdges(edges...); err != nil {
					t.Error(err)
				}
				return g
			},
			result: [][2]int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewGrid()
			vertices := tC.change(g).Vertices()
			if !reflect.DeepEqual(tC.result, vertices) {
				t.Errorf("expected: %v, got: %v", tC.result, vertices)
			}
		})
	}
}

func BenchmarkDeleteVertex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := graph.NewGrid()

		g.DeleteVertices([2]int{0, 0})
	}
}
