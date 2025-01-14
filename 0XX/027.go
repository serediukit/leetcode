package leetcode_ans

func removeElement(nums []int, val int) int {
    cnt := 0
    if len(nums) == 0 {
        return 0
    }
    for i := 0; i < len(nums); i++ {
        if (nums[i] != val) {
            nums[cnt] = nums[i]
            cnt++
        }
    }
    return cnt
}