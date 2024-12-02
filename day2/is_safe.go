package main

// isSafeDifference checks if the difference between adjacent levels is within the specified range.
func isSafeDifference(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}

// isMonotonic checks if the levels are either strictly increasing or strictly decreasing.
func isMonotonic(levels []int) (bool, bool) {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(levels); i++ {
		if levels[i] > levels[i-1] {
			isDecreasing = false
		} else if levels[i] < levels[i-1] {
			isIncreasing = false
		}
	}
	return isIncreasing, isDecreasing
}

func oneWrong(levels []int) bool {
	// is just one level wrong? (if so it is safe)
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff < -3 || diff > 3 {
			return false
		}
	}
	return true
}

// isSafeReport checks if a report is safe based on the specified rules.
func isSafeReport(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	if oneWrong(levels) {
		return true
	}

	if !isSafeDifference(levels) {
		return false
	}

	isIncreasing, isDecreasing := isMonotonic(levels)
	return isIncreasing || isDecreasing
}
