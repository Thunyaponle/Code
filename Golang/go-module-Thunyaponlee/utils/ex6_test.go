// utils/ex6_test.go
package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// ทดสอบโครงสร้างไฟล์
func TestToolsFileStructure(t *testing.T) {
	// ตรวจสอบว่ามีไฟล์ tools.go
	toolsPath := filepath.Join(".", "tools.go")
	if _, err := os.Stat(toolsPath); os.IsNotExist(err) {
		t.Error("tools.go not found")
	}

	// อ่านเนื้อหาไฟล์เพื่อตรวจสอบการ import uuid
	content, err := os.ReadFile(toolsPath)
	if err != nil {
		t.Errorf("Cannot read tools.go: %v", err)
		return
	}

	// ตรวจสอบว่ามีการ import github.com/google/uuid
	if !regexp.MustCompile(`"github\.com/google/uuid"`).Match(content) {
		t.Error("github.com/google/uuid import not found in tools.go")
	}
}

// ทดสอบฟังก์ชัน GenerateUUID
func TestGenerateUUID(t *testing.T) {
	// เช็คว่าฟังก์ชันสร้าง UUID ถูกต้องตามรูปแบบ
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	// ทดสอบการสร้าง UUID หลายๆ ครั้ง
	for i := 0; i < 5; i++ {
		uuid := GenerateUUID()

		// ตรวจสอบว่า UUID ไม่เป็นค่าว่าง
		if uuid == "" {
			t.Error("GenerateUUID() returned empty string")
			continue
		}

		// ตรวจสอบว่า UUID มีรูปแบบถูกต้อง
		if !uuidRegex.MatchString(uuid) {
			t.Errorf("GenerateUUID() = %q; does not match UUID format", uuid)
			continue
		}

		// ตรวจสอบว่า UUID แต่ละครั้งไม่ซ้ำกัน
		secondUUID := GenerateUUID()
		if uuid == secondUUID {
			t.Error("Generated UUIDs are not unique")
		}
	}
}

// ทดสอบว่าฟังก์ชันเรียกใช้ uuid.NewString()
func TestUUIDImplementation(t *testing.T) {
	toolsPath := filepath.Join(".", "tools.go")
	content, err := os.ReadFile(toolsPath)
	if err != nil {
		t.Errorf("Cannot read tools.go: %v", err)
		return
	}

	// ตรวจสอบว่ามีการเรียกใช้ uuid.NewString()
	if !regexp.MustCompile(`uuid\.NewString\(\)`).Match(content) {
		t.Error("uuid.NewString() call not found in implementation")
	}
}
