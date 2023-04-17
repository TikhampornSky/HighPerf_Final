package basic

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

func BinarySearch(arr []Pair, targetMin, targetMax float64) (int, int) {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid].Value < targetMin {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	lowUp, highUp := 0, len(arr)-1
	for lowUp <= highUp {
		mid := (lowUp + highUp) / 2
		if arr[mid].Value > targetMax {
			highUp = mid - 1
		} else {
			lowUp = mid + 1
		}
	}

	return low, lowUp
}
