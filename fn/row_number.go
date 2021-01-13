package fn

import (
	"dataman/cast"
	"fmt"
)

func (f *Fn) RowNumber(args []string, varMap map[string]string) string {
	startFrom := int64(0)
	caster := cast.New()
	if len(args) > 0 {
		startFrom = caster.ToInt(args[0], int64(0))
	}
	if val, ok := varMap["session.index"]; ok {
		rowNum := caster.ToInt(val, int64(0))
		return fmt.Sprintf("%d", rowNum+startFrom)
	}
	return "0"
}
