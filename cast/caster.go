package cast

import "strconv"

// ToInt casts an input to base64.
func ToInt(input string, defaultValue int64) int64 {
	result, e := strconv.ParseInt(input, 10, 64)
	if e != nil {
		return defaultValue
	}
	return result
}

// ToDecimal casts an input to base64.
func ToDecimal(input string, defaultValue float64) float64 {
	result, e := strconv.ParseFloat(input, 64)
	if e != nil {
		return defaultValue
	}
	return result
}
