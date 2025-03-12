package leetcode_ans

func isValidSudoku(board [][]byte) bool {
	return checkRows(board) && checkColumns(board) && checkSquares(board)
}

func checkRows(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		quantity := map[byte]int{}
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			quantity[board[i][j]]++
			if quantity[board[i][j]] > 1 {
				return false
			}
		}
	}
	return true
}

func checkColumns(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		quantity := map[byte]int{}
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			quantity[board[j][i]]++
			if quantity[board[j][i]] > 1 {
				return false
			}
		}
	}
	return true
}

func checkSquares(board [][]byte) bool {
	for n := 0; n < 9; n++ {
		quantity := map[byte]int{}
		row := n - n%3
		col := n % 3 * 3
		for i := row; i < row+3; i++ {
			for j := col; j < col+3; j++ {
				if board[i][j] == '.' {
					continue
				}
				quantity[board[i][j]]++
				if quantity[board[i][j]] > 1 {
					return false
				}
			}
		}
	}
	return true
}
