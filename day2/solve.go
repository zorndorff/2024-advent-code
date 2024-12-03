package day2

import (
	"github.com/zorndorff/2024-advent-code/util"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func Solution(leftlist []int, rightlist []int) int {
	sortedLeft := util.CopyAndSort(leftlist)
	sortedRight := util.CopyAndSort(rightlist)
	slog.Debug("Sorted", "leftlist", sortedLeft)
	slog.Debug("Sorted", "rightlist", sortedRight)
	totalDistance := 0
	for i, val := range sortedLeft {
		highest := max(val, sortedRight[i])
		lowest := min(val, sortedRight[i])

		slog.Debug("Comparing", "lowest", lowest, "highest", highest)

		totalDistance += highest - lowest
		slog.Debug("New total", "totaldistance", totalDistance)
	}

	return totalDistance
}

func Solution2(leftlist []int, rightlist []int) int {

	rightTotals := util.CountedListMap(rightlist)

	addDistance := 0
	totalDistance := 0

	var valFromMap int
	var ok bool

	for _, checkValue := range leftlist {
		valFromMap, ok = rightTotals[checkValue]
		if ok {
			addDistance = checkValue * valFromMap
			totalDistance += addDistance
		}
	}

	return totalDistance
}

func Run() {
	slog.Info("Loading input file from ./inputs/day-1.txt")
	lines, err := util.ReadInputLines("./inputs/day-1.txt")
	if err != nil {
		slog.Error("readLines error", "error", err)
		os.Exit(1)
	}

	slog.Info("Got %s lines from file", "lines", len(lines))

	leftListInput := make([]int, len(lines))
	rightListInput := make([]int, len(lines))
	for i, val := range lines {
		valParts := strings.Split(val, "   ")
		intValue, err := strconv.Atoi(valParts[0])

		leftListInput[i] = intValue

		if err != nil {
			slog.Error("Error", "error", err)
			os.Exit(1)
		}

		intValue, err = strconv.Atoi(valParts[1])
		if err != nil {
			slog.Error("Error", "error", err)
			os.Exit(1)
		}
		rightListInput[i] = intValue
	}

	totalDist := Solution(leftListInput, rightListInput)

	slog.Info("Solve for Day 1 Part 1", "solution_1", totalDist)

	totalDist2 := Solution2(leftListInput, rightListInput)

	slog.Info("Solve for Day 1 Part 2", "solution_2", totalDist2)
}
