package mergesort

func MergeInBlock(min, max float64, data [][]Pair) []Pair {
	var result []Pair
	startIndex := make([]int, len(data))
	stopIndex := make([]int, len(data))

	for i, d := range data {
		startIndex[i], stopIndex[i] = BinarySearch(d, min, max)
	}

	for {
		var min Pair
		var minIndex = -1
		for i, d := range data {
			if startIndex[i] < stopIndex[i] {
				if minIndex == -1 || d[startIndex[i]].Value < min.Value {
					min = d[startIndex[i]]
					minIndex = i
				}
			}
		}
		if minIndex == -1 {
			break
		}
		startIndex[minIndex] += 1
		result = append(result, min)
	}
	return result
}
