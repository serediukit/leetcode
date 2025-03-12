package leetcode_ans

func largestRectangleArea(heights []int) int {
    s := []int{0}
    maxSquare := 0
    l := len(heights)

    for i := 1; i <= l; i++ {
        var height int
        if i < l {
            height = heights[i]
        }
        
        for len(s) > 0 && heights[s[len(s) - 1]] > height {
            h := heights[s[len(s) - 1]]
            s = s[:len(s) - 1]
            w := i
            if len(s) > 0 {
                w = i - s[len(s)-1] - 1
            }
            maxSquare = max(maxSquare, h * w)
        }

        s = append(s, i)
    }

    return maxSquare
}
