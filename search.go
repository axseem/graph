package graph

// Each time a vertex is visited, the while function is triggered.
// The function exits if while returns false or there are no more vertices.
func BFS[K comparable](g Graph[K], entry K, while func(vertex K, depth uint) bool) error {
	queue := []K{entry}
	visited := make(map[K]struct{})
	var depth, depthCounter uint
	var depthThreshold uint = 1

	for len(queue) > 0 {
		bottom := queue[0]
		queue = queue[1:]

		if _, ok := visited[bottom]; ok {
			depthThreshold--
			continue
		}

		if !while(bottom, depth) {
			return nil
		}

		n := g.Adjacency(bottom)
		if n == nil {
			return ErrNilVertex
		}

		queue = append(queue, n...)
		visited[bottom] = struct{}{}

		depthCounter++
		if depthCounter == depthThreshold {
			depth++
			depthCounter = 0
			depthThreshold = uint(len(queue))
		}
	}

	return nil
}

// Each time a vertex is visited, the while function is triggered.
// The function exits if while returns false or there are no more vertices.
func DFS[K comparable](g Graph[K], entry K, while func(vertex K) bool) error {
	queue := []K{entry}
	visited := make(map[K]struct{})

	for len(queue) > 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if _, ok := visited[top]; ok {
			continue
		}

		if !while(top) {
			return nil
		}

		n := g.Adjacency(top)
		if n == nil {
			return ErrNilVertex
		}

		queue = append(queue, n...)
		visited[top] = struct{}{}
	}

	return nil
}
