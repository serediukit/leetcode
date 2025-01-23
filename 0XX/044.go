package leetcode_ans

func isMatch(s string, p string) bool {
    var lastStar, frontToBack bool
    var helper func(s, p string) bool
    helper = func(s, p string) bool {
        if len(p) == 0 {
            return len(s) == 0
        }
        var matched bool
        if frontToBack {
			if len(s) > 0 && (p[0] == '?' || s[0] == p[0]) {
				matched = helper(s[1:], p[1:])
			} else if p[0] == '*' {
				matched = len(p) == 1 || helper(s, p[1:])
				if !lastStar && !matched && len(s) > 0 {
					matched = helper(s[1:], p)
					lastStar = true
				}
			}
        } else {
            if len(s) > 0 && (p[len(p)-1] == '?' || s[len(s)-1] == p[len(p)-1]) {
                matched = helper(s[:len(s)-1], p[:len(p)-1])
            } else if p[len(p)-1] == '*' {
    			frontToBack = true
                matched = len(p) == 1 || helper(s, p) //Matching from front to back
            }
		}
        return matched
    }
    return helper(s, p)
}