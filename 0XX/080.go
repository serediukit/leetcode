package leetcode_ans

func removeDuplicates(nums []int) int {
    l := len(nums)
    if l < 2 {
        return l
    }

    j := 2

    for i := 2; i < l; i++ {
        if nums[i] != nums[j-1] || nums[j-1] != nums[j-2] {
            nums[j] = nums[i]
            j++
        }
    }

    return j
}