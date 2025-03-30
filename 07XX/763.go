package leetcode_ans

func partitionLabels(s string) []int {
	minMap := make(map[byte]int)
	maxMap := make(map[byte]int)

	order := ""

	for i, ch := range []byte(s) {
		if _, ok := minMap[ch]; !ok {
			minMap[ch] = i
			maxMap[ch] = i
			order = order + string(ch)
		} else {
			if i > maxMap[ch] {
				maxMap[ch] = i
			}
		}
	}

	res := make([]int, 0, len(order))

	start, end := 0, maxMap[order[0]]
	for i := 1; i < len(order); i++ {
		if minMap[order[i]] > end {
			res = append(res, end-start+1)
			start, end = minMap[order[i]], maxMap[order[i]]
		} else if maxMap[order[i]] > end {
			end = maxMap[order[i]]
		}
	}
	res = append(res, end-start+1)

	return res
}