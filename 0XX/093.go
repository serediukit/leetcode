package leetcode_ans

func restoreIpAddresses(s string) []string {
	ls := len(s)
	if ls < 4 {
		return []string{}
	}
	var res []string
	for i := 1; i <= min(3, ls-1); i++ {
		s1 := s[:i]
		if len(s1) > 1 && s1[0] == '0' {
			continue
		}
		n1, _ := strconv.Atoi(s1)
		if n1 >= 0 && n1 < 256 {
			for j := i + 1; j <= min(i+3, ls-1); j++ {
				s2 := s[i:j]
				if len(s2) > 1 && s2[0] == '0' {
					continue
				}
				n2, _ := strconv.Atoi(s2)
				if n2 >= 0 && n2 < 256 {
					for k := j + 1; k <= min(j+3, ls-1); k++ {
						s3 := s[j:k]
						if len(s3) > 1 && s3[0] == '0' {
							continue
						}
						n3, _ := strconv.Atoi(s3)
						if n3 >= 0 && n3 < 256 {
							s4 := s[k:ls]
							if len(s4) > 1 && s4[0] == '0' {
								continue
							}
							n4, _ := strconv.Atoi(s4)
							if n4 >= 0 && n4 < 256 {
								res = append(res, fmt.Sprintf("%d.%d.%d.%d", n1, n2, n3, n4))
							}
						}
					}
				}
			}
		}
	}
	return res
}
