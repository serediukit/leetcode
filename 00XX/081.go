package leetcode_ans

func search(nums []int, target int) bool {
    if len(nums) == 1 {
        return nums[0] == target
    }
    return search(nums[:len(nums)/2], target) || search(nums[len(nums)/2:], target)
}