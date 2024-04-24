package graph

import (
	"errors"
)

var ErrNilVertex = errors.New("nil vertex")
var ErrVertexExists = errors.New("vertex already exists")
var ErrLoop = errors.New("simple graph can't contain loops")

// Graph is the interface that wraps the basic Adjacency method.
type Graph[K comparable] interface {
	// There are three possible outputs Adjacency method has:
	// Nil slice: there in no a such vertex in the graph;
	// Empty slice: given vertex does not have neighbors;
	// Slice with values: all neighbors of the vertex.
	Adjacency(vertex K) []K
}

type Reader[K comparable] interface {
	Vertices() []K

	// If graph is infinite should return -1
	Order() int
}

type Writer[K comparable] interface {
	AddVertex(vertex K) error
	DeleteVertex(vertex K)
	AddEdge(vertex1, vertex2 K) error
	DeleteEdge(vertex1, vertex2 K)
	Grow(int)
}

type GraphReadWriter[K comparable] interface {
	Graph[K]
	Reader[K]
	Writer[K]
}
