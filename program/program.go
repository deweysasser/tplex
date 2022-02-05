package program

import (
	"github.com/deweysasser/tplex/data"
	"github.com/deweysasser/tplex/tpl"
	"github.com/deweysasser/tplex/ui"
)

type Options struct {
	Templates []string `short:"t" help:"templates or template directories"`
	Prefix    string   `short:"p" help:"Prefix all output files with this string"`
	Data      []string `type:"existing" arg:"" help:"data files to process"`
	OutputDir string   `default:"." short:"o" help:"Output directory"`
	Debug     bool     `help:"Print debugging information"`
	Verbose   bool     `help:"Be more verbose in output"`
}

// Run the main program
func (options Options) Run() {

	for _, f := range options.Data {
		if d, err := data.New(f); err != nil {
			ui.Error(err, "failed to read", d)
		} else {
			if templates, err := tpl.Gather(options.Prefix, options.Templates, d); err != nil {
				ui.Error(err, "failed to read templates")
			} else {
				for _, t := range templates {
					ui.Verbose("Generating", t.Name, "to", t.OutputName)
					if e := t.Generate(options.OutputDir, d); err != nil {
						ui.Error(e, "generating", t.Name)
					}
				}
			}
		}
	}
}
