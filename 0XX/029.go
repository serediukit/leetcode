package leetcode_ans

func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

    isMinus := (dividend > 0) != (divisor > 0)

	res := 0
	dividend = abs(dividend)
	divisor = abs(divisor)
	for dividend >= divisor {
		sub := divisor
		add := 1
		for dividend >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		dividend -= sub
		res += add
	}

	if isMinus {
		return -res
	}
	return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}