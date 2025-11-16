package sorting

import (
	"cmp"
)

// merge merges two sorted subarrays arr[start..mid] and arr[mid+1..end] into a single sorted subarray arr[start..end]
func merge[T cmp.Ordered](arr []T, start, mid, end int) {
	leftLen := mid - start + 1
	rightLen := end - mid

	left := make([]T, leftLen)
	right := make([]T, rightLen)

	copy(left, arr[start:mid+1])
	copy(right, arr[mid+1:end+1])

	leftIdx, rightIdx, mergeIdx := 0, 0, start

	for leftIdx < leftLen && rightIdx < rightLen {
		if left[leftIdx] <= right[rightIdx] {
			arr[mergeIdx] = left[leftIdx]
			leftIdx++
		} else {
			arr[mergeIdx] = right[rightIdx]
			rightIdx++
		}
		mergeIdx++
	}

	for leftIdx < leftLen {
		arr[mergeIdx] = left[leftIdx]
		leftIdx++
		mergeIdx++
	}
	for rightIdx < rightLen {
		arr[mergeIdx] = right[rightIdx]
		rightIdx++
		mergeIdx++
	}
}

// MergeSort sorts a slice using the merge sort algorithm.
// start is the starting index and end is the ending index (inclusive).
func MergeSort[T cmp.Ordered](arr []T, start, end int) {
	if start < end {
		mid := (start + end) / 2

		MergeSort(arr, start, mid)
		MergeSort(arr, mid+1, end)
		merge(arr, start, mid, end)
	}
}
