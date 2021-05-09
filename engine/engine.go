// Package engine is the core of OpenEngine that generates solutions
// nolint: wrapcheck
package engine

import (
	"crypto/sha256"

	// no other use than loading the default spec.
	_ "embed"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"

	//  to register builtins
	"golang.org/x/xerrors"
)

// SPEC is global variable to store OpenEngine default specifications
//go:embed spec.cue
// nolint
var SPEC string

// An Engine is the OpenEngine interface - all actions should be done using it.
type Engine struct {
	context *cue.Context
	value   cue.Value
}

func loadFile(path string) (string, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("file %s doesn't exist: %w", path, err)
	}

	if info.IsDir() {
		return "", xerrors.Errorf("Found directory instead of file: %s", path)
	}

	file, err := ioutil.ReadFile(path) // nolint: gosec
	if err != nil {
		return "", fmt.Errorf("unable to read file %s: %w", path, err)
	}

	return string(file), nil
}

func (e *Engine) loadSpec(path string) error {
	switch path {
	case "":
		if err := e.addDefinition(SPEC); err != nil {
			return err
		}

	default:
		file, err := loadFile(path)
		if err != nil {
			return err
		}

		if err := e.addDefinition(file); err != nil {
			return err
		}
	}

	return nil
}

// NewEngine creates an Engine and initialize it.
func NewEngine(spec string) (*Engine, error) {
	//TODO: there is no need for the sub struct "cue"
	// nolint
	e := Engine{context: cuecontext.New()}
	if err := e.loadSpec(spec); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Engine) addDefinition(source string) error {
	v := e.value.Unify(e.context.CompileString(source))

	if v.Err() != nil {
		return fmt.Errorf("bad definition %s: %w", source, v.Err())
	}

	e.value = v

	return nil
}

// System is a provider instance that contains matching values and other metadata such as credentials.
type System struct {
	Type       string                 `json:"type,omitempty"`
	Properties map[string]interface{} `json:"properties"`
}

// Resource is the user requested resource with its type and parameters.
type Resource struct {
	Type                   string                 `json:"type"`
	Name                   string                 `json:"name,omitempty"`
	Properties             map[string]interface{} `json:"properties"`
	System                 System                 `json:"system,omitempty"`
	Dependencies           []string               `json:"dependencies,omitempty"`
	InterfacesDependencies []string               `json:"interfacesDependencies,omitempty"`
}

func (e Engine) verifyValue(def, content string) error {
	v := e.context.CompileString(content)
	if v.Err() != nil {
		return v.Err()
	}

	defV := e.value.LookupPath(cue.ParsePath(fmt.Sprintf("#%s", def)))

	vTest := defV.Unify(v)
	if vTest.Err() != nil {
		return vTest.Err()
	}

	if _, err := vTest.MarshalJSON(); err != nil {
		return err
	}

	return nil
}

// Add cue definition to engine.
func (e *Engine) Add(def, content string) error {
	if len(content) == 0 {
		return xerrors.New("Content is empty")
	}

	re := regexp.MustCompile(`"_(.*)":`)
	content = re.ReplaceAllString(content, "_${1}:")

	if err := e.verifyValue(def, content); err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(content)) // nolint: gosec
	sha := hex.EncodeToString(h.Sum(nil))
	source := fmt.Sprintf("\n%ss:\"%s\":%s\n", strings.ToLower(def), sha, content)

	return e.addDefinition(source)
}

func (e *Engine) addObject(obj interface{}, def string) error {
	v := e.context.Encode(obj)
	if v.Err() != nil {
		return v.Err()
	}

	source, _ := format.Node(v.Syntax())

	return e.Add(def, string(source))
}

// AddSystem adds system to Engine.
func (e *Engine) AddSystem(system System) error {
	return e.addObject(system, "System")
}

// AddResource adds resource to Engine.
func (e *Engine) AddResource(resource Resource) error {
	return e.addObject(resource, "Resource")
}

// Solutions is the Engine purpose.
func (e *Engine) Solutions(action string) (results string, err error) {
	v := e.value.FillPath(cue.ParsePath("ACTION"), action)
	if v.Err() != nil {
		return "[]", v.Err()
	}

	vCue, _ := format.Node(v.Syntax(cue.All()))
	solutions := v.LookupPath(cue.ParsePath("DependencyGroupsSolutionsDecoupled"))

	if solutions.Err() != nil {
		return string(vCue), solutions.Err()
	}

	result, _ := format.Node(solutions.Syntax(cue.Concrete(true)))

	reEmpty := regexp.MustCompile("(?m)[\r\n]+^.*: (\\[]|{})")
	empty := reEmpty.ReplaceAllString(string(result), "")
	reSystem := regexp.MustCompile("(?m)System:")
	res := reSystem.ReplaceAllString(empty, "system:")

	return res, nil
}

// GetSpec return engine current specification for debugging.
func (e Engine) GetSpec() string {
	spec, _ := format.Node(e.value.Syntax(cue.All()))

	return string(spec)
}
