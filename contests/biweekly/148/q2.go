package contests

func minCost(arr []int, brr []int, k int64) int64 {
	n := len(arr)

	totalAdjustmentCost := 0
	for i := 0; i < n; i++ {
		totalAdjustmentCost += int(math.Abs(float64(arr[i] - brr[i])))
	}

	sort.Ints(arr)
	sort.Ints(brr)

	totalCost := 0
	for i := 0; i < n; i++ {
		totalCost += int(math.Abs(float64(arr[i] - brr[i])))
	}

	if int64(totalCost)+k < int64(totalAdjustmentCost) {
		return int64(totalCost) + k
	}

	return int64(totalAdjustmentCost)
}