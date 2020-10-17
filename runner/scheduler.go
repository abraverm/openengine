package runner

import (
	"reflect"
	"sort"

	"github.com/abraverm/openengine/engine"
)

// Schedule type.
type Schedule struct {
	Solutions []engine.Solution `json:"solutions"`
}

// Scheduler interface.
type Scheduler interface {
	Schedule(resource engine.Resource, action string) Schedule
}

// ResourceNumScheduler is a scheduler based on size of solutions to decide their order.
// Solution size is number of solution in its resolution tree (result of required implicit processes).
type ResourceNumScheduler struct {
	Solutions []engine.Solution `json:"solutions"`
}

func solutionSize(solution engine.Solution) int {
	size := 1

	for _, param := range solution.ResolutionTree {
		if param.ParamType == "implicit" {
			for _, task := range param.Tasks {
				if task.TaskType == "resource" {
					size += solutionSize(task.Solution)
				}
			}
		}
	}

	return size
}

// Len is number of solutions, sort Interface implementation.
func (r ResourceNumScheduler) Len() int { return len(r.Solutions) }

// Less compares size of two solutions, sort Interface implementation.
func (r ResourceNumScheduler) Less(i, j int) bool {
	return solutionSize(r.Solutions[i]) > solutionSize(r.Solutions[j])
}

// Swap two resources position by given index, sort Interface implementation.
func (r ResourceNumScheduler) Swap(i, j int) {
	r.Solutions[i], r.Solutions[j] = r.Solutions[j], r.Solutions[i]
}

// Schedule of solutions for given resource and action based, ordered by solution size - number of required solutions.
func (r ResourceNumScheduler) Schedule(resource engine.Resource, action string) Schedule {
	var solutions []engine.Solution

	for _, solution := range r.Solutions {
		if reflect.DeepEqual(resource, solution.Resource) && solution.Action == action {
			solutions = append(solutions, solution)
		}
	}

	tmp := ResourceNumScheduler{solutions}

	sort.Sort(tmp)

	return (Schedule)(tmp)
}
