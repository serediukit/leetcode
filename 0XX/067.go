package leetcode_ans

func addBinary(a string, b string) string {
	res := ""

	a = Reverse(a)
	b = Reverse(b)

	l := max(len(a), len(b))

	var ost byte = '0'
	var sum string
	for i := 0; i < l; i++ {
		if i < len(a) {
			if i < len(b) {
				sum, ost = sumBin(a[i], b[i], ost)
				res = sum + res
			} else {
				sum, ost = sumBin(a[i], '0', ost)
				res = sum + res
			}
		} else if i < len(b) {
			sum, ost = sumBin('0', b[i], ost)
			res = sum + res
		}
	}

	if ost == '1' {
		res = "1" + res
	}

	return res
}

func sumBin(a byte, b byte, ost byte) (string, byte) {
	n1 := a - 48
	n2 := b - 48
	n3 := ost - 48

	sum := n1 + n2 + n3

	r1 := (sum % 2) + 48
	r2 := (sum / 2) + 48

	return string(r1), r2
}

func Reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}
