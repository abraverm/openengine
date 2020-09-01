package common

import (
	"encoding/json"
	"fmt"
	"github.com/abraverm/openengine/engine"
	"github.com/goccy/go-yaml"
	"github.com/jedib0t/go-pretty/v6/table"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

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
	if _, err := url.ParseRequestURI(uri); err == nil {
		res, err := http.Get(uri)
		if err != nil {
			return nil, fmt.Errorf("unable to download %v: %v", uri, err)
		}
		defer func() {
			closeErr := res.Body.Close()
			if err == nil {
				err = closeErr
			}
		}()
		data, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
	} else if fileExists(uri) {
		data, err = ioutil.ReadFile(uri)
		if err != nil {
			return nil, fmt.Errorf("unable to read %v: %v", uri, err)
		}
	} else {
		return nil, fmt.Errorf("source %v was not found (neither a relative file or URL)", uri)
	}
	return data, nil
}

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
	log.Debugf("Match results:\n%v", listSolutions(e.GetSolutions()))
	e.Resolve()
	log.Debugf("Resolved solutions:\n%v", listSolutions(e.GetSolutions()))
	d.Engine = *e
}

func listSolutions(solutions []engine.Solution) string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Solution", "Resource", "System", "Provider", "Provisioner", "Debug"})
	autoMerge := table.RowConfig{AutoMerge: true}
	for _, solution := range solutions {
		t.AppendRow(table.Row{solution.ToJson(), toJson(solution.Resource), toJson(solution.System), toJson(solution.Provider), toJson(solution.Provisioner), solution.Output}, autoMerge)
		t.AppendSeparator()
	}
	return t.Render()
}

func toJson(object interface{}) string {
	oJSON, _ := json.MarshalIndent(object, "", "  ")
	return string(oJSON)
}

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
