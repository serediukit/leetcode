package leetcode_ans

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	sort.Ints(nums)

	backtrack(&res, []int{}, nums)
	return res
}

func backtrack(res *[][]int, cur []int, nums []int) {
	for i := 0; i < len(nums); i++ {
		if !(i > 0 && nums[i] == nums[i-1]) {
			cur = append(cur, nums[i])
			temp := make([]int, len(cur))
			copy(temp, cur)
			*res = append(*res, temp)

			backtrack(res, cur, nums[i+1:])
			cur = cur[:len(cur)-1]
		}
	}
}