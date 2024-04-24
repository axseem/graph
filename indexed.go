package graph

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Indexed[U unsigned] struct {
	vertices [][]U
}

func NewIndexed[U unsigned]() *Indexed[U] {
	return &Indexed[U]{
		vertices: [][]U{},
	}
}

func (g *Indexed[U]) Adjacency(vertex U) []U {
	if len(g.vertices) < int(vertex) {
		return nil
	}
	return g.vertices[vertex]
}
