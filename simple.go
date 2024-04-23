package graph

import (
	"cmp"
)

type Simple[K cmp.Ordered] struct {
	vertices map[K][]K
}

func NewSimple[K cmp.Ordered]() *Simple[K] {
	return &Simple[K]{
		vertices: make(map[K][]K),
	}
}

func (g *Simple[K]) Grow(addedLen uint) {
	newVertices := make(map[K][]K, len(g.vertices)+int(addedLen))
	for key, value := range g.vertices {
		newVertices[key] = value
	}
	g.vertices = newVertices
}

func (g *Simple[K]) Adjacency(vertex K) []K {
	neighbors, ok := g.vertices[vertex]
	if !ok {
		return nil
	}

	return neighbors
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

	g.vertices[vertex] = []K{}

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
