package engine

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"text/template"
)

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

type Task struct {
	taskType string
	resolved bool
	alternatives []Solution
	tool Tool
	solution Solution
}

type Param struct {
	resolved  bool
	tasks     []Task
	paramType string
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
	} else {
		return false
	}
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

func intersect(a []string, b []string) []string {
	set := make([]string, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j ++ {
			if 	a[i] == b[j] {
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
				paramType: "explicit",
				resolved: true,
				tasks:    nil,
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

func (s Solution) Run(solutionArgs map[string]interface{}) (string, error) {
	var args = make(map[string]interface{})
	re := regexp.MustCompile(`\$_[[:alpha:]]*`)
	if s.Resource.Args == nil {
		s.Resource.Args = args
	}
	for k, v := range solutionArgs {
		s.Resource.Args[k] = v
	}
	for key, def := range s.resolutionTree {
		if def.paramType == "explicit" {
			args[key] = s.Resource.Args[key]
		} else {
			taskResults := s.Resource.Args
			for i, task := range def.tasks {
				var store string
				implicitTask := s.Provider.Parameters[key].Implicit[i]
				if re.MatchString(implicitTask.Store) {
					store = implicitTask.Store[1:]
				} else {
					store = implicitTask.Store
				}
				if task.taskType == "tool" {
					taskArgs := implicitTask.resolve(taskResults)
					result, err := task.tool.Run(taskArgs)
					if err != nil {
						return "", err
					}
					taskResults[store] = result
				} else {
					result, err := task.solution.Run(taskResults)
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
		return "", fmt.Errorf("solution run failed creating temp file: %v", err)
	}
	defer func() {
		removeError := os.Remove(file.Name())
		if err == nil {
			err = fmt.Errorf("solution run failed to remove temp file: %v", removeError)
		}
	}()
	tmpl, err := template.ParseFiles(s.Provisioner.Logic)
	if err != nil {
		sJson, _ := json.MarshalIndent(s, "", "    ")
		return "", fmt.Errorf("solution run failed to parse provisioner logic: %v\n%v", err, string(sJson))
	}
	if err := tmpl.Execute(file, args); err != nil {
		return "", fmt.Errorf("solution run failed to execut script: %v", err)
	}
	out, err := exec.Command("/bin/sh", file.Name()).Output()
	if err != nil {
		return string(out), fmt.Errorf("solution run script failed: %v", err)
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

func (s Solution) decouple() []Solution {
	var decoupled []Solution
	var placeholder = make(map[string]map[int][]Solution)
	var paramMap []string
	var params [][]int
	var paramTasks [][][]int
	var paramTasksMap [][]int

	for a, param := range s.resolutionTree { // recursion stop condition
		var tasksMap []int
		var tasks [][]int
		for b, task := range param.tasks {   // recursion stop condition
			if task.taskType == "Resource" { // recursion stop condition
				placeholder[a] = map[int][]Solution{}
				for _, alt := range task.alternatives {
					for _, decoupledAlt := range alt.decouple(){
						tmp := s
						tmp.resolutionTree[a].tasks[b].solution = decoupledAlt
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
		params = append(params, makeRange(0,len(combTasks))) // [[0,1,2,3]]
		paramMap = append(paramMap, a)							  // [X]
		paramTasks = append(paramTasks, combTasks) 				  // [[{a, c}, {a, d}, {b, c}, {b, d}]]
		paramTasksMap = append(paramTasksMap, tasksMap)           // [[t1, t5]]
	}
	combParams := combIntSlices(params) // same as combTasks: [{X1, Y1}, {X1, Y2}, {X2, Y1}, {X2, Y2}]
	for _, paramComb := range combParams { // {X1,Y1} ...
		decoupledSolution := s // copy the original solution
		for paramId, paramTasksId := range paramComb { // X, 1
			paramName := paramMap[paramId] // "X"
			combTasks := paramTasks[paramTasksId] // [{a, c}, {a, d}, {b, c}, {b, d}]
			for _, taskComb := range combTasks {
				for taskMapId, taskAlt := range taskComb{ // 1, a
					taskPos := paramTasksMap[paramId][taskMapId]  // t1
					decoupledSolution.resolutionTree[paramName].tasks[taskPos].solution = placeholder[paramName][taskPos][taskAlt]
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

func (s Solution)  MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
	Debug       bool
	Resolved    bool
	Size        int
	Tree        map[string]Param
	Action      string
	Provider    Provider
	Provisioner Provisioner
	Resource    Resource
	System      System
}{
		Debug: s.debug,
		Resolved: s.resolved,
		Size: s.size,
		Tree: s.resolutionTree,
		Action: s.action,
		Provider: s.Provider,
		Provisioner: s.Provisioner,
		Resource: s.Resource,
		System: s.System,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (s Solution) ToJson() string {
	data := map[string]interface{}{
		"debug": s.debug,
		"resolved": s.resolved,
		"size": s.size,
		"tree": s.resolutionTree,
		"action": s.action,
	}
	sJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(sJSON)
}