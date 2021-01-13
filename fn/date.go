package fn

import (
	"time"
)

func (f *Fn) Date(args []string, varMap map[string]string) string {
	var val time.Time
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

	val = f.rnd.DateBetween(minTime, maxTime)

	return val.Format(time.RFC3339)
}
