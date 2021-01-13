package fn

import (
	"dataman/cast"
	"fmt"
	"time"
)

func (f *Fn) RowSequence(args []string, varMap map[string]string) string {
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

func (f *Fn) DateSequence(args []string, varMap map[string]string) string {
	startFrom := time.Now()

	if len(args) > 0 {
		startFrom, _ = time.Parse("2006-01-02", args[0])
	}

	caster := cast.New()
	if val, ok := varMap["session.index"]; ok {
		rowNum := caster.ToInt(val, int64(1)) - 1
		offset := time.Duration(rowNum * int64(time.Hour*24))
		return startFrom.Add(offset).Format("2006-01-02")
	}
	return startFrom.Format("2006-01-02")
}
