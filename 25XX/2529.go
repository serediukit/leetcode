package leetcode_ans

func maximumCount(nums []int) int {
    var zeroes, negatives int

    for _, n := range nums {
        if n < 0 {
            negatives++
        } else if n == 0 {
            zeroes++
        }
    }

    return max(negatives, len(nums) - negatives - zeroes)
}