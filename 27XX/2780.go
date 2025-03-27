package leetcode_ans

func minimumIndex(nums []int) int {
    dominant := dominantNum(slices.Clone(nums))
    countLeft := 0
    countRight := count(nums, dominant)
    left := 0
    right := len(nums)
    
    for i, num := range nums {
        left++
        right--

        if num == dominant {
            countLeft++
            countRight--
        }

        if 2 * countLeft > left && 2 * countRight > right {
            return i
        }
    }

    return -1
}

func dominantNum(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

func count(nums []int, n int) int {
    count := 0
    for _, num := range nums {
        if num == n {
            count++
        }
    }
    return count
}