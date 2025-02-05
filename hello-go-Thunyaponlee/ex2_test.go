package main

import "testing"

// Unit test for the sumArray function
func TestSumArray(t *testing.T) {
	// Test cases
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{-1, -2, -3, -4, -5}, -15},
		{[]int{10, 20, 30}, 60},
		{[]int{}, 0}, // Empty array case
	}

	for _, tt := range tests {
		result := sumArray(tt.input)
		if result != tt.expected {
			t.Errorf("sumArray(%v) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}
