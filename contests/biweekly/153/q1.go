package biweekly

func reverseDegree(s string) int {
	res := 0
	for i, c := range s {
		res += (123 - int(c)) * (i + 1)
	}
	return res
}