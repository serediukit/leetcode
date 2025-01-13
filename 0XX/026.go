package leetcode_ans

func removeDuplicates(nums []int) int {
    cnt := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] != nums[i-1] {
            nums[cnt] = nums[i]
            cnt++
        }
    }
    return cnt
}