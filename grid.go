package graph

import (
	"slices"
)

type Grid struct {
	vertices map[[2]int][][2]int
}

func NewGrid() *Grid {
	return &Grid{
		vertices: make(map[[2]int][][2]int),
	}
}

func (g *Grid) Adjacency(vertex [2]int) [][2]int {
	neighbors, custom := g.vertices[vertex]

	if !custom {
		return [][2]int{
			{vertex[0], vertex[1] + 1},
			{vertex[0] + 1, vertex[1]},
			{vertex[0], vertex[1] - 1},
			{vertex[0] - 1, vertex[1]},
		}
	}

	return neighbors
}

func (g *Grid) Vertices() [][2]int {
	vertices := make([][2]int, 0, len(g.vertices))
	for vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}

	slices.SortFunc(vertices, func(a [2]int, b [2]int) int {
		if b[1] == a[1] {
			return a[0] - b[0]
		}
		return b[1] - a[1]
	})
	return vertices
}

func (g *Grid) Order() int {
	return -1
}

func (g *Grid) AddVertices(vertices ...[2]int) error {
	for _, vertex := range vertices {
		neighbors, custom := g.vertices[vertex]
		if !custom || neighbors != nil {
			return ErrVertexExists
		}

		g.vertices[vertex] = [][2]int{}
	}
	return nil
}

func (g *Grid) DeleteVertices(vertices ...[2]int) {
	for _, vertex := range vertices {
		neighbors, custom := g.vertices[vertex]
		if custom && neighbors == nil {
			return
		}

		for v, neighbors := range g.vertices {
			for i, n := range neighbors {
				if n == vertex {
					g.vertices[v] = slices.Delete(g.vertices[v], i, i+1)
					g.tryMakePredictable(v)
				}
			}
		}

		neighborUp := [2]int{vertex[0], vertex[1] + 1}
		if _, custom := g.vertices[neighborUp]; !custom {
			g.vertices[neighborUp] = [][2]int{
				{neighborUp[0], neighborUp[1] + 1},
				{neighborUp[0] + 1, neighborUp[1]},
				{neighborUp[0] - 1, neighborUp[1]},
			}
		}

		neighborRight := [2]int{vertex[0] + 1, vertex[1]}
		if _, custom := g.vertices[neighborRight]; !custom {
			g.vertices[neighborRight] = [][2]int{
				{neighborRight[0], neighborRight[1] + 1},
				{neighborRight[0] + 1, neighborRight[1]},
				{neighborRight[0], neighborRight[1] - 1},
			}
		}

		neighborDown := [2]int{vertex[0], vertex[1] - 1}
		if _, custom := g.vertices[neighborDown]; !custom {
			g.vertices[neighborDown] = [][2]int{
				{neighborDown[0] + 1, neighborDown[1]},
				{neighborDown[0], neighborDown[1] - 1},
				{neighborDown[0] - 1, neighborDown[1]},
			}
		}

		neighborLeft := [2]int{vertex[0] - 1, vertex[1]}
		if _, custom := g.vertices[neighborLeft]; !custom {
			g.vertices[neighborLeft] = [][2]int{
				{neighborLeft[0], neighborLeft[1] + 1},
				{neighborLeft[0], neighborLeft[1] - 1},
				{neighborLeft[0] - 1, neighborLeft[1]},
			}
		}

		g.vertices[vertex] = nil
	}
}

func (g *Grid) AddEdges(edges ...[2][2]int) error {
	for _, edge := range edges {
		vertex1, vertex2 := edge[0], edge[1]

		if vertex1 == vertex2 {
			return ErrLoop
		}

		neighbors1, custom1 := g.vertices[vertex1]
		neighbors2, custom2 := g.vertices[vertex2]
		if (custom1 && neighbors1 == nil) || (custom2 && neighbors2 == nil) {
			return ErrNilVertex
		}

		if !custom1 {
			g.vertices[vertex1] = [][2]int{
				{vertex1[0], vertex1[1] + 1},
				{vertex1[0] + 1, vertex1[1]},
				{vertex1[0], vertex1[1] - 1},
				{vertex1[0] - 1, vertex1[1]},
				vertex2,
			}
		} else {
			g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
			g.tryMakePredictable(vertex1)
		}
	}
	return nil
}

func (g *Grid) DeleteEdges(edges ...[2][2]int) {
	for _, edge := range edges {
		vertex1, vertex2 := edge[0], edge[1]

		if vertex1 == vertex2 {
			return
		}

		neighbors1, custom1 := g.vertices[vertex1]
		neighbors2, custom2 := g.vertices[vertex2]
		if (custom1 && neighbors1 == nil) || (custom2 && neighbors2 == nil) {
			return
		}

		shift := [][2]int{
			{vertex1[0], vertex1[1] + 1},
			{vertex1[0] + 1, vertex1[1]},
			{vertex1[0], vertex1[1] - 1},
			{vertex1[0] - 1, vertex1[1]},
		}

		if !custom1 && slices.Contains(shift, vertex2) {
			g.vertices[vertex1] = shift
		}

		for i, neighbor := range g.vertices[vertex1] {
			if neighbor == vertex2 {
				g.vertices[vertex1] = slices.Delete(g.vertices[vertex1], i, i+1)
				g.tryMakePredictable(vertex1)
				return
			}
		}
	}
}

func (g *Grid) tryMakePredictable(vertex [2]int) {
	neighbors, custom := g.vertices[vertex]
	if !custom || len(neighbors) != 4 {
		return
	}

	shift := [...][2]int{
		{vertex[0], vertex[1] + 1},
		{vertex[0] + 1, vertex[1]},
		{vertex[0], vertex[1] - 1},
		{vertex[0] - 1, vertex[1]},
	}

	for i := range 4 {
		if !slices.Contains(neighbors, shift[i]) {
			return
		}
	}

	delete(g.vertices, vertex)
}
