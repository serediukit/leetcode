package leetcode_ans

func minPathSum(grid [][]int) int {
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[i]); j++{
            left := math.MaxInt64
            up := math.MaxInt64
            if j - 1 >= 0 {
                left = grid[i][j-1]
            }
            if i - 1 >= 0{
                up = grid[i-1][j]
            }
            if up != math.MaxInt64 || left != math.MaxInt64{
                grid[i][j] += min(up, left)
            }
        }
    }
    return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
    if a < b {return a}
    return b
}