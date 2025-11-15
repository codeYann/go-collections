package search

import "testing"

func TestIndexOf(t *testing.T) {
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
			name:     "target not found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
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
			name:     "duplicate elements - returns first occurrence",
			arr:      []int{1, 2, 2, 3, 2},
			target:   2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IndexOf(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("IndexOf(%v, %d) = %d, expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestIndexOfString(t *testing.T) {
	arr := []string{"apple", "banana", "cherry"}
	if IndexOf(arr, "banana") != 1 {
		t.Error("String search failed")
	}
	if IndexOf(arr, "grape") != -1 {
		t.Error("String not found should return -1")
	}
}
