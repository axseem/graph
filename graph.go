package graph

import (
	"errors"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

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

type WeightedGraph[K comparable, N Number] interface {
	Graph[K]
	VerticesValues(vertices ...K) []N
	EdgesValues(vertices ...[2]K) []N
}

// Reader is the interface that defines graph data read methods.
type Reader[K comparable] interface {
	// Returns all vertices in the graph.
	// Must not return nil.
	// Infinite graph should return vertices that can't be procedural calculated.
	Vertices() []K

	// Returns amount of vertices in graph.
	// Infinite graph should return -1.
	Order() int
}

// Reader is the interface that defines methods that allow modify graph data.
type Writer[K comparable] interface {
	AddVertices(vertices ...K) error
	DeleteVertices(vertices ...K)
	AddEdges(edges ...[2]K) error
	DeleteEdges(edges ...[2]K)
}

type GraphReadWriter[K comparable] interface {
	Graph[K]
	Reader[K]
	Writer[K]
}
