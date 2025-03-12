package leetcode_ans

func myPow(x float64, n int) float64 {
    var isMinus bool = n < 0

    if isMinus {
        n = -n
    }

    var res float64 = 1

    for n > 0 {
        if n % 2 == 0 {
            n /= 2
            x *= x
        } else {
            n--
            res *= x
        }
    }

    if isMinus {
        return 1 / res
    }

    return res
}