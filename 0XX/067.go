package leetcode_ans

func addBinary(a string, b string) string {
	res := []rune{}

    A := []rune(a)
    B := []rune(b)
	slices.Reverse(A)
    slices.Reverse(B)

	l := max(len(a), len(b))

	var ost rune = '0'
	var sum rune = '0'
	for i := 0; i < l; i++ {
		if i < len(A) {
			if i < len(B) {
				sum, ost = sumBin(A[i], B[i], ost)
			} else {
				sum, ost = sumBin(A[i], '0', ost)
			}
		} else if i < len(B) {
			sum, ost = sumBin('0', B[i], ost)
		}
        res = append(res, sum)
	}

	if ost == '1' {
		res = append(res, '1')
	}

    slices.Reverse(res)

	return string(res)
}

func sumBin(a rune, b rune, ost rune) (rune, rune) {
	n1 := a - 48
	n2 := b - 48
	n3 := ost - 48

	sum := n1 + n2 + n3

	r1 := (sum % 2) + 48
	r2 := (sum / 2) + 48

	return r1, r2
}