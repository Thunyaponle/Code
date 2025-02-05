package main

import (
	"testing"
)

func TestIncrementAge(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	incrementAge(&p)
	if p.Age != 31 {
		t.Errorf("Expected Age to be 31, got %d", p.Age)
	}

	p2 := Person{Name: "Bob", Age: 0}
	incrementAge(&p2)
	if p2.Age != 1 {
		t.Errorf("Expected Age to be 1, got %d", p2.Age)
	}
}
