package leetcode_ans

func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    backtrack(candidates, target, 0, 0, []int{}, &res)
    return res
}

func backtrack(candidates []int, target int, start int, sum int, path []int, res *[][]int) {
    if sum == target {
        *res = append(*res, append([]int{}, path...))
        return
    }

    for i := start; i < len(candidates); i++ {
        if sum+candidates[i] > target {
            continue
        }
        backtrack(candidates, target, i, sum+candidates[i], append(path, candidates[i]), res)
    }
}