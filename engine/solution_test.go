package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/nsf/jsondiff"
)

type TestData string

func (t TestData) Reader() ([]byte, error) {
	tJSON, err := ioutil.ReadFile(filepath.Join("testdata", string(t)+".json"))
	if err != nil {
		return nil, fmt.Errorf("failed reading: %s", err)
	}
	return tJSON, nil
}

type TestSolution struct {
	Resource       Resource         `json:"resource"`
	System         System           `json:"system"`
	Provider       Provider         `json:"provider"`
	Provisioner    Provisioner      `json:"provisioner"`
	Resolved       bool             `json:"resolved"`
	ResolutionTree map[string]Param `json:"resolution_tree"`
	Parent         *Solution        `json:"parent"`
	Size           int              `json:"size"`
	Action         string           `json:"action"`
	Output         string           `json:"output"`
	Debug          bool             `json:"debug"`
}

func (s *Solution) UnmarshalJSON(b []byte) error {
	temp := &TestSolution{}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	s.resolved = temp.Resolved
	s.ResolutionTree = temp.ResolutionTree
	s.Action = temp.Action
	s.debug = temp.Debug
	s.parent = temp.Parent
	s.System = temp.System
	s.Resource = temp.Resource
	s.Provisioner = temp.Provisioner
	s.Provider = temp.Provider
	s.Output = temp.Output

	return nil
}

func TestSolution_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		test    TestData
		want    TestData
		wantErr bool
	}{
		{"empty", "solution.empty", "solution.empty", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testJson, _ := tt.test.Reader()
			wantJson, _ := tt.want.Reader()
			testSolution := &Solution{}
			json.Unmarshal(testJson, &testSolution)
			opt := jsondiff.DefaultConsoleOptions()
			got, err := testSolution.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result, diff := jsondiff.Compare(wantJson, got, &opt); result.String() != "FullMatch" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestSolution_ToJson(t *testing.T) {
	tests := []struct {
		name string
		in   TestData
		out  TestData
	}{
		{"empty", "solution.empty", "solution.empty.ToJSON"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inJson, err := tt.in.Reader()
			if err != nil {
				t.Fatalf("%s", err)
			}
			outJson, err := tt.out.Reader()
			if err != nil {
				t.Fatalf("%s", err)
			}
			inSolution := &Solution{}
			json.Unmarshal(inJson, inSolution)
			opt := jsondiff.DefaultConsoleOptions()
			if result, diff := jsondiff.Compare([]byte(inSolution.ToJSON()), outJson, &opt); result.String() != "FullMatch" {
				t.Errorf("%v", diff)
			}
		})
	}
}

func TestSolution_decouple(t *testing.T) {
	tests := []struct {
		name string
		test TestData
		want []TestData
	}{
		{"empty", "solution.empty", []TestData{"solution.empty"}},
		{"implicit", "solution.implicit", []TestData{"solution.implicit.decouple"}},
	}
	opt := jsondiff.DefaultConsoleOptions()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want = []Solution{}
			for _, file := range tt.want {
				wantJSON, _ := file.Reader()
				wantSolution := &Solution{}
				json.Unmarshal(wantJSON, &wantSolution)
				want = append(want, *wantSolution)
			}
			testJson, _ := tt.test.Reader()
			testSolution := &Solution{}
			json.Unmarshal(testJson, &testSolution)
			got := testSolution.decouple()
			gJSON, _ := json.Marshal(got)
			wJSON, _ := json.Marshal(want)
			if result, diff := jsondiff.Compare(gJSON, wJSON, &opt); result.String() != "FullMatch" {
				t.Errorf("%v", diff)
			}
		})
	}
}

func TestSolution_equals(t *testing.T) {
	tests := []struct {
		name     string
		testData TestData
		want     TestData
	}{
		{"emtpy", "solution.empty", "solution.empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inJson, err := tt.testData.Reader()
			if err != nil {
				t.Fatalf("%s", err)
			}
			outJson, err := tt.want.Reader()
			if err != nil {
				t.Fatalf("%s", err)
			}
			inSolution := &Solution{}
			json.Unmarshal(inJson, inSolution)
			outSolution := &Solution{}
			json.Unmarshal(outJson, outSolution)
			if !inSolution.equals(*outSolution) {
				t.Errorf("Solution not equal")
			}
		})
	}
}

func TestSolution_inLoop(t *testing.T) {
	tests := []struct {
		name     string
		testCase TestData
		arg      TestData
		want     bool
	}{
		{"empty", "solution.empty", "solution.empty", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testJson, err := tt.testCase.Reader()
			if err != nil {
				t.Fatalf("%s", err)
			}
			argJson, err := tt.arg.Reader()
			if err != nil {
				t.Fatalf("%s", err)
			}
			testSolution := &Solution{}
			json.Unmarshal(testJson, testSolution)
			argSolution := &Solution{}
			json.Unmarshal(argJson, argSolution)
			if testSolution.inLoop(*argSolution) != tt.want {
				result := "'nt"
				if tt.want {
					result = ""
				}
				t.Errorf("Solution should%v be in loop", result)
			}
		})
	}
}

func TestSolution_resolveExplicit(t *testing.T) {
	tests := []struct {
		name string
		test TestData
		new  TestData
		want []string
	}{
		{"empty", "solution.empty", "solution.empty", nil},
		{"implicit", "solution.implicit", "solution.implicit", []string{"flavorRef", "name"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testJson, _ := tt.test.Reader()
			newJson, _ := tt.new.Reader()
			testSolution := &Solution{}
			err := json.Unmarshal(testJson, &testSolution)
			if err != nil {
				t.Fatalf("%v", err)
			}
			opt := jsondiff.DefaultConsoleOptions()
			got := testSolution.resolveExplicit()
			sort.Strings(got)
			sort.Strings(tt.want)
			newnew, _ := json.Marshal(testSolution)
			if result, diff := jsondiff.Compare(newJson, newnew, &opt); result.String() != "FullMatch" {
				t.Fatalf(diff)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("resolveExplicit() got %v but wanted %v", got, tt.want)
			}
		})
	}
}

func Test_combIntSlices(t *testing.T) {
	type args struct {
		seq [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantOut [][]int
	}{
		{
			name: "zero",
			args: args{
				seq: [][]int{},
			},
			wantOut: nil,
		},
		{
			name: "[[1]]",
			args: args{
				seq: [][]int{{1}},
			},
			wantOut: nil,
		},
		{
			name: "[[1][2]]",
			args: args{
				seq: [][]int{{1}, {2}},
			},
			wantOut: [][]int{{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := combIntSlices(tt.args.seq); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("combIntSlices() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_intersect(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				a: []string{},
				b: []string{},
			},
			want: []string{},
		},
		{
			name: "one",
			args: args{
				a: []string{"one"},
				b: []string{"one"},
			},
			want: []string{"one"},
		},
		{
			name: "none",
			args: args{
				a: []string{"one"},
				b: []string{"two"},
			},
			want: []string{},
		},
		{
			name: "empty a",
			args: args{
				a: []string{},
				b: []string{"two"},
			},
			want: []string{},
		},
		{
			name: "empty b",
			args: args{
				a: []string{"one"},
				b: []string{},
			},
			want: []string{},
		},
		{
			name: "intersect",
			args: args{
				a: []string{"one", "two"},
				b: []string{"two", "three"},
			},
			want: []string{"two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "zero",
			args: args{
				min: 0,
				max: 0,
			},
			want: []int{0},
		},
		{
			name: "min=0",
			args: args{
				min: 0,
				max: 1,
			},
			want: []int{0, 1},
		},
		{
			name: "min>0",
			args: args{
				min: 1,
				max: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "min<0",
			args: args{
				min: -1,
				max: 1,
			},
			want: []int{-1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRange(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
