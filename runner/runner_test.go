package runner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"openengine/engine"

	"github.com/imdario/mergo"

	"github.com/nsf/jsondiff"
)

type TestSet struct {
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Subject     interface{} `json:"subject"`
	Tests       []TestCase  `json:"tests"`
}

type TestCase struct {
	Description string    `json:"description"`
	Function    string    `json:"function"`
	Args        []TestArg `json:"args"`
	Want        []TestArg `json:"want"`
}

type TestArg struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
	URI   string      `json:"uri"`
}

type MockRunner struct {
	Solutions []engine.Solution `json:"solutions"`
}

func (m MockRunner) ResourceRun(solution engine.Solution, args map[string]interface{}) (string, error) {
	return "success", nil
}

func (m MockRunner) ToolRun(tool engine.Tool, args map[string]interface{}) (string, error) {
	return "success", nil
}

func (m MockRunner) Schedules() []Schedule {
	return []Schedule{{m.Solutions}}
}

func updateWant(mergeValue, mergeType, wantValue, wantType []interface{}) (interface{}, error) {
	tJSON, _ := json.Marshal(mergeValue)
	json.Unmarshal(tJSON, mergeType)
	wJSON, _ := json.Marshal(wantValue)
	json.Unmarshal(wJSON, wantType)
	if err := mergo.Merge(&wantType, mergeType, mergo.WithAppendSlice); err != nil {
		return nil, err
	}
	return wantType, nil
}

func (a TestArg) data() ([]byte, error) {
	if a.URI != "" {
		data, err := ioutil.ReadFile(a.URI)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		data, err := json.Marshal(a.Value)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

func testMethod(funcMap map[string]interface{}, test TestCase) error {
	f := reflect.ValueOf(funcMap[test.Function])
	if len(test.Args) != f.Type().NumIn() {
		return errors.New("The number of params is not adapted.")
	}
	in, err := castArg(test.Args)
	if err != nil {
		return err
	}

	out, err := castArg(test.Want)
	if err != nil {
		return err
	}

	results := f.Call(in)
	for i, result := range results {
		if result.CanInterface() && result.Kind() == reflect.ValueOf(map[string]interface{}{}).Kind() {
			opt := jsondiff.DefaultConsoleOptions()
			eJSON, _ := json.Marshal(result.Interface())
			wJSON, _ := json.Marshal(out[i].Interface())
			if result, diff := jsondiff.Compare(eJSON, wJSON, &opt); result.String() != "FullMatch" {
				return fmt.Errorf(diff)
			}
		} else if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", out[i]) {
			return errors.New(fmt.Sprintf("got %v, expected %v", result, out[i]))
		}
	}
	return nil
}

func toValue(data []byte, resource reflect.Value) (reflect.Value, error) {
	if err := json.Unmarshal(data, resource); err != nil {
		return reflect.Value{}, err
	}
	return resource, nil
}

func (a TestArg) value() (reflect.Value, error) {
	var value reflect.Value
	data, err := a.data()
	if err != nil {
		return reflect.Value{}, err
	}
	switch a.Type {
	case "Engine":
		e := engine.Engine{}
		if err := json.Unmarshal(data, &e); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(e)
	case "LocalRunner":
		runner := NewLocalRunner(engine.Engine{}, "", ResourceNumScheduler{})
		if err := json.Unmarshal(data, &runner); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(runner)
	case "Solution":
		solution := engine.Solution{}
		if err := json.Unmarshal(data, &solution); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(solution)
	case "Tool":
		tool := engine.Tool{}
		if err := json.Unmarshal(data, &tool); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(tool)
	case "Resource":
		resource := engine.Resource{}
		if err := json.Unmarshal(data, &resource); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(resource)
	case "ImplicitTask":
		task := engine.ImplicitTask{}
		if err := json.Unmarshal(data, &task); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(task)
	case "Schedule":
		schedule := Schedule{}
		if err := json.Unmarshal(data, &schedule); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(schedule)
	case "[]Schedule":
		schedule := []Schedule{}
		if err := json.Unmarshal(data, &schedule); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(schedule)
	case "[]Task":
		tasks := []engine.Task{}
		if err := json.Unmarshal(data, &tasks); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(tasks)
	case "ResourceNumScheduler":
		schedule := ResourceNumScheduler{}
		if err := json.Unmarshal(data, &schedule); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(schedule)
	case "MockRunner":
		mock := MockRunner{}
		if err := json.Unmarshal(data, &mock); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(mock)
	case "int":
		number := int(0)
		if err := json.Unmarshal(data, &number); err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(number)
	case "native":
		value = reflect.ValueOf(a.Value)
	case "error":
		value = reflect.ValueOf(errors.New(fmt.Sprint(a.Value)))
	default:
		return reflect.Value{}, errors.New(fmt.Sprintf("Unknown arg type %v", a.Type))
	}
	return value, nil
}

func castArg(args []TestArg) ([]reflect.Value, error) {
	values := make([]reflect.Value, len(args))
	for k, param := range args {
		value, err := param.value()
		if err != nil {
			return nil, err
		}
		values[k] = value
	}
	return values, nil
}

func testType(subject reflect.Value, test TestCase) error {
	method := subject.MethodByName(test.Function)
	if !method.IsValid() {
		return fmt.Errorf("Invalid function name")
	}
	if len(test.Args) != method.Type().NumIn() {
		return errors.New("The number of params is not adapted.")
	}
	in, err := castArg(test.Args)
	if err != nil {
		return err
	}

	out, err := castArg(test.Want)
	if err != nil {
		return err
	}

	results := method.Call(in)
	if len(results) == 0 {
		opt := jsondiff.DefaultConsoleOptions()
		eJSON, _ := json.Marshal(subject)
		wJSON, _ := json.Marshal(out[0])
		if result, diff := jsondiff.Compare(eJSON, wJSON, &opt); result.String() != "FullMatch" {
			return fmt.Errorf(diff)
		}
	} else {
		var errs string
		for i, result := range results {
			var want string
			if out[i].IsValid() {
				want = fmt.Sprintf("%v", out[i])
			}
			if fmt.Sprintf("%v", result) != want {
				errs = errs + fmt.Sprintf("got %v, expected %v\n", result, out[i])
			}
		}
		if len(errs) > 0 {
			return errors.New(errs)
		}
	}
	return nil
}

func runTests(t *testing.T, testFilePaths []string, funcMap map[string]interface{}, cast func(subject interface{}) reflect.Value) {
	tests := 0
	passed := 0
	for _, path := range testFilePaths {
		t.Run(path, func(t *testing.T) {
			base := filepath.Base(path)
			testSets := []*TestSet{}
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
					tests++
					switch ts.Type {
					case "method":
						if err := testMethod(funcMap, c); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %v", base, ts.Description, i, c.Description, err)
						} else {
							passed++
						}
					case "type":
						if err := testType(cast(ts.Subject), c); err != nil {
							t.Errorf("%s: %s test case %d: %s. error: %v", base, ts.Description, i, c.Description, err)
						} else {
							passed++
						}
					}
				}
			}
			t.Logf("%d/%d tests passed", passed, tests)
		})
	}
}

func Test_runner(t *testing.T) {
	runTests(t, []string{
		"testdata/resource_num_scheduler.json",
		"testdata/runner.json",
	}, map[string]interface{}{
		"solutionSize":       solutionSize,
		"variableKeyName":    variableKeyName,
		"updateArgs":         updateArgs,
		"resolveImplicitArg": resolveImplicitArg,
		"resolveArgs":        resolveArgs,
		"Run":                Run,
	},
		func(subject interface{}) reflect.Value {
			sJSON, _ := json.Marshal(subject)
			scheduler := ResourceNumScheduler{}
			json.Unmarshal(sJSON, &scheduler)
			return reflect.ValueOf(scheduler)
		})
	runTests(t, []string{
		"testdata/local.runner.json",
	}, map[string]interface{}{
		"NewLocalRunner": NewLocalRunner,
		"shell":          shell,
		"renderTemplate": renderTemplate,
	},
		func(subject interface{}) reflect.Value {
			sJSON, _ := json.Marshal(subject)
			runner := LocalRunner{Scheduler: ResourceNumScheduler{}}
			json.Unmarshal(sJSON, &runner)
			return reflect.ValueOf(&runner)
		})
}
