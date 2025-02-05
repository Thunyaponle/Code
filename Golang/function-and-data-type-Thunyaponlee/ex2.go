package main

// sumArray คำนวณผลรวมของตัวเลขในอาเรย์
func sumArray(arr [5]int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
