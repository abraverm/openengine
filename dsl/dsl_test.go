package dsl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"openengine/engine"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/imdario/mergo"
	"github.com/nsf/jsondiff"
)

func Test_getSource(t *testing.T) {
	srv := HTTPMock("/empty", http.StatusOK, "")
	defer srv.Close()
	tests := []struct {
		name    string
		uri     string
		want    []byte
		wantErr bool
	}{
		{"empty", "", nil, true},
		{"local file", "testdata/empty", []byte{}, false},
		{"unreachable remote file", "http://asdfasddf.com:9999/file.yaml", nil, true},
		{"unreachable remote file", fmt.Sprintf("%v/empty", srv.URL), []byte{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSource(tt.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSource() error = %v, wantErr %v", err, tt.wantErr)
				return

			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSource() got = %v, want %v", string(got), string(tt.want))

			}

		})

	}

}

func HTTPMock(pattern string, statusCode int, response interface{}) *httptest.Server {
	c := &ctrl{statusCode, response}

	handler := http.NewServeMux()
	handler.HandleFunc(pattern, c.mockHandler)

	return httptest.NewServer(handler)

}

type ctrl struct {
	statusCode int
	response   interface{}
}

func (c *ctrl) mockHandler(w http.ResponseWriter, r *http.Request) {
	resp := []byte{}

	rt := reflect.TypeOf(c.response)
	if rt.Kind() == reflect.String {
		resp = []byte(c.response.(string))

	} else if rt.Kind() == reflect.Struct || rt.Kind() == reflect.Ptr {
		resp, _ = json.Marshal(c.response)

	} else {
		resp = []byte("{}")

	}

	w.WriteHeader(c.statusCode)
	w.Write(resp)

}

type DSLTestSet struct {
	Description string
	DSL         DSL
	Tests       []struct {
		Description string
		Action      string
		Args        []interface{}
		Want        []interface{}
	}
}

func TestDSL(t *testing.T) {
	tests := 0
	passed := 0
	data, err := ioutil.ReadFile("testdata/dsl.json")
	if err != nil {
		t.Errorf("error loading test file: %s", err.Error())
		return

	}
	testSets := []*DSLTestSet{}
	if err := json.Unmarshal(data, &testSets); err != nil {
		t.Errorf("error unmarshaling test set from JSON: %s", err.Error())
		return

	}
	opt := jsondiff.DefaultConsoleOptions()
	t.Run("DSL", func(t *testing.T) {
		for _, ts := range testSets {
			for i, c := range ts.Tests {
				tests++
				d := ts.DSL
				var args []reflect.Value
				switch c.Action {
				case "CreateEngine":
					want := DSL{}
					wJSON, _ := json.Marshal(c.Want[0])
					tmpDSL := ts.DSL
					json.Unmarshal(wJSON, &want)
					if err := mergo.Merge(&tmpDSL, want, mergo.WithAppendSlice); err != nil {
						t.Errorf("%s test case %d: %s. error: %s", ts.Description, i, c.Description, err)

					}
					c.Want[0] = tmpDSL
				case "Run":
					d.CreateEngine()
					err := d.Run(fmt.Sprint(c.Args[0]))
					if fmt.Sprint(c.Want[0]) != fmt.Sprint(err) {
						t.Errorf("%s test case %d: %s. error: %s", ts.Description, i, c.Description, err)

					} else {
						passed++

					}
					continue
				case "GetSolutions":
					want := []engine.Solution{}
					wJSON, _ := json.Marshal(c.Want[0])
					json.Unmarshal(wJSON, &want)
					d.CreateEngine()
					solutions := d.GetSolutions()
					eJSON, _ := json.Marshal(solutions)
					if result, diff := jsondiff.Compare(eJSON, wJSON, &opt); result.String() != "FullMatch" {
						t.Errorf("%s test case %d: %s. error: %s", ts.Description, i, c.Description, diff)

					} else {
						passed++

					}
					continue

				}
				dsl := reflect.ValueOf(&d)
				method := dsl.MethodByName(c.Action)
				if !method.IsValid() {
					t.Errorf("%s test case %d: %s. error: Unknown method or private '%s'", ts.Description, i, c.Description, c.Action)
					continue

				}
				results := method.Call(args)
				if len(results) == 0 {
					eJSON, _ := json.Marshal(d)
					wJSON, _ := json.Marshal(c.Want[0])
					if result, diff := jsondiff.Compare(eJSON, wJSON, &opt); result.String() != "FullMatch" {
						t.Errorf("%s test case %d: %s. error: %s", ts.Description, i, c.Description, diff)

					} else {
						passed++

					}

				}

			}

		}

	})

}

func TestDSLCreation(t *testing.T) {

	var dsl DSL
	yamlFile, _ := ioutil.ReadFile(filepath.Clean("testdata/bdsl.yaml"))

	yaml.UnmarshalWithOptions(yamlFile, &dsl, yaml.Strict())

	dsl.CreateEngine()

}
