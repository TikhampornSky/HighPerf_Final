package usemergeconcurrent

import "math"

func MinMaxOfAllBlock(data [][]Pair) (float64, float64) {
	max := -1.000
	min := math.Inf(1)
	for _, d := range data {
		if d[len(d)-1].Value > max {
			max = d[len(d)-1].Value
		}
		if d[0].Value < min {
			min = d[0].Value
		}
	}
	return min, max
}

func LowerBoundIndex(arr []Pair, target float64) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid].Value < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func UpperBoundIndex(arr []Pair, target float64) int {
	low, high, mid := 0, len(arr)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if arr[mid].Value > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
