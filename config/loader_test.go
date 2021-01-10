package config_test

import (
	"dataman/config"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoadConfigShouldReturnErrorIfCannotReadConfigFile(t *testing.T) {
	mockReader := new(MockFileReader)

	mockReader.On("ReadAll", mock.AnythingOfType("string")).Return([]byte{}, errors.New("this file is not readable"))

	conf := config.New(mockReader)
	data, err := conf.Load("a_file_with_problem")

	assert.Nil(t, data)
	assert.NotNil(t, err)
}

func TestLoadConfigShouldReturnConfigurationIfDataIsUnmarshalled(t *testing.T) {
	mockConfig := []byte(`
datasets: ./somepath
export:
    target: console://csv
    count: 10
    fields:
        - name: some name
          value: 10
          type: integer
`)
	expectedConfig := &config.Config{
		DatasetPath: "./somepath",
		Export: config.ExportConfig{
			Target: "console://csv",
			Count:  int64(10),
			Fields: []config.FieldConfig{
				{
					Name:  "some name",
					Value: "10",
					Type:  "integer",
				},
			},
		},
	}

	mockReader := new(MockFileReader)

	mockReader.On("ReadAll", mock.AnythingOfType("string")).Return(mockConfig, nil)

	conf := config.New(mockReader)
	data, err := conf.Load("a_good_file")

	assert.Equal(t, expectedConfig, data)
	assert.Nil(t, err)
}

func TestLoadConfigShouldReturnErrorIfConfigCannotBeUnmarshalled(t *testing.T) {
	mockConfig := []byte(`
datasets: ./somepath
target console://csv
`)

	mockReader := new(MockFileReader)

	mockReader.On("ReadAll", mock.AnythingOfType("string")).Return(mockConfig, nil)

	conf := config.New(mockReader)
	data, err := conf.Load("a_good_file")

	assert.Nil(t, data)
	assert.NotNil(t, err)
}
