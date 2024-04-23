package graph

import (
	"cmp"
	"errors"
)

var ErrNilVertex = errors.New("nil vertex")
var ErrVertexExists = errors.New("vertex already exists")
var ErrLoop = errors.New("simple graph can't contain loops")

// Graph is the interface that wraps the basic Adjacency method.
//
// There are three possible outputs Adjacency method has:
// Nil slice: there in no a such vertex in the graph;
// Empty slice: given vertex does not have neighbors;
// Slice with values: all neighbors of the vertex.
type Graph[K cmp.Ordered] interface {
	Adjacency(vertex K) []K
}

// If graph is infinite Len should return -1
type Reader[K cmp.Ordered] interface {
	Vertices() []K
	Len() int
}

type Writer[K cmp.Ordered] interface {
	AddVertex(vertex K) error
	DeleteVertex(vertex K)
	AddEdge(vertex1 K, vertex2 K) error
	DeleteEdge(vertex1 K, vertex2 K)
	Grow(int)
}

type GraphReadWriter[K cmp.Ordered] interface {
	Graph[K]
	Reader[K]
	Writer[K]
}
