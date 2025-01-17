package leetcode_ans

func search(nums []int, target int) int {
    start, end := 0, len(nums) - 1

    for start <= end {
        mid := (start + end) / 2

        if nums[mid] == target {
            return mid
        }

        if nums[start] <= nums[mid] {
            if nums[start] <= target && target < nums[mid] {
                end = mid - 1
            } else {
                start = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[end] {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
    }

    return -1
}