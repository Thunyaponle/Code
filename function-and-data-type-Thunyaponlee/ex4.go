package main

func countOccurrences(slice []string) map[string]int {

	result := make(map[string]int)

	for _, word := range slice {
		result[word]++
	}
	return result
}
