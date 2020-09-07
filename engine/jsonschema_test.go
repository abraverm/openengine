package engine

import (
	"context"
	"github.com/qri-io/jsonpointer"
	"github.com/qri-io/jsonschema"
	"reflect"
	"testing"
)

func TestNewOeProperties(t *testing.T) {
	tests := []struct {
		name string
		want jsonschema.Keyword
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOeProperties(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOeProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOeRequired(t *testing.T) {
	tests := []struct {
		name string
		want jsonschema.Keyword
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOeRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOeProperties_JSONChildren(t *testing.T) {
	tests := []struct {
		name    string
		o       OeProperties
		wantRes map[string]jsonschema.JSONPather
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.o.JSONChildren(); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("JSONChildren() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestOeProperties_JSONProp(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		o    OeProperties
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.JSONProp(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONProp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOeProperties_Register(t *testing.T) {
	type args struct {
		uri      string
		registry *jsonschema.SchemaRegistry
	}
	tests := []struct {
		name string
		o    OeProperties
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestOeProperties_Resolve(t *testing.T) {
	type args struct {
		pointer jsonpointer.Pointer
		uri     string
	}
	tests := []struct {
		name string
		o    OeProperties
		args args
		want *jsonschema.Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Resolve(tt.args.pointer, tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOeProperties_ValidateKeyword(t *testing.T) {
	type args struct {
		ctx          context.Context
		currentState *jsonschema.ValidationState
		data         interface{}
	}
	tests := []struct {
		name string
		o    OeProperties
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_oeRequired_Register(t *testing.T) {
	type args struct {
		uri      string
		registry *jsonschema.SchemaRegistry
	}
	tests := []struct {
		name string
		f    oeRequired
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_oeRequired_Resolve(t *testing.T) {
	type args struct {
		pointer jsonpointer.Pointer
		uri     string
	}
	tests := []struct {
		name string
		f    oeRequired
		args args
		want *jsonschema.Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Resolve(tt.args.pointer, tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oeRequired_ValidateKeyword(t *testing.T) {
	type args struct {
		ctx          context.Context
		currentState *jsonschema.ValidationState
		data         interface{}
	}
	tests := []struct {
		name string
		f    oeRequired
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
