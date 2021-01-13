package fn

import (
	"dataman/cast"
	"fmt"
)

func (f *Fn) Decimal(args []string, varMap map[string]string) string {
	var val float64
	var maxRand float64 = float64(10000000.0)
	var minRand float64 = float64(0.0)
	var precision int64 = 10

	caster := cast.New()
	if len(args) == 1 {
		precision = caster.ToInt(args[0], precision)
	} else {
		if len(args) == 2 {
			precision = caster.ToInt(args[0], precision)
			maxRand = caster.ToDecimal(args[1], maxRand)
		} else {
			if len(args) > 2 {
				precision = caster.ToInt(args[0], precision)
				minRand = caster.ToDecimal(args[1], minRand)
				maxRand = caster.ToDecimal(args[2], maxRand)
			}
		}
	}

	val = f.rnd.DecimalBetween(minRand, maxRand)
	layout := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(layout, val)
}
