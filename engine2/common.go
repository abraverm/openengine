package engine2

type JSONSchema map[string]interface{}


type System map[string]interface{}

type Resource struct {
	Name string `json:"resource"`
	Type string
	Args map[string]interface{}
}

type Schedule struct{
	resource Resource
	solutions []Solution
}