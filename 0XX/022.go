package leetcode_ans

import "strings"

func generateParenthesis(n int) []string {
	return generateVariants("(", n)
}

func generateVariants(s string, n int) []string {
	variants := []string{}
	cntOpen := strings.Count(s, "(")
	cntClose := strings.Count(s, ")")

	if cntClose == n {
		return []string{s}
	}

	if cntOpen == n {
		return generateVariants(s+")", n)
	}

	variants = append(variants, generateVariants(s+"(", n)...)
	if cntClose < cntOpen {
		variants = append(variants, generateVariants(s+")", n)...)
	}

	return variants
}