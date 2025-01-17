package leetcode_ans

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    start, end := 0, len(nums) - 1
    startPos, endPos := start, end

    for start <= end {
        mid := (start + end) / 2

        if mid > 0 && nums[mid] == target && nums[mid] != nums[mid-1] {
            startPos = mid
            break
        }

        if nums[mid] >= target {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }

    if startPos == 0 && nums[0] != target {
        return []int{-1, -1}
    }

    start, end = startPos, len(nums) - 1

    for start <= end {
        mid := (start + end) / 2

        if mid == len(nums) - 1 {
            break
        }

        if mid < len(nums) - 1 && nums[mid] == target && nums[mid] != nums[mid+1] {
            endPos = mid
            break
        }

        if nums[mid] > target {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }

    return []int{startPos, endPos}
}