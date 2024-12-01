package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Pair struct to hold a row of two numbers
type Pair struct {
	first  int
	second int
}

// FileReader struct to handle reading and parsing the file
type FileReader struct {
	filePath string
	data     []Pair
}

// NewFileReader initializes the FileReader with the provided file path
func NewFileReader(filePath string) *FileReader {
	return &FileReader{
		filePath: filePath,
		data:     make([]Pair, 0),
	}
}

// ReadFile reads and parses the file line by line
func (reader *FileReader) ReadFile() error {
	file, err := os.Open(reader.filePath)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums, err := parseLine(line)
		if err != nil {
			return fmt.Errorf("error parsing line: %w", err)
		}
		reader.data = append(reader.data, nums)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// parseLine parses a line to extract two numbers
func parseLine(line string) (Pair, error) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		return Pair{}, fmt.Errorf("invalid line format")
	}

	first, err := strconv.Atoi(fields[0])
	if err != nil {
		return Pair{}, fmt.Errorf("error parsing first number: %w", err)
	}
	second, err := strconv.Atoi(fields[1])
	if err != nil {
		return Pair{}, fmt.Errorf("error parsing second number: %w", err)
	}

	return Pair{first, second}, nil
}

// GetData returns the parsed data
func (reader *FileReader) GetData() []Pair {
	return reader.data
}
