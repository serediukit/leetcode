package leetcode_ans

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			mergedArray, hasMerge := mergeTwoIntervals(intervals[i], intervals[j])
			if hasMerge {
				intervals[i] = mergedArray
				intervals = append(intervals[:j], intervals[j+1:]...)
				return merge(intervals)
			}
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
