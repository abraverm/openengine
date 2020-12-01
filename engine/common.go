// common is where all the common struct go
package engine

import (
	"fmt"
	"regexp"
)

// Schema is the top-level structure defining a json schema.
// TODO: how about using jsonschema directly instead?
type Schema map[string]interface{}

// System is a provider instance that contains matching values and other metadata such as credentials.
type System map[string]interface{}

// Resource is the user requested resource with its type and parameters.
type Resource struct {
	Name string                 `yaml:"resource"`
	Args map[string]interface{} `json:"args"`
}

// ToolAPI is a list of tools.
type ToolAPI map[string]Tool

// A Tool is a shell script that only requires parameters for proper execution.
type Tool struct {
	Name       string                 `json:"name"`
	Parameters map[string]interface{} `json:"parameters"`
	Script     string                 `json:"script"`
}

func (r Resource) getImplicitKeys() []string {
	re := regexp.MustCompile(`\$_[[:alpha:]]*`)

	// nolint: prealloc
	var keys []string

	for key, value := range r.Args {
		if re.MatchString(key) {
			keys = append(keys, key)
		}

		keys = append(keys, re.FindAllString(fmt.Sprint(value), -1)...)
	}

	return keys
}
