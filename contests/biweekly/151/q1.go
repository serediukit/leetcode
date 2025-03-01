package biweekly

func transformArray(nums []int) []int {
    c := 0
    for _, n := range nums {
        if n % 2 == 0 {
            c++
        }
    }
    return append(slices.Repeat([]int{0}, c), slices.Repeat([]int{1}, len(nums) - c)...)
}