package engine

import (
	"reflect"
	"testing"
)

func TestSolution_MarshalJSON(t *testing.T) {
	type fields struct {
		Resource       Resource
		System         System
		Provider       Provider
		Provisioner    Provisioner
		resolved       bool
		resolutionTree map[string]Param
		parent         *Solution
		size           int
		action         string
		Output         string
		debug          bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{
				Resource:       tt.fields.Resource,
				System:         tt.fields.System,
				Provider:       tt.fields.Provider,
				Provisioner:    tt.fields.Provisioner,
				resolved:       tt.fields.resolved,
				resolutionTree: tt.fields.resolutionTree,
				parent:         tt.fields.parent,
				size:           tt.fields.size,
				action:         tt.fields.action,
				Output:         tt.fields.Output,
				debug:          tt.fields.debug,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Run(t *testing.T) {
	type fields struct {
		Resource       Resource
		System         System
		Provider       Provider
		Provisioner    Provisioner
		resolved       bool
		resolutionTree map[string]Param
		parent         *Solution
		size           int
		action         string
		Output         string
		debug          bool
	}
	type args struct {
		solutionArgs map[string]interface{}
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
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{
				Resource:       tt.fields.Resource,
				System:         tt.fields.System,
				Provider:       tt.fields.Provider,
				Provisioner:    tt.fields.Provisioner,
				resolved:       tt.fields.resolved,
				resolutionTree: tt.fields.resolutionTree,
				parent:         tt.fields.parent,
				size:           tt.fields.size,
				action:         tt.fields.action,
				Output:         tt.fields.Output,
				debug:          tt.fields.debug,
			}
			got, err := s.Run(tt.args.solutionArgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_ToJson(t *testing.T) {
	type fields struct {
		Resource       Resource
		System         System
		Provider       Provider
		Provisioner    Provisioner
		resolved       bool
		resolutionTree map[string]Param
		parent         *Solution
		size           int
		action         string
		Output         string
		debug          bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{
				Resource:       tt.fields.Resource,
				System:         tt.fields.System,
				Provider:       tt.fields.Provider,
				Provisioner:    tt.fields.Provisioner,
				resolved:       tt.fields.resolved,
				resolutionTree: tt.fields.resolutionTree,
				parent:         tt.fields.parent,
				size:           tt.fields.size,
				action:         tt.fields.action,
				Output:         tt.fields.Output,
				debug:          tt.fields.debug,
			}
			if got := s.ToJSON(); got != tt.want {
				t.Errorf("ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_decouple(t *testing.T) {
	type fields struct {
		Resource       Resource
		System         System
		Provider       Provider
		Provisioner    Provisioner
		resolved       bool
		resolutionTree map[string]Param
		parent         *Solution
		size           int
		action         string
		Output         string
		debug          bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{
				Resource:       tt.fields.Resource,
				System:         tt.fields.System,
				Provider:       tt.fields.Provider,
				Provisioner:    tt.fields.Provisioner,
				resolved:       tt.fields.resolved,
				resolutionTree: tt.fields.resolutionTree,
				parent:         tt.fields.parent,
				size:           tt.fields.size,
				action:         tt.fields.action,
				Output:         tt.fields.Output,
				debug:          tt.fields.debug,
			}
			if got := s.decouple(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decouple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_equals(t *testing.T) {
	type fields struct {
		Resource       Resource
		System         System
		Provider       Provider
		Provisioner    Provisioner
		resolved       bool
		resolutionTree map[string]Param
		parent         *Solution
		size           int
		action         string
		Output         string
		debug          bool
	}
	type args struct {
		solution Solution
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{
				Resource:       tt.fields.Resource,
				System:         tt.fields.System,
				Provider:       tt.fields.Provider,
				Provisioner:    tt.fields.Provisioner,
				resolved:       tt.fields.resolved,
				resolutionTree: tt.fields.resolutionTree,
				parent:         tt.fields.parent,
				size:           tt.fields.size,
				action:         tt.fields.action,
				Output:         tt.fields.Output,
				debug:          tt.fields.debug,
			}
			if got := s.equals(tt.args.solution); got != tt.want {
				t.Errorf("equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_inLoop(t *testing.T) {
	type fields struct {
		Resource       Resource
		System         System
		Provider       Provider
		Provisioner    Provisioner
		resolved       bool
		resolutionTree map[string]Param
		parent         *Solution
		size           int
		action         string
		Output         string
		debug          bool
	}
	type args struct {
		solution Solution
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{
				Resource:       tt.fields.Resource,
				System:         tt.fields.System,
				Provider:       tt.fields.Provider,
				Provisioner:    tt.fields.Provisioner,
				resolved:       tt.fields.resolved,
				resolutionTree: tt.fields.resolutionTree,
				parent:         tt.fields.parent,
				size:           tt.fields.size,
				action:         tt.fields.action,
				Output:         tt.fields.Output,
				debug:          tt.fields.debug,
			}
			if got := s.inLoop(tt.args.solution); got != tt.want {
				t.Errorf("inLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_resolveExplicit(t *testing.T) {
	type fields struct {
		Resource       Resource
		System         System
		Provider       Provider
		Provisioner    Provisioner
		resolved       bool
		resolutionTree map[string]Param
		parent         *Solution
		size           int
		action         string
		Output         string
		debug          bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{
				Resource:       tt.fields.Resource,
				System:         tt.fields.System,
				Provider:       tt.fields.Provider,
				Provisioner:    tt.fields.Provisioner,
				resolved:       tt.fields.resolved,
				resolutionTree: tt.fields.resolutionTree,
				parent:         tt.fields.parent,
				size:           tt.fields.size,
				action:         tt.fields.action,
				Output:         tt.fields.Output,
				debug:          tt.fields.debug,
			}
			if got := s.resolveExplicit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveExplicit() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRange(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solutionList_Len(t *testing.T) {
	tests := []struct {
		name string
		s    solutionList
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solutionList_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    solutionList
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solutionList_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    solutionList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
