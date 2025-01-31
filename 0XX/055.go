package leetcode_ans

func canJump(nums []int) bool {
    maxJump := nums[0]
    for i := 1; i < len(nums) - 1 && i <= maxJump; i++ {
        jump := i + nums[i]
        if jump > maxJump {
            maxJump = jump
        }
    }
    return maxJump >= len(nums) - 1
}