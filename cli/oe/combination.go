// Package common contains shared functions, types and const used by oe commands
// nolint: forcetypeassert, wrapcheck
package main

import (
	"context"
	"encoding/json"
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

// combination is the result of processing the config file and manages the engine operations.
type combination struct {
	Providers    []string          `yaml:"providers"`
	Provisioners []string          `yaml:"provisioners"`
	Systems      []engine.System   `yaml:"systems"`
	Resources    []engine.Resource `yaml:"resources"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func getSource(uri, basepath string) ([]byte, error) {
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

	default:
		if !filepath.IsAbs(uri) {
			uri = filepath.Join(basepath, uri)
		}

		if fileExists(uri) {
			data, err := ioutil.ReadFile(filepath.Clean(uri))
			if err != nil {
				return nil, fmt.Errorf("unable to read %w", err)
			}

			return data, nil
		}

		return nil, xerrors.Errorf("source was not found (neither a relative file or URL)")
	}

	return data, nil
}

func add(e *engine.Engine, obj interface{}, def, basepath string) {
	switch def {
	case "System":
		system := obj.(engine.System)

		if err := e.AddSystem(system); err != nil {
			jsystem, _ := json.Marshal(system)
			log.Errorf("unable to add system:\n%s\nerror: %v", jsystem, err)
		}
	case "Resource":
		resource := obj.(engine.Resource)

		if err := e.AddResource(resource); err != nil {
			jresource, _ := json.Marshal(resource)
			log.Errorf("unable to add resource \n%s\nerror: %v", jresource, err)
		}
	case "Provider", "Provisioner", "Tool":
		target := obj.(string)

		data, err := getSource(target, basepath)
		if err != nil {
			log.Errorf("unable to load get source of %v:\n%v", target, err)

			break
		}

		err = e.Add(target, def, string(data))
		if err != nil {
			log.Errorf("unable to add %v %s:\n%v", def, target, err)
		}
	}
}

func createEngine(c combination, basePath string) engine.Engine {
	e, _ := engine.NewEngine("")

	for _, system := range c.Systems {
		add(e, system, "System", "")
	}

	for _, resource := range c.Resources {
		add(e, resource, "Resource", "")
	}

	for _, provider := range c.Providers {
		add(e, provider, "Provider", basePath)
	}

	for _, provisioner := range c.Provisioners {
		add(e, provisioner, "Provisioner", basePath)
	}

	return *e
}
