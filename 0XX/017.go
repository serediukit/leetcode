package leetcode_ans

func letterCombinations(digits string) []string {
    nums := map[int]string{
        2: "abc",
        3: "def",
        4: "ghi",
        5: "jkl",
        6: "mno",
        7: "pqrs",
        8: "tuv",
        9: "wxyz",
    }

    l := len(digits)

    res := make([]string, 0)

    if 1 <= l {
        for _, s1 := range nums[int(digits[0]) - 48] {
            if 2 <= l {
                for _, s2 := range nums[int(digits[1]) - 48] {
                    if 3 <= l {
                        for _, s3 := range nums[int(digits[2]) - 48] {
                            if 4 <= l {
                                for _, s4 := range nums[int(digits[3]) - 48] {
                                    res = append(res, string(s1)+string(s2)+string(s3)+string(s4))
                                }
                            } else {
                                res = append(res, string(s1)+string(s2)+string(s3))
                            }
                        }
                    } else {
                        res = append(res, string(s1)+string(s2))
                    }
                }
            } else {
                res = append(res, string(s1))
            }
        }
    }

    return res
}