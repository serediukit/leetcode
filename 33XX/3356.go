package leetcode_ans

func minZeroArray(nums []int, queries [][]int) int {
    cntZeroes := 0
    for _, n := range nums {
        if n == 0 {
            cntZeroes++
        }
    }
    k := 0
    for _, q := range queries {
        if cntZeroes < len(nums) {
            for i := q[0]; i <= q[1]; i++ {
                if nums[i] != 0 {
                    if nums[i] <= q[2] {
                        nums[i] = 0
                        cntZeroes++
                    } else {
                        nums[i] -= q[2]
                    }
                }
            }
            k++
        }
        if cntZeroes == len(nums) {
            return k
        }
    }
    return -1
}