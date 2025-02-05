package main

import "testing"

func TestHorseMove(t *testing.T) {
	h := Horse{}
	expected := "Horse is galloping"
	if got := h.Move(); got != expected {
		t.Errorf("Horse.Move() = %q, want %q", got, expected)
	}
}

func TestHorseMakeSound(t *testing.T) {
	h := Horse{}
	expected := "Neigh!"
	if got := h.MakeSound(); got != expected {
		t.Errorf("Horse.MakeSound() = %q, want %q", got, expected)
	}
}

func TestDescribeCreature(t *testing.T) {
	h := Horse{}
	expected := "Horse is galloping and Neigh!"
	if got := describeCreature(h); got != expected {
		t.Errorf("describeCreature(Horse{}) = %q, want %q", got, expected)
	}
}

// TestHorseImplementsCreature verifies that Horse implements Creature interface
func TestHorseImplementsCreature(t *testing.T) {
	var _ Creature = Horse{} // This will fail to compile if Horse doesn't implement Creature
}
