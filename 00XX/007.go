package leetcode_ans

func reverse(x int) int {
    sign := 1
    if x < 0 {
        sign = -1
    }
    x = abs(x)
    var res int64 = 0
    for x > 0 {
        lastDigit := x % 10
        x /= 10
        res *= 10
        res += int64(lastDigit)
    }
    res *= int64(sign)
    if res >= leftBound() && res <= rightBound() {
        return int(res)
    }
    return 0
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func leftBound() int64 {
    var res int64 = 1
    for i := 0; i < 31; i++ {
        res *= 2
    }
    return -res
}

func rightBound() int64 {
    var res int64 = 1
    for i := 0; i < 31; i++ {
        res *= 2
    }
    return res - 1
}