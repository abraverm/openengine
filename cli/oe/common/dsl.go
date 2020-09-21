// Package common contains shared functions, types and const used by oe commands
package common

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/abraverm/openengine/engine"
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func getSource(uri string) ([]byte, error) {
	var data []byte

	urlParsed, _ := url.ParseRequestURI(uri)

	switch {
	case urlParsed != nil:
		res, err := http.NewRequestWithContext(context.Background(), http.MethodGet, uri, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to download %w", err)
		}

		defer func() {
			if closeErr := res.Body.Close(); closeErr != nil {
				err = closeErr
			}
		}()

		data, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
	case fileExists(uri):
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

// CreateEngine process dsl data and initialize the engine
// nolint: funlen
// TODO: function too long (70 > 60) .
func (d *DSL) CreateEngine() {
	e := engine.NewEngine()

	for _, system := range d.Systems {
		e.AddSystem(system)
	}

	for _, resource := range d.Resources {
		e.AddResource(resource)
	}

	for _, provider := range d.API {
		data, err := getSource(provider)
		if err != nil {
			log.Errorf("unable to load get source of %v:\n%v", provider, err)
		}

		var parsedData engine.ProviderAPI
		err = yaml.UnmarshalWithOptions(data, &parsedData, yaml.Strict())

		if err != nil {
			log.Errorf("unable to parse %v as a provider:\n %v", provider, err)
		}

		e.AddProvider(parsedData)
	}

	for _, provisioner := range d.Provisioners {
		data, err := getSource(provisioner)
		if err != nil {
			log.Errorf("unable to load get source of %v:\n%v", provisioner, err)
		}

		var parsedData engine.ProvisionerAPI
		if err := yaml.UnmarshalWithOptions(data, &parsedData, yaml.Strict()); err != nil {
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

	for _, tool := range d.Tools {
		data, err := getSource(tool)
		if err != nil {
			log.Errorf("unable to load get source of %v:\n%v", tool, err)
		}

		var parsedData engine.ToolAPI
		err = yaml.UnmarshalWithOptions(data, &parsedData, yaml.Strict())

		if err != nil {
			log.Errorf("unable to parse %v as a tool:\n %v", tool, err)
		}

		e.AddTool(parsedData)
	}

	if err := e.Match(); err != nil {
		log.Fatalf("New engine match process failed:\n%v", err)
	}

	e.Resolve()
	d.Engine = *e
}

// Run ignites the engine and get it to run found solutions for give action.
func (d DSL) Run(action string) error {
	if err := d.Engine.Schedule(action); err != nil {
		return err
	}

	results, err := d.Engine.Run()
	if err != nil {
		return err
	}

	for _, result := range results {
		log.Debugln("\n", result)
	}

	return nil
}

// GetSolutions wrapper function?
func (d *DSL) GetSolutions() []engine.Solution {
	return d.Engine.GetSolutions()
}
