package leetcode_ans

func fullJustify(words []string, maxWidth int) []string {
	output := []string{}

	tempLen := 0
	tempStr := []string{}
	for _, str := range words {
		if tempLen == 0 {
			tempStr = append(tempStr, str)
			tempLen = len(str)
		} else if tempLen+len(str)+1 <= maxWidth {
			tempStr = append(tempStr, str)
			tempLen += len(str) + 1
		} else {
			output = append(output, alignWords(tempStr, tempLen, maxWidth))
			tempStr = []string{str}
			tempLen = len(str)
		}
	}

	output = append(output, alignLastWords(tempStr, maxWidth))

	return output
}

func alignWords(words []string, l, max int) string {
	n := len(words)

	if n == 1 {
		o := max - l
		return words[0] + strings.Repeat(" ", o)
	}

	c := (max - l) / (n - 1)
	o := (max - l) % (n - 1)

	res := words[0]

	for i := 1; i < n; i++ {
		res += strings.Repeat(" ", c+1)
		if i <= o {
			res += " "
		}
		res += words[i]
	}

	return res
}

func alignLastWords(words []string, maxWidth int) string {
	res := words[0]
	for i := 1; i < len(words); i++ {
		res += " " + words[i]
	}

	return res + strings.Repeat(" ", maxWidth-len(res))
}
