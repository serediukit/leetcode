package biweekly

func maxActiveSectionsAfterTrade(s string) int {
	n := len(s)
	maxC := strings.Count(s, "1")

	ans := 0
	i := 0
	for i < n {
		if s[i] == '1' {
			start := i
			for i < n && s[i] == '1' {
				i++
			}
			end := i - 1
			if start > 0 && end < n-1 {
				leftZeros := 0
				j := start - 1
				for j >= 0 && s[j] == '0' {
					leftZeros++
					j--
				}
				rightZeros := 0
				j = end + 1
				for j < n && s[j] == '0' {
					rightZeros++
					j++
				}
				more := leftZeros + rightZeros
				if more > ans {
					ans = more
				}
			}
		} else {
			i++
		}
	}
	return maxC + ans
}