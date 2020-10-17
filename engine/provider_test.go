package engine

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
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
		{
			name: "empty",
			fields: fields{
				Name:   "",
				Args:   nil,
				Type:   "",
				Store:  "",
				Action: "",
			},
			want: nil,
		},
		{
			name: "non_match",
			fields: fields{
				Name: "",
				Args: map[string]interface{}{
					"key": "value",
				},
				Type:   "",
				Store:  "",
				Action: "",
			},
			want: nil,
		},
		{
			name: "match_one",
			fields: fields{
				Name: "",
				Args: map[string]interface{}{
					"key": "$_value",
				},
				Type:   "",
				Store:  "",
				Action: "",
			},
			want: []string{"$_value"},
		},
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
		{
			name: "empty",
			fields: fields{
				Required: false,
				Explicit: nil,
				Implicit: nil,
			},
			want: nil,
		},
		{
			name: "empty_task",
			fields: fields{
				Required: false,
				Explicit: nil,
				Implicit: []ImplicitTask{
					{
						Name:   "",
						Args:   nil,
						Type:   "",
						Store:  "",
						Action: "",
					},
				},
			},
			want: nil,
		},
		{
			name: "one_key",
			fields: fields{
				Required: false,
				Explicit: nil,
				Implicit: []ImplicitTask{
					{
						Name: "",
						Args: map[string]interface{}{
							"key": "$_value",
						},
						Type:   "",
						Store:  "",
						Action: "",
					},
				},
			},
			want: []string{"$_value"},
		},
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
		{
			name: "empty",
			fields: fields{
				Required: false,
				Explicit: Schema{},
				Implicit: []ImplicitTask{},
			},
			args: args{
				resourceImplicit: map[string]Schema{},
				name:             "",
			},
			want: Schema{
				"$id": "",
			},
		},
		{
			name: "emptyRequired",
			fields: fields{
				Required: true,
				Explicit: Schema{},
				Implicit: []ImplicitTask{},
			},
			args: args{
				resourceImplicit: map[string]Schema{},
				name:             "",
			},
			want: Schema{
				"$id":        "",
				"oeRequired": true,
			},
		},
		{
			name: "emptyRequiredAndOne",
			fields: fields{
				Required: true,
				Explicit: Schema{},
				Implicit: []ImplicitTask{
					{
						Name:   "",
						Args:   nil,
						Type:   "",
						Store:  "",
						Action: "",
					},
				},
			},
			args: args{
				resourceImplicit: map[string]Schema{},
				name:             "",
			},
			want: Schema{
				"$id":        "",
				"oeRequired": true,
				"oneOf": []Schema{
					{},
					{
						"type":       "object",
						"$anchor":    "implicit",
						"properties": Schema{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProviderAPIResourcesParam{
				Required: tt.fields.Required,
				Explicit: tt.fields.Explicit,
				Implicit: tt.fields.Implicit,
			}
			got := p.toJSONSchema(tt.args.resourceImplicit, tt.args.name)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
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
		{
			name: "empty",
			fields: fields{
				Match:      nil,
				Implicit:   nil,
				Resource:   "",
				Action:     "",
				Parameters: nil,
				Debug:      false,
			},
			want: nil,
		},
		{
			name: "one_key",
			fields: fields{
				Match: nil,
				Implicit: map[string]Schema{
					"key": {},
				},
				Resource:   "",
				Action:     "",
				Parameters: nil,
				Debug:      false,
			},
			want: []string{"key"},
		},
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
		{
			name: "empty",
			fields: fields{
				Match:      Schema{},
				Implicit:   nil,
				Resource:   "",
				Action:     "",
				Parameters: nil,
				Debug:      false,
			},
			want: Schema{
				"title":    "Provider",
				"type":     "object",
				"required": []string{"Resource", "System"},
				"properties": map[string]interface{}{
					"Resource": Schema{
						"type": "object",
						"properties": Schema{
							"Name": Schema{
								"const": "",
							},
							"Type": Schema{
								"type": "string",
							},
							"args": Schema{
								"type":         "object",
								"oeProperties": Schema{},
							},
						},
						"additionalProperties": false,
					},
					"System": Schema{},
				},
			},
		},
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
			got := p.toJSONSchema()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
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
		{
			name: "empty",
			fields: fields{
				Match:      nil,
				Implicit:   nil,
				Resource:   "",
				Action:     "",
				Parameters: ProviderParameters{},
				Debug:      false,
			},
			want: Schema{},
		},
		{
			name: "one",
			fields: fields{
				Match:    nil,
				Implicit: nil,
				Resource: "",
				Action:   "",
				Parameters: ProviderParameters{
					"key": {
						Required: false,
						Explicit: Schema{},
						Implicit: []ImplicitTask{},
					},
				},
				Debug: false,
			},
			want: Schema{
				"key": Schema{
					"$id": "key",
				},
			},
		},
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
			got := p.toJSONSchemaDefs()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
