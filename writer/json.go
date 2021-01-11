package writer

import (
	"encoding/json"
	"os"
)

// JSONWriter is console writer
type JSONWriter struct {
	writer *os.File
}

// NewJSONWriter create a new writer writing to console
func NewJSONWriter(writer *os.File) IWriter {
	return &JSONWriter{
		writer: writer,
	}
}

// Close console writer
func (w *JSONWriter) Close() error {
	return w.writer.Close()
}

// Write content to console
func (w *JSONWriter) Write(header []string, data []map[string]interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.writer.WriteString(string(b))
	return err
}
