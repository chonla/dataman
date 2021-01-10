package text_test

import (
	"dataman/text"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartWithShouldReturnTrue(t *testing.T) {
	str := "hello world"
	testStr := "hello"

	result := text.StartWith(str, testStr)

	assert.True(t, result)
}

func TestStartWithEmptyStringShouldReturnTrue(t *testing.T) {
	str := "hello world"
	testStr := ""

	result := text.StartWith(str, testStr)

	assert.True(t, result)
}

func TestStartWithShouldReturnFalse(t *testing.T) {
	str := "hello world"
	testStr := "helly"

	result := text.StartWith(str, testStr)

	assert.False(t, result)
}

func TestStartWithShouldReturnFalseWhenTestStringIsLonger(t *testing.T) {
	str := "hello world"
	testStr := "hello worlds"

	result := text.StartWith(str, testStr)

	assert.False(t, result)
}
