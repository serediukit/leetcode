package leetcode_ans

func maximumCandies(candies []int, k int64) int {
    var sum int64 = 0
    for _, c := range candies {
        sum += int64(c)
    }

    left, right := 0, int(sum / k)

    for left < right {
        mid := (left + right) / 2 + 1

        var kCount int64

        for _, c := range candies {
            kCount += int64(c / mid)
        }

        if kCount >= k {
            left = mid
        } else {
            right = mid - 1
        }
    }

    return left
}