package cast_test

import (
	"dataman/cast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCastToInt(t *testing.T) {
	val := "10"
	expected := int64(10)

	result := cast.ToInt(val, int64(20))

	assert.Equal(t, expected, result)
}

func TestCastToIntFailover(t *testing.T) {
	val := "t10"
	expected := int64(20)

	result := cast.ToInt(val, int64(20))

	assert.Equal(t, expected, result)
}

func TestCastToDecimal(t *testing.T) {
	val := "10.123"
	expected := float64(10.123)

	result := cast.ToDecimal(val, float64(20.7712))

	assert.Equal(t, expected, result)
}

func TestCastToDecimalFailover(t *testing.T) {
	val := "t10.123"
	expected := float64(20.7712)

	result := cast.ToDecimal(val, float64(20.7712))

	assert.Equal(t, expected, result)
}
