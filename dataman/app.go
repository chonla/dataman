package dataman

import (
	"dataman/cast"
	"dataman/config"
	"dataman/filesys"
	"dataman/varmap"
	"dataman/writer"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// Dataman is an application
type Dataman struct {
	configLoader *config.Loader
	writer       writer.IWriter
	caster       cast.ICaster
}

// New creates a new app instance
func New() *Dataman {
	file := filesys.New()
	caster := cast.New()
	return &Dataman{
		configLoader: config.New(file),
		writer:       nil,
		caster:       caster,
	}
}

// Validate to validate mandatory fields in configFile
func (d *Dataman) Validate(configFile string) error {
	conf, err := d.configLoader.Load(configFile)
	if err != nil {
		return err
	}

	return conf.Validate()
}

// Generate a random content from configFile
func (d *Dataman) Generate(configFile string) error {
	conf, err := d.configLoader.Load(configFile)
	if err != nil {
		return err
	}

	if err = conf.Validate(); err != nil {
		return err
	}

	target, _ := conf.ParseTarget(conf.Export.Target)

	var output *os.File
	var contentType string
	switch {
	case target.Console != nil:
		output = os.Stdout
		contentType = target.Console.Type
	}

	switch contentType {
	case "json":
		d.writer = writer.NewJSONWriter(output)
	case "csv":
		d.writer = writer.NewXSVWriter(output, ",")
	case "tsv":
		d.writer = writer.NewXSVWriter(output, "\t")
	case "sql":
		ext := filepath.Ext(configFile)
		objectName := filepath.Base(configFile[:len(configFile)-len(ext)])

		d.writer = writer.NewSQLWriter(output, objectName)
	}

	return d.generate(conf)
}

func (d *Dataman) generate(config *config.Config) error {
	var rowIndex int64
	records := []map[string]interface{}{}

	for rowIndex = 0; rowIndex < config.Export.Count; rowIndex++ {
		row := d.newRow(rowIndex, config)
		records = append(records, row)
	}

	headers := []string{}
	for _, field := range config.Export.Fields {
		headers = append(headers, field.Name)
	}

	err := d.writer.Write(headers, records)
	if err != nil {
		return err
	}
	d.writer.Close()

	return nil
}

func (d *Dataman) newRow(index int64, config *config.Config) map[string]interface{} {
	sessionVars := map[string]string{
		"session.index": fmt.Sprintf("%d", index+1),
	}

	varMap := varmap.Import(sessionVars, config.Export.Variables)

	output := map[string]interface{}{}
	for _, field := range config.Export.Fields {
		output[field.Name] = d.createValue(field, varMap)
	}
	return output
}

func (d *Dataman) createValue(field config.FieldConfig, varMap map[string]string) interface{} {
	variables := d.captureVariables(field.Value)
	resolvedVariables := d.resolveVariables(variables, varMap)
	result := field.Value

	for varName, varValue := range resolvedVariables {
		targetVarName := fmt.Sprintf("${%s}", varName)
		result = strings.ReplaceAll(result, targetVarName, varValue)
	}

	var resolvedValue interface{}
	switch field.Type {
	case "integer":
		resolvedValue = d.caster.ToInt(result, int64(0))
	case "decimal":
		resolvedValue = d.caster.ToDecimal(result, float64(0.0))
	default:
		resolvedValue = result
	}

	return resolvedValue
}

func (d *Dataman) captureVariables(template string) []string {
	re := regexp.MustCompile("\\$\\{([^}]+)\\}")
	m := re.FindAllStringSubmatch(template, -1)
	if len(m) > 0 && len(m[0]) == 2 {
		matched := []string{}
		for i, ml := 0, len(m); i < ml; i++ {
			matched = append(matched, m[i][1])
		}
		return matched
	}
	return []string{}
}

func (d *Dataman) resolveVariables(variables []string, varMap map[string]string) map[string]string {
	return varMap
}

// Err colorizes standard error message
func (d *Dataman) Err(msg error) error {
	colorize := color.New(color.FgRed)
	return errors.New(colorize.Sprint(msg.Error()))
}
