// ex2_test.go
package main

import "testing"

func TestSumArray(t *testing.T) {
	tests := []struct {
		name     string
		input    [5]int
		expected int
	}{
		{"AllPositive", [5]int{1, 2, 3, 4, 5}, 15},
		{"WithZero", [5]int{0, 2, 4, 6, 8}, 20},
		{"AllNegative", [5]int{-1, -2, -3, -4, -5}, -15},
		{"Mixed", [5]int{-10, 10, 0, 5, -5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumArray(tt.input)
			if result != tt.expected {
				t.Errorf("sumArray(%v) got %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
