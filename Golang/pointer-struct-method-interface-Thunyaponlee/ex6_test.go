package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCat(t *testing.T) {
	myCat := Cat{Animal: Animal{Name: "ตุนตัง"}, Breed: "มันช์กิ้น"}

	// เตรียมตัวแปรไว้เก็บ stdout เดิม
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// เรียกใช้ method ที่จะ print
	myCat.Breathe()
	myCat.Meow()

	// ปิด writer เพื่อให้การอ่านเสร็จสิ้น
	w.Close()

	// อ่านค่าจาก pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// คืนค่า stdout กลับ
	os.Stdout = originalStdout

	// ได้ output ทั้งหมดอยู่ใน buf แล้ว
	output := buf.String()

	// ตรวจสอบข้อความที่คาดหวัง
	expected := "ตุนตัง กำลังหายใจ\nตุนตัง กำลังร้อง: เมี้ยว! เมี้ยว!\n"
	if output != expected {
		t.Errorf("Expected output:\n%q\nBut got:\n%q", expected, output)
	}
}
