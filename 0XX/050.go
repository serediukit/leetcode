package leetcode_ans

func myPow(x float64, n int) float64 {
    if x == 0 || x == 1 {
        return x
    }
    var isMinus bool
    n, isMinus = abs(n)

    var res float64 = 1

    for i := 1; i <= n; i++ {
        res *= x
    }

    if isMinus {
        return 1 / res
    }

    return res
}