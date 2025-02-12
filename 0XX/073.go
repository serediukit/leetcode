package leetcode_ans

func setZeroes(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])

    zeroRow := make([]int, n)

    rows := make([]bool, m)
    cols := make([]bool, n)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                rows[i] = true
                cols[j] = true
            }
        }
    }

    for i := 0; i < m; i++ {
        if rows[i] {
            copy(matrix[i], zeroRow)
        }
    }

    for j := 0; j < n; j++ {
        if cols[j] {
            for i := 0; i < m; i++ {
                matrix[i][j] = 0
            }
        }
    }
}