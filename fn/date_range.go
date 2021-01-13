package fn

import (
	"dataman/cast"
	"dataman/random"
	"time"
)

func (f *Fn) DateRange(args []string, varMap map[string]string) string {
	var val random.Period
	var maxDate string = "2599-12-31"
	var minDate string = "1970-01-01"

	if len(args) == 1 {
		maxDate = args[0]
	} else {
		if len(args) > 1 {
			minDate = args[0]
			maxDate = args[1]
		}
	}

	minTime, _ := time.Parse("2006-01-02", minDate)
	maxTime, _ := time.Parse("2006-01-02", maxDate)

	val = f.rnd.DateRangeBetween(minTime, maxTime)

	return val.String()
}

func (f *Fn) DateRangeOffset(args []string, varMap map[string]string) string {
	var val random.Period
	var maxDate string = "2599-12-31"
	var minDate string = "1970-01-01"
	var offsetDays string = "30"

	if len(args) == 1 {
		offsetDays = args[0]
	} else {
		if len(args) == 2 {
			offsetDays = args[0]
			maxDate = args[1]
		} else {
			if len(args) > 2 {
				offsetDays = args[0]
				minDate = args[1]
				maxDate = args[2]
			}
		}
	}

	caster := cast.New()
	minTime, _ := time.Parse("2006-01-02", minDate)
	maxTime, _ := time.Parse("2006-01-02", maxDate)
	maxOffset := caster.ToInt(offsetDays, int64(30))

	val = f.rnd.DateRangeBetweenByOffset(minTime, maxTime, maxOffset)

	return val.String()
}
