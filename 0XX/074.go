package leetcode_ans

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left, right := 0, m-1

	for left <= right {
		midRow := (left + right) / 2

		if matrix[midRow][0] <= target {
			if midRow+1 == m || matrix[midRow+1][0] > target {
				left, right = 0, n-1

				for left <= right {
					mid := (left + right) / 2

					if matrix[midRow][mid] == target {
						return true
					} else if matrix[midRow][mid] < target {
						left = mid + 1
					} else {
						right = mid - 1
					}
				}

				break
			} else {
				left = midRow + 1
			}
		} else {
			right = midRow - 1
		}
	}

	return false
}