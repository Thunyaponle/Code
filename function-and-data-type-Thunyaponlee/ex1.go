package main

import (
	"fmt"
)

// applyOperation รับฟังก์ชัน op และเลข a, b เพื่อคำนวณผลลัพธ์ของ op(a, b)
func applyOperation(op func(float64, float64) float64, a, b float64) float64 {
	return op(a, b)
}

func main() {
	// การดำเนินการทางคณิตศาสตร์
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
		if y != 0 {
			return x / y
		}
		return 0
	}
	result := applyOperation(add, 5, 3)
	fmt.Println(result)
	result2 := applyOperation(subtract, 5, 3)
	fmt.Println(result2)
	result3 := applyOperation(multiply, 5, 3)
	fmt.Println(result3)
	result4 := applyOperation(divide, 5, 3)
	fmt.Println(result4)
}
