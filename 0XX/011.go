package leetcode_ans

func maxArea(height []int) int {
    maxArea := 0
    start := 0
    end := len(height) - 1
    for start < end {
        width := end - start
        heightMin := min(height[start], height[end])
        area := width * heightMin
        if maxArea < area {
            maxArea = area
        }

        if height[start] <= height[end] {
            start++
        } else {
            end--
        }
    }
    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}