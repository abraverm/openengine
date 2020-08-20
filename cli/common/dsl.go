package common

import (
	"fmt"
	"github.com/abraverm/engine/engine2"
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type DSL struct {
	API          []string  `yaml:"api"`
	Provisioners []string `yaml:"provisioners"`
	Systems    	 []engine2.System `yaml:"systems"`
	Tools        []string `yaml:"tools"`
	Resources    []engine2.Resource `yaml:"resources"`
	Engine 		 engine2.Engine
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

func (d *DSL) CreateEngine() error {
	e := engine2.NewEngine()
	for _, system := range d.Systems {
		e.AddSystem(system)
	}
	for _, resource := range d.Resources {
		e.AddResource(resource)
	}
	for _, provider := range d.API {
		data, err := getSource(provider)
		if err != nil {
			return err
		}
		var parsedData engine2.ProviderAPI
		err = yaml.UnmarshalWithOptions(data, &parsedData, yaml.Strict())
		if err != nil {
			return err
		}
		e.AddProvider(parsedData)
	}
	for _, provisioner := range d.Provisioners {
		data, err := getSource(provisioner)
		if err != nil {
			return err
		}
		var parsedData engine2.Provisioner
		if err := yaml.UnmarshalWithOptions(data, &parsedData, yaml.Strict()); err != nil {
			return err
		}
		e.AddProvisioner(parsedData)
	}
	if err := e.Match(); err != nil {
		return err
	}
	e.Resolve()
	d.Engine = *e
	return nil
}

func (d DSL) Run(action string) error {
	if err := d.Engine.Schedule(action); err != nil {
		return err
	}
	if _, err := d.Engine.Run(); err != nil {
		return err
	}
	return nil
}
