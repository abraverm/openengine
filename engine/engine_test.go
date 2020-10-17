package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/imdario/mergo"
	"github.com/nsf/jsondiff"
)

type EngineTestSet struct {
	Description string           `json:"descrption"`
	Engine      Engine           `json:"engine"`
	Tests       []EngineTestCase `json:"tests"`
}

type EngineTestCase struct {
	Description string        `json:"description"`
	Action      string        `json:"action"`
	Args        []interface{} `json:"args"`
	Want        []interface{} `json:"want"`
}

func runEngineTests(t *testing.T, testFilePaths []string) {
	tests := 0
	passed := 0
	for _, path := range testFilePaths {
		t.Run(path, func(t *testing.T) {
			base := filepath.Base(path)
			testSets := []*EngineTestSet{}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				t.Errorf("error loading test file: %s", err.Error())
				return
			}

			if err := json.Unmarshal(data, &testSets); err != nil {
				t.Errorf("error unmarshaling test set %s from JSON: %s", base, err.Error())
				return
			}

			for _, ts := range testSets {
				for i, c := range ts.Tests {
					e := ts.Engine
					tests++
					var args []reflect.Value
					switch c.Action {
					case "AddProvider":
						provider := &ProviderAPI{}
						tJSON, _ := json.Marshal(c.Args[0])
						json.Unmarshal(tJSON, provider)
						args = append(args, reflect.ValueOf(*provider))
						want := Engine{}
						wJSON, _ := json.Marshal(c.Want[0])
						json.Unmarshal(wJSON, &want)
						tmpEngine := e
						if err := mergo.Merge(&tmpEngine, want, mergo.WithAppendSlice); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, err)
						}
						c.Want[0] = tmpEngine
					case "AddProvisioner":
						provisioner := &Provisioner{}
						tJSON, _ := json.Marshal(c.Args[0])
						json.Unmarshal(tJSON, provisioner)
						args = append(args, reflect.ValueOf(*provisioner))
						want := Engine{}
						wJSON, _ := json.Marshal(c.Want[0])
						tmpEngine := e
						json.Unmarshal(wJSON, &want)
						if err := mergo.Merge(&tmpEngine, want, mergo.WithAppendSlice); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, err)
						}
						c.Want[0] = tmpEngine
					case "AddSystem":
						system := &System{}
						tJSON, _ := json.Marshal(c.Args[0])
						json.Unmarshal(tJSON, system)
						args = append(args, reflect.ValueOf(*system))
						want := Engine{}
						wJSON, _ := json.Marshal(c.Want[0])
						tmpEngine := e
						json.Unmarshal(wJSON, &want)
						if err := mergo.Merge(&tmpEngine, want, mergo.WithAppendSlice); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, err)
						}
						c.Want[0] = tmpEngine
					case "AddResource":
						resource := &Resource{}
						tJSON, _ := json.Marshal(c.Args[0])
						json.Unmarshal(tJSON, resource)
						args = append(args, reflect.ValueOf(*resource))
						want := Engine{}
						wJSON, _ := json.Marshal(c.Want[0])
						tmpEngine := e
						json.Unmarshal(wJSON, &want)
						if err := mergo.Merge(&tmpEngine, want, mergo.WithAppendSlice); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, err)
						}
						c.Want[0] = tmpEngine
					case "AddTool":
						tool := &ToolAPI{}
						tJSON, _ := json.Marshal(c.Args[0])
						json.Unmarshal(tJSON, tool)
						args = append(args, reflect.ValueOf(*tool))
						want := Engine{}
						wJSON, _ := json.Marshal(c.Want[0])
						tmpEngine := e
						json.Unmarshal(wJSON, &want)
						if err := mergo.Merge(&tmpEngine, want, mergo.WithAppendSlice); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, err)
						}
						c.Want[0] = tmpEngine
					case "Schedule":
						args = append(args, reflect.ValueOf(c.Args[0]))
					case "getTool":
						implicit := &ImplicitTask{}
						iJSON, _ := json.Marshal(c.Args[0])
						json.Unmarshal(iJSON, implicit)
						tool, err := e.getTool(*implicit)
						if c.Want[1] != fmt.Sprint(err) {
							t.Errorf("%s: %s test case %d: %s. error: expected %s, got %s", base, ts.Description, i, c.Description, c.Want[1], err)
						}
						opt := jsondiff.DefaultConsoleOptions()
						tJSON, _ := json.Marshal(tool)
						wJSON, _ := json.Marshal(c.Want[0])
						if result, diff := jsondiff.Compare(wJSON, tJSON, &opt); result.String() != "FullMatch" {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, diff)
						} else {
							passed++
						}
						continue
					case "matchProvidersProvisioners":
						resource := &Resource{}
						system := &System{}
						rJSON, _ := json.Marshal(c.Args[0])
						sJSON, _ := json.Marshal(c.Args[1])
						json.Unmarshal(rJSON, resource)
						json.Unmarshal(sJSON, system)
						solutions, err := e.matchProvidersProvisioners(*resource, *system)
						if fmt.Sprint(c.Want[1]) != fmt.Sprint(err) {
							t.Errorf("%s: %s test case %d: %s. error: expected %s, got %s", base, ts.Description, i, c.Description, c.Want[1], err)
						}
						opt := jsondiff.DefaultConsoleOptions()
						tJSON, _ := json.Marshal(solutions)
						wJSON, _ := json.Marshal(c.Want[0])
						if result, diff := jsondiff.Compare(wJSON, tJSON, &opt); result.String() != "FullMatch" {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, diff)
						} else {
							passed++
						}
						continue
					case "resolveDependencies":
						solution := &Solution{}
						sJSON, _ := json.Marshal(c.Args[0])
						json.Unmarshal(sJSON, solution)
						resolved := e.resolveDependencies(*solution)
						opt := jsondiff.DefaultConsoleOptions()
						tJSON, _ := json.Marshal(resolved)
						wJSON, _ := json.Marshal(c.Want[0])
						if result, diff := jsondiff.Compare(wJSON, tJSON, &opt); result.String() != "FullMatch" {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, diff)
						} else {
							passed++
						}
						continue
					case "Resolve":
						want := Engine{}
						wJSON, _ := json.Marshal(c.Want[0])
						tmpEngine := e
						json.Unmarshal(wJSON, &want)
						if err := mergo.Merge(&tmpEngine, want, mergo.WithAppendSlice); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, err)
						}
						c.Want[0] = tmpEngine
					}
					engine := reflect.ValueOf(&e)
					method := engine.MethodByName(c.Action)
					if !method.IsValid() {
						t.Errorf("%s: %s test case %d: %s. error: Unknown method or private '%s'", base, ts.Description, i, c.Description, c.Action)
						continue
					}
					results := method.Call(args)
					if len(results) == 0 {
						opt := jsondiff.DefaultConsoleOptions()
						eJSON, _ := json.Marshal(e)
						wJSON, _ := json.Marshal(c.Want[0])
						if result, diff := jsondiff.Compare(eJSON, wJSON, &opt); result.String() != "FullMatch" {
							t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, diff)
						} else {
							passed++
						}
					}
				}
			}
		})
	}
	t.Logf("%d/%d tests passed", passed, tests)
}

func TestEngine(t *testing.T) {
	runEngineTests(t, []string{
		"testdata/engine.empty.json",
		"testdata/engine.implicitTool.json",
		"testdata/engine.implicitResource.json",
	})
}

func TestNewEngine(t *testing.T) {
	tests := []struct {
		name string
		want *Engine
	}{
		{
			name: "empty",
			want: &Engine{
				Systems:      nil,
				Resources:    nil,
				Providers:    nil,
				Provisioners: nil,
				Solutions:    nil,
				Tools:        nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
