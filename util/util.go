package util

import (
	"bufio"
	"os"
	"slices"
)

func CopyAndSort(list []int) []int {
	sorted := make([]int, len(list))
	for i, val := range list {
		sorted[i] = val
	}
	slices.Sort(sorted)
	return sorted
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadInputLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func CountedListMap(list []int) map[int]int {
	intMap := make(map[int]int)
	for _, val := range list {
		valTotal := intMap[val]
		valTotal += 1
		intMap[val] = valTotal
	}
	return intMap
}
