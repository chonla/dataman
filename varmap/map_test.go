package varmap_test

import (
	"dataman/varmap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportShouldImportSourceMapToDestMapWithPrefix(t *testing.T) {
	src := map[string]string{
		"a": "1",
		"b": "2",
	}
	dest := map[string]string{
		"A": "11",
		"B": "22",
	}
	expected := map[string]string{
		"var.a": "1",
		"var.b": "2",
		"A":     "11",
		"B":     "22",
	}

	result := varmap.Import(dest, src)

	assert.Equal(t, result, expected)
}

func TestImportShouldReplaceSourceMapIfKeyCollides(t *testing.T) {
	src := map[string]string{
		"a": "1",
		"c": "2",
	}
	dest := map[string]string{
		"var.a": "11",
		"var.b": "22",
	}
	expected := map[string]string{
		"var.a": "1",
		"var.b": "22",
		"var.c": "2",
	}

	result := varmap.Import(dest, src)

	assert.Equal(t, result, expected)
}
