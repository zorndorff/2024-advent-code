package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1Solve(t *testing.T) {
	inputs := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	result := Solution(inputs[0], inputs[1])

	assert.Equal(t, 11, result)

}
