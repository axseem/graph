package graph

import (
	"cmp"
	"errors"
)

var ErrNilVertex = errors.New("nil vertex")
var ErrVertexExists = errors.New("vertex already exists")
var ErrLoop = errors.New("simple graph can't contain loops")

type Graph[K cmp.Ordered] interface {
	Adjacency(vertex K) ([]K, error)
}

type Reader[K cmp.Ordered] interface {
	Vertices() []K
	Edges() [][2]K
}

type Writer[K cmp.Ordered] interface {
	AddVertex(vertex K) error
	DeleteVertex(vertex K)
	AddEdge(vertex1 K, vertex2 K) error
	DeleteEdge(vertex1 K, vertex2 K)
}

type GraphReadWriter[K cmp.Ordered] interface {
	Graph[K]
	Reader[K]
	Writer[K]
}

type Simple[K cmp.Ordered] struct {
	vertices map[K][]K
}

func NewSimple[K cmp.Ordered]() *Simple[K] {
	return &Simple[K]{
		vertices: make(map[K][]K),
	}
}

func (g *Simple[K]) Adjacency(vertex K) ([]K, error) {
	neighbors, ok := g.vertices[vertex]
	if !ok {
		return nil, ErrNilVertex
	}

	return neighbors, nil
}

func (g *Simple[K]) Vertices() []K {
	vertices := make([]K, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g *Simple[K]) AddVertex(vertex K) error {
	_, ok := g.vertices[vertex]
	if ok {
		return ErrVertexExists
	}

	g.vertices[vertex] = make([]K, 0)

	return nil
}

func (g *Simple[K]) DeleteVertex(vertex K) {
	neighbors, ok := g.vertices[vertex]
	if !ok {
		return
	}

	for _, n := range neighbors {
		for i, v := range g.vertices[n] {
			if v == vertex {
				g.vertices[n] = append(g.vertices[n][:i], g.vertices[n][i+1:]...)
			}
		}
	}

	delete(g.vertices, vertex)
}

func (g *Simple[K]) AddEdge(vertex1, vertex2 K) error {
	if vertex1 == vertex2 {
		return ErrLoop
	}

	_, ok := g.vertices[vertex1]
	if !ok {
		return ErrNilVertex
	}
	_, ok = g.vertices[vertex2]
	if !ok {
		return ErrNilVertex
	}

	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)

	return nil
}

func (g *Simple[K]) DeleteEdge(vertex1, vertex2 K) {
	_, ok1 := g.vertices[vertex1]
	_, ok2 := g.vertices[vertex2]

	if !ok1 || !ok2 || vertex1 == vertex2 {
		return
	}

	for i, v := range g.vertices[vertex1] {
		if v == vertex2 {
			g.vertices[vertex1] = append(g.vertices[vertex1][:i], g.vertices[vertex1][i+1:]...)
		}
	}
}
