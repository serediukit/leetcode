package leetcode_ans

func trap(height []int) int {
    index := maxIndex(height)

    return leftSideWater(height[:index]) + rightSideWater(height[index+1:])
}

func leftSideWater(height []int) int {
    l := len(height)

    if l <= 1 {
        return 0
    }

    res := 0
    minHeight := height[0]

    for i := 1; i < len(height); i++ {
        if minHeight >= height[i] {
            res += minHeight - height[i]
        } else {
            minHeight = height[i]
        }
    }

    return res
}

func rightSideWater(height []int) int {
    l := len(height)

    if l <= 1 {
        return 0
    }

    res := 0
    minHeight := height[l-1]

    for i := l-2; i >= 0; i-- {
        if minHeight >= height[i] {
            res += minHeight - height[i]
        } else {
            minHeight = height[i]
        }
    }

    return res
}

func maxIndex(array []int) int {
	index := 0
	for i := 1; i < len(array); i++ {
		if array[i] > array[index] {
			index = i
		}
	}
	return index
}