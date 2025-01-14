package leetcode_ans

import "sort"

func threeSumClosest(nums []int, target int) int {
    closest := nums[0] + nums[1] + nums[2]
	l := len(nums)
    sort.Ints(nums)

    for i := 0; i < l; i++ {
        j := i + 1
        k := l - 1

        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == target {
                return target
            }
            if abs(target - sum) < abs(target - closest) {
                closest = sum
            }
            if sum < target {
                j++
            }
            if sum > target {
                k--
            }
        }
    }
    return closest
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}