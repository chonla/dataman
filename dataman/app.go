package dataman

import (
	"dataman/args"
	"dataman/cast"
	"dataman/config"
	"dataman/filesys"
	"dataman/fn"
	"dataman/pipe"
	"dataman/random"
	"dataman/varmap"
	"dataman/writer"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Dataman is an application
type Dataman struct {
	configLoader *config.Loader
	writer       writer.IWriter
	caster       cast.ICaster
	functions    map[string]fn.ResolverFn
	pipes        map[string]pipe.ResolverPipe
	datasets     map[string][]string
	rnd          *random.Random
}

// New creates a new app instance
func New() *Dataman {
	file := filesys.New()
	caster := cast.New()
	rnd := random.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	fh := fn.New(rnd)
	ph := pipe.New()

	return &Dataman{
		configLoader: config.New(file),
		writer:       nil,
		caster:       caster,
		functions: map[string]fn.ResolverFn{
			"fn.rowSequence":     fh.RowSequence,
			"fn.dateSequence":    fh.DateSequence,
			"fn.number":          fh.Number,
			"fn.decimal":         fh.Decimal,
			"fn.date":            fh.Date,
			"fn.dateRange":       fh.DateRange,
			"fn.dateRangeOffset": fh.DateRangeOffset,
		},
		pipes: map[string]pipe.ResolverPipe{
			"pipe.decimalFormat": ph.DecimalFormat,
			"pipe.dateFormat":    ph.DateFormat,
		},
		datasets: map[string][]string{},
		rnd:      rnd,
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

func (d *Dataman) LoadDatasets(path string) error {
	files, err := filepath.Glob(fmt.Sprintf("%s/*.txt", path))
	if err != nil {
		return err
	}

	for _, file := range files {
		_, fileName := filepath.Split(file)
		datasetName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
		if err = d.LoadDataset(datasetName, file); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dataman) LoadDataset(name, path string) error {
	dataByte, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	data := strings.Split(string(dataByte), "\n")
	d.datasets[fmt.Sprintf("data.%s", name)] = data
	return nil
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

	if err = d.LoadDatasets(conf.DatasetPath); err != nil {
		return err
	}

	target, _ := conf.ParseTarget(conf.Export.Target)

	var output *os.File
	var contentType string
	switch {
	case target.Console != nil:
		output = os.Stdout
		contentType = target.Console.Type
	case target.File != nil:
		writer, err := os.Create(target.File.FileName)
		if err != nil {
			return err
		}
		output = writer
		contentType = target.File.Type
	}

	switch contentType {
	case "json":
		d.writer = writer.NewJSONWriter(output, conf.Export.Pretty)
	case "csv":
		d.writer = writer.NewXSVWriter(output, ",")
	case "tsv":
		d.writer = writer.NewXSVWriter(output, "\t")
	case "sql":
		ext := filepath.Ext(configFile)
		objectName := filepath.Base(configFile[:len(configFile)-len(ext)])

		d.writer = writer.NewSQLWriter(output, objectName, conf.Export.Pretty)
	}

	return d.generate(conf)
}

func (d *Dataman) generate(config *config.Config) error {
	var rowIndex int64
	records := []map[string]interface{}{}

	for rowIndex = 0; rowIndex < config.Export.Count; rowIndex++ {
		row, err := d.newRow(rowIndex, config)
		if err != nil {
			return err
		}
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

func (d *Dataman) newRow(index int64, config *config.Config) (map[string]interface{}, error) {
	sessionVars := map[string]string{
		"session.index": fmt.Sprintf("%d", index+1),
	}

	varMap := varmap.Import(sessionVars, config.Export.Variables)

	var err error
	output := map[string]interface{}{}
	for _, field := range config.Export.Fields {
		output[field.Name], err = d.createValue(field, varMap)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func (d *Dataman) createValue(field config.FieldConfig, varMap map[string]string) (interface{}, error) {
	variables := d.captureVariables(field.Value)
	funcs := d.captureFunctions(field.Value)
	data := d.captureData(field.Value)

	resolvedVariables, err := d.resolveVariables(varMap)
	if err != nil {
		return nil, err
	}

	result := field.Value

	for _, varName := range variables {
		if varValue, ok := resolvedVariables[varName]; ok {
			targetVarName := fmt.Sprintf("${%s}", varName)
			result = strings.ReplaceAll(result, targetVarName, varValue)
		} else {
			return nil, fmt.Errorf("Unable to resolve variable %s", varName)
		}
	}

	if len(data) > 0 {
		for _, dataName := range data {
			if dataValue, err := d.rnd.Element(d.datasets[dataName]); err == nil {
				targetVarName := fmt.Sprintf("${%s}", dataName)
				result = strings.ReplaceAll(result, targetVarName, dataValue.(string))
			} else {
				return nil, err
			}
		}
	}

	if len(funcs) > 0 {
		for _, fnName := range funcs {
			if fnValue, fnArgs, err := d.parseFunc(fnName); err == nil {
				targetVarName := fmt.Sprintf("${%s}", fnName)
				result = strings.ReplaceAll(result, targetVarName, fnValue(fnArgs, varMap))
			} else {
				return nil, err
			}
		}
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

	for _, layout := range field.Layout {
		resolvedValue, err = d.pipe(resolvedValue, layout)
		if err != nil {
			return nil, err
		}
	}

	return resolvedValue, nil
}

func (d *Dataman) captureData(template string) []string {
	return d.capture(template, "data")
}

func (d *Dataman) captureVariables(template string) []string {
	return d.capture(template, "var")
}

func (d *Dataman) captureFunctions(template string) []string {
	return d.capture(template, "fn")
}

func (d *Dataman) capture(template string, prefix string) []string {
	re := regexp.MustCompile(fmt.Sprintf("\\$\\{(%s\\.[^}]+)\\}", prefix))
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

func (d *Dataman) resolveVariables(varMap map[string]string) (map[string]string, error) {
	variablesFound := false
	for k, v := range varMap {
		variables := d.captureVariables(v)
		funcs := d.captureFunctions(v)
		data := d.captureData(v)
		result := v
		if len(variables) > 0 {
			for _, varName := range variables {
				if varValue, ok := varMap[varName]; ok {
					targetVarName := fmt.Sprintf("${%s}", varName)
					result = strings.ReplaceAll(result, targetVarName, varValue)
				} else {
					return nil, fmt.Errorf("Unable to resolve variable %s", varName)
				}
			}
			varMap[k] = result
			variablesFound = true
		}

		if len(data) > 0 {
			for _, dataName := range data {
				if dataValue, err := d.rnd.Element(d.datasets[dataName]); err == nil {
					targetVarName := fmt.Sprintf("${%s}", dataName)
					result = strings.ReplaceAll(result, targetVarName, dataValue.(string))
				} else {
					return nil, err
				}
			}
			varMap[k] = result
			variablesFound = true
		}

		if len(funcs) > 0 {
			var targetValue string

			for _, fnName := range funcs {
				if fnValue, fnArgs, err := d.parseFunc(fnName); err == nil {
					targetVarName := fmt.Sprintf("${%s}", fnName)
					targetValue = fnValue(fnArgs, varMap)
					result = strings.ReplaceAll(result, targetVarName, targetValue)
				} else {
					return nil, err
				}
			}

			// Special hook when there is only one function declared and function return objects
			// Apply only in variables part
			if len(funcs) == 1 {
				fnCall := strings.SplitN(funcs[0], ":", 2)
				switch fnCall[0] {
				case "fn.dateRange", "fn.dateRangeOffset":
					dateRange := strings.Split(targetValue, " - ")
					varMap[fmt.Sprintf("%s.from", k)] = dateRange[0]
					varMap[fmt.Sprintf("%s.to", k)] = dateRange[1]
				case "fn.numberRange":
					numberRange := strings.Split(targetValue, " - ")
					varMap[fmt.Sprintf("%s.from", k)] = numberRange[0]
					varMap[fmt.Sprintf("%s.to", k)] = numberRange[1]
				}
			}

			varMap[k] = result
			variablesFound = true
		}
	}
	if variablesFound {
		// resolve repeated var references
		return d.resolveVariables(varMap)
	}
	return varMap, nil
}

func (d *Dataman) pipe(s interface{}, pipeFn string) (interface{}, error) {
	value := s
	pipeCall := strings.SplitN(pipeFn, ":", 2)
	pipeName := fmt.Sprintf("pipe.%s", strings.TrimSpace(pipeCall[0]))

	if resolver, ok := d.pipes[pipeName]; ok {
		if len(pipeCall) > 1 {
			pipeArgs := strings.TrimSpace(pipeCall[1])
			value = resolver(value, args.Parse(pipeArgs))
		} else {
			value = resolver(value, []string{})
		}
	} else {
		return "", fmt.Errorf("Unable to resolve pipe %s", pipeName)
	}

	return value, nil
}

func (d *Dataman) parseFunc(fn string) (fn.ResolverFn, []string, error) {
	result := strings.SplitN(fn, ":", 2)

	if resolver, ok := d.functions[result[0]]; ok {
		if len(result) > 1 {
			return resolver, args.Parse(result[1]), nil
		}
		return resolver, []string{}, nil
	}
	return nil, nil, fmt.Errorf("Unable to resolve function %s", result[0])
}

// Err colorizes standard error message
func (d *Dataman) Err(msg error) error {
	colorize := color.New(color.FgRed)
	return errors.New(colorize.Sprint(msg.Error()))
}
