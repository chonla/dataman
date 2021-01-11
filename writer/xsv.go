package writer

import (
	"fmt"
	"os"
	"strings"
)

// XSVWriter is something separated value writer
type XSVWriter struct {
	writer    *os.File
	separator string
}

// NewXSVWriter create a new writer writing as something separated value
func NewXSVWriter(writer *os.File, separator string) IWriter {
	return &XSVWriter{
		writer:    writer,
		separator: separator,
	}
}

// Close writer
func (w *XSVWriter) Close() error {
	return w.writer.Close()
}

// Write content to writer
func (w *XSVWriter) Write(header []string, data []map[string]interface{}) error {
	_, err := w.writer.WriteString(fmt.Sprintf("%s\n", strings.Join(header, w.separator)))
	if err != nil {
		return err
	}
	for _, row := range data {
		buffer := []string{}
		for _, col := range header {
			intVal, intOk := row[col].(int64)
			if intOk {
				buffer = append(buffer, fmt.Sprintf("%d", intVal))
			} else {
				floatVal, floatOk := row[col].(float64)
				if floatOk {
					buffer = append(buffer, fmt.Sprintf("%f", floatVal))
				} else {
					buffer = append(buffer, fmt.Sprintf("\"%s\"", row[col].(string)))
				}
			}
		}
		_, err = w.writer.WriteString(fmt.Sprintf("%s\n", strings.Join(buffer, w.separator)))
		if err != nil {
			return err
		}
	}
	return nil
}
