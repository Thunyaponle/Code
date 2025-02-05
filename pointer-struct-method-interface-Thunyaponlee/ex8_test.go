package main

import (
	"testing"
)

// ประกาศ Struct Unknown ที่ Implement Soundmaker สำหรับการทดสอบ
type Unknown struct{}

// Implement Method MakeSound() สำหรับ Unknown
func (u Unknown) MakeSound() string {
	return ""
}

// ทดสอบฟังก์ชัน identifyAnimal กับ Dog
func TestIdentifyAnimalDog(t *testing.T) {
	d := Dog{}
	res := identifyAnimal(d)
	expected := "This is a Dog!"
	if res != expected {
		t.Errorf("Expected '%s', got '%s'", expected, res)
	}
}

// ทดสอบฟังก์ชัน identifyAnimal กับ Cow
func TestIdentifyAnimalCow(t *testing.T) {
	c := Cow{}
	res := identifyAnimal(c)
	expected := "This is a Cow!"
	if res != expected {
		t.Errorf("Expected '%s', got '%s'", expected, res)
	}
}

// ทดสอบฟังก์ชัน identifyAnimal กับ Unknown
func TestIdentifyAnimalUnknown(t *testing.T) {
	u := Unknown{}
	res := identifyAnimal(u)
	expected := "Unknown Animal!"
	if res != expected {
		t.Errorf("Expected '%s', got '%s'", expected, res)
	}
}

// ทดสอบฟังก์ชัน identifyAnimal กับประเภทอื่นๆ ที่ Implement Soundmaker
// ประกาศ Struct Cat ที่ Implement Soundmaker
type Ant struct{}

// Implement Method MakeSound() สำหรับ Cat
func (a Ant) MakeSound() string {
	return "Meow!"
}

func TestIdentifyAnimalOther(t *testing.T) {
	ant := Ant{}
	res := identifyAnimal(ant)
	expected := "Unknown Animal!"
	if res != expected {
		t.Errorf("Expected '%s', got '%s'", expected, res)
	}
}
