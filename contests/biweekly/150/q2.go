package biweekly

func separateSquares(squares [][]int) float64 {
    calculateArea := func(y float64) float64 {
		above := 0.0
		for _, square := range squares {
			ySquare, l := float64(square[1]), float64(square[2])
			yTop := ySquare + l
			if ySquare >= y {
				above += l * l
			} else if yTop >= y {
				above += (yTop - y) * l
			}
		}
		return above
	}

	totalArea := 0.0
	for _, square := range squares {
		l := float64(square[2])
		totalArea += l * l
	}

	low := float64(squares[0][1])
	high := low
	for _, square := range squares {
		y, l := float64(square[1]), float64(square[2])
		if y < low {
			low = y
		}
		if y+l > high {
			high = y + l
		}
	}

	for high-low > 1e-6 {
		mid := (low + high) / 2
		above := calculateArea(mid)
		if above > totalArea-above {
			low = mid
		} else {
			high = mid
		}
	}

	return (low + high) / 2
}©leetcode