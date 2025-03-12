package leetcode_ans

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(board, word, i, j, 1) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j, pos int) bool {
	if pos == len(word) {
		return true
	}

	res := false
	temp := board[i][j]
	board[i][j] = '#'

	if i > 0 {
		if board[i-1][j] == word[pos] {
			res = dfs(board, word, i-1, j, pos+1)
			if res {
				return true
			}
		}
	}
	if i < len(board)-1 {
		if board[i+1][j] == word[pos] {
			res = dfs(board, word, i+1, j, pos+1)
			if res {
				return true
			}
		}
	}
	if j > 0 {
		if board[i][j-1] == word[pos] {
			res = dfs(board, word, i, j-1, pos+1)
			if res {
				return true
			}
		}
	}
	if j < len(board[0])-1 {
		if board[i][j+1] == word[pos] {
			res = dfs(board, word, i, j+1, pos+1)
			if res {
				return true
			}
		}
	}
	board[i][j] = temp

	return false
}