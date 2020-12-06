package cli

import (
	"fmt"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
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
			err := Run(args)
			if fmt.Sprint(tt.wantErr) != fmt.Sprint(err) {
				t.Errorf("CLI test %v of '%v' expected '%v' but got '%v'", tt.name, tt.args, tt.wantErr, err)
			}
		})
	}
}

func TestMain(m *testing.M) {
	exitcode := m.Run()
	os.RemoveAll("oe.log")
	os.Exit(exitcode)
}
