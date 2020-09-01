package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)
type ToolAPI map[string]Tool

type Tool struct {
	Name string
	Parameters map[string]interface{}
	Script string
}

func (t Tool) Run(args map[string]interface{}) (string, error) {
	file, err := ioutil.TempFile("", "script.*.sh")
	if err != nil {
		return "", fmt.Errorf("tool Run failed: %v", err)
	}
	defer func() {
		removeError := os.Remove(file.Name())
		if err == nil {
			err = removeError
		}
	}()
	tmpl, err := template.ParseFiles(t.Script)
	if err != nil {
		return "", fmt.Errorf("tool Run failed: %v", err)
	}
	if err := tmpl.Execute(file, args); err != nil {
		return "", fmt.Errorf("tool Run failed: %v", err)
	}
	out, err := exec.Command("/bin/sh", file.Name()).Output()
	if err != nil {
		return string(out), fmt.Errorf("tool Run failed: %v", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func (t Tool) toJsonSchema() Schema {
	return Schema{
		"type": "object",
		"properties": Schema{
			"Resource": Schema{
				"type": "string",
				"const": t.Name,
			},
			"parameters": t.Parameters,
		},
	}
}
