package biweekly

func shortestMatchingSubstring(s string, p string) int {
    parts := strings.Split(p, "*")
	if len(parts) != 3 {
		return -1
	}

	prefix, mid, suffix := parts[0], parts[1], parts[2]

	minLength := math.MaxInt32

	for i := 0; i <= len(s); i++ {
		if !strings.HasPrefix(s[i:], prefix) {
			continue
		}

		startMid := i + len(prefix)
		indexMid := strings.Index(s[startMid:], mid)
		if indexMid == -1 {
			continue
		}
		indexMid += startMid

		startSuffix := indexMid + len(mid)
		indexSuffix := strings.Index(s[startSuffix:], suffix)
		if indexSuffix == -1 {
			continue
		}
		indexSuffix += startSuffix + len(suffix)

		matchLength := indexSuffix - i
		if matchLength < minLength {
			minLength = matchLength
		}
	}

	if minLength == math.MaxInt32 {
		return -1
	}
	return minLength
}