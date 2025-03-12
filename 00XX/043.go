package leetcode_ans

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    l1 := len(num1)
    l2 := len(num2)

    res := make([]byte, l1 + l2)

    for i := l1-1; i >= 0; i-- {
        for j := l2-1; j >= 0; j-- {
            temp := (num1[i] - 48) * (num2[j] - 48)
            res[i+j+1] += temp

            if res[i+j+1] >= 10 {
                res[i+j] += res[i+j+1] / 10
                res[i+j+1] %= 10
            }
        }
    }

    if res[0] == 0 {
        res = res[1:]
    }

    for i := 0; i < len(res); i++ {
        res[i] += 48
    }

    return string(res)
}