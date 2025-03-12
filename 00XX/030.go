package leetcode_ans

func findSubstring(s string, words []string) []int {

	length := len(words[0])
	wordsCount := len(words)
	wordsQuantity := make(map[string]int, wordsCount)

	for _, word := range words {
		wordsQuantity[word]++
	}

	temp := make(map[string]int, wordsCount)
	var found bool
	res := make([]int, 0)

	for i := 0; i+length*wordsCount <= len(s); {

		found = true
		temp = make(map[string]int, wordsCount)

		for j := i; j < i+length*wordsCount; j += length {
			if _, ok := wordsQuantity[s[j:j+length]]; ok {
				temp[s[j:j+length]]++
			} else {
				found = false
				break
			}
		}

		if found {
			for key, _ := range wordsQuantity {
				if val, ok := temp[key]; !ok || val != wordsQuantity[key] {
					i++
					found = false
					break
				}
			}

			if found {
				res = append(res, i)
				i++
			}
		} else {
			i++
		}
	}

	return res
}