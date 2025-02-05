// ex4_test.go
package main

import (
	"reflect"
	"testing"
)

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name:     "EmptySlice",
			input:    []string{},
			expected: map[string]int{},
		},
		{
			name:     "SingleItem",
			input:    []string{"apple"},
			expected: map[string]int{"apple": 1},
		},
		{
			name:     "MultipleItems",
			input:    []string{"apple", "banana", "apple", "orange", "banana", "apple"},
			expected: map[string]int{"apple": 3, "banana": 2, "orange": 1},
		},
		{
			name:     "AllSameItems",
			input:    []string{"cat", "cat", "cat"},
			expected: map[string]int{"cat": 3},
		},
		{
			name:     "MixedCase",
			input:    []string{"Apple", "apple", "APPLE"},
			expected: map[string]int{"Apple": 1, "apple": 1, "APPLE": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countOccurrences(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("countOccurrences(%v) got %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
