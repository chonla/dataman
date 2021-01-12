package config

// Config represents configuration file structure
type Config struct {
	DatasetPath string       `yaml:"datasets,omitempty"`
	Export      ExportConfig `yaml:"export"`
}

// ExportConfig represents configuration in export section
type ExportConfig struct {
	Target    string            `yaml:"target"`
	Count     int64             `yaml:"count"`
	Variables map[string]string `yml:"variables,omitempty"`
	Fields    []FieldConfig     `yaml:"fields"`
}

// FieldConfig represents configuration for exported field
type FieldConfig struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
	Type  string `yaml:"type,omitempty"`
}

// ParsedTarget is structured target
type ParsedTarget struct {
	Console *ConsoleTarget
	File    *FileTarget
}

// ConsoleTarget is target for console
type ConsoleTarget struct {
	Type string
}

// FileTarget is target for file
type FileTarget struct {
	FileName string
	Type     string
}
