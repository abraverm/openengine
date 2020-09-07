package engine

import (
	"fmt"
	"regexp"
)

type Schema map[string]interface{}


type System map[string]interface{}


type Resource struct {
	Name string `yaml:"resource"`
	Args map[string]interface{} `json:"args"`
}

type Schedule struct{
	resource Resource
	solutions []Solution
}

func (r Resource) getImplicitKeys() []string {
	re := regexp.MustCompile(`\$_[[:alpha:]]*`)
	var keys []string
	for key, value := range r.Args{
		if re.MatchString(key) {
			keys = append(keys, key)
		}
		keys = append(keys, re.FindAllString(fmt.Sprint(value), -1)...)
	}
	return keys
}