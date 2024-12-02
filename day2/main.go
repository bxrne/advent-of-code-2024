package main

import (
	"fmt"
)

// isSafeReport checks if a report is safe based on the specified rules.
func isSafeReport(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		if diff > 0 {
			isDecreasing = false
		} else {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

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
