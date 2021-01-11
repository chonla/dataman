package writer

import (
	"fmt"
	"os"
	"strings"
)

// SQLWriter is console writer
type SQLWriter struct {
	writer     *os.File
	objectName string
}

// NewSQLWriter create a new writer writing to console
func NewSQLWriter(writer *os.File, objectName string) IWriter {
	return &SQLWriter{
		writer:     writer,
		objectName: objectName,
	}
}

// Close console writer
func (w *SQLWriter) Close() error {
	return w.writer.Close()
}

// Write content to console
func (w *SQLWriter) Write(header []string, data []map[string]interface{}) error {
	initialSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES", w.objectName, strings.Join(header, ","))

	inserts := []string{}

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
					buffer = append(buffer, fmt.Sprintf("'%s'", w.escapeQuote(row[col].(string))))
				}
			}
		}
		inserts = append(inserts, fmt.Sprintf("(%s)", strings.Join(buffer, ",")))
	}

	sql := fmt.Sprintf("%s %s;", initialSQL, strings.Join(inserts, ","))
	_, err := w.writer.WriteString(sql)
	if err != nil {
		return err
	}

	return nil
}

func (w *SQLWriter) escapeQuote(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
