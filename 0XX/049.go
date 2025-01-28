package leetcode_ans

func groupAnagrams(strs []string) [][]string {
    templates := make(map[string][]string)

    for _, str := range strs {
        chars := strings.Split(str, "")
        sort.Strings(chars)
        sortedString := strings.Join(chars, "")
        templates[sortedString] = append(templates[sortedString], str)
    }

    res := [][]string{}
    for _, v := range templates {
        res = append(res, v)
    }

    return res
}