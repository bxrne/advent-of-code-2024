package main

import (
	"fmt"
	"math"
	"sort"
)

func SortAndMerge(pairs []Pair) []Pair {
	firsts := make([]int, len(pairs))
	seconds := make([]int, len(pairs))

	// Seperate out first and second
	for i, pair := range pairs {
		firsts[i] = pair.first
		seconds[i] = pair.second
	}

	// Sort first and second independently
	sort.Ints(firsts)
	sort.Ints(seconds)

	// Merge back sorted first and second into pairs
	for i := range pairs {
		pairs[i].first = firsts[i]
		pairs[i].second = seconds[i]
	}

	return pairs
}

func main() {
	fileReader := NewFileReader("input.txt") // INFO: Available here: https://adventofcode.com/2024/day/1/input

	if err := fileReader.ReadFile(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// To match them up, we can vertically sort them beside each other
	data := fileReader.GetData()
	sortedData := SortAndMerge(data)

	fmt.Printf("Min (%d, %d)\n", sortedData[0].first, sortedData[0].second)
	fmt.Printf("Max (%d, %d)\n", sortedData[len(sortedData)-1].first, sortedData[len(sortedData)-1].second)

	// within each pair calculate the difference between the two numbers
	differences := make([]int, len(sortedData))
	for i, pair := range sortedData {
		differences[i] = int(math.Abs(float64(pair.second) - float64(pair.first)))
	}

	totalDifference := 0
	for _, diff := range differences {
		totalDifference += diff
	}

	fmt.Println("Total difference:", totalDifference)

}
