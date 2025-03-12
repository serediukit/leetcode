package leetcode_ans

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    res := ""
    for row := 0; row < numRows; row++ {
        for i := row; i < len(s); i += numRows * 2 - 2 {
            res += string(s[i])
            if row > 0 && row < numRows - 1 {
                nextI := i + (numRows - row - 1) * 2
                if nextI < len(s) {
                    res += string(s[nextI])
                }
            }
        }
    }
    return res
}