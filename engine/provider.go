package engine

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// ProviderAPI is a list of resource types and their matching providers.
type ProviderAPI map[string]ProviderAPIResources

// ProviderAPIResources is a list of providers with common implicit parameters.
type ProviderAPIResources struct {
	Implicit  map[string]Schema `json:"implicit"`
	Providers []Provider        `json:"providers"`
}

// Provider is an parameter API of a matching provider for a specific action on resource.
type Provider struct {
	Match      Schema             `json:"match"`
	Implicit   map[string]Schema  `json:"implicit"`
	Resource   string             `json:"resource"`
	Action     string             `json:"action"`
	Parameters ProviderParameters `json:"parameters"`
	Debug      bool               `json:"debug"`
}

// ProviderParameters is a list of provider parameters and their conditions.
type ProviderParameters map[string]ProviderAPIResourcesParam

// A ProviderAPIResourcesParam is a set of conditions for the parameter to match.
// The explicit conditions defined by the provider original API and implicit to explicit process that produces
// a fitting explicit value using implicit (other) parameters.
type ProviderAPIResourcesParam struct {
	Required bool           `json:"required"`
	Explicit Schema         `json:"explicit"`
	Implicit []ImplicitTask `json:"implicit"`
}

// ImplicitTask is a task part of the implicit to explicit procedure of explicit parameter value resolution.
type ImplicitTask struct {
	Name   string                 `yaml:"resource" json:"resource"`
	Args   map[string]interface{} `json:"args"`
	Type   string                 `json:"type"`
	Store  string                 `json:"store"`
	Action string                 `json:"action"`
}

func (t ImplicitTask) resolve(values map[string]interface{}) map[string]interface{} {
	args := make(map[string]interface{})

	for paramName, paramValue := range t.Args {
		result := fmt.Sprint(paramValue)
		for resolvedParam, resolvedValue := range values {
			result = strings.Replace(
				result,
				fmt.Sprintf("$%v", resolvedParam),
				fmt.Sprint(resolvedValue),
				-1,
			)
		}

		args[paramName] = result
	}

	return args
}

func (t ImplicitTask) getImplicitKeys() []string {
	re := regexp.MustCompile(`\$_[[:alpha:]]*`)

	// nolint: prealloc
	var keys []string

	for _, value := range t.Args {
		keys = append(keys, re.FindAllString(fmt.Sprint(value), -1)...)
	}

	return keys
}

func (p Provider) toJSONSchemaDefs() Schema {
	argsSchema := make(Schema)
	for param, def := range p.Parameters {
		argsSchema[param] = def.toJSONSchema(p.Implicit, param)
	}

	return argsSchema
}

func (p Provider) toJSONSchema() Schema {
	// var required []string
	properties := make(map[string]interface{})
	argsSchema := make(Schema)
	for param := range p.Parameters {
		argsSchema[param] = Schema{"$ref": fmt.Sprintf("%v", param)}
	}

	properties["Resource"] = Schema{
		"type": "object",
		"properties": Schema{
			"Name": Schema{
				"const": p.Resource,
			},
			"Type": Schema{
				"type": "string",
			},
			"args": Schema{
				"type":         "object",
				"oeProperties": argsSchema,
			},
		},
		"additionalProperties": false,
	}
	properties["System"] = p.Match

	return Schema{
		//"$id": "provider.json",
		"title":      "Provider",
		"type":       "object",
		"required":   []string{"Resource", "System"},
		"properties": properties,
	}
}

func (p ProviderAPIResourcesParam) toJSONSchema(resourceImplicit map[string]Schema, name string) Schema {
	if len(p.Implicit) > 0 {
		implicitProperties := make(Schema)

		var implicitArgs []string

		for _, implicitTask := range p.Implicit {
			for _, arg := range implicitTask.Args {
				re := regexp.MustCompile(`\$_[[:alpha:]]*`)
				for _, match := range re.FindAll([]byte(fmt.Sprint(arg)), -1) {
					implicitArgs = append(implicitArgs, string(match[1:]))
				}
			}
		}

		var required []string

		for param, def := range resourceImplicit {
			if sort.SearchStrings(implicitArgs, param) != len(implicitArgs) {
				implicitProperties[param] = def

				required = append(required, param)
			}
		}

		implicit := Schema{
			"type":       "object",
			"$anchor":    "implicit",
			"properties": implicitProperties,
		}

		if len(required) > 0 {
			implicit["required"] = required
		}

		result := Schema{
			"$id": fmt.Sprintf("%v", name),
			"oneOf": []Schema{
				p.Explicit,
				implicit,
			},
		}

		if p.Required {
			result["oeRequired"] = true
		}

		return result
	}

	p.Explicit["$id"] = fmt.Sprintf("%v", name)

	if p.Required {
		p.Explicit["oeRequired"] = true
	}

	return p.Explicit
}

func (p Provider) getImplicitKeys() []string {
	// nolint: prealloc
	var keys []string
	for key := range p.Implicit {
		keys = append(keys, key)
	}

	return keys
}

func (p ProviderAPIResourcesParam) getImplicitKeys() []string {
	// nolint: prealloc
	var keys []string
	for _, task := range p.Implicit {
		keys = append(keys, task.getImplicitKeys()...)
	}

	return keys
}
