package leetcode_ans

func isValid(s string) bool {
    if len(s)%2 == 1 {
		return false
	}

	stack := []rune{}
    brackets := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, char := range s {
        if char == ')' || char == ']' || char == '}' {
            if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, char)
        }
    }
    return len(stack) == 0
}