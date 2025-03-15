package leetcode_ans

func minCapability(nums []int, k int) int {
    left, right := 0, 1000000000

    for left < right {
        mid := (left + right) / 2

        countAvailable := 0

        for index := 0; index < len(nums); index++ {
            if nums[index] <= mid {
                countAvailable++
                index++
            }
        }

        if countAvailable >= k {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}