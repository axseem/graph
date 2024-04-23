package graph

import (
	"cmp"
)

func DFS[K cmp.Ordered](g Graph[K], entry K) []K {
	if g.Adjacency(entry) == nil {
		return nil
	}

	// empty struct doesn't take space
	visited := make(map[K]struct{})
	output := []K{entry}
	dfs(g, entry, &visited, &output)
	return output
}

func dfs[K cmp.Ordered](g Graph[K], entry K, visited *map[K]struct{}, output *[]K) {
	(*visited)[entry] = struct{}{}
	n := g.Adjacency(entry)
	for _, v := range n {
		if _, ok := (*visited)[v]; !ok {
			*output = append(*output, v)
			dfs(g, v, visited, output)
		}
	}
}

func BFS[K cmp.Ordered](g Graph[K], entry K) []K {
	queue := []K{entry}
	visited := make(map[K]struct{})
	var output []K

	for len(queue) > 0 {
		bot := queue[0]
		queue = queue[1:]

		if _, ok := visited[bot]; !ok {
			continue
		}

		output = append(output, bot)
		n := g.Adjacency(bot)
		if n == nil {
			return nil
		}

		queue = append(queue, n...)
		visited[bot] = struct{}{}
	}

	return output
}
