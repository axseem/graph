package graph

import "slices"

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Unstable - vertex deletion, moves the last vertex to the place of the deleted one.
type Indexed[U unsigned] struct {
	vertices [][]U
}

func NewIndexed[U unsigned]() *Indexed[U] {
	return &Indexed[U]{
		vertices: [][]U{},
	}
}

func (g *Indexed[U]) Adjacency(vertex U) []U {
	if len(g.vertices) <= int(vertex) {
		return nil
	}
	if g.vertices[vertex] == nil {
		return []U{}
	}
	return g.vertices[vertex]
}

func (g *Indexed[U]) Vertices() []U {
	vertices := make([]U, len(g.vertices))
	for i := range g.vertices {
		vertices = append(vertices, U(i))
	}
	return vertices
}

func (g *Indexed[U]) Order() int {
	return len(g.vertices)
}

// If only one vertex passed, increases slice size by its value.
// If more than one passed, increases slice size by amount of vertices passed.
func (g *Indexed[U]) AddVertices(vertices ...U) error {
	if len(vertices) == 0 {
		return nil
	}

	length := len(g.vertices)
	if len(vertices) == 1 {
		length += int(vertices[0])
	}
	if len(vertices) > 1 {
		length += len(vertices)
	}

	if length > cap(g.vertices) {
		newVertices := make([][]U, length)
		copy(newVertices, g.vertices)
		g.vertices = newVertices
	} else {
		g.vertices = g.vertices[:length]
	}

	return nil
}

// Deleted vertices get replaced by last ones.
func (g *Indexed[U]) DeleteVertices(vertices ...U) {
	for _, vertex := range vertices {
		if int(vertex) >= len(g.vertices) {
			continue
		}

		g.vertices[vertex] = g.vertices[len(g.vertices)-1]
		g.vertices = g.vertices[:len(g.vertices)-1]

	loopVertices:
		for vertexIndex, neighbors := range g.vertices {

			var foundDeleted, foundLast bool
			for neighborIndex, neighbor := range neighbors {
				if neighbor == vertex {
					foundDeleted = true
					g.vertices[vertexIndex] = slices.Delete(g.vertices[vertexIndex], neighborIndex, neighborIndex+1)
				}
				if int(neighbor) == len(g.vertices) {
					foundLast = true
					g.vertices[vertexIndex][neighborIndex] = vertex
				}

				// no need to check all neighbors once these two are found
				if foundDeleted && foundLast {
					continue loopVertices
				}
			}

		}

	}
}

func (g *Indexed[U]) AddEdges(edges ...[2]U) error {
	for _, edge := range edges {
		vertex1, vertex2 := edge[0], edge[1]

		if int(vertex1) >= len(g.vertices) || int(vertex2) >= len(g.vertices) {
			return ErrNilVertex
		}

		if vertex1 == vertex2 {
			return ErrLoop
		}

		if slices.Contains(g.vertices[vertex1], vertex2) {
			continue
		}

		g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	}
	return nil
}

func (g *Indexed[U]) DeleteEdges(edges ...[2]U) {
	for _, edge := range edges {
		vertex1, vertex2 := edge[0], edge[1]

		if int(vertex1) >= len(g.vertices) || int(vertex2) >= len(g.vertices) {
			continue
		}

		g.vertices[vertex1] = slices.DeleteFunc(g.vertices[vertex1], func(v U) bool {
			return v == vertex2
		})
	}
}
