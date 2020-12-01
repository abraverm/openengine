// Package dsl provides all processing functionality for the DSL (YAML parsing)
package dsl

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"openengine/engine"
	"openengine/runner"
	"openengine/util"
	"path/filepath"

	"github.com/goccy/go-yaml"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// DSL is the result of processing the ie dsl file and manages the engine operations.
type DSL struct {
	API          []string          `yaml:"api"`
	Provisioners []string          `yaml:"provisioners"`
	Systems      []engine.System   `yaml:"systems"`
	Tools        []string          `yaml:"tools"`
	Resources    []engine.Resource `yaml:"resources"`
	Engine       engine.Engine
}

//New will returns a DSL struct with the input parameters.
//func New(api []string, provisioners []string, systems []engine.System,
//tools []string, resources []engine.Resource, engine engine.Engine) DSL {
//return DSL{
//API:          api,
//Provisioners: provisioners,
//Systems:      systems,
//Tools:        tools,
//Resources:    resources,
//Engine:       engine,
//}
//}

func getSource(uri string) ([]byte, error) {
	var data []byte

	urlParsed, _ := url.ParseRequestURI(uri)

	switch {
	case urlParsed != nil:
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, uri, nil)

		client := http.Client{}

		res, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("unable to download %w", err)
		}

		defer func() {
			if closeErr := res.Body.Close(); closeErr != nil {
				err = closeErr
			}
		}()

		data, _ = ioutil.ReadAll(res.Body)

	case util.FileExists(uri):
		data, err := ioutil.ReadFile(filepath.Clean(uri))
		if err != nil {
			return nil, fmt.Errorf("unable to read %w", err)
		}

		return data, nil
	default:
		return nil, xerrors.Errorf("source was not found (neither a relative file or URL)")
	}

	return data, nil
}

func (d *DSL) addSystem(e *engine.Engine) {
	for _, system := range d.Systems {
		e.AddSystem(system)
	}
}

func (d *DSL) addResource(e *engine.Engine) {
	for _, resource := range d.Resources {
		e.AddResource(resource)
	}
}

func (d *DSL) addTool(e *engine.Engine) {
	for _, tool := range d.Tools {
		data, err := getSource(tool)
		unableToGetSourceError(err, tool)

		var parsedData engine.ToolAPI

		err = parseData(data, tool)

		if err != nil {
			log.Errorf("unable to parse %v as a tool:\n %v", tool, err)
		}

		e.AddTool(parsedData)
	}
}

func parseData(data []byte, pData interface{}) error {
	err := yaml.UnmarshalWithOptions(data, &pData, yaml.Strict())

	return err
}

func unableToGetSourceError(err error, resouceToBeAdd interface{}) {
	if err != nil {
		log.Errorf("unable to load get source of %v:\n%v", resouceToBeAdd, err)
	}
}

func (d *DSL) addProvider(e *engine.Engine) {
	for _, provider := range d.API {
		data, err := getSource(provider)
		unableToGetSourceError(err, provider)

		var parsedData engine.ProviderAPI

		err = parseData(data, parsedData)

		unableToGetSourceError(err, provider)

		e.AddProvider(parsedData)
	}
}

func (d *DSL) addProvisioner(e *engine.Engine) {
	for _, provisioner := range d.Provisioners {
		data, err := getSource(provisioner)

		unableToGetSourceError(err, provisioner)

		var parsedData engine.ProvisionerAPI

		if err := parseData(data, parseData); err != nil {
			log.Errorf("unable to parse %v as a provisioner:\n %v", provisioner, err)
		}

		for resourceName, resourceActions := range parsedData {
			for actionName, actionProvisioners := range resourceActions {
				for _, actionProvisioner := range actionProvisioners {
					actionProvisioner.Resource = resourceName
					actionProvisioner.Action = actionName
					e.AddProvisioner(actionProvisioner)
				}
			}
		}
	}
}

// CreateEngine process dsl data and initialize the engine
// nolint: funlen
// TODO: function too long (70 > 60) .
func (d *DSL) CreateEngine() {
	e := engine.NewEngine()

	d.addSystem(e)
	d.addResource(e)
	d.addProvider(e)
	d.addProvisioner(e)
	d.addTool(e)

	if err := e.Match(); err != nil {
		log.Fatalf("New engine match process failed:\n%v", err)
	}

	e.Resolve()
	d.Engine = *e
}

// Run ignites the engine and get it to run found solutions for give action.
func (d DSL) Run(action string) error {
	scheduler := runner.ResourceNumScheduler{
		Solutions: d.Engine.Solutions,
	}
	local := runner.NewLocalRunner(d.Engine, action, scheduler)

	result, err := runner.Run(local)
	if err != nil {
		return err
	}

	log.Info(result)

	return nil
}

// GetSolutions wrapper function?
func (d *DSL) GetSolutions() []engine.Solution {
	return d.Engine.GetSolutions()
}
