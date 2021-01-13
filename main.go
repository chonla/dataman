package main

import (
	"dataman/dataman"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

// AppName for application name
var AppName = "dataman"

// AppVersion for application version
var AppVersion = "development"

// GitCommitID for git commit id
var GitCommitID = "xxxxxx"

// Args for application argument
type Args struct {
	Generate *GenerateCmd `arg:"subcommand:generate" help:"generate random data from given configuration"`
	Validate *ValidateCmd `arg:"subcommand:validate" help:"validate configuration file"`
}

// GenerateCmd for `gen` command
type GenerateCmd struct {
	FileName string `arg:"positional,required"`
}

// ValidateCmd for `validate` command
type ValidateCmd struct {
	FileName string `arg:"positional,required"`
}

// Version returns an application version
func (Args) Version() string {
	return fmt.Sprintf("%s version %s (%s)", AppName, AppVersion, GitCommitID[0:6])
}

func main() {
	app := dataman.New()

	var args Args

	arg.MustParse(&args)

	switch {
	case args.Generate != nil:
		err := app.Generate(args.Generate.FileName)
		if err != nil {
			fmt.Println(app.Err(err))
			os.Exit(1)
		}
	case args.Validate != nil:
		err := app.Validate(args.Validate.FileName)
		if err != nil {
			fmt.Println(app.Err(err))
			os.Exit(1)
		}
		fmt.Printf("File %s is valid.\n", args.Validate.FileName)
	}
}
