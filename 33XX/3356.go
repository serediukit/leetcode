package leetcode_ans

func minZeroArray(nums []int, queries [][]int) int {
    n := len(nums)
    sum, k := 0, 0
    cnt := make([]int, n+1)
    for i := 0; i < n; i++ {
        for sum+cnt[i] < nums[i] {
            if k == len(queries) {
                return -1
            }
            l, r, val := queries[k][0], queries[k][1], queries[k][2]
            k++
            if r < i {
                continue
            }
            cnt[max(l, i)] += val
            cnt[r+1] -= val
        }
        sum += cnt[i]
    }
    return k
}