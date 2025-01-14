package leetcode_ans

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    minLength := findMinLength(strs)
    length := findPrefix(strs, 0, minLength)
    return strs[0][:length]
}

func findPrefix(strs []string, low int, high int) int {
    for low < high {
        mid := (low + high + 1) / 2
        if isCommonPrefix(strs, mid) {
            low = mid
        } else {
            high = mid - 1
        }
    }
    return low
}

func isCommonPrefix(strs []string, length int) bool {
    for i := 1; i < len(strs); i++ {
        for j := 0; j < length; j++ {
            if strs[i][j] != strs[0][j] {
                return false
            }
        }
    }
    return true
}

func findMinLength(strs []string) int {
    minLength := len(strs[0])
    for _, str := range strs {
        if len(str) < minLength {
            minLength = len(str)
        }
    }
    return minLength
}