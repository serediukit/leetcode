package leetcode_ans

func intToRoman(num int) string {
    letter := [7]string{"M", "D", "C", "L", "X", "V", "I"}
    output := ""
    for i := 3; i > 0; i-- {
        for num >= pow(10, i) {
            output += letter[6 - 2*i]
            num -= pow(10, i)
        }
        if num >= pow(10, i) - pow(10, i - 1) {
            output += letter[8 - 2 * i] + letter[6 - 2 * i]
            num -= pow(10, i) - pow(10, i - 1)
        }
        for num >= pow(10, i) / 2 {
            output += letter[7 - 2*i]
            num -= pow(10, i) / 2
        }
        if num >= pow(10, i) / 2 - pow(10, i - 1) {
            output += letter[8 - 2 * i] + letter[7 - 2 * i]
            num -= pow(10, i) / 2 - pow(10, i - 1)
        }
    }
    for num > 0 {
        output += "I"
        num--
    }
    return output
}

func pow(a, b int) int {
    res := 1
    for i := 0; i < b; i++ {
        res *= a
    }
    return res
}