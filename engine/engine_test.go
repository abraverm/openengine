package engine

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"cuelang.org/go/cue/cuecontext"

	"github.com/nsf/jsondiff"
)

func TestEngine_AddResource(t *testing.T) {
	type P map[string]interface{}
	tests := []struct {
		name     string
		resource Resource
		wantErr  bool
	}{
		{"empty", Resource{}, false},
		{"minimal", Resource{Name: "Server"}, false},
		{"with args", Resource{Name: "Server", Properties: P{"key": "value"}}, false},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := NewEngine("")
			if err := e.AddResource(tt.resource); (err != nil) != tt.wantErr {
				t.Errorf("AddResource() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("[AddResource] completed test (%d/%d) - %s", i, len(tests), tt.name)
		})
	}
}

func TestEngine_AddSystem(t *testing.T) {
	tests := []struct {
		name    string
		system  System
		wantErr bool
	}{
		{"minimal", System{Type: "Openstack"}, false},
		{"with args", System{Type: "Openstack", Properties: map[string]interface{}{"key": "value"}}, false},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := NewEngine("")
			if err := e.AddSystem(tt.system); (err != nil) != tt.wantErr {
				t.Errorf("AddSystem() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("[AddSystem] completed test (%d/%d) - %s", i, len(tests), tt.name)
		})
	}
}

func TestEngine_addDefinition(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		definition string
		wantErr    bool
	}{
		{"empty", "", "", false},
		{"empty path", "", "test: true", false},
		{"empty definition", "test", "", false},
		{"bad definition", "test", "test:", true},
		{"good definition", "test", "test: true", false},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := NewEngine("")
			if err := e.addDefinition(tt.definition); (err != nil) != tt.wantErr {
				t.Errorf("addDefinition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		t.Logf("[addDefinition] completed test (%d/%d) - %s", i, len(tests), tt.name)
	}
}

func TestEngine_loadSpec(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"empty", "", false},
		{"missing", "testdata/missing.cue", true},
		{"different", "testdata/empty.cue", false},
		{"broken", "testdata/broken.cue", true},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := NewEngine("")
			if err := e.loadSpec(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("loadSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("[loadSpec] completed test (%d/%d) - %s", i, len(tests), tt.name)
		})
	}
}

func TestNewEngine(t *testing.T) {
	type args struct {
		spec string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"default", args{""}, false},
		{"missing file", args{"testdata/missing.cue"}, true},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEngine(tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEngine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("[NewEngine] completed test (%d/%d) - %s", i, len(tests), tt.name)
		})
	}
}

func Test_loadFile(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{"empty", "", "", true},
		{"dir", "testdata", "", true},
		{"missing", "testdata/missing.cue", "", true},
		{"empty file", "testdata/empty.cue", "", false},
		{"file with content", "testdata/loadFile.txt", "test", false},
		{"file with content", "testdata/badPermissions.txt", "", true},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadFile(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadFile() got = %v, want %v", got, tt.want)
			}
			t.Logf("[loadFile] completed test (%d/%d) - %s", i, len(tests), tt.name)
		})
	}
}

type files []string

type SolutionTest struct {
	name         string
	action       string
	resources    files
	systems      files
	providers    files
	provisioners files
	solutions    string
	skip         bool
	wantErr      bool
}

func (s SolutionTest) update(name, id, action, resource, provider, provisioner, system, solution string) SolutionTest {
	r := s
	if id != "" {
		r.resources = files{fmt.Sprintf("testdata/resource_%s.cue", id)}
		r.systems = files{fmt.Sprintf("testdata/system_%s.cue", id)}
		r.providers = files{fmt.Sprintf("testdata/provider_%s.cue", id)}
		r.provisioners = files{fmt.Sprintf("testdata/provisioner_%s.cue", id)}
		r.solutions = fmt.Sprintf("testdata/solution_%s.cue", id)
	}
	if action != "" {
		r.action = action
	}
	if name != "" {
		r.name = name
	}
	if resource != "" {
		r.resources = files{fmt.Sprintf("testdata/resource_%s.cue", resource)}
	}
	if system != "" {
		r.systems = files{fmt.Sprintf("testdata/system_%s.cue", system)}
	}
	if provider != "" {
		r.providers = files{fmt.Sprintf("testdata/provider_%s.cue", provider)}
	}
	if provisioner != "" {
		r.provisioners = files{fmt.Sprintf("testdata/provisioner_%s.cue", provisioner)}
	}
	if solution != "" {
		r.solutions = fmt.Sprintf("testdata/solution_%s.cue", solution)
	}

	return r
}

func (s SolutionTest) merge(resources, providers, provisioners, systems []string) SolutionTest {
	newS := s
	for _, resource := range resources {
		newS.resources = append(newS.resources, fmt.Sprintf("testdata/resource_%s.cue", resource))
	}
	for _, provider := range providers {
		newS.providers = append(newS.providers, fmt.Sprintf("testdata/provider_%s.cue", provider))
	}
	for _, provisioner := range provisioners {
		newS.provisioners = append(newS.provisioners, fmt.Sprintf("testdata/provisioner_%s.cue", provisioner))
	}
	for _, system := range systems {
		newS.systems = append(newS.systems, fmt.Sprintf("testdata/system_%s.cue", system))
	}
	return newS
}

func generateResource(name, rType bool) string {
	newS := "{\n"
	if name {
		newS += "name: \"Test\""
	}
	if rType {
		newS += "type: \"Server\""
	}
	newS += "}"
	return newS
}

// TODO: Complete other definition testing
func TestEngine_Spec(t *testing.T) {
	// Resource tests
	for _, tt := range []struct {
		name    bool
		rType   bool
		wantErr bool
	}{
		{true, false, true},
	} {
		t.Run("Resource validation", func(t *testing.T) {
			r := generateResource(tt.name, tt.rType)
			e, _ := NewEngine("")
			err := e.Add("Resource", r)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Resource validation failed: %v, wanted Error %v\n%s", err, tt.wantErr, e.GetSpec())
			}
		})
	}

	// Action tests
	for _, tt := range []struct {
		action  string
		wantErr bool
	}{
		{"", true},
		{"wrong", true},
		{"create", false},
		{"delete", false},
		{"update", false},
		{"read", false},
	} {
		t.Run("Action validation", func(t *testing.T) {
			e, _ := NewEngine("")
			err := e.Add("Action", "\""+tt.action+"\"")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Action validation failed: %v, wanted Error %v", err, tt.wantErr)
			}
			t.Logf("[Spec - Actions] completed test  - %s", tt.action)
		})
	}
}

func TestEngine_SolutionsCombinations(t *testing.T) {
	t.Log("Test: Solutions definitions combinations")
	var tests = []SolutionTest{}

	// Combination tests - all ingredients are needed to make a solution
	for _, resource := range []bool{false, true} {
		for _, system := range []bool{false, true} {
			for _, provider := range []bool{false, true} {
				for _, provisioner := range []bool{false, true} {
					s := SolutionTest{
						name:      "combination testing",
						action:    "create",
						wantErr:   false,
						solutions: "testdata/solution_empty.cue",
					}
					if resource && system && provider && provisioner {
						s.wantErr = false
						s.solutions = "testdata/solution_one.cue"
					}
					if resource {
						s.resources = files{"testdata/resource_minimal.cue"}
					}
					if system {
						s.systems = files{"testdata/system_minimal.cue"}
					}
					if provider {
						s.providers = files{"testdata/provider_minimal.cue"}
					}
					if provisioner {
						s.provisioners = files{"testdata/provisioner_minimal.cue"}
					}
					tests = append(tests, s)
				}
			}
		}
	}

	runSolutionTests(t, tests)
}

func TestEngine_SolutionsPropertiesCombinations(t *testing.T) {
	t.Log("Test: Solutions properties combinations")
	var tests []SolutionTest
	testPropertiesTemplate := SolutionTest{
		name:         "Properties",
		action:       "create",
		systems:      files{"testdata/system_minimal.cue"},
		provisioners: files{"testdata/provisioner_mixed_properties.cue"},
		providers:    files{"testdata/provider_mixed_properties.cue"},
		wantErr:      false,
	}

	for _, t := range []OneToOne{
		// m - minimal, mx - mixed, i - implicit, e - explicit, o - optional
		// Test *independent* properties of different type when used *together*
		// mx1 - 1 explicit without implicit, 1 implicit without explicit, 1 optional
		{name: "(0e0i0o)R-mx1[PD,PR]-mSU-_SO", resource: "minimal", solution: "empty"},
		{name: "(0e0i1o)R-mx1[PD,PR]-mSU-_SO", resource: "0e0i1o", solution: "empty"},
		{name: "(0e1i0o)R-mx1[PD,PR]-mSU-_SO", resource: "0e1i0o", solution: "empty"},
		{name: "(0e1i1o)R-mx1[PD,PR]-mSY-_SO", resource: "0e1i1o", solution: "empty"},
		{name: "(1e0i0o)R-mx1[PD,PR]-mSY-_SO", resource: "1e0i0o", solution: "empty"},
		{name: "(1e0i1o)R-mx1[PD,PR]-mSU-_SO", resource: "1e0i1o", solution: "empty"},
		{name: "(1e1i0o)R-mx1[PD,PR]-mSY-(1e1i0o)SO", resource: "1e1i0o", solution: "1e1i0o"},
		{name: "(1e1i1o)R-mx1[PD,PR]-mSY-(1e1i1o)SO", resource: "1e1i1o", solution: "1e1i1o"},
		// Test properties of different type when used *alone*
		{name: "(1e)R-e[PD,PR]-mSY-eSO", id: "explicit", system: "minimal"},
		{name: "(1i)R-iPD-ePR-mSY-iSO", id: "implicit", provisioner: "explicit", system: "minimal"},
		{name: "(1o)R-o[PD,PR]-mSY-eSO", id: "optional", system: "minimal", resource: "explicit", solution: "explicit"},
		{name: "(0o)R-o[PD,PR]-mSY-oneSO", id: "optional", system: "minimal", resource: "minimal"},
		// Test properties override
		{name: "(1ei)[R,PD]-ePR-mSY-1SO", id: "override", provider: "implicit", provisioner: "explicit", system: "minimal"},
	} {
		tests = append(tests, testPropertiesTemplate.update(t.name, t.id, t.action, t.resource, t.provider, t.provisioner, t.system, t.solution))
	}

	runSolutionTests(t, tests)
}

type OneToOne struct {
	name        string
	id          string
	resource    string
	solution    string
	provider    string
	provisioner string
	system      string
	action      string
}

func TestEngine_Solutions_N1(t *testing.T) {
	t.Log("Test: Solutions N+1")
	var tests = []SolutionTest{}

	// N+1 (N1) - how increasing ingredients will produce same or more solutions
	// R - Resource, S - System, PD - Provider, PR - Provisioner, SO - Solutions
	// s - same, dm - different matching, n - not, i - implicit
	// #SO - Number of solutions
	for _, t := range []SolutionTest{
		// The tests definition only require what to add and expected result
		// The "base" set is the "minimal" combination
		// * Test increasing (s)ame ingredients will result in same number of solutions
		{name: "N1-sR-1SO", resources: files{"minimal"}, solutions: "one"},
		{name: "N1-sPD-1SO", providers: files{"minimal"}, solutions: "one"},
		{name: "N1-sPR-1SO", provisioners: files{"minimal"}, solutions: "one"},
		{name: "N1-sS-1SO", systems: files{"minimal"}, solutions: "one"},
		// * Test increasing different matching ingredients will increase number of solutions
		{name: "N1-dmR-2SO", resources: files{"minimal2"}, solutions: "N1-dmR"},
		{name: "N1-dmPD-2SO", providers: files{"minimal2"}, solutions: "N1-dmPD"},
		{name: "N1-dmPR-2SO", provisioners: files{"minimal2"}, solutions: "N1-dmPR"},
		{name: "N1-dmS-2SO", systems: files{"minimal2"}, solutions: "N1-dmS"},
		// * Test increasing different but not matching ingredients will not increase number of solutions
		{name: "N1-dnmR-1SO", resources: files{"other"}, solutions: "one"},
		{name: "N1-dnmPD-1SO", providers: files{"other"}, solutions: "one"},
		{name: "N1-dnmPR-1SO", provisioners: files{"other"}, solutions: "one"},
		{name: "N1-dnmS-1SO", systems: files{"other"}, solutions: "one"},
	} {
		tn := SolutionTest{}.update(t.name, "minimal", "create", "", "", "", "", t.solutions)
		tests = append(tests, tn.merge(t.resources, t.providers, t.provisioners, t.systems))
	}
	runSolutionTests(t, tests)
}

func TestEngine_Solutions_N1_ImplicitOne(t *testing.T) {
	t.Log("Test: Solutions N+1 Implicit one")
	var tests = []SolutionTest{}
	// This is a case where the OpenEngine finds solutions for implicit tasks that use resources
	// TODO: It takes very long time for Cue to resolve such cases (bug in spec or cue)

	// N+1 (N1) - how increasing ingredients will produce same or more solutions
	// R - Resource, S - System, PD - Provider, PR - Provisioner, SO - Solutions
	// n - not matched, m - matched, r - resource task, s - script task
	// #SO - Number of solutions
	for _, t := range []SolutionTest{
		{name: "N1-mrPD-2SO", providers: files{"implicit2"}, solutions: "N1-mrPD", skip: true}, // todo: fix memory leak
		{name: "N1-nrPD-1SO", providers: files{"implicit3"}, solutions: "one"},
		{name: "N1-msPD-2SO", providers: files{"implicit4"}, solutions: "N1-msPD"},
		{name: "N1-nsPD-1SO", providers: files{"implicit"}, solutions: "one"},
	} {
		tn := SolutionTest{}.update(t.name, "minimal", "create", "", "", "", "", t.solutions)
		if t.skip {
			tn.skip = true
		}
		tests = append(tests, tn.merge(t.resources, t.providers, t.provisioners, t.systems))
	}
	runSolutionTests(t, tests)
}

func TestEngine_Solutions_N1_ImplicitTwo(t *testing.T) {
	t.Log("Test: Solutions N+1 Implicit two")
	var tests = []SolutionTest{}
	// This is a case where the OpenEngine finds solutions for mix of implicit tasks
	// TODO: It takes very long time for Cue to resolve such cases (bug in spec or cue)

	// N+1 (N1) - how increasing ingredients will produce same or more solutions
	// R - Resource, S - System, PD - Provider, PR - Provisioner, SO - Solutions
	// m - match, n - not matched, r - resource task, s - script task
	// #SO - Number of solutions
	for _, t := range []SolutionTest{
		{name: "N1- mrmsPD-2SO", providers: files{"mrms"}, solutions: "N1-mrmsPD", skip: true}, // matched - todo: fix memory leak
		{name: "N1- mrnsPD-1SO", providers: files{"mrns"}, solutions: "one"},                   // unmatched
		{name: "N1- nrmsPD-1SO", providers: files{"nrms"}, solutions: "one"},                   // unmatched
		{name: "N1- nrnsPD-1SO", providers: files{"nrns"}, solutions: "one"},                   // unmatched
	} {
		tn := SolutionTest{}.update(t.name, "minimal", "create", "", "", "", "", t.solutions)
		if t.skip {
			tn.skip = true
		}
		tests = append(tests, tn.merge(t.resources, t.providers, t.provisioners, t.systems))
	}
	runSolutionTests(t, tests)
}

func runSolutionTests(t *testing.T, tests []SolutionTest) {
	jsondiffOpts := jsondiff.DefaultConsoleOptions()
	for i, tt := range tests {
		if tt.skip && testing.Short() {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			var errors []error
			e, err := NewEngine("")
			if err != nil {
				t.Errorf("Failed to initialize new Engine: %v", err)
				return
			}
			for _, path := range tt.resources {
				def, _ := loadFile(path)
				if err = e.Add("Resource", def); err != nil {
					errors = append(errors, fmt.Errorf("%s - %v", path, err))
				}
			}
			for _, path := range tt.systems {
				def, _ := loadFile(path)
				if err = e.Add("System", def); err != nil {
					errors = append(errors, fmt.Errorf("%s - %v", path, err))
				}
			}
			for _, path := range tt.providers {
				def, _ := loadFile(path)
				e.Add("Provider", def)
				if err = e.Add("Provider", def); err != nil {
					errors = append(errors, fmt.Errorf("%s - %v", path, err))
				}
			}
			for _, path := range tt.provisioners {
				def, _ := loadFile(path)
				if err = e.Add("Provisioner", def); err != nil {
					errors = append(errors, fmt.Errorf("%s - %v", path, err))
				}
			}
			if len(errors) > 0 {
				for _, e := range errors {
					t.Error(e)
				}
				t.Fatalf("%v", e.GetSpec())
			}
			want, _ := loadFile(tt.solutions)
			c := cuecontext.New()
			wantValue := c.CompileString(want)
			wantJSON, err := wantValue.MarshalJSON()
			if err != nil {
				t.Fatalf("Unable to create want json:%v\n%s", err, wantJSON)
			}
			c1 := make(chan string, 1)
			go func() {
				got, err := e.Solutions(tt.action)
				if (err != nil) != tt.wantErr {
					t.Logf("Solutions() error = %v, wantErr %v, spec:\n %s", err, tt.wantErr, e.GetSpec())
					t.Errorf("Solutions() error = %v, wantErr %v, spec:\n %s", err, tt.wantErr, e.GetSpec())
					return
				}
				c1 <- got
			}()

			select {
			case got := <-c1:
				gotValue := c.CompileString(got)
				if gotValue.Err() != nil {
					t.Fatalf("%v\nSpec:\n%s", gotValue.Err(), e.GetSpec())
				}

				gotJSON, _ := gotValue.MarshalJSON()
				if result, diff := jsondiff.Compare(wantJSON, gotJSON, &jsondiffOpts); result.String() != "FullMatch" {
					t.Logf("Got:\n%s\nDifference:\n%s\nSpec:\n%s", got, diff, e.GetSpec())
					t.Fatalf("Got:\n%s\nDifference:\n%s\nSpec:\n%s", got, diff, e.GetSpec())
				}
				endTime := time.Now()
				t.Logf("completed test (%d/%d) in %s - %s", i+1, len(tests), endTime.Sub(startTime).String(), tt.name)
			case <-time.After(15 * time.Minute):
				t.Error("Out of time (15 minutes)")
			}
		})
	}
}

// Explicit dependencies are cases when resource depends on another defined resources
// Example: Server depends on a Network with the id "1" for the property "network", the network doesn't exist and will be created too.
//          The network is a defined resource just as the server and it is not part of Server implicit process.
//          The network creation response doesn't need to be processed as the Server resource has its explicit id in its definition.
//          The only critical part is the existence of a solution with the the right dependency structure
// The followings cases are tested:
//  - basic: two resources, same type and one depends on another, one system, provisioner and provider - one solution
//  - cycle: basic and the second resource depends on first - no solution
//  - two systems - basic with additional matching system as alternative - four solutions
//  - two different resources and two different systems - no solution
//  - two providers - basic with additional provider that matches - four solutions
//  - two provisioners - basic with additional provisioner that matches - four solutions
//  - there resources deep nested  - basic with additional resource so a->b->c - one solution
//  - there resources two brothers, both resolved - basic with additional resource so b<-a->c - one solution
//  - there resources two brothers, one not resolved - basic with additional resource so b<-a->c - no solution
//  - there resources two parents  - basic with additional resource so b->a<-c - one solution
//  - there resources two parents, one not resolved - basic with additional resource so b->a<-c - no solution
//  - three resources cycle - basic with additional resource so a->b->c->a - no solutions
//  - double basic - two groups of resources (a->b, c->d) - two solutions
func TestEngine_Explicit_Dependencies(t *testing.T) {
	t.Log("Test: Solutions explicit dependencies")
	var tests = []SolutionTest{}
	for _, t := range []SolutionTest{
		{name: "ED-2R-2R-2dS-2SO", resources: files{"ED-minimal", "minimal", "ED-other2", "other"}, systems: files{"other"}, provisioners: files{"other"}, providers: files{"other"}, solutions: "ED-2R-2R-2dS"},
		{name: "ED-3R-cycle-0SO", resources: files{"ED-grampa", "ED-minimal", "ED-cycle3"}, solutions: "empty"},
		{name: "ED-3R-parents-all-1SO", resources: files{"ED-minimal4", "ED-minimal", "minimal"}, solutions: "ED-3R-parents-all"},
		{name: "ED-3R-brothers-one-0SO", resources: files{"ED-minimal3", "minimal", "other"}, solutions: "empty"},
		{name: "ED-3R-brothers-all-1SO", resources: files{"ED-minimal2", "minimal", "minimal2"}, solutions: "ED-3R-brothers-all"},
		{name: "ED-3R-nested-1SO", resources: files{"ED-grampa", "ED-minimal", "minimal"}, solutions: "ED-3R-nested"},
		{name: "ED-basic-1SO", resources: files{"ED-minimal", "minimal"}, solutions: "ED-basic"},
		{name: "ED-cycle-0SO", resources: files{"ED-cycle1", "ED-cycle2"}, solutions: "empty"},
		{name: "ED-2S-4SO", resources: files{"ED-minimal", "minimal"}, systems: files{"minimal2"}, solutions: "ED-basic-2S"},
		{name: "ED-2dR-2dS-0SO", resources: files{"ED-other", "other"}, systems: files{"other"}, provisioners: files{"other"}, providers: files{"other"}, solutions: "empty"},
		{name: "ED-2PD-4SO", resources: files{"ED-minimal", "minimal"}, providers: files{"minimal2"}, solutions: "ED-basic-2PD"},
		{name: "ED-2PR-4SO", resources: files{"ED-minimal", "minimal"}, provisioners: files{"minimal2"}, solutions: "ED-basic-2PR"},
	} {
		tn := SolutionTest{}.update(t.name, "", "create", "", "minimal", "minimal", "minimal", t.solutions)
		tests = append(tests, tn.merge(t.resources, t.providers, t.provisioners, t.systems))
	}
	runSolutionTests(t, tests)
}

// User defined implicit dependencies are cases when resource depends on another defined resource without explicit "connection"
// Example: Two resources will be created Server "s" and Network "n", "s" depends on "n".
//          In some providers, "s" has property "network" that expects a network ID. OpenEngine will able to find solution(s)
//          so the result(response) of network creation will have network ID and it will be passed to "s".
// The following cases are tested:
//  - ... just like explicit dependency tests
//  - no interface - same as basic but without appropriate interface - no solution
func TestEngine_Implicit_Dependencies(t *testing.T) {
	t.Log("Test: Solutions implicit dependencies")
	var tests = []SolutionTest{}
	for _, t := range []SolutionTest{
		{name: "ID-basic-1SO", resources: files{"ID-minimal", "ID-network"}, provisioners: files{"ID-minimal", "ID-network"}, providers: files{"ID-minimal", "ID-network"}, solutions: "ID-basic"},
		{name: "ID-cycle-0SO", resources: files{"ID-cycle1", "ID-cycle2"}, provisioners: files{"ID-cycle"}, providers: files{"ID-cycle"}, solutions: "empty"},
		{name: "ID-2S-4SO", resources: files{"ID-minimal", "ID-network"}, provisioners: files{"ID-minimal", "ID-network"}, providers: files{"ID-minimal", "ID-network"}, systems: files{"minimal2"}, solutions: "ID-basic-2S"},
		{name: "ID-2dR-2dS-0SO", resources: files{"ID-minimal", "ID-network"}, provisioners: files{"ID-minimal", "ID-network2"}, providers: files{"ID-minimal", "ID-network2"}, systems: files{"other"}, solutions: "empty"},
		{name: "ID-2PD-2SO", resources: files{"ID-minimal", "ID-network"}, provisioners: files{"ID-minimal", "ID-network"}, providers: files{"ID-minimal", "ID-network", "ID-minimal2"}, solutions: "ID-basic-2PD"},
		{name: "ID-2PR-2SO", resources: files{"ID-minimal", "ID-network"}, provisioners: files{"ID-minimal", "ID-network", "ID-minimal2"}, providers: files{"ID-minimal", "ID-network"}, solutions: "ID-basic-2PR"},
		{name: "ID-3R-cycle-0SO", resources: files{"ID-3R-cycle1", "ID-3R-cycle2", "ID-3R-cycle3"}, provisioners: files{"ID-cycle"}, providers: files{"ID-cycle"}, solutions: "empty"},
		{name: "ID-3R-brothers-all-1SO", resources: files{"ID-minimal3", "ID-network", "ID-storage"}, provisioners: files{"ID-minimal3", "ID-network", "ID-storage"}, providers: files{"ID-minimal3", "ID-network", "ID-storage"}, solutions: "ID-3R-brothers-all"},
		{name: "ID-3R-brothers-one-0SO", resources: files{"ID-minimal3", "ID-network", "ID-storage"}, provisioners: files{"ID-minimal3", "ID-network"}, providers: files{"ID-minimal3", "ID-network"}, solutions: "empty"},
		{name: "ID-3R-parents-all-1SO", resources: files{"ID-minimal", "ID-minimal2", "ID-network"}, provisioners: files{"ID-minimal", "ID-network"}, providers: files{"ID-minimal", "ID-network"}, solutions: "ID-3R-parents-all"},
		{name: "ID-3R-parents-one-0SO", resources: files{"ID-minimal", "ID-storage2", "ID-network"}, provisioners: files{"ID-minimal", "ID-network", "ID-storage"}, providers: files{"ID-minimal", "ID-network"}, solutions: "empty"},
		{name: "ID-no-basic-0SO", resources: files{"ID-minimal", "ID-network"}, provisioners: files{"ID-minimal", "ID-network"}, providers: files{"ID-no-minimal", "ID-network"}, solutions: "empty"},
	} {
		tn := SolutionTest{}.update(t.name, "", "create", "", "", "", "minimal", t.solutions)
		tests = append(tests, tn.merge(t.resources, t.providers, t.provisioners, t.systems))
	}
	runSolutionTests(t, tests)
}

// Constrains are solutions with additional steps before and\or after the action requests upon a resource.
// The steps might contain actions on resource or running scripts. The additional steps might require their own solutions,
// just like implicit properties. OpenEngine matches constrains and finds solutions for them too.
// the following cases are tested:
// - basic - a working constrain that has resources, scripts, pre and post actions - one solution
// todo:
// - missing dependency - constrain failed as dependent resource doesn't have solution - no solution
// - multiple constrains - multiple constrains applied, *conflicts are out of scope* - one solution
// - disabled - constrains exist in OpenEngine but disabled by resource request - one solution
// - partial disabled - specific constrains are disabled - one solution
// - partial enabled - specific constrains are enabled - one solution
func TestEngine_Constrains(t *testing.T) {
	t.Log("Test: Solutions that require constrains")
	var tests = []SolutionTest{}
	for _, t := range []SolutionTest{
		{name: "C-0S-0S-0C-0R-1SO", providers: files{"C-0S-0S-0C-0R"}, solutions: "C-0S-0S"},                                                        // the specification allows empty list of pre\post actions
		{name: "C-0S-0S-0C-1nmR-1SO", providers: files{"C-0S-0S-0C-1nmR"}, solutions: "C-empty"},                                                    // not matching response
		{name: "C-0S-0S-0C-1mR-1SO", providers: files{"C-0S-0S-0C-1mR"}, solutions: "C-0S-0S"},                                                      // matching response
		{name: "C-0S-0S-1nmC-0R-1SO", resources: files{"explicit"}, providers: files{"C-0S-0S-1nmC-0R"}, solutions: "C-0S-0S-1nmC-0R"},              // not matching properties
		{name: "C-0S-0S-1mC-0R-1SO", resources: files{"explicit"}, providers: files{"C-0S-0S-1mC-0R"}, solutions: "C-0S-0S-1mC-0R"},                 // matching properties
		{name: "C-1S-1S-0C-0R-1SO", providers: files{"C-1S-1S-0C-0R"}, solutions: "C-1S-1S"},                                                        // Simplest script case
		{name: "C-1R-1R-0C-0R-1SO", providers: files{"C-1R-1R-0C-0R", "ID-storage"}, provisioners: files{"ID-storage"}, solutions: "C-1R-1R-0C-0R"}, // Simplest resource case
	} {
		tn := SolutionTest{}.update(t.name, "", "create", "minimal", "", "minimal", "minimal", t.solutions)
		tests = append(tests, tn.merge(t.resources, t.providers, t.provisioners, t.systems))
	}
	runSolutionTests(t, tests)
}
