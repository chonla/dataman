package random

import (
	"fmt"
	"time"
)

// Period represents date range
type Period struct {
	From time.Time
	To   time.Time
}

func (p Period) String() string {
	return fmt.Sprintf("%s - %s", p.From.Format(time.RFC3339), p.To.Format(time.RFC3339))
}
