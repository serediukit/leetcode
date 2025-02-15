package biweekly

func sumOfGoodNumbers(nums []int, k int) int {
    l := len(nums)

    sum := 0

    for i := 0; i < l; i++ {
        f := true

        if i - k >= 0 {
            f = (nums[i] > nums[i-k]) && f
        }

        if i + k < l {
            f = (nums[i] > nums[i+k]) && f
        }

        if f {
            sum += nums[i]
        }
    }

    return sum
}©leetcode