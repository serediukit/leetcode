package leetcode_ans

func myAtoi(s string) int {
    if len(s) == 0 {
        return 0
    }
    for s[0] == ' ' {
        if len(s) > 1 {
            s = s[1:]
        } else {
            break
        }
    }
    sign := int64(1)
    if s[0] == '+' {
        s = s[1:]
    } else if s[0] == '-' {
        s = s[1:]
        sign = int64(-1)
    }
    var result int64
    for len(s) > 0 {
        a := int64(s[0]) - 48
        if a < 0 || a > 9 {
            break
        }

        result *= 10
        result += sign * a

        if result > math.MaxInt32 {
            return math.MaxInt32
        }

        if result < math.MinInt32 {
            return math.MinInt32
        }

        s = s[1:]
    }
    return int(result)
}