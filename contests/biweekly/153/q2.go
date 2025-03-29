package biweekly

func maxActiveSectionsAfterTrade(s string) int {
	s = "1" + s + "1"
	cMax := strings.Count(s, "1")

	count := make([]int, 0)

	c := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			c++
		} else {
			count = append(count, c)
			c = 1
		}
	}
	count = append(count, c)

	if len(count) < 5 {
		return cMax - 2
	}

	for i := 2; i < len(count)-2; i += 2 {
		mergedZeroes := make([]int, 0, len(count)-2)
		mergedZeroes = append(mergedZeroes, count[:i-1]...)
		mergedZeroes = append(mergedZeroes, count[i-1]+count[i]+count[i+1])
		mergedZeroes = append(mergedZeroes, count[i+2:]...)

		for j := 1; j < len(mergedZeroes)-1; j += 2 {
			mergedOnes := make([]int, 0, len(mergedZeroes)-2)
			mergedOnes = append(mergedOnes, mergedZeroes[:j-1]...)
			mergedOnes = append(mergedOnes, mergedZeroes[j-1]+mergedZeroes[j]+mergedZeroes[j+1])
			if j < len(mergedZeroes)-2 {
				mergedOnes = append(mergedOnes, mergedZeroes[j+2:]...)
			}

			sum := 0
			for k := 0; k < len(mergedOnes); k += 2 {
				sum += mergedOnes[k]
			}
			if sum > cMax {
				cMax = sum
			}
		}
	}

	return cMax - 2
}