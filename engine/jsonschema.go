package engine

import (
	"context"
	"regexp"

	jptr "github.com/qri-io/jsonpointer"
	"github.com/qri-io/jsonschema"
)

// OeRequired is a customize customized required keyword used with OeProperties to support implicit values.
// TODO: I don't think oeRequired is really being used (ValidateKeyword function is empty)
// 	or maybe I'm missing something.
type oeRequired bool

// NewOeRequired creates an instance of OeRequired.
func NewOeRequired() jsonschema.Keyword {
	return new(oeRequired)
}

// Register is an empty function, just as the original required keyword implementation, required by the interface.
func (f *oeRequired) Register(uri string, registry *jsonschema.SchemaRegistry) {}

// Resolve is an empty function, just as the original required keyword implementation, required by the interface.
func (f *oeRequired) Resolve(pointer jptr.Pointer, uri string) *jsonschema.Schema {
	return nil
}

// ValidateKeyword is an empty function, just as the original required keyword implementation,
//	required by the interface.
func (f *oeRequired) ValidateKeyword(ctx context.Context, currentState *jsonschema.ValidationState, data interface{}) {
}

// OeProperties is a customized properties keyword to support implicit values.
type OeProperties map[string]*jsonschema.Schema

// NewOeProperties creates an instance of OeProperties.
func NewOeProperties() jsonschema.Keyword {
	return new(OeProperties)
}

// ValidateKeyword OeProperties is what makes it different from the original properties, it will use implicit parameters
// as a alternative validation method for given keyword.
// nolint: nestif
// TODO: function too comlicated (nested if).
func (o OeProperties) ValidateKeyword(ctx context.Context, currentState *jsonschema.ValidationState, data interface{}) {
	if obj, ok := data.(map[string]interface{}); ok {
		subState := currentState.NewSubState()
		re := regexp.MustCompile(`^_[[:alpha:]]*`)
		implicitParams := make(map[string]interface{})

		for key := range obj {
			if re.MatchString(key) {
				if _, ok := obj[key]; ok {
					implicitParams[key] = obj[key]
				}
			}
		}

		for key := range o {
			if re.MatchString(key) {
				continue
			}

			currentState.SetEvaluatedKey(key)
			subState.ClearState()
			subState.DescendBaseFromState(currentState, "oeProperties", key)
			subState.DescendRelativeFromState(currentState, "oeProperties", key)
			subState.DescendInstanceFromState(currentState, key)

			if _, ok := obj[key]; ok { // Explicit case
				o[key].ValidateKeyword(ctx, subState, obj[key])
			} else if len(implicitParams) > 0 { // Implicit case
				o[key].ValidateKeyword(ctx, subState, implicitParams)
			}

			if o[key].HasKeyword("oeRequired") {
				currentState.AddSubErrors(*subState.Errs...)
			}

			if subState.IsValid() {
				currentState.UpdateEvaluatedPropsAndItems(subState)
			}
		}
	}
}

// Register 'OeProperties' keyword is just like resolving the original 'properties' keyword.
func (o *OeProperties) Register(uri string, registry *jsonschema.SchemaRegistry) {
	for _, v := range *o {
		v.Register(uri, registry)
	}
}

// Resolve 'OeProperties' keyword is just like resolving the original 'properties' keyword.
func (o *OeProperties) Resolve(pointer jptr.Pointer, uri string) *jsonschema.Schema {
	if pointer == nil {
		return nil
	}

	current := pointer.Head()
	if current == nil {
		return nil
	}

	if schema, ok := (*o)[*current]; ok {
		return schema.Resolve(pointer.Tail(), uri)
	}

	return nil
}

// JSONProp implements the JSONPather for Properties.
func (o OeProperties) JSONProp(name string) interface{} {
	return o[name]
}

// JSONChildren implements the JSONContainer interface for Properties.
func (o OeProperties) JSONChildren() (res map[string]jsonschema.JSONPather) {
	res = map[string]jsonschema.JSONPather{}
	for key, sch := range o {
		res[key] = sch
	}

	return
}
