package leetcode_ans

func maxSubArray(nums []int) int {
	maxSum, tempSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > tempSum+nums[i] {
			tempSum = nums[i]
		} else {
			tempSum += nums[i]
		}
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}

	return maxSum
}