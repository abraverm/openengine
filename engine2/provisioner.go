package engine2

type Provisioner struct {
	Resource string
	Parameters JSONSchema
	Match JSONSchema
	Action string
	Logic string
}

func (p Provisioner) toJsonSchema() JSONSchema {
	// TODO
	return JSONSchema{}
}
