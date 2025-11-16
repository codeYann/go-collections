package search

import "cmp"

// BinarySearch performs a binary search to find the index of target in a sorted array.
// Returns -1 if target is not found. The array must be sorted in ascending order.
func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	a, b := 0, len(arr)-1

	for a <= b {
		middle := (a + b) / 2
		if arr[middle] == target {
			return middle
		}

		if arr[middle] > target {
			b = middle - 1
		} else {
			a = middle + 1
		}
	}

	return -1
}
