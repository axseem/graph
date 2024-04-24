package graph

import (
	"slices"
)

type Grid struct {
	vertices map[[2]int][][2]int
	deleted  map[[2]int]struct{}
}

func NewGrid() *Grid {
	return &Grid{
		vertices: make(map[[2]int][][2]int),
		deleted:  make(map[[2]int]struct{}),
	}
}

func (g *Grid) Adjacency(vertex [2]int) [][2]int {
	if _, deleted := g.deleted[vertex]; deleted {
		return nil
	}

	v, custom := g.vertices[vertex]
	if !custom {
		return [][2]int{
			{vertex[0], vertex[1] + 1},
			{vertex[0] + 1, vertex[1]},
			{vertex[0], vertex[1] - 1},
			{vertex[0] - 1, vertex[1]},
		}
	}

	return v
}

func (g *Grid) AddVertex(vertex [2]int) error {
	if _, deleted := g.deleted[vertex]; !deleted {
		return ErrVertexExists
	}

	delete(g.deleted, vertex)

	// TODO procedural neighbors if possible

	return nil
}

func (g *Grid) DeleteVertex(vertex [2]int) {
	if _, deleted := g.deleted[vertex]; deleted {
		return
	}

	for v, neighbors := range g.vertices {
		for i, n := range neighbors {
			if n == vertex {
				g.vertices[v] = slices.Delete(g.vertices[v], i, i+1)
			}
		}
	}

	neighborhood := [...][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i, shift := range neighborhood {
		n := [2]int{vertex[0] + shift[0], vertex[1] + shift[1]}

		_, custom := g.vertices[n]
		_, deleted := g.deleted[n]

		if !custom && !deleted {
			switch i {
			case 0:
				g.vertices[n] = [][2]int{
					{vertex[0] + shift[0], vertex[1] + shift[1] + 1},
					{vertex[0] + shift[0] + 1, vertex[1] + shift[1]},
					{vertex[0] + shift[0] - 1, vertex[1] + shift[1]},
				}
			case 1:
				g.vertices[n] = [][2]int{
					{vertex[0] + shift[0], vertex[1] + shift[1] + 1},
					{vertex[0] + shift[0] + 1, vertex[1] + shift[1]},
					{vertex[0] + shift[0], vertex[1] + shift[1] - 1},
				}
			case 2:
				g.vertices[n] = [][2]int{
					{vertex[0] + shift[0] + 1, vertex[1] + shift[1]},
					{vertex[0] + shift[0], vertex[1] + shift[1] - 1},
					{vertex[0] + shift[0] - 1, vertex[1] + shift[1]},
				}
			case 3:
				g.vertices[n] = [][2]int{
					{vertex[0] + shift[0], vertex[1] + shift[1] + 1},
					{vertex[0] + shift[0], vertex[1] + shift[1] - 1},
					{vertex[0] + shift[0] - 1, vertex[1] + shift[1]},
				}
			}

		}
	}

	g.deleted[vertex] = struct{}{}
}

func (g *Grid) AddEdge(vertex1, vertex2 [2]int) error {
	if vertex1 == vertex2 {
		return ErrLoop
	}

	_, deleted1 := g.deleted[vertex1]
	_, deleted2 := g.deleted[vertex2]
	if deleted1 || deleted2 {
		return ErrNilVertex
	}

	if _, custom := g.vertices[vertex1]; !custom {
		g.vertices[vertex1] = [][2]int{
			{vertex1[0], vertex1[1] + 1},
			{vertex1[0] + 1, vertex1[1]},
			{vertex1[0], vertex1[1] - 1},
			{vertex1[0] - 1, vertex1[1]},
			vertex2,
		}
	} else {
		g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	}

	return nil
}

func (g *Grid) DeleteEdge(vertex1, vertex2 [2]int) {

}
