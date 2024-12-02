package main

import (
	"fmt"
)



func main() {
	filePath := "input.txt" // Make sure this is the correct path to your input file
	records, err := ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	safeCount := 0
	for _, record := range records {
		if isSafeReport(record) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}
