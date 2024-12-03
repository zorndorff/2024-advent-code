package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyAndSort(t *testing.T) {
	inputs := []int{1, 3, 2, 3, 20, 4, 5}

	result := CopyAndSort(inputs)
	expected := []int{1, 2, 3, 3, 4, 5, 20}

	assert.Equal(t, expected, result)

}

func TestCountedListMap(t *testing.T) {
	inputs := []int{1, 3, 2, 3, 20, 4, 5}

	result := CountedListMap(inputs)
	expected := map[int]int{1: 1, 2: 1, 3: 2, 4: 1, 5: 1, 20: 1}

	assert.Equal(t, expected, result)

}
