package leetcode_ans

func divideArray(nums []int) bool {
	pairs := make(map[int]int)

	for _, n := range nums {
		pairs[n]++
	}

	for _, p := range pairs {
		if p % 2 != 0 {
			return false
		}
	}
	return true
}