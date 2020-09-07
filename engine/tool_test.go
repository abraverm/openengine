package engine

import "testing"

func TestTool_Run(t1 *testing.T) {
	type fields struct {
		Name       string
		Parameters map[string]interface{}
		Script     string
	}
	type args struct {
		args map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tool{
				Name:       tt.fields.Name,
				Parameters: tt.fields.Parameters,
				Script:     tt.fields.Script,
			}
			got, err := t.Run(tt.args.args)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
