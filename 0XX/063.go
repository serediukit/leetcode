package leetcode_ans

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[0][0] == 1 {
        return 0
    }

    m := len(obstacleGrid)
    n := len(obstacleGrid[0])

    dp := make([][]int, m)

    for i := range m {
        dp[i] = make([]int, n)
    }

    for i := range m {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = 1
    }

    for j := range n {
        if obstacleGrid[0][j] == 1 {
            break
        }
        dp[0][j] = 1
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] == 0 {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }

    return dp[m-1][n-1]
}