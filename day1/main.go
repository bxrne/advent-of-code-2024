package main

import (
	"fmt"
	"math"
	"sort"
)

// Builds a frequency map of integers from the right list
func buildFrequencyMap(nums []int) map[int]int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}
	return frequencyMap
}

// Calculates the similarity score based on frequencies of numbers from the right list
func calculateSimilarityScore(leftList, rightList []int) int {
	frequencyMap := buildFrequencyMap(rightList)
	similarityScore := 0

	for _, num := range leftList {
		if count, found := frequencyMap[num]; found {
			similarityScore += num * count
		}
	}

	return similarityScore
}

func sortAndMerge(pairs []Pair) []Pair {
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
	sortedData := sortAndMerge(data)

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

	leftList := make([]int, len(sortedData))
	rightList := make([]int, len(sortedData))

	for i, pair := range sortedData {
		leftList[i] = pair.first
		rightList[i] = pair.second
	}

	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Printf("Similarity Score: %d\n", similarityScore)

}
