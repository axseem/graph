package graph_test

import (
	"reflect"
	"testing"

	"github.com/axseem/graph"
)

func TestAdjacency(t *testing.T) {
	testCases := []struct {
		desc   string
		entry  int
		result []int
	}{
		{
			desc:   "one neighbor",
			entry:  0,
			result: []int{1},
		},
		{
			desc:   "zero neighbors",
			entry:  1,
			result: []int{},
		},
		{
			desc:   "nil vertex",
			entry:  2,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewMapped[int]()
			g.AddVertex(0)
			g.AddVertex(1)
			g.AddEdge(0, 1)

			n := g.Adjacency(tC.entry)
			if !reflect.DeepEqual(n, tC.result) {
				t.Error(tC.desc)
			}
		})
	}
}

func BenchmarkMappedGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := graph.NewMapped[int]()
		g.Grow(4096)
	}
}

func BenchmarkSimpleAdd4096Vertices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const l = 4096
		g := graph.NewMapped[int]()
		for i := range l {
			if err := g.AddVertex(i); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkSimpleGrowAdd4096Vertices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const l = 4096
		g := graph.NewMapped[int]()
		for j := range l {
			if err := g.AddVertex(j); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkSimpleAdd4096Edges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		const l = 4097
		g := graph.NewMapped[int]()
		g.Grow(l)
		for i := range l {
			if err := g.AddVertex(i); err != nil {
				b.Error(err)
			}
		}
		b.StartTimer()

		for i := 1; i < l; i++ {
			if err := g.AddEdge(0, i); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const l = 4097
		g := graph.NewMapped[int]()
		for i := range l {
			if err := g.AddVertex(i); err != nil {
				b.Error(err)
			}
		}

		for i := 1; i < l; i++ {
			if err := g.AddEdge(0, i); err != nil {
				b.Error(err)
			}
		}

		for i := 0; i < l; i += 2 {
			g.DeleteVertex(i)
		}

		for i := 1; i < l; i += 2 {
			g.DeleteEdge(0, i)
		}

	}
}
