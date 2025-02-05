// ex7_test.go
package main

import (
	"os"
	"regexp"
	"testing"
)

func TestGoModTidy(t *testing.T) {
	// ตรวจสอบการมีอยู่ของไฟล์ go.mod
	goModPath := "go.mod"
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		t.Error("go.mod file not found")
		return
	}

	// อ่านเนื้อหาของไฟล์ go.mod
	content, err := os.ReadFile(goModPath)
	if err != nil {
		t.Errorf("Cannot read go.mod: %v", err)
		return
	}

	contentStr := string(content)

	// ตรวจสอบว่ามีการระบุ module name
	if !regexp.MustCompile(`module\s+[a-zA-Z0-9/_-]+`).Match(content) {
		t.Error("Module name not found in go.mod")
	}

	// ตรวจสอบว่ามี github.com/google/uuid เป็น dependency
	if !regexp.MustCompile(`github\.com/google/uuid\s+v[0-9.]+`).Match(content) {
		t.Error("github.com/google/uuid dependency not found in go.mod")
	}

	// ตรวจสอบว่ามีไฟล์ go.sum
	goSumPath := "go.sum"
	if _, err := os.Stat(goSumPath); os.IsNotExist(err) {
		t.Error("go.sum file not found")
		return
	}

	// ตรวจสอบการไม่มี dependency ที่ไม่ได้ใช้ (indirect)
	indirectDeps := regexp.MustCompile(`//\s+indirect`).FindAllString(contentStr, -1)
	if len(indirectDeps) > 0 {
		t.Error("Found unused (indirect) dependencies. Run 'go mod tidy' to clean up")
	}

	// ตรวจสอบ go version
	if !regexp.MustCompile(`go\s+1\.[0-9]+`).Match(content) {
		t.Error("Go version not specified in go.mod")
	}
}

func TestProjectStructure(t *testing.T) {
	// รายการไฟล์และโฟลเดอร์ที่ควรมี
	requiredPaths := []string{
		"greetings/greetings.go",
		"greetings/farewell.go",
		"calc/calculator.go",
		"utils/converter.go",
		"utils/tools.go",
		"go.mod",
		"go.sum",
	}

	for _, path := range requiredPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Required file/directory not found: %s", path)
		}
	}
}

func TestGoModRequirements(t *testing.T) {
	// อ่านเนื้อหาของ go.mod
	content, err := os.ReadFile("go.mod")
	if err != nil {
		t.Errorf("Cannot read go.mod: %v", err)
		return
	}

	contentStr := string(content)

	// ตรวจสอบ requirements ที่สำคัญ
	requirements := []struct {
		name        string
		pattern     string
		shouldExist bool
	}{
		{
			name:        "uuid package",
			pattern:     `github\.com/google/uuid`,
			shouldExist: true,
		},
	}

	for _, req := range requirements {
		exists := regexp.MustCompile(req.pattern).MatchString(contentStr)
		if exists != req.shouldExist {
			if req.shouldExist {
				t.Errorf("Required dependency %s not found in go.mod", req.name)
			} else {
				t.Errorf("Unexpected dependency %s found in go.mod", req.name)
			}
		}
	}
}
