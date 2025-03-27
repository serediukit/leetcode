package leetcode_ans

func minimumIndex(nums []int) int {
    countRight := make(map[int]int, 0)
    countLeft := make(map[int]int, 0)

    for _, num := range nums {
        countRight[num]++
    }

    for i, num := range nums {
        countLeft[num]++
        countRight[num]--

        if countLeft[num] > (i + 1) / 2 && countRight[num] > (len(nums) - i - 1) / 2 {
            return i
        }
    }

    return -1
}