package leetcode_ans

import "sort"

func fourSum(nums []int, target int) [][]int {
	var res = make([][]int, 0)
	l := len(nums)
    sort.Ints(nums)
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
        }
        for h := i+1; h < l; h++ {
			if h > i+1 && nums[h] == nums[h-1] {
            	continue
	        }

	        j := h + 1
	        k := l - 1

	        for j < k {
	            sum := nums[i]+nums[h]+nums[j]+nums[k]
	            if sum == target {
	                a := []int{nums[i], nums[h], nums[j], nums[k]}
	                sort.Ints(a)
	                res = append(res, a)
	                j++
	                for j < k && nums[j] == nums[j-1] {
	                    j++
	                }
	                k--
	                for j < k && nums[k] == nums[k+1] {
	                    k--
	                }
	            } else if sum < target {
	                j++
	            } else {
	                k--
	            }
	        }
		}
	}
    return res
}