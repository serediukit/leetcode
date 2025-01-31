package leetcode_ans

func spiralOrder(matrix [][]int) []int {
	res := []int{}

	m, n := len(matrix), len(matrix[0])
    minMN := m
    if n < m {
        minMN = n
    }

	for i := 0; i < (minMN+1)/2; i++ {
		rowUp := i
		for j := i; j < n-i; j++ {
			res = append(res, matrix[rowUp][j])
		}
		columnRight := n - i - 1
		for j := i + 1; j < m-i-1; j++ {
			res = append(res, matrix[j][columnRight])
		}
		rowDown := m - i - 1
		if rowUp != rowDown {
			for j := n - i - 1; j >= i; j-- {
				res = append(res, matrix[rowDown][j])
			}
		}
		columnLeft := i
		if columnRight != columnLeft {
			for j := m - i - 2; j > i; j-- {
				res = append(res, matrix[j][i])
			}
		}
	}

	return res
}