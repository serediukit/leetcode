package leetcode_ans

func combine(n int, k int) [][]int {
    res := [][]int{}

    backtrack(&res, []int{}, 1, k, n)

    return res
}

func backtrack(res *[][]int, temp []int, num, k, n int) {
    if len(temp) == k {
        tempRes := make([]int, len(temp))
        copy(tempRes, temp)
        *res = append(*res, tempRes)
        return
    }
    for i := num; i <= n; i++ {
        backtrack(res, append(temp, i), i + 1, k, n)
    }
}