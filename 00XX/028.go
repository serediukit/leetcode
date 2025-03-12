package leetcode_ans

func strStr(haystack string, needle string) int {
    c := len(haystack) - len(needle) + 1
    l := len(needle)
    for i := 0; i < c; i++ {
        if haystack[i:i+l] == needle {
            return i
        }
    }
    return -1
}