package leetcode

func countGoodTriplets(arr []int, a int, b int, c int) int {
    num := 0
    for i := 0; i < len(arr) - 2; i++ {
        for j := i + 1; j < len(arr) - 1; j++ {
            for k := j + 1; k < len(arr); k++ {
                if abs(arr[i] - arr[j]) <= a && abs(arr[j] - arr[k]) <= b && abs(arr[k] - arr[i]) <= c {
                    num++
                }
            }
        }
    }
    return num
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}