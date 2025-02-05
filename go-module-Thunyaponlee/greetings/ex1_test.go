package greetings

import "testing"

func TestHello(t *testing.T) {
	// Define test cases
	tests := []struct {
		name  string // ชื่อ test case
		input string // input ที่จะทดสอบ
		want  string // ผลลัพธ์ที่ต้องการ
	}{
		{
			name:  "normal greeting",
			input: "Alice",
			want:  "Hello, Alice",
		},
		{
			name:  "empty name",
			input: "",
			want:  "Hello, World",
		},
		{
			name:  "with special characters",
			input: "Bob!@#",
			want:  "Hello, Bob!@#",
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := Hello(tt.input)

			// Assert
			if result != tt.want {
				t.Errorf("Hello(%q) = %q; want %q", tt.input, result, tt.want)
			}
		})
	}
}
