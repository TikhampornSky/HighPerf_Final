package basic

func quickSort(arr []Pair, low int, high int) {
	if low < high {
		pi := partion(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partion(arr []Pair, low int, high int) int {
	pivot := arr[high].Value
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j].Value < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}