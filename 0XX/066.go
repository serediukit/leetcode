package leetcode_ans

func plusOne(digits []int) []int {
    res := []int{}

    x := 1
    for i := len(digits) - 1; i >= 0; i-- {
        temp := digits[i] + x
        if temp == 10 {
            x = 1
            res = append(res, temp % 10)
        } else {
            x = 0
            res = append(res, temp)
        }
    }
    if x == 1 {
        res = append(res, 1)
    }

    return reverse(res)
}

func reverse(array []int) []int {
    l := len(array)
    for i := 0; i < l / 2; i++ {
        array[i], array[l - i - 1] = array[l - i - 1], array[i]
    }
    return array
}