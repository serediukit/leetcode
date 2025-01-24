package leetcode_ans

func restoreIpAddresses(s string) []string {
	ls := len(s)
	if ls < 4 {
		return []string{}
	}
	var res []string
	for i := 1; i <= min(3, ls-1); i++ {
		n1, _ := strconv.Atoi(s[:i])
		if n1 >= 0 && n1 < 256 {
			for j := i + 1; j <= min(i+3, ls-1); j++ {
				n2, _ := strconv.Atoi(s[i:j])
				if n2 >= 0 && n2 < 256 {
					for k := j + 1; k <= min(j+3, ls-1); k++ {
						n3, _ := strconv.Atoi(s[j:k])
						if n3 >= 0 && n3 < 256 {
							n4, _ := strconv.Atoi(s[k:ls])
							if n4 >= 0 && n4 < 256 {
								r := fmt.Sprintf("%d.%d.%d.%d", n1, n2, n3, n4)
								if len(r) == len(s) + 3 {
									res = append(res, r)
								}
							}
						}
					}
				}
			}
		}
	}
	return res
}