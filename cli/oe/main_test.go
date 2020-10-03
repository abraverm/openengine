package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	args := []string{} //os.Args[0:1] // Name of the program.
	tests := []struct {
		name    string
		args    []string
		wantErr error
	}{
		{"help", []string{"help"}, nil},
		{"deploy", []string{"-n", "deploy", "testdata/empty"}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args = append(args, tt.args...) // Append a flag
			err := run(args)
			if fmt.Sprint(tt.wantErr) != fmt.Sprint(err) {
				t.Errorf("CLI test %v of '%v' expected '%v' but got '%v'", tt.name, tt.args, tt.wantErr, err)
			}
		})
	}
}

func Test_deploy(t *testing.T) {
	tests := []struct {
		name    string
		dsl     string
		wantErr error
	}{
		{"empty", "testdata/empty", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := deploy(tt.dsl, true)
			if fmt.Sprint(err) != fmt.Sprint(tt.wantErr) {
				t.Errorf("CLI test %v of '%v' expected '%v' but got '%v'", tt.name, tt.dsl, tt.wantErr, err)
			}
		})
	}
}
