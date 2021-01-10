package config

import (
	"dataman/array"
	"dataman/text"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// SupportedExportTypes are supported export file type
var SupportedExportTypes = []string{"csv", "tsv", "json", "sql"}

// SupportedFieldTypes are supported export field type
var SupportedFieldTypes = []string{"integer", "decimal", "string"}

// Validate performs self-validate
func (c *Config) Validate() error {
	target, err := c.ParseTarget(c.Export.Target)
	if err != nil {
		return err
	}

	if err := c.validateExportTarget(target); err != nil {
		return err
	}

	if err := c.validateExportCount(c.Export.Count); err != nil {
		return err
	}

	if err := c.validateExportFields(c.Export.Fields); err != nil {
		return err
	}

	return nil
}

// ParseTarget parse a given target in to a structure
func (c *Config) ParseTarget(target string) (*ParsedTarget, error) {
	var parsedTarget ParsedTarget

	if target == "" {
		return nil, errors.New("Invalid or missing field: export.target")
	}

	if text.StartWith(target, "console://") {
		exportType := target[10:]
		parsedTarget.Console = &ConsoleTarget{
			Type: exportType,
		}
	}

	if text.StartWith(target, "file://") {
		exportFile := target[7:]
		exportType := ""
		if strings.Contains(exportFile, ".") {
			exportType = filepath.Ext(exportFile)[1:]
		}
		parsedTarget.File = &FileTarget{
			FileName: exportFile,
			Type:     exportType,
		}
	}

	return &parsedTarget, nil
}

func (c *Config) validateExportType(exportType string) error {
	if array.IndexOf(SupportedExportTypes, exportType) == array.ErrorNotFound {
		return fmt.Errorf("Unsupported export type: %s", exportType)
	}
	return nil
}

func (c *Config) validateExportCount(count int64) error {
	if count <= 0 {
		return errors.New("Invalid or missing field: export.count")
	}
	return nil
}

func (c *Config) validateExportTarget(target *ParsedTarget) error {
	if target.Console != nil {
		if err := c.validateExportType(target.Console.Type); err != nil {
			return err
		}
		return nil
	}

	if target.File != nil {
		if err := c.validateExportType(target.File.Type); err != nil {
			return err
		}
		return nil
	}

	return errors.New("Invalid or missing field: export.target")
}

func (c *Config) validateExportFields(fields []FieldConfig) error {
	if len(fields) == 0 {
		return errors.New("Invalid or missing field: export.fields")
	}

	for _, field := range fields {
		if err := c.validateExportFieldType(field.Type); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) validateExportFieldType(fieldType string) error {
	if array.IndexOf(SupportedFieldTypes, fieldType) == array.ErrorNotFound {
		return fmt.Errorf("Unsupported field type: %s", fieldType)
	}
	return nil
}
