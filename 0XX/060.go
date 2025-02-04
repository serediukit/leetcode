package leetcode_ans

func getPermutation(n int, k int) string {
	if n == 1 {
		return strconv.Itoa(n)
	}
	factorials := make([]int, n)
	factorials[1] = 1
	for i := 2; i < n; i++ {
		factorials[i] = factorials[i-1] * i
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte(i + 49)
	}

	for i := 0; i < n-1; i++ {
		if k == 1 {
			break
		}
		index := i
		for k > factorials[n-i-1] {
			k -= factorials[n-i-1]
			index++
		}
		var temp []byte
		temp = append(temp, res[:i]...)
		temp = append(temp, res[index])
		temp = append(temp, res[i:index]...)
		res = append(temp, res[index+1:]...)
	}

	return string(res)
}