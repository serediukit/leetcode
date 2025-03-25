package leetcode_ans

func minOperations(grid [][]int, x int) int {
    rest := grid[0][0] % x
    grids := make([]int, len(grid) * len(grid[0]))

    index := 0
    for _, g := range grid {
        for _, i := range g {
            if rest != i % x {
                return -1
            }
            grids[index] = i
            index++
        }
    }

    sort.Ints(grids)

    mid := grids[len(grids) / 2]

    c := 0
    for _, g := range grids {
        c += abs(g - mid) / x
    }

    return c
}

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}