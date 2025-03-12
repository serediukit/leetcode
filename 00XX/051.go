package leetcode_ans

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
	c1 := col - row
	c2 := col + row
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
		if c1+i >= 0 && board[i][c1+i] == 'Q' {
			return false
		}
		if c2-i < n && board[i][c2-i] == 'Q' {
			return false
		}
	}
	return true
}