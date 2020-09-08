package engine

import (
	"reflect"
	"testing"
)

func TestImplicitTask_getImplicitKeys(t *testing.T) {
	type fields struct {
		Name   string
		Args   map[string]interface{}
		Type   string
		Store  string
		Action string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := ImplicitTask{
				Name:   tt.fields.Name,
				Args:   tt.fields.Args,
				Type:   tt.fields.Type,
				Store:  tt.fields.Store,
				Action: tt.fields.Action,
			}
			if got := k.getImplicitKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getImplicitKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImplicitTask_resolve(t1 *testing.T) {
	type fields struct {
		Name   string
		Args   map[string]interface{}
		Type   string
		Store  string
		Action string
	}
	type args struct {
		values map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := ImplicitTask{
				Name:   tt.fields.Name,
				Args:   tt.fields.Args,
				Type:   tt.fields.Type,
				Store:  tt.fields.Store,
				Action: tt.fields.Action,
			}
			if got := t.resolve(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProviderAPIResourcesParam_getImplicitKeys(t *testing.T) {
	type fields struct {
		Required bool
		Explicit Schema
		Implicit []ImplicitTask
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProviderAPIResourcesParam{
				Required: tt.fields.Required,
				Explicit: tt.fields.Explicit,
				Implicit: tt.fields.Implicit,
			}
			if got := p.getImplicitKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getImplicitKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProviderAPIResourcesParam_toJsonSchema(t *testing.T) {
	type fields struct {
		Required bool
		Explicit Schema
		Implicit []ImplicitTask
	}
	type args struct {
		resourceImplicit map[string]Schema
		name             string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProviderAPIResourcesParam{
				Required: tt.fields.Required,
				Explicit: tt.fields.Explicit,
				Implicit: tt.fields.Implicit,
			}
			if got := p.toJsonSchema(tt.args.resourceImplicit, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toJSONSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_getImplicitKeys(t *testing.T) {
	type fields struct {
		Match      Schema
		Implicit   map[string]Schema
		Resource   string
		Action     string
		Parameters ProviderParameters
		Debug      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Provider{
				Match:      tt.fields.Match,
				Implicit:   tt.fields.Implicit,
				Resource:   tt.fields.Resource,
				Action:     tt.fields.Action,
				Parameters: tt.fields.Parameters,
				Debug:      tt.fields.Debug,
			}
			if got := p.getImplicitKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getImplicitKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_toJsonSchema(t *testing.T) {
	type fields struct {
		Match      Schema
		Implicit   map[string]Schema
		Resource   string
		Action     string
		Parameters ProviderParameters
		Debug      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Provider{
				Match:      tt.fields.Match,
				Implicit:   tt.fields.Implicit,
				Resource:   tt.fields.Resource,
				Action:     tt.fields.Action,
				Parameters: tt.fields.Parameters,
				Debug:      tt.fields.Debug,
			}
			if got := p.toJSONSchema(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toJSONSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_toJsonSchemaDefs(t *testing.T) {
	type fields struct {
		Match      Schema
		Implicit   map[string]Schema
		Resource   string
		Action     string
		Parameters ProviderParameters
		Debug      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Provider{
				Match:      tt.fields.Match,
				Implicit:   tt.fields.Implicit,
				Resource:   tt.fields.Resource,
				Action:     tt.fields.Action,
				Parameters: tt.fields.Parameters,
				Debug:      tt.fields.Debug,
			}
			if got := p.toJSONSchemaDefs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toJSONSchemaDefs() = %v, want %v", got, tt.want)
			}
		})
	}
}
