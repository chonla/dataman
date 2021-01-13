package dataman

import "dataman/fn"

var supportedFunctions map[string]fn.Fn = map[string]fn.Fn{
	"fn.rowNumber": fn.RowNumber,
}
