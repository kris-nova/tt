package interpreter

import (
	"bytes"
	"github.com/kris-nova/tt/library"
	"github.com/pkg/errors"
	"strings"
	"text/template"
)

var Vars map[interface{}]interface{}

func ParseFile(path string) error {
	file, err := LoadFile(path)
	if err != nil {
		return err
	}
	var funcMap template.FuncMap
	funcMap = library.StandardLibrary()
	tmpl, err := template.New(file.AbsolutePath).Funcs(funcMap).Parse(file.String())
	if err != nil {
		return prettyTemplateError(err)
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, "")
	if err != nil {
		return prettyTemplateError(err)
	}
	return nil
}

func prettyTemplateError(err error) error {
	nerr := strings.Replace(err.Error(), "template: ", "", -1)
	return errors.New(nerr)
}
