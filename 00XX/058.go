package leetcode_ans

func lengthOfLastWord(s string) int {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            s = s[:i+1]
            break
        }
    }
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' ' {
            return len(s) - i - 1
        }
    }
    return len(s)
}