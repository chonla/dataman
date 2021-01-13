package array_test

import (
	"dataman/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetElementInArrayByIndex(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	expected := "d"

	result := array.Get(arr, 3, "z")

	assert.Equal(t, expected, result)
}

func TestGetElementInArrayByIndexLowerThanLowerBound(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	expected := "z"

	result := array.Get(arr, -1, "z")

	assert.Equal(t, expected, result)
}

func TestGetElementInArrayByIndexOverThanUpperBound(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	expected := "z"

	result := array.Get(arr, 5, "z")

	assert.Equal(t, expected, result)
}

func TestGetElementIndex(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	expected := 3

	result := array.IndexOf(arr, "d")

	assert.Equal(t, expected, result)
}

func TestGetElementIndexReturnErrorNotFoundIfNotFound(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	expected := array.ErrorNotFound

	result := array.IndexOf(arr, "z")

	assert.Equal(t, expected, result)
}
