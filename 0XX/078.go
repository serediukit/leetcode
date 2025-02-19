package leetcode_ans

func subsets(nums []int) [][]int {
    res := [][]int{{}}
    l := len(nums)

    for i := 1; i <= l; i++ {
        temp := combine(nums, l, i)
        res = append(res, temp...)
    }

    return res
}

func combine(nums []int, n int, k int) [][]int {
    size := 1
    for i := 0; i < k; i++ {
        size = size * (n - i) / (i + 1)
    }

    res := make([][]int, 0, size)
    temp := make([]int, k)

    var backtrack func(int, int)

    backtrack = func(start, pos int) {
        if pos == k {
            res = append(res, append([]int(nil), temp...))
            return
        }

        for i := start; i <= n - k + pos; i++ {
            temp[pos] = nums[i]
            backtrack(i + 1, pos + 1)
        }
    }

    backtrack(0, 0)
    return res
}