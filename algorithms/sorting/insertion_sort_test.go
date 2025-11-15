package sorting

import (
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "already sorted array",
			arr:      []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted array",
			arr:      []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random array",
			arr:      []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			arr:      []int{42},
			expected: []int{42},
		},
		{
			name:     "duplicate elements",
			arr:      []int{3, 3, 3, 3},
			expected: []int{3, 3, 3, 3},
		},
		{
			name:     "negative numbers",
			arr:      []int{-3, -1, -4, -1, -5},
			expected: []int{-5, -4, -3, -1, -1},
		},
		{
			name:     "mixed positive and negative",
			arr:      []int{-2, 3, -1, 0, 1, -3},
			expected: []int{-3, -2, -1, 0, 1, 3},
		},
		{
			name:     "two elements - already sorted",
			arr:      []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "two elements - need sorting",
			arr:      []int{2, 1},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since InsertionSort modifies in place
			arr := make([]int, len(tt.arr))
			copy(arr, tt.arr)

			InsertionSort(arr)

			if !slices.Equal(arr, tt.expected) {
				t.Errorf("InsertionSort(%v) = %v, expected %v", tt.arr, arr, tt.expected)
			}
		})
	}
}

func TestInsertionSortString(t *testing.T) {
	tests := []struct {
		name     string
		arr      []string
		expected []string
	}{
		{
			name:     "string array",
			arr:      []string{"banana", "apple", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "already sorted strings",
			arr:      []string{"apple", "banana", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "reverse sorted strings",
			arr:      []string{"zebra", "dog", "cat"},
			expected: []string{"cat", "dog", "zebra"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]string, len(tt.arr))
			copy(arr, tt.arr)

			InsertionSort(arr)

			if !slices.Equal(arr, tt.expected) {
				t.Errorf("InsertionSort(%v) = %v, expected %v", tt.arr, arr, tt.expected)
			}
		})
	}
}

func TestInsertionSortFloat64(t *testing.T) {
	arr := []float64{3.14, 2.71, 1.41, 4.67}
	expected := []float64{1.41, 2.71, 3.14, 4.67}

	InsertionSort(arr)

	if !slices.Equal(arr, expected) {
		t.Errorf("InsertionSort(%v) = %v, expected %v", []float64{3.14, 2.71, 1.41, 4.67}, arr, expected)
	}
}
