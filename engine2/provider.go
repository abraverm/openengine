package engine2

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type ProviderAPI map[string]ProviderAPIResources

type ProviderAPIResources struct {
	Implicit map[string]JSONSchema
	Providers []Provider
}

type Provider struct {
	Match JSONSchema
	Implicit map[string]JSONSchema
	Resource string
	Action string
	Parameters ProviderParameters
}

type ProviderParameters map[string]ProviderAPIResourcesParam

type ProviderAPIResourcesParam struct {
	Required bool
	Explicit JSONSchema
	Implicit []ImplicitTask
}

type ImplicitTask struct {
	Resource
	Store  string
	Action string
}

func (t ImplicitTask) resolve(values map[string]interface{}) map[string]interface{} {
	var args = make(map[string]interface{})
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
		args[paramName]	= result
	}
	return args
}

func (p Provider) toJsonSchema() JSONSchema {
	var required []string
	properties := make(map[string]interface{})
	argsSchema := make(JSONSchema)
	for param, def := range p.Parameters {
		if def.Required {
			required = append(required, param)
		}
		argsSchema[param] = def.toJsonSchema(p.Implicit)
	}
	properties["resource"] = JSONSchema{
		"type": "object",
		"properties": JSONSchema{
			"resource": JSONSchema{
				"type": "string",
				"enum": []string{p.Resource},
			},
			"args": argsSchema,
		},
	}
	properties["system"] = p.Match
	return JSONSchema{
		"type":     "object",
		"required": required,
		"properties": properties,
	}
}

func (p ProviderAPIResourcesParam) toJsonSchema(resourceImplicit map[string]JSONSchema) JSONSchema {
	if len(p.Implicit) > 0 {
		implicitProperties := make(JSONSchema)
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
		return JSONSchema{
			"type": "object",
			"properties": JSONSchema{
				"oneOf": []JSONSchema{
					p.Explicit,
					{
						"type": "object",
						"properties": implicitProperties,
						"required": required,
					},
				},
			},
		}
	}
	return p.Explicit
}

