package main

import (
	"testing"
)

func TestSafeDereference(t *testing.T) {
	var p *int
	val, err := safeDereference(p)
	if val != 0 {
		t.Errorf("Expected val = 0 for nil pointer, got %d", val)
	}
	if err == nil {
		t.Errorf("Expected non-nil error for nil pointer")
	}

	x := 42
	val2, err2 := safeDereference(&x)
	if val2 != 42 {
		t.Errorf("Expected val = 42, got %d", val2)
	}
	if err2 != nil {
		t.Errorf("Expected err = nil, got error: %v", err2)
	}
}
