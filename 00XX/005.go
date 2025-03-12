package leetcode_ans

func longestPalindrome(s string) string {
    l := len(s)
    for length := l; length > 0; length-- {
        for start := 0; start <= l - length; start++ {
            str := s[start:start+length]
            if isPalindromic(str) {
                return str
            }
        }
    }
    return ""
}

func isPalindromic(s string) bool {
    for i := 0; i < len(s) / 2; i++ {
        if s[i] != s[len(s) - i - 1] {
            return false
        }
    }
    return true
}