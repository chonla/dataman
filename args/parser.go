package args

import (
	"encoding/csv"
	"strings"
)

const separator = ','

// Parse parses a given string argument into argument list
// argument separator:
// ,
// argument format:
// a -> [a]
// a,b -> [a, b]
// a,"b,c",d -> [a, "b,c", d]
func Parse(arg string) []string {
	reader := csv.NewReader(strings.NewReader(arg))
	reader.Comma = separator
	argv, err := reader.Read()
	if err != nil {
		return []string{arg}
	}
	return argv
}
