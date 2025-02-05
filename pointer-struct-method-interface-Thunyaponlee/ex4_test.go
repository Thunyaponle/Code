package main

import (
	"math"
	"testing"
)

func TestShapeArea(t *testing.T) {
	c := Circle{Radius: 5}
	sq := Square{Side: 4}

	cArea := c.Area()
	expectedCArea := math.Pi * 5 * 5
	if cArea != expectedCArea {
		t.Errorf("Expected circle area to be %f, got %f", expectedCArea, cArea)
	}

	sArea := sq.Area()
	expectedSArea := float64(16)
	if sArea != expectedSArea {
		t.Errorf("Expected square area to be %f, got %f", expectedSArea, sArea)
	}
}

func TestPrintArea(t *testing.T) {
	c := Circle{Radius: 5}
	sq := Square{Side: 4}
	printArea(c)
	printArea(sq)
	// Manual check: run `go test -v` and see printed output
}
