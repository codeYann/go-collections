package sorting

import "cmp"

// InsertionSort sorts a slice in place using the insertion sort algorithm.
func InsertionSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for j := 1; j < n; j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}
}
