package dataman

import (
	"dataman/fn"
	"dataman/pipe"
)

var funcHandler = fn.New()

var supportedFunctions map[string]fn.ResolverFn = map[string]fn.ResolverFn{
	"fn.rowNumber": funcHandler.RowNumber,
	"fn.number":    funcHandler.Number,
	"fn.decimal":   funcHandler.Decimal,
	"fn.date":      funcHandler.Date,
}

var pipeHandler = pipe.New()

var supportedPipes map[string]pipe.ResolverPipe = map[string]pipe.ResolverPipe{
	"pipe.decimalFormat": pipeHandler.DecimalFormat,
	"pipe.dateFormat":    pipeHandler.DateFormat,
}
