package main

import (
	"testing"
)

func TestDoubleValue(t *testing.T) {
	num := 5
	doubleValue(&num)
	if num != 10 {
		t.Errorf("Expected num to be 10, got %d", num)
	}

	num2 := 0
	doubleValue(&num2)
	if num2 != 0 {
		t.Errorf("Expected num2 to be 0, got %d", num2)
	}

	num3 := -3
	doubleValue(&num3)
	if num3 != -6 {
		t.Errorf("Expected num3 to be -6, got %d", num3)
	}
}
