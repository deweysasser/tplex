package tpl

import (
	"github.com/deweysasser/tplex/data"
	"github.com/deweysasser/tplex/ui"
	"html/template"
	"os"
	"path"
	"strings"
)

type Template struct {
	*template.Template
	Name, Path string
	OutputName string
}

func (t *Template) Generate(outputDir string, d data.Data) error {
	info, err := os.Stat(outputDir)
	switch {
	case os.IsNotExist(err):
		if err = os.MkdirAll(outputDir, os.ModeDir|os.ModePerm); err != nil {
			return err
		}
	case info.IsDir():
		// all good
	default:
		return err
	}

	out := path.Join(outputDir, t.OutputName)

	f, err := os.Create(out)

	if err != nil {
		return err
	}

	defer f.Close()

	err = t.ExecuteTemplate(f, t.Name, d.Content)

	return err
}

type errors []error

func (e errors) Error() string {
	s := make([]string, len(e))

	for i := 0; i < len(e); i++ {
		s[i] = e[i].Error()
	}

	return strings.Join(s, "\n")
}

func New(prefix, file string) (Template, error) {
	if t, err := template.New(file).ParseFiles(file); err != nil {
		return Template{}, err
	} else {
		var p string

		switch prefix {
		case "":
			p = path.Base(file)
		default:
			p = path.Join(prefix, path.Base(file))
		}

		return Template{
			Template:   t,
			Name:       path.Base(file),
			Path:       file,
			OutputName: p,
		}, nil
	}
}

func Gather(prefix string, paths []string, data data.Data) ([]Template, error) {
	var templates []Template
	var e errors

	for _, p := range paths {
		ui.Debug("Creating template from", p)
		if t, err := New(prefix, p); err != nil {
			e = append(e, err)
		} else {
			templates = append(templates, t)
		}
	}

	for _, p := range data.Templates() {
		ui.Debug("Creating template from data", p)

		if t, err := New(prefix, p); err != nil {
			e = append(e, err)
		} else {
			templates = append(templates, t)
		}
	}

	if len(e) > 0 {
		return templates, e
	} else {
		return templates, nil
	}
}
