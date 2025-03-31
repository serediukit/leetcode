package leetcode_ans

func putMarbles(weights []int, k int) int64 {
    if k == 1 {
        return 0
    }

    pairs := make([]int, len(weights) - 1)

    for i := 0; i < len(weights) - 1; i++ {
        pairs[i] = weights[i] + weights[i + 1]
    }

    sort.Ints(pairs)

    var left, right int64

    for i := 0; i < k - 1; i++ {
        left += int64(pairs[i])
        right += int64(pairs[len(pairs) - i - 1])
    }

    return right - left
}