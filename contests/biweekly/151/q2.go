package biweekly

func countArrays(original []int, bounds [][]int) int {
	var calc func([]int, int) int
	calc = func(temp []int, step int) int {
		if step == len(original) {
			return len(temp)
		}

		diff := original[step] - original[step-1]

		for i := 0; i < len(temp); i++ {
			temp[i] += diff
			if !(temp[i] >= bounds[step][0] && temp[i] <= bounds[step][1]) {
				temp = append(temp[:i], temp[i+1:]...)
				i--
			}
		}

		return calc(temp, step+1)
	}

	temp := make([]int, bounds[0][1]-bounds[0][0]+1)
	for i := 0; i < len(temp); i++ {
		temp[i] = bounds[0][0] + i
	}
	return calc(temp, 1)
}
