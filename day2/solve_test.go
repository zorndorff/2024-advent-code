package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay2Solve(t *testing.T) {
	inputs := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := Solution(inputs[0], inputs[1])

	assert.Equal(t, 11, result)

}

func TestDay2Part2Solve(t *testing.T) {
	inputs := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	result := Solution2(inputs[0], inputs[1])

	assert.Equal(t, 31, result)

}
