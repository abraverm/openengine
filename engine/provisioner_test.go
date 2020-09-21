package engine

import (
	"testing"

	"github.com/go-test/deep"
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
		{
			name: "empty",
			fields: fields{
				Resource:   "empty",
				Parameters: nil,
				Match:      Schema{},
				Action:     "",
				Logic:      "",
				Debug:      false,
				Required:   nil,
			},
			want: Schema{
				"title":                "Provisioner",
				"type":                 "object",
				"$id":                  "provisioner.json",
				"additionalProperties": false,
				"properties": map[string]interface{}{
					"Resource": Schema{
						"type": "object",
						"properties": Schema{
							"Name": Schema{
								"const": "empty",
							},
							"args": Schema{
								"oeProperties": map[string]interface{}{},
								"type":         "object",
							},
						},
					},
					"System": Schema{},
				},
				"required": []string{"Resource", "System"},
			},
		},
		{
			name: "simple",
			fields: fields{
				Resource: "empty",
				Parameters: map[string]Schema{
					"key": {"type": "string"},
				},
				Match:    Schema{},
				Action:   "",
				Logic:    "",
				Debug:    false,
				Required: nil,
			},
			want: Schema{
				"title":                "Provisioner",
				"type":                 "object",
				"$id":                  "provisioner.json",
				"additionalProperties": false,
				"properties": map[string]interface{}{
					"Resource": Schema{
						"type": "object",
						"properties": Schema{
							"Name": Schema{
								"const": "empty",
							},
							"args": Schema{
								"oeProperties": map[string]interface{}{
									"key": Schema{
										"oneOf": []Schema{
											{"type": "string"},
											{"$ref": "key#implicit"},
										},
									},
								},
								"type": "object",
							},
						},
					},
					"System": Schema{},
				},
				"required": []string{"Resource", "System"},
			},
		},
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
			got := p.toJSONSchema()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
