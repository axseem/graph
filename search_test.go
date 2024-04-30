package graph_test

import (
	"reflect"
	"testing"

	"github.com/axseem/graph"
)

type test struct {
	desc   string
	order  uint
	edges  [][2]uint
	entry  uint
	output []uint
	err    error
}

func commonTests() []test {
	return []test{
		{
			desc:   "empty graph",
			order:  0,
			edges:  nil,
			entry:  0,
			output: nil,
			err:    graph.ErrNilVertex,
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
	}
}

func TestDFS(t *testing.T) {
	testCases := append(commonTests(), []test{
		{
			desc:   "graph:4←3←0→1→2",
			order:  7,
			edges:  [][2]uint{{3, 4}, {0, 3}, {0, 1}, {1, 2}},
			entry:  0,
			output: []uint{0, 1, 2, 3, 4},
		},
	}...)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewIndexed[uint]()

			if err := g.AddVertices(tC.order); err != nil {
				panic(err)
			}

			if err := g.AddEdges(tC.edges...); err != nil {
				panic(err)
			}

			path := []uint{}
			err := graph.DFS(g, tC.entry, func(vertex uint) bool {
				path = append(path, vertex)
				return true
			})
			if err != nil {
				if err == tC.err {
					return
				}
				panic(err)
			}

			if !reflect.DeepEqual(tC.output, path) {
				t.Errorf("expected: %v, got: %v", tC.output, path)
			}
		})
	}
}

func TestBFS(t *testing.T) {
	testCases := append(commonTests(), []test{
		{
			desc:   "graph:3←1←0→2→4",
			order:  7,
			edges:  [][2]uint{{1, 3}, {0, 1}, {0, 2}, {2, 4}},
			entry:  0,
			output: []uint{0, 1, 2, 3, 4},
		},
	}...)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewIndexed[uint]()

			if err := g.AddVertices(tC.order); err != nil {
				panic(err)
			}

			if err := g.AddEdges(tC.edges...); err != nil {
				panic(err)
			}

			path := []uint{}
			err := graph.BFS(g, tC.entry, func(vertex, depth uint) bool {
				path = append(path, vertex)
				return true
			})
			if err != nil {
				if err == tC.err {
					return
				}
				panic(err)
			}

			if !reflect.DeepEqual(tC.output, path) {
				t.Errorf("expected: %v, got: %v", tC.output, path)
			}
		})
	}
}

func TestBFSDepth(t *testing.T) {
	g := graph.NewIndexed[uint]()

	if err := g.AddVertices(7); err != nil {
		panic(err)
	}

	if err := g.AddEdges([][2]uint{{1, 2}, {0, 1}, {0, 3}, {3, 4}}...); err != nil {
		panic(err)
	}

	path := []uint{}
	err := graph.BFS(g, 0, func(vertex, depth uint) bool {
		if depth > 1 {
			return false
		}

		path = append(path, vertex)
		return true
	})
	if err != nil {
		panic(err)
	}

	expect := []uint{0, 1, 3}
	if !reflect.DeepEqual(expect, path) {
		t.Errorf("expected: %v, got: %v", expect, path)
	}
}

func TestBFSDepthGrid(t *testing.T) {
	g := graph.NewGrid()

	amount := 0
	err := graph.BFS(g, [2]int{0, 0}, func(vertex [2]int, depth uint) bool {
		if depth > 2 {
			return false
		}

		// log.Println(vertex)
		amount++
		return true
	})
	if err != nil {
		panic(err)
	}

	expect := 13
	if !reflect.DeepEqual(expect, amount) {
		t.Errorf("expected: %v, got: %v", expect, amount)
	}
}
