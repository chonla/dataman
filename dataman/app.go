package dataman

import (
	"dataman/config"
	"dataman/filesys"
	"errors"

	"github.com/fatih/color"
)

// Dataman is an application
type Dataman struct {
	configLoader *config.Loader
}

// New creates a new app instance
func New() *Dataman {
	file := filesys.New()
	return &Dataman{
		configLoader: config.New(file),
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

	return d.generate(conf)
}

func (d *Dataman) generate(config *config.Config) error {
	return nil
}

// Err colorizes standard error message
func (d *Dataman) Err(msg error) error {
	colorize := color.New(color.FgRed)
	return errors.New(colorize.Sprint(msg.Error()))
}
