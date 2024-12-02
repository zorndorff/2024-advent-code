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
