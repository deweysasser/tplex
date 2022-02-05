package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/deweysasser/tplex/program"
	"github.com/deweysasser/tplex/ui"
	"os"
)

var options program.Options

func main() {
	context := kong.Parse(&options,
		kong.Description("Expand templates into files"),
	)

	ui.DebugOn = options.Debug
	ui.VerboseOn = options.Verbose

	if err := context.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
