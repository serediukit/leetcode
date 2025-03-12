package leetcode_ans

func permuteUnique(nums []int) [][]int {
    cnt := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        cnt[nums[i]]++
    }

    res := [][]int{}
    permute(&res, cnt, len(nums), []int{})
    return res
}

func permute(res *[][]int, cnt map[int]int, length int, permutation []int) {
    if len(permutation) == length {
        temp := make([]int, length)
        copy(temp, permutation)

        *res = append(*res, temp)
        return
    }

    for num, count := range cnt {
        if count > 0 {
            cnt[num]--

            tempPermutation := append(permutation, num)

            permute(res, cnt, length, tempPermutation)

            cnt[num]++
        }
    }
}