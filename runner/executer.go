// Package runner is defines the interfaces and different methods for scheduling solutions and their execution
package runner

import (
	"fmt"
	"io/ioutil"
	"openengine/engine"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/imdario/mergo"
	"golang.org/x/xerrors"
)

// Runner interface standardizes the execution process no matter how it is actually computed.
type Runner interface {
	ResourceRun(solution engine.Solution, args map[string]interface{}) (string, error)
	ToolRun(tool engine.Tool, args map[string]interface{}) (string, error)
	Schedules() []Schedule
}

func variableKeyName(variable string) string {
	key := variable

	re := regexp.MustCompile(`\$_[[:alpha:]]*`)
	if re.MatchString(variable) {
		key = variable[1:]
	}

	return key
}

func updateArgs(task engine.ImplicitTask, values map[string]interface{}) map[string]interface{} {
	args := make(map[string]interface{})

	for paramName, paramValue := range task.Args {
		result := fmt.Sprint(paramValue)
		for resolvedParam, resolvedValue := range values {
			result = strings.Replace(
				result,
				fmt.Sprintf("$%v", resolvedParam),
				fmt.Sprint(resolvedValue),
				-1,
			)
		}

		args[paramName] = result
	}

	return args
}

// nolint: lll
func resolveImplicitArg(runner Runner, tasks []engine.Task, args map[string]interface{}, key string) (interface{}, error) {
	for _, task := range tasks {
		name := variableKeyName(task.ImplicitTask.Store)

		if task.TaskType == "tool" {
			taskArgs := updateArgs(task.ImplicitTask, args)

			result, err := runner.ToolRun(task.Tool, taskArgs)
			if err != nil {
				return nil, err
			}

			args[name] = result
		} else {
			result, err := runner.ResourceRun(task.Solution, args)
			if err != nil {
				return nil, err
			}

			args[name] = result
		}
	}

	value, ok := args[key]
	if !ok {
		return "", xerrors.Errorf("key '%v' not found", key)
	}

	return value, nil
}

func resolveArgs(runner Runner, solution engine.Solution, args map[string]interface{}) (map[string]interface{}, error) {
	mergedArgs := make(map[string]interface{})
	if err := mergo.Merge(&mergedArgs, solution.Resource.Args); err != nil {
		return nil, err
	}

	if err := mergo.Merge(&mergedArgs, args, mergo.WithOverride); err != nil {
		return nil, err
	}

	for key, def := range solution.ResolutionTree {
		if def.ParamType == "explicit" {
			mergedArgs[key] = solution.Resource.Args[key]
		} else {
			value, err := resolveImplicitArg(runner, def.Tasks, mergedArgs, key)
			if err != nil {
				return nil, err
			}
			mergedArgs[key] = value
		}
	}

	return mergedArgs, nil
}

// Run executes the given runner.
func Run(runner Runner) ([]string, error) {
	var (
		results []string
		errors  []string
	)

	failed := false

OUTER:
	for _, schedule := range runner.Schedules() {
		for _, solution := range schedule.Solutions {
			args, err := resolveArgs(runner, solution, map[string]interface{}{})
			if err != nil {
				errors = append(errors, fmt.Sprint(err))

				continue
			}
			if result, err := runner.ResourceRun(solution, args); err == nil {
				results = append(results, result)

				continue OUTER
			} else {
				errors = append(errors, fmt.Sprint(err))
			}
		}
		failed = true

		break
	}

	if failed {
		return nil, xerrors.Errorf("failed to provision Resource:\n%+v\nresults:\n%v", errors, results)
	}

	return results, nil
}

// LocalRunner implements Runner interface to execute on local shell (bash).
type LocalRunner struct {
	Scheduler
	Action    string            `json:"action"`
	Solutions []engine.Solution `json:"solutions"`
	Resources []engine.Resource `json:"resources"`
}

// NewLocalRunner creates a new LocalRunner
// TODO: is it a go best practice?
func NewLocalRunner(e engine.Engine, action string, scheduler Scheduler) LocalRunner {
	return LocalRunner{
		Scheduler: scheduler,
		Action:    action,
		Solutions: e.Solutions,
		Resources: e.Resources,
	}
}

func renderTemplate(templateFile string, args map[string]interface{}) (string, error) {
	file, err := ioutil.TempFile("", "script.*.sh")
	if err != nil {
		return "", fmt.Errorf("failed creating temp file: %w", err)
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w\n%v", err, templateFile)
	}

	if err := tmpl.Execute(file, args); err != nil {
		return "", fmt.Errorf("failed to fill parameters into template: %w", err)
	}

	return file.Name(), nil
}

func shell(scriptTemplate string, args map[string]interface{}) (string, error) {
	file, err := renderTemplate(scriptTemplate, args)
	if err != nil {
		return "", err
	}

	defer func() {
		if removeError := os.Remove(file); removeError != nil {
			err = fmt.Errorf("failed to remove temp file: %w", removeError)
		}
	}()

	// nolint: gosec
	out, err := exec.Command("/bin/sh", file).Output()
	if err != nil {
		return string(out), fmt.Errorf("failed shell script: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}

// ResourceRun implements the Runner interface for resource execution.
func (l LocalRunner) ResourceRun(resource engine.Solution, args map[string]interface{}) (string, error) {
	return shell(resource.Provisioner.Logic, args)
}

// ToolRun implements the Runner interface for tool execution.
func (l LocalRunner) ToolRun(tool engine.Tool, args map[string]interface{}) (string, error) {
	return shell(tool.Script, args)
}

// Schedules implements the Runner interface for generating Schedules.
func (l LocalRunner) Schedules() []Schedule {
	schedules := make([]Schedule, 0, len(l.Resources))

	for _, resource := range l.Resources {
		schedule := l.Scheduler.Schedule(resource, l.Action)
		schedules = append(schedules, schedule)
	}

	return schedules
}
