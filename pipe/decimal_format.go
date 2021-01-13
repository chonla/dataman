package pipe

import (
	"fmt"
	"reflect"
)

func (p *Pipe) DecimalFormat(val interface{}, args []string) interface{} {
	if len(args) == 0 {
		return val
	}

	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.Float64:
		buff := fmt.Sprintf(args[0], value.Interface().(float64))
		return p.caster.ToDecimal(buff, float64(0.0))
	}
	return val
}
