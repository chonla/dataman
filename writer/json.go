package writer

import (
	"encoding/json"
	"os"
)

// JSONWriter is console writer
type JSONWriter struct {
	writer *os.File
	pretty bool
}

// NewJSONWriter create a new writer writing to console
func NewJSONWriter(writer *os.File, pretty bool) IWriter {
	return &JSONWriter{
		writer: writer,
		pretty: pretty,
	}
}

// Close console writer
func (w *JSONWriter) Close() error {
	return w.writer.Close()
}

// Write content to console
func (w *JSONWriter) Write(header []string, data []map[string]interface{}) error {
	var b []byte
	var err error
	if w.pretty {
		b, err = json.MarshalIndent(data, "", "    ")
	} else {
		b, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	_, err = w.writer.WriteString(string(b))
	return err
}
