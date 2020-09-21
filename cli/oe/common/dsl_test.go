package common

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/abraverm/openengine/engine"
)

func TestDSL_CreateEngine(t *testing.T) {
	type fields struct {
		API          []string
		Provisioners []string
		Systems      []engine.System
		Tools        []string
		Resources    []engine.Resource
		Engine       engine.Engine
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DSL{
				API:          tt.fields.API,
				Provisioners: tt.fields.Provisioners,
				Systems:      tt.fields.Systems,
				Tools:        tt.fields.Tools,
				Resources:    tt.fields.Resources,
				Engine:       tt.fields.Engine,
			}
			fmt.Print(d)
		})
	}
}

func TestDSL_GetSolutions(t *testing.T) {
	type fields struct {
		API          []string
		Provisioners []string
		Systems      []engine.System
		Tools        []string
		Resources    []engine.Resource
		Engine       engine.Engine
	}
	tests := []struct {
		name   string
		fields fields
		want   []engine.Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DSL{
				API:          tt.fields.API,
				Provisioners: tt.fields.Provisioners,
				Systems:      tt.fields.Systems,
				Tools:        tt.fields.Tools,
				Resources:    tt.fields.Resources,
				Engine:       tt.fields.Engine,
			}
			if got := d.GetSolutions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSolutions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDSL_Run(t *testing.T) {
	type fields struct {
		API          []string
		Provisioners []string
		Systems      []engine.System
		Tools        []string
		Resources    []engine.Resource
		Engine       engine.Engine
	}
	type args struct {
		action string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DSL{
				API:          tt.fields.API,
				Provisioners: tt.fields.Provisioners,
				Systems:      tt.fields.Systems,
				Tools:        tt.fields.Tools,
				Resources:    tt.fields.Resources,
				Engine:       tt.fields.Engine,
			}
			if err := d.Run(tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExists(tt.args.filename); got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSource(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSource(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSource() got = %v, want %v", got, tt.want)
			}
		})
	}
}
