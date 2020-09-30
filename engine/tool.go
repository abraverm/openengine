package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// ToolAPI is a list of tools.
type ToolAPI map[string]Tool

// A Tool is a shell script that only requires parameters for proper execution.
type Tool struct {
	Name       string                 `json:"name"`
	Parameters map[string]interface{} `json:"parameters"`
	Script     string                 `json:"script"`
}

// Run tool script with given parameters.
func (t Tool) Run(args map[string]interface{}) (string, error) {
	file, err := ioutil.TempFile("", "script.*.sh")
	if err != nil {
		return "", fmt.Errorf("tool Run failed: %w", err)
	}

	defer func() {
		if removeError := os.Remove(file.Name()); removeError != nil {
			err = removeError
		}
	}()

	tmpl, err := template.ParseFiles(t.Script)
	if err != nil {
		return "", fmt.Errorf("tool Run failed: %w", err)
	}

	if err := tmpl.Execute(file, args); err != nil {
		return "", fmt.Errorf("tool Run failed: %w", err)
	}

	// nolint: gosec
	out, err := exec.Command("/bin/sh", file.Name()).Output()
	if err != nil {
		return string(out), fmt.Errorf("tool Run failed: %w\n%v", err, string(out))
	}

	return strings.TrimSpace(string(out)), nil
}
