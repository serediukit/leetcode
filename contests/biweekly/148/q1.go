package contests

func maxAdjacentDistance(nums []int) int {
	res := abs(nums[0] - nums[len(nums)-1])
	for i := 1; i < len(nums); i++ {
		res = max(res, abs(nums[i]-nums[i-1]))
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}