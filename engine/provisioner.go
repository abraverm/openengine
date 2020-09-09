package engine

import "fmt"

// ProvisionerAPI is a list of resource types and their provisioners.
type ProvisionerAPI map[string]map[string][]Provisioner

// Provisioner is a provisioning script that supports specific provider, resource type and action upon it.
type Provisioner struct {
	Resource   string
	Parameters map[string]Schema
	Match      Schema
	Action     string
	Logic      string
	Debug      bool
	Required   []string
}

func (p Provisioner) toJSONSchema() Schema {
	properties := make(map[string]interface{})
	parameters := make(map[string]interface{})

	for param, def := range p.Parameters {
		parameters[param] = Schema{
			"oneOf": []Schema{
				def,
				{"$ref": fmt.Sprintf("%v#implicit", param)},
			},
		}
	}

	args := Schema{
		"oeProperties": parameters,
		//"oeRequired": p.Required,
		"type": "object",
	}
	properties["Resource"] = Schema{
		"type": "object",
		"properties": Schema{
			"Name": Schema{
				"const": p.Resource,
			},
			"args": args,
		},
	}
	properties["System"] = p.Match

	return Schema{
		"$id":                  "provisioner.json",
		"title":                "Provisioner",
		"type":                 "object",
		"required":             []string{"Resource", "System"},
		"properties":           properties,
		"additionalProperties": false,
	}
}
