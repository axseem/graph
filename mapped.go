package graph

import "slices"

type Mapped[K comparable] struct {
	vertices map[K][]K
}

func NewMapped[K comparable]() *Mapped[K] {
	return &Mapped[K]{
		vertices: make(map[K][]K),
	}
}

func (g *Mapped[K]) Adjacency(vertex K) []K {
	neighbors, ok := g.vertices[vertex]
	if !ok {
		return nil
	}

	return neighbors
}

func (g *Mapped[K]) Vertices() []K {
	vertices := make([]K, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g *Mapped[K]) AddVertices(vertices ...K) error {
	for _, vertex := range vertices {
		_, ok := g.vertices[vertex]
		if ok {
			return ErrVertexExists
		}

		g.vertices[vertex] = []K{}
	}
	return nil
}

func (g *Mapped[K]) DeleteVertices(vertices ...K) {
	for _, vertex := range vertices {
		if _, ok := g.vertices[vertex]; !ok {
			return
		}

		for v, neighbors := range g.vertices {
			for i, n := range neighbors {
				if n == vertex {
					g.vertices[v] = slices.Delete(g.vertices[v], i, i+1)
				}
			}
		}

		delete(g.vertices, vertex)
	}
}

func (g *Mapped[K]) AddEdges(edges ...[2]K) error {
	for _, edge := range edges {
		vertex1, vertex2 := edge[0], edge[1]

		if vertex1 == vertex2 {
			return ErrLoop
		}

		_, exists1 := g.vertices[vertex1]
		_, exists2 := g.vertices[vertex2]
		if !exists1 || !exists2 {
			return ErrNilVertex
		}

		g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	}
	return nil
}

func (g *Mapped[K]) DeleteEdges(edges ...[2]K) {
	for _, edge := range edges {
		vertex1, vertex2 := edge[0], edge[1]
		_, exists1 := g.vertices[vertex1]
		_, exists2 := g.vertices[vertex2]

		if !exists1 || !exists2 || vertex1 == vertex2 {
			return
		}

		for i, v := range g.vertices[vertex1] {
			if v == vertex2 {
				g.vertices[vertex1] = append(g.vertices[vertex1][:i], g.vertices[vertex1][i+1:]...)
			}
		}
	}
}
