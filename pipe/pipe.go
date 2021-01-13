package pipe

import "dataman/cast"

// ResolverPipe is resolver function
type ResolverPipe func(interface{}, []string) interface{}

// Pipe is function handler
type Pipe struct {
	caster cast.ICaster
}

// New creates a new function handler
func New() *Pipe {
	return &Pipe{
		caster: cast.New(),
	}
}
