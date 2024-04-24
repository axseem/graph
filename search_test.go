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
		edges  [][2]int
		entry  int
		output []int
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
			output: []int{0},
		},
		{
			desc:   "null graph",
			order:  2,
			edges:  nil,
			entry:  0,
			output: []int{0},
		},
		{
			desc:   "graph:0→1",
			order:  2,
			edges:  [][2]int{{0, 1}},
			entry:  0,
			output: []int{0, 1},
		},
		{
			desc:   "graph:0←1",
			order:  2,
			edges:  [][2]int{{1, 0}},
			entry:  0,
			output: []int{0},
		},
		{
			desc:   "2 vertices cycled undirected graph",
			order:  2,
			edges:  [][2]int{{0, 1}, {1, 0}},
			entry:  0,
			output: []int{0, 1},
		},
		{
			desc:   "3 vertices cycled directed graph",
			order:  3,
			edges:  [][2]int{{0, 1}, {1, 2}, {2, 0}},
			entry:  0,
			output: []int{0, 1, 2},
		},
		{
			desc:   "graph:2←1←0→3→4",
			order:  7,
			edges:  [][2]int{{1, 2}, {0, 1}, {0, 3}, {3, 4}},
			entry:  0,
			output: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewMapped[int]()

			for i := range int(tC.order) {
				if err := g.AddVertex(i); err != nil {
					panic(err)
				}
			}

			for _, e := range tC.edges {
				if err := g.AddEdge(e[0], e[1]); err != nil {
					panic(err)
				}
			}

			output := graph.DFS(g, tC.entry)
			if !reflect.DeepEqual(output, tC.output) {
				t.Errorf("expected: %v, got: %v", tC.output, output)
			}
		})
	}
}

func BenchmarkDFS1024Vertices8Branches(b *testing.B) {
	const l = 1024
	branches := 8
	g := graph.NewMapped[int]()

	for i := range l {
		g.AddVertex(i)
	}

	for i := 0; i < branches; i++ {
		for j := i * (l / branches); j < (i+1)*(l/branches); j++ {
			if i != j {
				if err := g.AddEdge(i, j); err != nil {
					panic(err)
				}
			}
		}
	}
	g.AddEdge(l, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := graph.DFS(g, 0)
		if v == nil {
			b.Error("DFS nil slice")
		}
	}
}

func BenchmarkDFS1024Vertices512Branches(b *testing.B) {
	const l = 1024
	branches := 512
	g := graph.NewMapped[int]()

	for i := range l {
		g.AddVertex(i)
	}

	for i := 0; i < branches; i++ {
		for j := i * (l / branches); j < (i+1)*(l/branches); j++ {
			if i != j {
				if err := g.AddEdge(i, j); err != nil {
					panic(err)
				}
			}
		}
	}
	g.AddEdge(l, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := graph.DFS(g, 0)
		if v == nil {
			b.Error("DFS nil slice")
		}
	}
}

func BenchmarkBFS1024Vertices8Branches(b *testing.B) {
	const l = 1024
	g := graph.NewMapped[int]()
	branches := 8

	for i := range l {
		g.AddVertex(i)
	}

	for i := 0; i < branches; i++ {
		for j := i * (l / branches); j < (i+1)*(l/branches); j++ {
			if i != j {
				if err := g.AddEdge(i, j); err != nil {
					panic(err)
				}
			}
		}
	}
	g.AddEdge(l, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := graph.BFS(g, 0)
		if v == nil {
			b.Error("DFS nil slice")
		}
	}
}

func BenchmarkBFS1024Vertices512Branches(b *testing.B) {
	const l = 1024
	g := graph.NewMapped[int]()
	branches := 512

	for i := range l {
		g.AddVertex(i)
	}

	for i := 0; i < branches; i++ {
		for j := i * (l / branches); j < (i+1)*(l/branches); j++ {
			if i != j {
				if err := g.AddEdge(i, j); err != nil {
					panic(err)
				}
			}
		}
	}
	g.AddEdge(l, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := graph.BFS(g, 0)
		if v == nil {
			b.Error("DFS nil slice")
		}
	}
}
