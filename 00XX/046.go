package leetcode_ans

func permute(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{{nums[0]}}
    }

    var res [][]int

    for i := 0; i < len(nums); i++ {
        temp := append([]int{}, nums[:i]...)
        temp = append(temp, nums[i+1:]...)

        variants := permute(temp)

        for j := 0; j < len(variants); j++ {
            tempArray := append([]int{}, nums[i])
            tempArray = append(tempArray, variants[j]...)

            res = append(res, tempArray)
        }
    }

    return res
}