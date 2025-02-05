// greetings/ex3_test.go
package greetings

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// ทดสอบ FormatHello (Exported function)
func TestFormatHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "normal greeting",
			input: "Alice",
			want:  "Hello, *** Alice ***",
		},
		{
			name:  "empty name",
			input: "",
			want:  "Hello, *** World ***", // สมมติว่าใช้ "World" เมื่อชื่อว่าง
		},
		{
			name:  "with special characters",
			input: "Bob!@#",
			want:  "Hello, *** Bob!@# ***",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatHello(tt.input)
			if result != tt.want {
				t.Errorf("FormatHello(%q) = %q; want %q", tt.input, result, tt.want)
			}
		})
	}
}

// ทดสอบการมีอยู่และการเป็น unexported ของ formatMessage
func TestFormatMessageExists(t *testing.T) {
	// ใช้ reflection เพื่อตรวจสอบฟังก์ชัน formatMessage
	greetingsType := reflect.TypeOf(formatMessage)
	if greetingsType == nil {
		t.Error("formatMessage function does not exist")
		return
	}

	// ตรวจสอบว่าฟังก์ชันเป็น unexported (ขึ้นต้นด้วยตัวพิมพ์เล็ก)
	functionName := runtime.FuncForPC(reflect.ValueOf(formatMessage).Pointer()).Name()
	if !strings.Contains(functionName, "formatMessage") {
		t.Error("formatMessage should be unexported (start with lowercase)")
	}

	// ทดสอบการทำงานของ formatMessage
	result := formatMessage("Test")
	expected := "*** Test ***"
	if result != expected {
		t.Errorf("formatMessage(%q) = %q; want %q", "Test", result, expected)
	}
}

// ทดสอบว่า FormatHello ใช้ formatMessage จริง
func TestFormatHelloUsesFormatMessage(t *testing.T) {
	// ทดสอบโดยเช็คผลลัพธ์ว่ามีรูปแบบที่ต้องผ่าน formatMessage
	input := "TestName"
	result := FormatHello(input)
	formattedPart := strings.TrimPrefix(result, "Hello, ")

	// ตรวจสอบว่าส่วนที่เหลือตรงกับรูปแบบของ formatMessage
	expectedFormat := formatMessage(input)
	if formattedPart != expectedFormat {
		t.Errorf("FormatHello does not seem to use formatMessage properly.\nGot: %q\nExpected format: %q", formattedPart, expectedFormat)
	}
}
