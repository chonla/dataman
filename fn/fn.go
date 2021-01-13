package fn

import (
	"dataman/random"
)

// ResolverFn is resolver function
type ResolverFn func([]string, map[string]string) string

// Fn is function handler
type Fn struct {
	rnd *random.Random
}

// New creates a new function handler
func New(rnd *random.Random) *Fn {
	return &Fn{
		rnd: rnd,
	}
}
