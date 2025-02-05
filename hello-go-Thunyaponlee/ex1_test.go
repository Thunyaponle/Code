package main

import "testing"

// Unit test for calculateArea function
func TestCalculateArea(t *testing.T) {
	tests := []struct {
		width, height, expected float64
	}{
		{5, 10, 50},
		{7, 3, 21},
		{0, 10, 0},
		{10, 10, 100},
	}

	for _, tt := range tests {
		result := calculateArea(tt.width, tt.height)
		if result != tt.expected {
			t.Errorf("calculateArea(%f, %f) = %f; want %f", tt.width, tt.height, result, tt.expected)
		}
	}
}
