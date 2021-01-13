package fn

import (
	"dataman/cast"
	"fmt"
	"math"
)

func (f *Fn) Number(args []string, varMap map[string]string) string {
	var val int64
	var maxRand int64 = math.MaxInt64 - 1
	var minRand int64 = int64(0)

	caster := cast.New()
	if len(args) == 1 {
		maxRand = caster.ToInt(args[0], maxRand)
	} else {
		if len(args) > 1 {
			minRand = caster.ToInt(args[0], minRand)
			maxRand = caster.ToInt(args[1], maxRand)
		}
	}
	val = f.rnd.IntBetween(minRand, maxRand)

	return fmt.Sprintf("%d", val)
}
