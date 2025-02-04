package leetcode_ans

func insert(intervals [][]int, newInterval []int) [][]int {
    return merge(append(intervals, newInterval))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(a, b int) bool {
		return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]))
	})
	
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}