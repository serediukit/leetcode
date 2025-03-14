package leetcode_ans

func maximumCandies(candies []int, k int64) int {
    var sum int64 = 0
    for _, c := range candies {
        sum += int64(c)
    }

    left, right := 0, int(sum / k)
    res := 0

    for left <= right {
        mid := (left + right) / 2
        if mid == 0 {
            left++
            continue
        }

        var kCount int64

        for _, c := range candies {
            kCount += int64(c / mid)
        }

        if kCount >= k {
            left = mid + 1
            res = mid
        } else {
            right = mid - 1
        }
    }

    return res
}