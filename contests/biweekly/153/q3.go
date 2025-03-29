package biweekly

func minimumCost(nums []int, cost []int, k int) int64 {
    n := len(nums)

    preNum := make([]int64, n+1)
    preCost := make([]int64, n+1)
    for i := 0; i < n; i++ {
        preNum[i+1] = preNum[i] + int64(nums[i])
        preCost[i+1] = preCost[i] + int64(cost[i])
    }

    dp := make([][]int64, n+1)
    for i := range dp {
        dp[i] = make([]int64, n+1)
        for j := range dp[i] {
            dp[i][j] = 1 << 60
        }
    }
    dp[0][0] = 0

    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            for l := j - 1; l < i; l++ {
                diff := (preNum[i] + int64(k)*int64(j)) * (preCost[i] - preCost[l])
                if dp[l][j-1]+diff < dp[i][j] {
                    dp[i][j] = dp[l][j-1] + diff
                }
            }
        }
    }

    minCost := int64(1 << 60)
    for j := 1; j <= n; j++ {
        if dp[n][j] < minCost {
            minCost = dp[n][j]
        }
    }

    return minCost
}©leetcode