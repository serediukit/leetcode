package leetcode_ans

func jump(nums []int) int {
    l := len(nums)

    if l == 1 {
        return 0
    }

    jumps := 1
    currentLength := 0
    maxLength := nums[0]
    for maxLength < l - 1 {
        temp := maxLength
        for pos := currentLength + 1; pos <= maxLength; pos++ {
            if pos + nums[pos] > temp {
                temp = pos + nums[pos]
            }
        }
        currentLength = maxLength
        maxLength = temp

        jumps++
    }

    return jumps
}