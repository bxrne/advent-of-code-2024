package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFile is responsible for reading the input file and processing each row as a record of integers.
func ReadFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	var records [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, parseLine(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return records, nil
}

// parseLine converts a line of text into a slice of integers.
func parseLine(line string) []int {
	stringFields := strings.Fields(line)
	intFields := make([]int, len(stringFields))

	for i, str := range stringFields {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("error converting string to int: %v\n", err)
			continue // Skip invalid integers but process the rest
		}
		intFields[i] = num
	}

	return intFields
}
