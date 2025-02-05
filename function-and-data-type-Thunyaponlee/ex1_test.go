// ex1_test.go
package main

import (
	"testing"
)

func TestApplyOperation(t *testing.T) {
	add := func(x, y float64) float64 {
		return x + y
	}
	subtract := func(x, y float64) float64 {
		return x - y
	}
	multiply := func(x, y float64) float64 {
		return x * y
	}
	divide := func(x, y float64) float64 {
		if y == 0 {
			// สมมติให้ถ้าหารด้วยศูนย์แล้วได้ 0
			return 0
		}
		return x / y
	}

	tests := []struct {
		name     string
		op       func(float64, float64) float64
		a, b     float64
		expected float64
	}{
		// ทดสอบการบวก
		{"AddPositive", add, 5, 3, 8},
		{"AddNegative", add, -5, -3, -8},
		{"AddMixed", add, -5, 3, -2},
		{"AddFloat", add, 1.5, 2.3, 3.8},

		// ทดสอบการลบ
		{"SubtractPositive", subtract, 10, 4, 6},
		{"SubtractNegative", subtract, -10, -4, -6},
		{"SubtractMixed", subtract, 5, -3, 8},
		{"SubtractFloat", subtract, 5.5, 2.2, 3.3},

		// ทดสอบการคูณ
		{"MultiplyPositive", multiply, 2, 7, 14},
		{"MultiplyNegative", multiply, -2, 7, -14},
		{"MultiplyMixed", multiply, -3, -3, 9},
		{"MultiplyFloat", multiply, 1.2, 2.5, 3.0},

		// ทดสอบการหาร
		{"DividePositive", divide, 20, 5, 4},
		{"DivideNegative", divide, -20, 5, -4},
		{"DivideMixed", divide, -20, -5, 4},
		{"DivideFloat", divide, 10.0, 2.5, 4.0},

		// ทดสอบเคสพิเศษ การหารด้วยศูนย์
		{"DivideByZero", divide, 10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applyOperation(tt.op, tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("operation %s(%v, %v) got %v, want %v", tt.name, tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
