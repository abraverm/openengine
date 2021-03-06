package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func Test_create(t *testing.T) {
	tests := []struct {
		name    string
		dsl     string
		wantErr error
	}{
		{"empty", "testdata/empty", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := op(tt.dsl, true, "create")
			if fmt.Sprint(err) != fmt.Sprint(tt.wantErr) {
				t.Errorf("CLI test %v of '%v' expected '%v' but got '%v'", tt.name, tt.dsl, tt.wantErr, err)
			}
		})
	}
}

func Test_Examples(t *testing.T) {
	// dmp := diffmatchpatch.New()

	type step struct {
		action string
		result string
	}

	tests := []struct {
		name  string
		dsl   string
		steps []step
	}{
		{"empty", "empty.yaml", []step{{"create", "empty_create.log"}}},
		{"new", "new.yaml", []step{{"create", "empty_create.log"}}},
		{"explicit generic", "getting_started_generic.yaml", []step{{"create", "getting_started_generic.log"}}},
		{"explicit aws", "getting_started_aws.yaml", []step{{"create", "getting_started_aws.log"}}},
		{"explicit openstack", "getting_started_openstack.yaml", []step{{"create", "getting_started_openstack.log"}}},
		{"explicit beaker", "getting_started_beaker.yaml", []step{{"create", "getting_started_beaker.log"}}},
		{"implicit generic", "implicit_generic.yaml", []step{{"create", "implicit_generic.log"}}},
		//		{"implicit aws os bkr", "implicit_aws_os_bkr.yaml", []step{{"create", "implicit_aws_os_bkr.log"}}}, // TODO: Fix memory leak
		{"constrains generic", "constrains_generic.yaml", []step{{"create", "constrains_generic.log"}}},
		{"constrains os resize", "constrains_os_resize.yaml", []step{{"update", "constrains_os_resize.log"}}},
		{"explicit dependency generic", "dependencies_generic_explicit.yaml", []step{{"create", "dependencies_generic_explicit.log"}}},
		{"implicit dependency generic", "dependencies_generic_implicit.yaml", []step{{"create", "dependencies_generic_implicit.log"}}},
		{"dependencies os disk", "dependencies_os_disk.yaml", []step{{"create", "dependencies_os_disk.log"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.steps {
				logfile, err := ioutil.TempFile("/tmp", "oe.*.log")
				if err != nil {
					t.Fatal(err)
				}
				defer os.Remove(logfile.Name())
				command := []string{"oe", "--noop", "--log", logfile.Name(), s.action, fmt.Sprintf("examples/%s", tt.dsl)}
				t.Log(command)
				if e := run(command); e != nil {
					t.Fatalf("Example %s failed on %s: %v", tt.name, s.action, e)
				}
				actual := readFile(logfile.Name())
				expected := readFile(fmt.Sprintf("examples/results/%s", s.result))

				if actual != expected {
					edits := myers.ComputeEdits(span.URIFromPath(logfile.Name()), actual, expected)
					diffs := fmt.Sprint(gotextdiff.ToUnified(logfile.Name(), fmt.Sprintf("examples/results/%s", s.result), actual, edits))

					// diffs := dmp.DiffMain(expected, actual, true)
					t.Errorf("Example %s has different result than what was expected:\n%s", tt.name, diffs)
				}
			}
		})
	}
}

func readFile(path string) string {
	data, _ := ioutil.ReadFile(filepath.Clean(path))
	return string(data)
}
