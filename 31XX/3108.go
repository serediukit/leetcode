package leetcode_ans

func minimumCost(n int, edges [][]int, query [][]int) []int {
    parent := make([]int, n)
	minPathCost := make([]int, n)

    var findRoot func(node int) int
    findRoot = func(node int) int {
        if parent[node] != node {
            parent[node] = findRoot(parent[node])
        }
        return parent[node]
    }

	for i := 0; i < n; i++ {
		parent[i] = i
		minPathCost[i] = -1
	}

	for _, edge := range edges {
		source, target, weight := edge[0], edge[1], edge[2]
		sourceRoot := findRoot(source)
		targetRoot := findRoot(target)

		minPathCost[targetRoot] &= weight

		if sourceRoot != targetRoot {
			minPathCost[targetRoot] &= minPathCost[sourceRoot]
			parent[sourceRoot] = targetRoot
		}
	}

	result := make([]int, len(query))
	for i, q := range query {
		start, end := q[0], q[1]
		if start == end {
			result[i] = 0
		} else if findRoot(start) != findRoot(end) {
			result[i] = -1
		} else {
			result[i] = minPathCost[findRoot(start)]
		}
	}

	return result
}