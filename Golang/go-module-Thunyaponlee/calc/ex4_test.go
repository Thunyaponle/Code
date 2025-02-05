// calc/ex4_test.go
package calc

import (
	"os"
	"path/filepath"
	"testing"
)

// ทดสอบการมีอยู่ของไฟล์และโครงสร้าง
func TestFileStructure(t *testing.T) {
	// ตรวจสอบว่ามีไฟล์ calculator.go
	calculatorPath := filepath.Join(".", "calculator.go")
	if _, err := os.Stat(calculatorPath); os.IsNotExist(err) {
		t.Error("calculator.go not found")
	}
}

// ทดสอบฟังก์ชัน Add
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "positive numbers",
			a:    2,
			b:    3,
			want: 5,
		},
		{
			name: "negative numbers",
			a:    -2,
			b:    -3,
			want: -5,
		},
		{
			name: "zero and positive",
			a:    0,
			b:    5,
			want: 5,
		},
		{
			name: "positive and negative",
			a:    5,
			b:    -3,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.want)
			}
		})
	}
}

// ทดสอบฟังก์ชัน Multiply
func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "positive numbers",
			a:    2,
			b:    3,
			want: 6,
		},
		{
			name: "negative numbers",
			a:    -2,
			b:    -3,
			want: 6,
		},
		{
			name: "zero and number",
			a:    0,
			b:    5,
			want: 0,
		},
		{
			name: "positive and negative",
			a:    5,
			b:    -3,
			want: -15,
		},
		{
			name: "multiply by one",
			a:    7,
			b:    1,
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.want)
			}
		})
	}
}
