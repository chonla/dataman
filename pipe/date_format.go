package pipe

import (
	"reflect"
	"time"
)

func (p *Pipe) DateFormat(val interface{}, args []string) interface{} {
	if len(args) == 0 {
		return val
	}

	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.String:
		date, err := time.Parse(time.RFC3339, value.Interface().(string))
		if err != nil {
			return val
		}
		return date.Format(args[0])
	}
	return val
}
