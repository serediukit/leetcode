package leetcode_ans

func uniquePaths(m int, n int) int {
    dp := make([][]int, m)

    for i := range m {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }

    for j := range n {
        dp[0][j] = 1
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}