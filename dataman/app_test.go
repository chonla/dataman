package dataman_test

import (
	"dataman/dataman"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := dataman.New()

	assert.NotNil(t, app)
}
