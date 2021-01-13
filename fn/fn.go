package fn

import (
	"dataman/random"
	"math/rand"
	"time"
)

// ResolverFn is resolver function
type ResolverFn func([]string, map[string]string) string

// Fn is function handler
type Fn struct {
	rnd *random.Random
}

// New creates a new function handler
func New() *Fn {
	return &Fn{
		rnd: random.New(rand.New(rand.NewSource(time.Now().UnixNano()))),
	}
}
