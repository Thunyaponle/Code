// ex3_test.go
package main

import (
	"reflect"
	"testing"
)

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"EmptySlice", []int{}, []int{}},
		{"AllEven", []int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{"AllOdd", []int{1, 3, 5, 7}, []int{}},
		{"Mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"IncludingNegative", []int{-3, -2, -1, 0, 1, 2}, []int{-2, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterEven(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("filterEven(%v) got %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
