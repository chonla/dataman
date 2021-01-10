package args_test

import (
	"testing"

	"dataman/args"

	"github.com/stretchr/testify/assert"
)

func TestParseOneArg(t *testing.T) {
	arg := `hello`
	expected := []string{"hello"}

	result := args.Parse(arg)

	assert.Equal(t, expected, result)
}

func TestParseTwoArgs(t *testing.T) {
	arg := `hello,world`
	expected := []string{"hello", "world"}

	result := args.Parse(arg)

	assert.Equal(t, expected, result)
}

func TestParseMultipleArgs(t *testing.T) {
	arg := `hello,world,ok,then,now`
	expected := []string{"hello", "world", "ok", "then", "now"}

	result := args.Parse(arg)

	assert.Equal(t, expected, result)
}

func TestParseMultipleArgsWithQuotes(t *testing.T) {
	arg := `hello,world,"ok,then",now`
	expected := []string{"hello", "world", "ok,then", "now"}

	result := args.Parse(arg)

	assert.Equal(t, expected, result)
}

func TestParseMalformArg(t *testing.T) {
	arg := `hello",world,"ok,then",now`
	expected := []string{`hello",world,"ok,then",now`}

	result := args.Parse(arg)

	assert.Equal(t, expected, result)
}
