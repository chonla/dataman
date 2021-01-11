package writer

// IWriter is writer interface
type IWriter interface {
	Write([]string, []map[string]interface{}) error
	Close() error
}
