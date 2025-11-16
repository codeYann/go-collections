package search

// IndexOf performs a linear search to find the index of target in arr.
// Returns -1 if target is not found.
func IndexOf[T comparable](arr []T, target T) int {
	n, i := len(arr), 0

	for i < n {
		if arr[i] == target {
			return i
		}
		i = i + 1
	}

	return -1
}
