package leetcode_ans

func romanToInt(s string) int {
    romanMap := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

    res := 0
    for i := 0; i < len(s); i++ {
        if i < len(s) - 1 {
            if romanMap[s[i]] < romanMap[s[i+1]] {
                res += romanMap[s[i+1]] - romanMap[s[i]]
                i++
                continue
            }
        }
        res += romanMap[s[i]]
    }

    return res
}