package main

func filterEven(slice []int) []int {
	newslice := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			newslice = append(newslice, num)
		}
	}
	return newslice
}
