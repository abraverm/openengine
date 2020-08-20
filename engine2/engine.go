package engine2

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"reflect"
	"sort"
)

type Engine struct {
	systems      []System
	resources    []Resource
	providers    []Provider
	provisioners []Provisioner
	solutions    []Solution
	tools        []Tool
	resolved     bool
	schedule     []Schedule
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) AddSystem(system System) {
	e.systems = append(e.systems, system)
}

func (e *Engine) AddResource(resource Resource) {
	e.resources = append(e.resources, resource)
}

func (e *Engine) AddProvider(api ProviderAPI) {
	for resourceType, resource := range api {
		for _, provider := range resource.Providers {
			provider.Implicit = resource.Implicit
			provider.Resource = resourceType
			e.providers = append(e.providers, provider)
		}
	}
}

func (e *Engine) AddProvisioner(provisioner Provisioner) {
	e.provisioners = append(e.provisioners, provisioner)
}

func (e *Engine) Match() error {
	/*
	Match engine's resources, systems, providers and provisioners, save the results, a.k.a "solutions"
	for later use. The resource and systems are given "facts" - something that the user has or wishes.
	The matching process finds a provider and a provisioner that support the resource and system together,
	as resource actions are done on a system (create, delete, get, update)
	 */
	for _, resource := range e.resources {
		for _, system := range e.systems {
			solutions, err := e.matchProvidersProvisioners(resource, system)
			if err != nil {
				return err
			}
			e.solutions = append(e.solutions, solutions...)
		}
	}
	return nil
}

func (e Engine) matchProvidersProvisioners(resource Resource, system System) ([]Solution, error) {
	/*
	Resource and System are joined to a single object that transforms to a Json document, same thing
	happens with each provider and provisioner. The provider and provisioner document has the structure
	of Json Schema which validates the resource and system document. Successful validation means all
	parties match. If the resource has implicit parameter, then the provisioner trusts the provider if
	implicit parameter fulfils the explicit parameter with the same name as the provisioner allows.
	The trust works by using the Json Schema reference functionality.
	 */
	var solutions []Solution
	data := gojsonschema.NewGoLoader(map[string]interface{}{
		"resource": resource,
		"system":   system,
	})
	for _, provider := range e.providers {
		for _, provisioner := range e.provisioners {
			pnpSchema := JSONSchema{
				"type": "object",
				"allOf": []JSONSchema{
					provider.toJsonSchema(),
					provisioner.toJsonSchema(),
				},
			}
			loader := gojsonschema.NewGoLoader(pnpSchema)
			result, err := gojsonschema.Validate(loader, data)
			if err != nil {
				return nil, err
			}
			if result.Valid() {
				solutions = append(solutions, Solution{
					resource:    resource,
					system:      system,
					provider:    provider,
					provisioner: provisioner,
				})
			}
		}
	}
	return solutions, nil
}

func (e *Engine) Resolve() {
	/*
	Resolve engine's solutions dependencies of implicit parameters. The dependencies might be tools or other resources.
	In case of resources, other solutions are needed, and might be more than one alternative. The dependent solutions
	are also resolved recursively. Unresolved solutions are removed from engine's solutions list.
	 */
	var solutions []Solution
	for _, solution := range e.solutions {
		newSolution := e.resolveDependencies(solution)
		if newSolution.resolved {
			solutions = append(solutions, newSolution)
		}
	}
	e.solutions = solutions
}

func (e Engine) resolveDependencies(solution Solution) Solution {
	/*
	Resolving dependencies of a solution is a recursive process that identifies if the parameters are explicit or
	implicit, if the implicit task is fulfilled by a tool or another resource, finds new solutions for dependent
	resource and resolves its dependencies too. The process might find multiple solutions for implicit task and saves
	them as alternative for later use in the scheduling process. The process eliminates loops and unresolved solutions.
	The recursion ends when a solution parameters are all explicit or implicit with only tools are used.
	 */
	solutionResolved := true
	for _, param := range solution.resolveExplicit() {
		var tasks []Task
		resolved := true
		for _, task := range solution.provider.Parameters[param].Implicit {
			if e.isTool(task) {
				tasks = append(tasks, Task{
					taskType: "tool",
					resolved: true,
					tool:     e.getTool(task),
				})
			} else {
				matches, _ := e.matchProvidersProvisioners(task.Resource, solution.system)
				var alternatives []Solution
				for _, match := range matches {
					match.parent = &solution
					if solution.inLoop(match) {
						continue
					}
					match = e.resolveDependencies(match)
					if match.resolved {
						alternatives = append(alternatives, match)
					}
				}
				taskResolved := true
				if len(alternatives) == 0 {
					solutionResolved = false
					resolved = false
					taskResolved = false
				}
				tasks = append(tasks, Task{
					taskType:     "resource",
					resolved:     taskResolved,
					alternatives: alternatives,
				})
			}
		}
		solution.resolutionTree[param] = Param{
			paramType: "implicit",
			resolved:  resolved,
			tasks:     tasks,
		}
	}
	solution.resolved = solutionResolved
	return solution
}

func (e Engine) isTool(task ImplicitTask) bool {
	// Checks if the given task has a tool that matches its description
	for _, tool := range e.tools {
		if tool.match(task) {
			return true
		}
	}
	return false
}

func (e Engine) getTool(task ImplicitTask) Tool {
	// Gets a tool that matches the Implicit task
	for _, tool := range e.tools {
		if tool.match(task) {
			return tool
		}
	}
	return Tool{}
}

func (e *Engine) Schedule(action string) error {
	/*
	For all requested resources and given action, find solutions that can fulfil the request and order them by size.
	Size of a solution is number of its dependent solutions.
	 */
	for _, resource := range e.resources {
		var solutions []Solution
		for _, solution := range e.solutions {
			if reflect.DeepEqual(resource, solution.resource) && solution.action == action {
				solutions = append(solutions, solution)
			}
		}
		if len(solutions) == 0 {
			return fmt.Errorf("no solution found for resource")
		}
		var decoupled []Solution
		for _, solution := range solutions {
			decoupled = append(decoupled, solution.decouple()...)
		}
		sort.Sort(solutionList(decoupled))
		e.schedule = append(e.schedule, Schedule{
			resource:  resource,
			solutions: decoupled,
		})
	}
	return nil
}

func (e Engine) Run() ([]string, error) {
	/*
	Engine will run the scheduled solutions and tries the alternatives when needed.
	 */
	var results []string
	failed := false
	OUTER:
	for _, schedule := range e.schedule {
		for _, solution := range schedule.solutions {
			if result, err := solution.Run(map[string]string{}); err == nil {
				results = append(results, result)
				continue OUTER
			}
		}
		failed = true
		break
	}
	if failed {
		return nil, fmt.Errorf("failed to provision resource")
	}
	return results, nil
}
