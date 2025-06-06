package leetcode_ans

func minOperations(nums []int) int {
    res := 0

    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] == 0 {
            nums[i] = 1
            nums[i+1] ^= 1
            nums[i+2] ^= 1
            res++
        }
    }

    if nums[len(nums)-1] + nums[len(nums)-2] != 2 {
        return -1
    }

    return res
}