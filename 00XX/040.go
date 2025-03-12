package leetcode_ans

func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int
    sort.Ints(candidates)
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
            break
        }
        if i > start && candidates[i] == candidates[i-1] {
            continue
        }
        backtrack(candidates, target, i+1, sum+candidates[i], append(path, candidates[i]), res)
    }
}