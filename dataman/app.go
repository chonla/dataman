package dataman

import (
	"dataman/config"
	"dataman/filesys"
	"dataman/writer"
	"errors"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

// Dataman is an application
type Dataman struct {
	configLoader *config.Loader
	writer       writer.IWriter
}

// New creates a new app instance
func New() *Dataman {
	file := filesys.New()
	return &Dataman{
		configLoader: config.New(file),
		writer:       nil,
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

	for rowIndex = 0; rowIndex < config.Export.Count; rowIndex++ {

	}

	err := d.writer.Write([]string{"hello", "test"}, []map[string]interface{}{{"hello": "world", "test": int64(1)}})
	if err != nil {
		return err
	}
	d.writer.Close()

	return nil
}

// Err colorizes standard error message
func (d *Dataman) Err(msg error) error {
	colorize := color.New(color.FgRed)
	return errors.New(colorize.Sprint(msg.Error()))
}
