// Package engine is for finding solutions for given resources, systems, providers, provisioners and tools
package engine

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qri-io/jsonschema"
	"golang.org/x/xerrors"
)

// An Engine is the OpenEngine interface - all actions should be done using it.
type Engine struct {
	Systems      []System      `json:"systems"`
	Resources    []Resource    `json:"resources"`
	Providers    []Provider    `json:"providers"`
	Provisioners []Provisioner `json:"provisioners"`
	Solutions    []Solution    `json:"solutions"`
	Tools        []Tool        `json:"tools"`
}

// NewEngine creates a new engine instance.
func NewEngine() *Engine {
	return &Engine{}
}

// AddSystem will add a system to the engine.
func (e *Engine) AddSystem(system System) {
	e.Systems = append(e.Systems, system)
}

// AddResource will add a resource to the engine.
func (e *Engine) AddResource(resource Resource) {
	e.Resources = append(e.Resources, resource)
}

// AddProvider will add a provider to the engine.
func (e *Engine) AddProvider(api ProviderAPI) {
	for resourceType, resource := range api {
		for _, provider := range resource.Providers {
			provider.Implicit = resource.Implicit
			provider.Resource = resourceType
			e.Providers = append(e.Providers, provider)
		}
	}
}

// AddProvisioner will add a provisioner to the engine.
func (e *Engine) AddProvisioner(provisioner Provisioner) {
	e.Provisioners = append(e.Provisioners, provisioner)
}

// AddTool will add a tool to the engine.
func (e *Engine) AddTool(api ToolAPI) {
	for name, tool := range api {
		tool.Name = name
		e.Tools = append(e.Tools, tool)
	}
}

// Match engine's resources, systems, providers and provisioners, save the results, a.k.a "solutions"
// for later use. The Resource and systems are given "facts" - something that the user has or wishes.
// The matching process finds a Provider and a Provisioner that support the Resource and System together,
// as Resource actions are done on a System (create, delete, get, update).
func (e *Engine) Match() error {
	jsonschema.LoadDraft2019_09()
	jsonschema.RegisterKeyword("oeProperties", NewOeProperties)
	jsonschema.RegisterKeyword("oeRequired", NewOeRequired)

	for _, resource := range e.Resources {
		for _, system := range e.Systems {
			solutions, err := e.matchProvidersProvisioners(resource, system)
			if err != nil {
				return err
			}

			e.Solutions = append(e.Solutions, solutions...)
		}
	}

	return nil
}

// matchProvidersProvisioners is the magic behind OpenEngine.
// Resource and System are joined to a single object that transforms to a Json document, same thing
// happens with each Provider and Provisioner. The Provider and Provisioner document has the structure
// of Json Schema which validates the Resource and System document. Successful validation means all
// parties match. If the Resource has implicit parameter, then the Provisioner trusts the Provider if
// implicit parameter fulfils the explicit parameter with the same name as the Provisioner allows.
// The trust works by using the Json Schema reference functionality.
// nolint: funlen
// TODO: function is too long.
func (e Engine) matchProvidersProvisioners(resource Resource, system System) ([]Solution, error) {
	var solutions []Solution

	ctx := context.Background()
	dataRaw := map[string]interface{}{
		"Resource": resource,
		"System":   system,
	}

	for _, provider := range e.Providers {
		for _, provisioner := range e.Provisioners {
			pnpSchema := Schema{
				"$id":   "engine.json",
				"$defs": provider.toJSONSchemaDefs(),
				"type":  "object",
				"allOf": []Schema{
					provider.toJSONSchema(),
					provisioner.toJSONSchema(),
				},
			}

			dJSON, _ := json.MarshalIndent(dataRaw, "", "  ")
			sJSON, _ := json.MarshalIndent(pnpSchema, "", "  ")
			errors := fmt.Sprintf("Data:\n%v\nSchema:\n%v\n", string(dJSON), string(sJSON))
			loader := new(jsonschema.Schema)

			if err := json.Unmarshal(sJSON, loader); err != nil {
				return nil, xerrors.Errorf("unmarshal schema: %v\n%v", err.Error(), string(sJSON))
			}

			errs, err := loader.ValidateBytes(ctx, dJSON)
			if err != nil {
				return nil, err
			}

			if len(errs) > 0 && (provisioner.Debug && provider.Debug) {
				for _, err := range errs {
					vJSON, _ := json.MarshalIndent(err.InvalidValue, "", "  ")
					errors = fmt.Sprintf("%v\n%v at %v:\n%v", errors, err.Message, err.PropertyPath, string(vJSON))
				}

				solutions = append(solutions, Solution{
					Resource:    resource,
					System:      system,
					Provider:    provider,
					Provisioner: provisioner,
					debug:       provider.Debug || provisioner.Debug,
					Output:      errors,
					Action:      provider.Action,
				})

				continue
			} else if len(errs) == 0 {
				solutions = append(solutions, Solution{
					Resource:    resource,
					System:      system,
					Provider:    provider,
					Provisioner: provisioner,
					debug:       provider.Debug || provisioner.Debug,
					Output:      errors,
					Action:      provider.Action,
				})
			}
		}
	}

	return solutions, nil
}

// Resolve engine's solutions dependencies of implicit parameters. The dependencies might be tools or other resources.
// In case of resources, other solutions are needed, and might be more than one alternative. The dependent solutions
// are also resolved recursively. Unresolved solutions are removed from engine's solutions list.
func (e *Engine) Resolve() {
	solutions := []Solution{}

	for _, solution := range e.Solutions {
		newSolution := e.resolveDependencies(solution)
		if newSolution.resolved || newSolution.debug {
			solutions = append(solutions, newSolution)
		}
	}

	decoupled := make([]Solution, 0, len(solutions))
	for _, solution := range solutions {
		decoupled = append(decoupled, solution.decouple()...)
	}

	e.Solutions = decoupled
}

// nolint: funlen, nestif
// TODO: function is too long and complicated.
/*
	Resolving dependencies of a solution is a recursive process that identifies if the parameters are explicit or
	implicit, if the implicit task is fulfilled by a tool or another Resource, finds new solutions for dependent
	Resource and resolves its dependencies too. The process might find multiple solutions for implicit task and saves
	them as alternative for later use in the scheduling process. The process eliminates loops and unresolved solutions.
	The recursion ends when a solution parameters are all explicit or implicit with only tools are used.
*/
func (e Engine) resolveDependencies(solution Solution) Solution {
	solutionResolved := true
	solution.ResolutionTree = make(map[string]Param)
	// resolveExplicit populates resolutionTree with explicit params and returns implicit params to be handled here
	for _, param := range solution.resolveExplicit() {
		var tasks []Task

		resolved := true

		for _, task := range solution.Provider.Parameters[param].Implicit {
			if task.Type == "tool" {
				tool, err := e.getTool(task)
				if err != nil {
					solution.Output = fmt.Sprint(err)
					resolved = false
				} else {
					tasks = append(tasks, Task{
						TaskType:     "tool",
						Resolved:     true,
						Tool:         tool,
						ImplicitTask: task,
					})
				}
			} else {
				resource := Resource{
					Name: task.Name,
					Args: task.Args,
				}
				matches, _ := e.matchProvidersProvisioners(resource, solution.System)
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
					TaskType:     "resource",
					Resolved:     taskResolved,
					Alternatives: alternatives,
					ImplicitTask: task,
				})
			}
		}

		solution.ResolutionTree[param] = Param{
			ParamType: "implicit",
			Resolved:  resolved,
			Tasks:     tasks,
		}
	}

	solution.resolved = solutionResolved

	return solution
}

// Gets a tool that matches the Implicit task.
func (e Engine) getTool(task ImplicitTask) (Tool, error) {
	for _, tool := range e.Tools {
		if tool.Name == task.Name {
			return tool, nil
		}
	}

	return Tool{}, xerrors.Errorf("tool %v not found", task.Name)
}

// GetSolutions returns engine current solutions.
func (e *Engine) GetSolutions() []Solution {
	return e.Solutions
}
