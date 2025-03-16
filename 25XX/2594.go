package leetcode_ans

func repairCars(ranks []int, cars int) int64 {
    var left, right int64 = 0, 1_000_000_000_000_000

    for left < right {
        mid := (left + right) / 2

        countRepair := 0

        for _, r := range ranks {
            countRepair += int(math.Sqrt(float64(mid) / float64(r)))

            if countRepair >= cars {
                break
            }
        }

        if countRepair >= cars {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}