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
			g.AddVertices(0, 1)
			g.AddEdges([2]int{0, 1})

			n := g.Adjacency(tC.entry)
			if !reflect.DeepEqual(n, tC.result) {
				t.Error(tC.desc)
			}
		})
	}
}
