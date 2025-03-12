package leetcode_ans

func mySqrt(x int) int {
    start, end := 0, 65536

    for start <= end {
        mid := (start + end) / 2

        square := mid * mid

        if square == x {
            return mid
        } else if square < x {
            if square + 2 * mid + 1 > x {
                return mid
            }
            start = mid + 1
        } else {
            end = mid - 1
        }
    }

    return 0
}