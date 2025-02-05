package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	c := Counter{Count: 10}
	c.Increment()
	c.Increment()
	if c.GetCount() != 12 {
		t.Errorf("Expected count to be 12, got %d", c.GetCount())
	}

	c2 := Counter{Count: 0}
	for i := 0; i < 5; i++ {
		c2.Increment()
	}
	if c2.GetCount() != 5 {
		t.Errorf("Expected count to be 5, got %d", c2.GetCount())
	}
}
