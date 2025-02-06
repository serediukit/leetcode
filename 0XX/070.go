package leetcode_ans

var history = map[int]int{}

func climbStairs(n int) int {
    if n <= 2 {
        return n
    }

    if value, ok := history[n]; ok {
        return value
    }

    res := climbStairs(n-1) + climbStairs(n-2)
    history[n] = res
    return res
}