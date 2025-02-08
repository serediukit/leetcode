package leetcode_ans

func isNumber(s string) bool {
	strs := []byte(s)
	i, n := 0, len(strs)

	if i < n && (strs[i] == '+' || strs[i] == '-') {
		i++
	}

	isNumeric := false
	for i < n && isDigit(strs[i]) {
		i++
		isNumeric = true
	}

	if i < n && strs[i] == '.' {
		i++
		for i < n && isDigit(strs[i]) {
			i++
			isNumeric = true
		}
	}

	if isNumeric && i < n && (strs[i] == 'e' || strs[i] == 'E') {
		i++
		isNumeric = false
		if i < n && (strs[i] == '+' || strs[i] == '-') {
			i++
		}
		for i < n && isDigit(strs[i]) {
			i++
			isNumeric = true
		}
	}

	return isNumeric && i == n
}

func isDigit(d byte) bool {
	return d >= '0' && d <= '9'
}