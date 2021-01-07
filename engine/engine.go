// Package engine is the core of OpenEngine that generates solutions
package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/encoding/gocode/gocodec"

	//  to register builtins
	_ "cuelang.org/go/pkg"
	"golang.org/x/xerrors"
)

// An Engine is the OpenEngine interface - all actions should be done using it.
type Engine struct {
	cue struct {
		runtime  *cue.Runtime
		instance *cue.Instance
		codec    *gocodec.Codec
		spec     string
	}
}

func loadFile(path string) (string, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}

	if info.IsDir() {
		return "", xerrors.Errorf("Found directory instead of file: %s", path)
	}

	file, err := ioutil.ReadFile(path) // nolint: gosec
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func (e *Engine) loadSpec(path string) error {
	switch path {
	case "":
		if err := e.addDefinition("spec.cue", spec); err != nil {
			return err
		}

	default:
		file, err := loadFile(path)
		if err != nil {
			return err
		}

		if err := e.addDefinition("spec.cue", file); err != nil {
			return err
		}
	}

	return nil
}

// NewEngine creates an Engine and initialize it.
func NewEngine(spec string) (*Engine, error) {
	e := Engine{}
	e.cue.runtime = &cue.Runtime{}

	if err := e.loadSpec(spec); err != nil {
		return nil, err
	}

	e.cue.codec = gocodec.New(e.cue.runtime, nil)

	return &e, nil
}

func (e *Engine) addDefinition(path string, definition string) error {
	r := e.cue.runtime

	instance, err := r.Compile(path, e.cue.spec+definition)
	if err != nil {
		return err
	}

	e.cue.instance = instance
	e.cue.runtime = r
	e.cue.spec += definition

	return err
}

func (e Engine) validateValue(value cue.Value, def string) (err error) {
	var x interface{}

	defValue := e.cue.instance.LookupDef(def)
	_ = e.cue.codec.Encode(value, &x)
	err = e.cue.codec.Validate(defValue, x)

	if _, err := defValue.Unify(value).MarshalJSON(); err != nil {
		return err
	}

	return
}

func (e Engine) validateRaw(path, def, content string) error {
	r := cue.Runtime{}

	instance, err := r.Compile(path, content)
	if err != nil {
		return err
	}

	return e.validateValue(instance.Value(), def)
}

// Add cue definition to engine.
func (e *Engine) Add(path, def, content string) error {
	if len(content) == 0 {
		return xerrors.New("Content is empty")
	}

	defType := fmt.Sprintf("#%s", def)
	if err := e.validateRaw(path, defType, content); err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(content)) // nolint: errcheck, gosec
	sha := hex.EncodeToString(h.Sum(nil))
	source := fmt.Sprintf("\n%ss:\"%s\":%s\n", strings.ToLower(def), sha, content)
	e.addDefinition(path, source) // nolint: errcheck, gosec

	return nil
}

// System is a provider instance that contains matching values and other metadata such as credentials.
type System struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}

// Resource is the user requested resource with its type and parameters.
type Resource struct {
	Name         string                 `json:"type"`
	Properties   map[string]interface{} `json:"properties"`
	Dependencies map[string]interface{}
}

func (e Engine) addObject(obj interface{}, def string) error {
	defType := fmt.Sprintf("#%s", def)
	defValue := e.cue.instance.LookupDef(defType)

	objValue, _ := e.cue.codec.Decode(obj)
	if err := defValue.Unify(objValue).Validate(); err != nil {
		return err
	}

	objCue, _ := format.Node(objValue.Syntax())

	return e.Add("i wonder", def, string(objCue))
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
	actionValue := e.cue.instance.Lookup("ACTION")
	if err := e.cue.codec.Validate(actionValue, action); err != nil {
		return "[]", err
	}

	instance, _ := e.cue.instance.Fill(action, "ACTION")
	instanceCue, _ := format.Node(instance.Value().Syntax(cue.All()))
	solutions := instance.Lookup("DependencyGroupsSolutionsDecoupled")

	result, err := format.Node(solutions.Syntax(cue.Concrete(true)))
	if err != nil {
		return "", err
	}

	if strings.Contains(string(result), "_|_") {
		return string(instanceCue), xerrors.New(string(result))
	}

	return string(result), nil
}
