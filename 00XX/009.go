package leetcode_ans

func isPalindrome(x int) bool {
    if (x < 0) {
        return false
    }

    return x == reverseNum(x)
}

func reverseNum(x int) int {
    res := 0
    for x > 0 {
        res *= 10
        res += x % 10
        x /= 10
    }
    return res
}