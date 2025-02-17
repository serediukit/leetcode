package leetcode_ans

func combine(n int, k int) [][]int {
    res := [][]int{}

    backtrack(&res, []int{}, 1, k, n)

    return res
}

func backtrack(res *[][]int, temp []int, num, k, n int) {
    if len(temp) == k {
        *res = append(*res, append([]int{}, temp...))
        return
    }
    for i := num; i <= n; i++ {
        backtrack(res, append(temp, i), i + 1, k, n)
    }
}