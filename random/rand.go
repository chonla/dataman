package random

import (
	"errors"
	"math"
	"reflect"
	"time"
)

const (
	// MaxRandomSec is maximum second since epoch for date randomization
	MaxRandomSec int64 = 253402300799

	// MaxRandomNanosec is maximum nano second for date randomization
	MaxRandomNanosec int64 = 999999999
)

// IRand is randomizer interface
type IRand interface {
	Seed(int64)
	Int63n(int64) int64
	Float64() float64
}

// Random is random struct
type Random struct {
	random IRand
}

// New creates a new random
func New(r IRand) *Random {
	r.Seed(time.Now().UnixNano())
	return &Random{
		random: r,
	}
}

// CloseInt returns random integer range [0, max]
func (r *Random) CloseInt(max int64) int64 {
	if max >= math.MaxInt64 {
		max = math.MaxInt64 - 2
	}
	return r.random.Int63n(max + 1)
}

// Int returns random integer range [0, max)
func (r *Random) Int(max int64) int64 {
	return r.random.Int63n(max)
}

// IntBetween returns a random integer range [min, max]
func (r *Random) IntBetween(min, max int64) int64 {
	if max < min {
		tmp := max
		max = min
		min = tmp
	}
	delta := max - min + 1
	if delta == int64(1) || delta == int64(2) {
		return min
	}
	val := r.random.Int63n(delta)
	return min + val
}

// Decimal returns random decimal range [0.0, max]
func (r *Random) Decimal(max float64) float64 {
	return r.DecimalBetween(float64(0.0), max)
}

// DecimalBetween returns a random decimal in range [min, max]
func (r *Random) DecimalBetween(min, max float64) float64 {
	if max < min {
		tmp := max
		max = min
		min = tmp
	}
	delta := max - min
	return min + (delta * r.CloseFloat())
}

// CloseFloat returns a random float between [0, 1]
func (r *Random) CloseFloat() float64 {
	var divisor = r.random.Float64()
	var dividend = r.random.Float64()

	if divisor < dividend {
		tmp := divisor
		divisor = dividend
		dividend = tmp
	}

	if divisor == 0.0 {
		return float64(1.0)
	}
	return dividend / divisor
}

// Element returns a random element from given array
func (r *Random) Element(value interface{}) (interface{}, error) {
	array := reflect.ValueOf(value)
	if array.Kind() != reflect.Slice {
		return nil, errors.New("value is not a slice")
	}

	arrLen := array.Len()

	if arrLen == 0 {
		return nil, errors.New("array is empty")
	}
	return array.Index(int(r.Int(int64(arrLen)))).Interface(), nil
}

// Date returns a random date since epoch
func (r *Random) Date() time.Time {
	secOffset := r.CloseInt(MaxRandomSec)
	nanoOffset := r.CloseInt(MaxRandomNanosec)

	return time.Unix(secOffset, nanoOffset)
}

// DateBetween returns a random date in range [min, max]
func (r *Random) DateBetween(min, max time.Time) time.Time {
	if max.Before(min) {
		tmp := max
		max = min
		min = tmp
	}
	delta := max.Sub(min)
	var duration int64

	duration = r.CloseInt(delta.Nanoseconds())
	offset := time.Duration(duration)
	return min.Add(offset)
}

// DateRange return a random date range
func (r *Random) DateRange() Period {
	date1 := r.Date()
	date2 := r.Date()

	if date1.Before(date2) {
		return Period{
			From: date1,
			To:   date2,
		}
	}
	return Period{
		From: date2,
		To:   date1,
	}
}

// DateRangeBetween return a random date range in range [min, max]
func (r *Random) DateRangeBetween(min, max time.Time) Period {
	date1 := r.DateBetween(min, max)
	date2 := r.DateBetween(min, max)

	if date1.Before(date2) {
		return Period{
			From: date1,
			To:   date2,
		}
	}
	return Period{
		From: date2,
		To:   date1,
	}
}
