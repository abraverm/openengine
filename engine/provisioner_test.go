package engine

import (
	"reflect"
	"testing"
)

func TestProvisioner_toJsonSchema(t *testing.T) {
	type fields struct {
		Resource   string
		Parameters map[string]Schema
		Match      Schema
		Action     string
		Logic      string
		Debug      bool
		Required   []string
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
			p := Provisioner{
				Resource:   tt.fields.Resource,
				Parameters: tt.fields.Parameters,
				Match:      tt.fields.Match,
				Action:     tt.fields.Action,
				Logic:      tt.fields.Logic,
				Debug:      tt.fields.Debug,
				Required:   tt.fields.Required,
			}
			if got := p.toJsonSchema(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toJsonSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
