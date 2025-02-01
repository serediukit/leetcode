package leetcode_ans

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(a, b int) bool {
		return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]))
	})

	for i := 1; i < len(intervals); i++ {
		mergedArray, hasMerge := mergeTwoIntervals(intervals[i-1], intervals[i])
		if hasMerge {
			intervals[i-1] = mergedArray
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		}
	}

	return intervals
}

func mergeTwoIntervals(a []int, b []int) ([]int, bool) {
	if a[0] <= b[0] && a[1] >= b[0] {
		if a[1] >= b[1] {
			return a, true
		}
		return []int{a[0], b[1]}, true
	} else if b[0] <= a[0] && b[1] >= a[0] {
		if b[1] >= a[1] {
			return b, true
		}
		return []int{b[0], a[1]}, true
	}

	return nil, false
}