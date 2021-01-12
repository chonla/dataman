package cast

import "strconv"

// Caster is type caster
type Caster struct{}

// ICaster is caster interface
type ICaster interface {
	ToInt(string, int64) int64
	ToDecimal(string, float64) float64
}

// New creates new caster
func New() ICaster {
	return &Caster{}
}

// ToInt casts an input to base64.
func (c *Caster) ToInt(input string, defaultValue int64) int64 {
	result, e := strconv.ParseInt(input, 10, 64)
	if e != nil {
		return defaultValue
	}
	return result
}

// ToDecimal casts an input to base64.
func (c *Caster) ToDecimal(input string, defaultValue float64) float64 {
	result, e := strconv.ParseFloat(input, 64)
	if e != nil {
		return defaultValue
	}
	return result
}
