package main

import (
	"testing"
)

func TestRectangleMethods(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	area := r.Area()
	if area != 50 {
		t.Errorf("Expected area to be 50, got %f", area)
	}

	r.Resize(2, 2)
	newArea := r.Area()
	if newArea != 4 {
		t.Errorf("Expected area after resize to be 4, got %f", newArea)
	}

	r.Resize(1.5, 3.5)
	resizedArea := r.Area()
	expected := 1.5 * 3.5
	if resizedArea != expected {
		t.Errorf("Expected area after resize to be %f, got %f", expected, resizedArea)
	}
}
