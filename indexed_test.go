package graph_test

import (
	"reflect"
	"testing"

	"github.com/axseem/graph"
)

func TestIndexedAdjacency(t *testing.T) {
	testCases := []struct {
		desc   string
		entry  uint
		result []uint
	}{
		{
			desc:   "one neighbor",
			entry:  0,
			result: []uint{1},
		},
		{
			desc:   "zero neighbors",
			entry:  1,
			result: []uint{},
		},
		{
			desc:   "nil vertex",
			entry:  2,
			result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := graph.NewIndexed[uint]()
			g.AddVertices(0, 1)
			g.AddEdges([2]uint{0, 1})

			n := g.Adjacency(tC.entry)
			if !reflect.DeepEqual(n, tC.result) {
				t.Error(tC.desc)
			}
		})
	}
}
