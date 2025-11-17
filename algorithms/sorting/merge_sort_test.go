package sorting

import (
	"slices"
	"testing"
)

func TestMergeSortInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random order",
			input:    []int{3, 1, 4, 5, 2},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "with duplicates",
			input:    []int{3, 1, 3, 2, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			if len(arr) > 0 {
				MergeSort(arr, 0, len(arr)-1)
			} else {
				MergeSort(arr, 0, -1)
			}

			if !slices.Equal(arr, tt.expected) {
				t.Errorf("got %v, expected %v", arr, tt.expected)
			}
		})
	}
}

func TestMergeSortString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "letters",
			input:    []string{"d", "a", "b", "c"},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "already sorted",
			input:    []string{"cat", "dog", "fish"},
			expected: []string{"cat", "dog", "fish"},
		},
		{
			name:     "empty array",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "duplicates",
			input:    []string{"bob", "alice", "bob"},
			expected: []string{"alice", "bob", "bob"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]string, len(tt.input))
			copy(arr, tt.input)

			if len(arr) > 0 {
				MergeSort(arr, 0, len(arr)-1)
			} else {
				MergeSort(arr, 0, -1)
			}

			if !slices.Equal(arr, tt.expected) {
				t.Errorf("got %v, expected %v", arr, tt.expected)
			}
		})
	}
}
