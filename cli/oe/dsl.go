// Package common contains shared functions, types and const used by oe commands
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/abraverm/openengine/engine"
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
	Engine       *engine.Engine
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
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, uri, nil)

		// nolint: exhaustivestruct
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

func (d *DSL) add(obj interface{}, def string) {
	switch def {
	case "System":
		err := d.Engine.AddSystem(obj.(engine.System))
		if err != nil {
			log.Errorf("unable to add system #{system}: \n #{err}")
		}
	case "Resource":
		err := d.Engine.AddResource(obj.(engine.Resource))
		if err != nil {
			log.Errorf("unable to add resource #{resource}: \n #{err}")
		}
	case "Provider", "Provisioner", "Tool":
		target := obj.(string)

		data, err := getSource(target)
		if err != nil {
			log.Errorf("unable to load get source of %v:\n%v", target, err)
		}

		err = d.Engine.Add(target, def, string(data))
		if err != nil {
			log.Errorf("unable to add #{def} #{target}:\n#{err}")
		}
	}
}

// CreateEngine process dsl data and initialize the engine.
func (d *DSL) CreateEngine() {
	d.Engine, _ = engine.NewEngine("")

	for _, system := range d.Systems {
		d.add(system, "system")
	}

	for _, resource := range d.Resources {
		d.add(resource, "resource")
	}

	for _, provider := range d.API {
		d.add(provider, "Provider")
	}

	for _, provisioner := range d.Provisioners {
		d.add(provisioner, "Provisioner")
	}
}

// Run ignites the engine and get it to run found solutions for give action.
func (d DSL) Run(action string) error {
	return nil
}
