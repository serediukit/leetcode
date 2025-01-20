package leetcode_ans

func countAndSay(n int) string {
	res := "1"

	for i := 1; i < n; i++ {
		res = RLE(res)
	}

	return res
}

func RLE(s string) string {
	if len(s) == 1 {
		return "1" + s
	}

	res := ""

	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			res += strconv.Itoa(cnt) + string(s[i-1])
			cnt = 1
		}
	}
    res += strconv.Itoa(cnt) + string(s[len(s)-1])

	return res
}
