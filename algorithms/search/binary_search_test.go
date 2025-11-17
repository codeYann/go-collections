package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "target found at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "target found in middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "target found at end",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "target not found - too large",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "target not found - too small",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: -1,
		},
		{
			name:     "target not found - in between",
			arr:      []int{1, 3, 5, 7, 9},
			target:   4,
			expected: -1,
		},
		{
			name:     "empty array",
			arr:      []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "single element found",
			arr:      []int{42},
			target:   42,
			expected: 0,
		},
		{
			name:     "single element not found",
			arr:      []int{42},
			target:   1,
			expected: -1,
		},
		{
			name:     "duplicate elements - finds one occurrence",
			arr:      []int{1, 2, 2, 2, 3},
			target:   2,
			expected: 2,
		},
		{
			name:     "large array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   7,
			expected: 6,
		},
		{
			name:     "negative numbers",
			arr:      []int{-5, -4, -3, -2, -1},
			target:   -3,
			expected: 2,
		},
		{
			name:     "mixed positive and negative",
			arr:      []int{-3, -2, -1, 0, 1, 2, 3},
			target:   0,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d, expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestBinarySearchString(t *testing.T) {
	tests := []struct {
		name     string
		arr      []string
		target   string
		expected int
	}{
		{
			name:     "string array - found",
			arr:      []string{"apple", "banana", "cherry"},
			target:   "banana",
			expected: 1,
		},
		{
			name:     "string array - not found",
			arr:      []string{"apple", "banana", "cherry"},
			target:   "grape",
			expected: -1,
		},
		{
			name:     "string array - found at beginning",
			arr:      []string{"apple", "banana", "cherry"},
			target:   "apple",
			expected: 0,
		},
		{
			name:     "string array - found at end",
			arr:      []string{"apple", "banana", "cherry"},
			target:   "cherry",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %q) = %d, expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestBinarySearchFloat64(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	if BinarySearch(arr, 3.3) != 2 {
		t.Error("Float64 search failed")
	}
	if BinarySearch(arr, 6.6) != -1 {
		t.Error("Float64 not found should return -1")
	}
}
