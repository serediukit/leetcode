package leetcode_ans

func generateMatrix(n int) [][]int {
    res := make([][]int, n)

    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }

    number := 1

    for i := 0; i < (n+1)/2; i++ {
		rowUp := i
		for j := i; j < n-i; j++ {
			res[rowUp][j] = number
            number++
		}
		columnRight := n - i - 1
		for j := i + 1; j < n-i-1; j++ {
            res[j][columnRight] = number
            number++
		}
		rowDown := n - i - 1
		if rowUp != rowDown {
			for j := n - i - 1; j >= i; j-- {
                res[rowDown][j] = number
                number++
			}
		}
		columnLeft := i
		if columnRight != columnLeft {
			for j := n - i - 2; j > i; j-- {
                res[j][i] = number
                number++
			}
		}
	}

	return res
}
