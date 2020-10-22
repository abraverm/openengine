package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// A Solution is a match of requested resource, given system, provider definition and a provisioner logic.
// nolint: maligned
// TODO: Fix order.
type Solution struct {
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

// A Param is a resolution tree metadata about a solution parameter used in the resolving process.
type Param struct {
	Resolved  bool   `json:"resolved"`
	Tasks     []Task `json:"tasks"`
	ParamType string `json:"param_type"`
}

// A Task is a resolution tree metadata about a solution parameter task used in the resolving process.
type Task struct {
	TaskType     string     `json:"task_type"`
	Resolved     bool       `json:"resolved"`
	Alternatives []Solution `json:"alternatives"`
	Tool         Tool       `json:"tool"`
	Solution     Solution   `json:"solution"`
}

type solutionList []Solution

func (s solutionList) Len() int {
	return len(s)
}

func (s solutionList) Less(i, j int) bool {
	return s[i].size > s[j].size
}

func (s solutionList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Solution) equals(solution Solution) bool {
	provisionerMatch := reflect.DeepEqual(s.Provisioner, solution.Provisioner)
	providerMatch := reflect.DeepEqual(s.Provider, solution.Provider)
	resourceMatch := reflect.DeepEqual(s.Resource, solution.Resource)
	systemMatch := reflect.DeepEqual(s.System, solution.System)

	if provisionerMatch && providerMatch && resourceMatch && systemMatch {
		return true
	}

	return false
}

func (s Solution) inLoop(solution Solution) bool {
	if s.equals(solution) {
		return true
	}

	if s.parent == nil {
		return false
	}

	return s.parent.inLoop(solution)
}

//TODO: Can be improved by using map
func intersect(a []string, b []string) []string {
	set := make([]string, 0)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				set = append(set, a[i])
			}
		}
	}

	return set
}

func (s *Solution) resolveExplicit() []string {
	var implicit []string

	resourceImplicit := s.Resource.getImplicitKeys()
	providerImplicit := s.Provider.getImplicitKeys()

	for param := range s.Provider.Parameters {
		if _, ok := s.Resource.Args[param]; ok { // Explicit
			s.resolutionTree[param] = Param{
				ParamType: "explicit",
				Resolved:  true,
				Tasks:     nil,
			}

			continue
		}

		paramImplicit := intersect(s.Provider.Parameters[param].getImplicitKeys(), providerImplicit)
		if len(intersect(paramImplicit, resourceImplicit)) == len(paramImplicit) { // Supported Implicit
			implicit = append(implicit, param)
		}
	}

	return implicit
}

// Run solution will resolve implicit arguments and execute solution script
// nolint: funlen, gocognit, nestif
// TODO: function is too long and complicatd.
func (s *Solution) Run(solutionArgs map[string]interface{}) (string, error) {
	args := make(map[string]interface{})
	re := regexp.MustCompile(`\$_[[:alpha:]]*`)

	if s.Resource.Args == nil {
		s.Resource.Args = args
	}

	for k, v := range solutionArgs {
		s.Resource.Args[k] = v
	}

	for key, def := range s.resolutionTree {
		if def.ParamType == "explicit" {
			args[key] = s.Resource.Args[key]
		} else {
			taskResults := s.Resource.Args
			for i, task := range def.Tasks {
				var store string
				implicitTask := s.Provider.Parameters[key].Implicit[i]
				if re.MatchString(implicitTask.Store) {
					store = implicitTask.Store[1:]
				} else {
					store = implicitTask.Store
				}
				if task.TaskType == "tool" {
					taskArgs := implicitTask.resolve(taskResults)
					result, err := task.Tool.Run(taskArgs)
					if err != nil {
						return "", err
					}
					taskResults[store] = result
				} else {
					result, err := task.Solution.Run(taskResults)
					if err != nil {
						return "", err
					}
					taskResults[store] = result
				}
			}
			args[key] = taskResults[key]
		}
	}

	file, err := ioutil.TempFile("", "script.*.sh")
	if err != nil {
		return "", fmt.Errorf("solution run failed creating temp file: %w", err)
	}

	defer func() {
		if removeError := os.Remove(file.Name()); removeError != nil {
			err = fmt.Errorf("solution run failed to remove temp file: %w", removeError)
		}
	}()

	tmpl, err := template.ParseFiles(s.Provisioner.Logic)
	if err != nil {
		sJSON, _ := json.MarshalIndent(s, "", "    ")

		return "", fmt.Errorf("solution run failed to parse provisioner logic: %w\n%v", err, string(sJSON))
	}

	if err := tmpl.Execute(file, args); err != nil {
		return "", fmt.Errorf("solution run failed to execut script: %w", err)
	}

	// nolint: gosec
	out, err := exec.Command("/bin/sh", file.Name()).Output()
	if err != nil {
		return string(out), fmt.Errorf("solution run script failed: %w", err)
	}

	return string(out), nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}

	return a
}

// function to build the combinations
// https://groups.google.com/g/golang-nuts/c/UcJ5h0P2jc0
func combIntSlices(seq [][]int) (out [][]int) {
	if len(seq) == 0 {
		return nil
	}

	if len(seq) == 1 {
		return nil
	}
	// fill combSeq with the first slice of seq
	// nolint: prealloc
	var combSeq [][]int // combSeq is the initial [][]int, for example [[1] [2] [3]]
	for _, i := range seq[0] {
		combSeq = append(combSeq, []int{i})
	}

	seq = seq[1:] // seq is the [][]slice to combine with combSeq [[4 5] [6 7 8]]

	// rec recursive function
	var rec func(int, [][]int, [][]int)
	rec = func(i int, seq, combSeq [][]int) {
		var temp [][]int // temporary slice to append combinations

		last := len(seq) - 1

		for _, c := range combSeq { // for each slice in combSeq slice
			for _, s := range seq[i] { // append each element of the slice i in seq
				c1 := append([]int{}, c...)
				c1 = append(c1, s)
				temp = append(temp, c1)
			}

			combSeq = temp // at this step temp has recorded one round of combination
		}

		// nolint: golint
		// TODO: fix lint issue
		if i == last { // if the length of seq is reached, the solution is returned
			out = combSeq

			return
		} else {
			rec(i+1, seq, combSeq) // if the length of seq is not reached, rec is called to perform another step of combinations
		}
	}
	rec(0, seq, combSeq) // start the first step of combinations

	return out
}

// nolint: funlen, prealloc
// TODO: function is too long and complicated.
func (s Solution) decouple() []Solution {
	var (
		decoupled     []Solution
		paramMap      []string
		params        [][]int
		paramTasks    [][][]int
		paramTasksMap [][]int
	)

	placeholder := make(map[string]map[int][]Solution)

	for a, param := range s.resolutionTree { // recursion stop condition
		var (
			tasksMap []int
			tasks    [][]int
		)

		for b, task := range param.Tasks { // recursion stop condition
			if task.TaskType == "resource" { // recursion stop condition
				placeholder[a] = map[int][]Solution{}

				for _, alt := range task.Alternatives {
					for _, decoupledAlt := range alt.decouple() {
						tmp := s
						tmp.resolutionTree[a].Tasks[b].Solution = decoupledAlt
						placeholder[a][b] = append(placeholder[a][b], decoupledAlt)
					}
				}
				// enumerate all the alternatives of a task for combination matching
				tasks = append(tasks, makeRange(0, len(placeholder[a][b])))
				// map the combination position to original task index, see next comment
				tasksMap = append(tasksMap, b)
			}
		}
		// implicit param X has N > 1 Resource type tasks
		// the Resource type task requires a solution by itself, and might be implicit itself
		// thus is each Resource task might have alternative solutions and they are too need to be decoupled
		// combTasks is the all possible combinations of decoupled alternatives of Resource tasks
		// tasksMap restores the original position of combTask set
		// for example task t1,t5 in param X have alternatives {a, b} and {c, d}
		// combTasks : [{a, c}, {a, d}, {b, c}, {b, d}] , tasksMap: [t1, t5]
		combTasks := combIntSlices(tasks)
		// same as "tasks" variable - enumerate all the task combinations for specific param
		// continue the previous example, where X is the first implicit parameter processed:
		params = append(params, makeRange(0, len(combTasks))) // [[0,1,2,3]]
		paramMap = append(paramMap, a)                        // [X]
		paramTasks = append(paramTasks, combTasks)            // [[{a, c}, {a, d}, {b, c}, {b, d}]]
		paramTasksMap = append(paramTasksMap, tasksMap)       // [[t1, t5]]
	}

	combParams := combIntSlices(params)    // same as combTasks: [{X1, Y1}, {X1, Y2}, {X2, Y1}, {X2, Y2}]
	for _, paramComb := range combParams { // {X1,Y1} ...
		decoupledSolution := s // copy the original solution

		for paramID, paramTasksID := range paramComb { // X, 1
			paramName := paramMap[paramID] // "X"

			combTasks := paramTasks[paramTasksID] // [{a, c}, {a, d}, {b, c}, {b, d}]
			for _, taskComb := range combTasks {
				for taskMapID, taskAlt := range taskComb { // 1, a
					taskPos := paramTasksMap[paramID][taskMapID] // t1
					decoupledSolution.resolutionTree[paramName].Tasks[taskPos].Solution = placeholder[paramName][taskPos][taskAlt]
				}
			}
		}

		decoupled = append(decoupled, decoupledSolution)
	}
	// If there is implicit Resource tasks, then there is nothing to decouple and the original
	// solution should be returned - recursion end logic
	if len(decoupled) == 0 {
		decoupled = append(decoupled, s)
	}

	return decoupled
}

// MarshalJSON converts the solution to a JSON.
func (s Solution) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"debug":           s.debug,
		"resolved":        s.resolved,
		"size":            s.size,
		"resolution_tree": s.resolutionTree,
		"action":          s.action,
		"provider":        s.Provider,
		"provisioner":     s.Provisioner,
		"resource":        s.Resource,
		"system":          s.System,
	})
}

// ToJSON converts the solution to a JSON string.
func (s Solution) ToJSON() string {
	data := map[string]interface{}{
		"debug":    s.debug,
		"resolved": s.resolved,
		"size":     s.size,
		"tree":     s.resolutionTree,
		"action":   s.action,
	}

	sJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(sJSON)
}
