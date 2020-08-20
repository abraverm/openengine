package engine2

import (
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

type Tool struct {
	Name string
	Parameters map[string]interface{}
	Script string
}

func (t Tool) match(task ImplicitTask) bool {
	data := gojsonschema.NewGoLoader(task)
	loader := gojsonschema.NewGoLoader(t.toJsonSchema())
	result, err := gojsonschema.Validate(loader, data)
	if err != nil {
		return false
	}
	if result.Valid() {
		return true
	}
	return false
}

func (t Tool) Run(args map[string]interface{}) (string, error) {
	file, err := ioutil.TempFile("", "script.*.sh")
	if err != nil {
		return "", err
	}
	defer func() {
		removeError := os.Remove(file.Name())
		if err == nil {
			err = removeError
		}
	}()
	tmpl, err := template.ParseFiles(t.Script)
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(file, args); err != nil {
		return "", err
	}
	out, err := exec.Command("/bin/sh", file.Name()).Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func (t Tool) toJsonSchema() JSONSchema {
	return JSONSchema{
		"type": "object",
		"properties": JSONSchema{
			"resource": JSONSchema{
				"type": "string",
				"const": t.Name,
			},
			"parameters": t.Parameters,
		},
	}
}
