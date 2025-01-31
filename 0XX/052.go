package leetcode_ans

func totalNQueens(n int) int {
    return len(solveNQueens(n))
}

func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([]string, n)
	for i := range board {
		board[i] = strings.Repeat(".", n)
	}
	backtrack(&res, board, 0, n)

	return res
}

func backtrack(res *[][]string, board []string, row int, n int) {
	if row == n {
		temp := make([]string, n)
		copy(temp, board)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < n; i++ {
		if isValidPosition(board, row, i, n) {
			tempRow := []byte(board[row])
			tempRow[i] = 'Q'
			board[row] = string(tempRow)

			backtrack(res, board, row+1, n)

			tempRow[i] = '.'
			board[row] = string(tempRow)
		}
	}
}

func isValidPosition(board []string, row int, col int, n int) bool {
	for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }

    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }

    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
	return true
}
