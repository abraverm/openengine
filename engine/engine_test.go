package engine

import (
	"reflect"
	"testing"
)

func TestEngine_AddProvider(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		api ProviderAPI
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			e.Match()
		})
	}
}

func TestEngine_AddProvisioner(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		provisioner Provisioner
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			e.Match()
		})
	}
}

func TestEngine_AddResource(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		resource Resource
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			e.Match()
		})
	}
}

func TestEngine_AddSystem(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		system System
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			e.Match()
		})
	}
}

func TestEngine_AddTool(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		api ToolAPI
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			e.Match()
		})
	}
}

func TestEngine_GetSolutions(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
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
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			if got := e.GetSolutions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSolutions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_Match(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			if err := e.Match(); (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_Resolve(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			e.Match()
			e.Resolve()
		})
	}
}

func TestEngine_Run(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			got, err := e.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_Schedule(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
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
			e := &Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			if err := e.Schedule(tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("Schedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_getTool(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		task ImplicitTask
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Tool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			got, err := e.getTool(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_matchProvidersProvisioners(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		resource Resource
		system   System
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Solution
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			got, err := e.matchProvidersProvisioners(tt.args.resource, tt.args.system)
			if (err != nil) != tt.wantErr {
				t.Errorf("matchProvidersProvisioners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchProvidersProvisioners() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_resolveDependencies(t *testing.T) {
	type fields struct {
		systems      []System
		resources    []Resource
		providers    []Provider
		provisioners []Provisioner
		solutions    []Solution
		tools        []Tool
		schedule     []Schedule
	}
	type args struct {
		solution Solution
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				systems:      tt.fields.systems,
				resources:    tt.fields.resources,
				providers:    tt.fields.providers,
				provisioners: tt.fields.provisioners,
				solutions:    tt.fields.solutions,
				tools:        tt.fields.tools,
				schedule:     tt.fields.schedule,
			}
			if got := e.resolveDependencies(tt.args.solution); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveDependencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEngine(t *testing.T) {
	tests := []struct {
		name string
		want *Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
