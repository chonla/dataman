package dataman

import "dataman/fn"

var funcHandler = fn.New()

var supportedFunctions map[string]fn.ResolverFn = map[string]fn.ResolverFn{
	"fn.rowNumber": funcHandler.RowNumber,
	"fn.number":    funcHandler.Number,
}
