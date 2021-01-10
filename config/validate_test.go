package config_test

import (
	"dataman/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConsoleTarget(t *testing.T) {
	expected := &config.ParsedTarget{
		Console: &config.ConsoleTarget{
			Type: "csv",
		},
		File: nil,
	}

	conf := &config.Config{}

	result, err := conf.ParseTarget("console://csv")

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParseFileTarget(t *testing.T) {
	expected := &config.ParsedTarget{
		Console: nil,
		File: &config.FileTarget{
			Type:     "csv",
			FileName: "some-path/some-file.csv",
		},
	}

	conf := &config.Config{}

	result, err := conf.ParseTarget("file://some-path/some-file.csv")

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParseUnsupportedTarget(t *testing.T) {
	expected := &config.ParsedTarget{
		Console: nil,
		File:    nil,
	}

	conf := &config.Config{}

	result, err := conf.ParseTarget("unsupported://some-path/some-file.csv")

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestValidateFileTargetSuccess(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.csv",
			Count:  1,
			Fields: []config.FieldConfig{
				{
					Name:  "Some Name",
					Type:  "integer",
					Value: "1",
				},
			},
		},
	}

	err := conf.Validate()

	assert.Nil(t, err)
}

func TestValidateConsoleTargetSuccess(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "console://csv",
			Count:  1,
			Fields: []config.FieldConfig{
				{
					Name:  "Some Name",
					Type:  "integer",
					Value: "1",
				},
			},
		},
	}

	err := conf.Validate()

	assert.Nil(t, err)
}

func TestValidateShouldReturnErrIfTargetIsMissing(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfTargetIsUnsupported(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "unsupported://some-path/some-file.csv",
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportTypeIsUnsupportedForConsole(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "console://unsupported",
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportTypeIsUnsupportedForFile(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.unsupported",
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportCountIsMissing(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.csv",
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportCountIsNegative(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.csv",
			Count:  -1,
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportCountIsZero(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.csv",
			Count:  0,
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportFieldsIsMissing(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.csv",
			Count:  1,
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportFieldsIsEmpty(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.csv",
			Count:  1,
			Fields: []config.FieldConfig{},
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}

func TestValidateShouldReturnErrIfExportFieldTypeIsUnsupported(t *testing.T) {
	conf := &config.Config{
		Export: config.ExportConfig{
			Target: "file://some-file.csv",
			Count:  1,
			Fields: []config.FieldConfig{
				{
					Name:  "Some field",
					Value: "1",
					Type:  "unsupported",
				},
			},
		},
	}

	err := conf.Validate()

	assert.NotNil(t, err)
}
