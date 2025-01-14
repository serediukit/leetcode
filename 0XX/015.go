package leetcode_ans

import "sort"

func threeSum(nums []int) [][]int {
	var res = make([][]int, 0)
	l := len(nums)
    sort.Ints(nums)
	for i := 0; i < l; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j := i + 1
        k := l - 1

        for j < k {
            sum := nums[i]+nums[j]+nums[k]
            if sum == 0 {
                a := []int{nums[i], nums[j], nums[k]}
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
            } else if sum < 0 {
                j++
            } else {
                k--
            }
        }
	}
    return res
}

func contain(a [][]int, b []int) bool {
	for _, array := range a {
		f := true
		for i := 0; i < 3; i++ {
			if array[i] != b[i] {
				f = false
                break
			}
		}
		if f {
			return true
		}
	}
	return false
}