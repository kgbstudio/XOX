package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"apple", "banana", "cherry", "date", "elderberry"}
	searchTerm := "er"

	results := search(data, searchTerm)

	fmt.Println("Search results:")
	for _, result := range results {
		fmt.Println(result)
	}
}

func search(data []string, searchTerm string) []string {
	var results []string

	for _, item := range data {
		if strings.Contains(item, searchTerm) {
			results = append(results, item)
		}
	}

	return results
}