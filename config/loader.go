package config

import (
	"dataman/filesys"

	"gopkg.in/yaml.v2"
)

// Loader is configuration loader
type Loader struct {
	file filesys.IFileSys
}

// New creates a configuration loader
func New(f filesys.IFileSys) *Loader {
	return &Loader{
		file: f,
	}
}

// Load configuration file from filename and return Config object
func (l *Loader) Load(filename string) (*Config, error) {
	configData, err := l.file.ReadAll(filename)
	if err != nil {
		return nil, err
	}

	var conf = Config{}

	err = yaml.Unmarshal(configData, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
