package leetcode_ans

func sortColors(nums []int)  {
    var countZeros, countOnes int

    l := len(nums)

    for _, value := range nums {
        if value == 0 {
            countZeros++
        } else if value == 1 {
            countOnes++
        }
    }

    for i := 0; i < countZeros; i++ {
        nums[i] = 0
    }

    for i := countZeros; i < countZeros + countOnes; i++ {
        nums[i] = 1
    }

    for i := countZeros + countOnes; i < l; i++ {
        nums[i] = 2
    }
}