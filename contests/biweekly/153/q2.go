package biweekly

func maxActiveSections(s string) int {
	s = "1" + s + "1"
	count := make([]int, 0)
	cMax := strings.Count(s, "1")

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
		temp := make([]int, len(count))
		copy(temp, count)
		newCount := make([]int, len(temp)-2)
		newCount = append(append(temp[:i-1], temp[i-1]+temp[i]+temp[i+1]), temp[i+2:]...)

		for j := 1; j < len(newCount)-1; j += 2 {
			temp2 := make([]int, len(newCount))
			copy(temp2, newCount)
			newCount2 := make([]int, len(temp2)-2)
			newCount2 = append(temp2[:j-1], temp2[j-1]+temp2[j]+temp2[j+1])
			if j < len(newCount)-2 {
				newCount2 = append(newCount2, temp2[j+2:]...)
			}

			sum := 0
			for k := 0; k < len(newCount2); k += 2 {
				sum += newCount2[k]
			}
			if sum > cMax {
				cMax = sum
			}
		}
	}

	return cMax - 2
}