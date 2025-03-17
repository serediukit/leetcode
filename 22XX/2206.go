package leetcode_ans

func divideArray(nums []int) bool {
	pairs := make(map[int]bool)
	for _, n := range nums {
		pairs[n] = !pairs[n]
	}

	for _, p := range pairs {
		if p {
			return false
		}
	}
	return true
}