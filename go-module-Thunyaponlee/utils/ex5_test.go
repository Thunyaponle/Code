// utils/ex5_test.go
package utils

import (
	"os"
	"path/filepath"
	"testing"
)

// ทดสอบโครงสร้างไฟล์
func TestConverterFileStructure(t *testing.T) {
	// ตรวจสอบว่ามีไฟล์ converter.go
	converterPath := filepath.Join(".", "converter.go")
	if _, err := os.Stat(converterPath); os.IsNotExist(err) {
		t.Error("converter.go not found")
	}
}

// ทดสอบฟังก์ชัน NumberToThai
func TestNumberToThai(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    string
		wantErr bool
	}{
		{
			name:    "number 0",
			input:   0,
			want:    "ศูนย์",
			wantErr: false,
		},
		{
			name:    "number 1",
			input:   1,
			want:    "หนึ่ง",
			wantErr: false,
		},
		{
			name:    "number 2",
			input:   2,
			want:    "สอง",
			wantErr: false,
		},
		{
			name:    "number 3",
			input:   3,
			want:    "สาม",
			wantErr: false,
		},
		{
			name:    "number 4",
			input:   4,
			want:    "สี่",
			wantErr: false,
		},
		{
			name:    "number 5",
			input:   5,
			want:    "ห้า",
			wantErr: false,
		},
		{
			name:    "number 6",
			input:   6,
			want:    "หก",
			wantErr: false,
		},
		{
			name:    "number 7",
			input:   7,
			want:    "เจ็ด",
			wantErr: false,
		},
		{
			name:    "number 8",
			input:   8,
			want:    "แปด",
			wantErr: false,
		},
		{
			name:    "number 9",
			input:   9,
			want:    "เก้า",
			wantErr: false,
		},
		{
			name:    "invalid number - negative",
			input:   -1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid number - too large",
			input:   10,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumberToThai(tt.input)
			if tt.wantErr {
				if result != "" {
					t.Errorf("NumberToThai(%d) = %q; want empty string for invalid input", tt.input, result)
				}
			} else {
				if result != tt.want {
					t.Errorf("NumberToThai(%d) = %q; want %q", tt.input, result, tt.want)
				}
			}
		})
	}
}
