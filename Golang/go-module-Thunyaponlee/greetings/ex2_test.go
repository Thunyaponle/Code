// greetings/ex2_test.go
package greetings

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"testing"
)

func TestGoodbye(t *testing.T) {
	// Define test cases
	tests := []struct {
		name  string // ชื่อ test case
		input string // input ที่จะทดสอบ
		want  string // ผลลัพธ์ที่ต้องการ
	}{
		{
			name:  "normal farewell",
			input: "Alice",
			want:  "Goodbye, Alice",
		},
		{
			name:  "with special characters",
			input: "Bob!@#",
			want:  "Goodbye, Bob!@#",
		},
		{
			name:  "with number",
			input: "User123",
			want:  "Goodbye, User123",
		},
		{
			name:  "empty name",
			input: "",
			want:  "Goodbye, ", // ทดสอบกรณี input เป็นค่าว่าง
		},
		{
			name:  "with space",
			input: "John Doe",
			want:  "Goodbye, John Doe",
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := Goodbye(tt.input)

			// Assert
			if result != tt.want {
				t.Errorf("Goodbye(%q) = %q; want %q", tt.input, result, tt.want)
			}
		})
	}
}

func TestGoodbyeLocation(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	if err != nil {
		t.Fatalf("Failed to parse package: %v", err)
	}

	found := false
	for _, pkg := range pkgs {
		for filename, file := range pkg.Files {
			if filepath.Base(filename) == "farewell.go" {
				ast.Inspect(file, func(n ast.Node) bool {
					if fn, ok := n.(*ast.FuncDecl); ok {
						if fn.Name.Name == "Goodbye" {
							found = true
							return false
						}
					}
					return true
				})
			}
		}
	}

	if !found {
		t.Error("Goodbye function not found in farewell.go")
	}
}
