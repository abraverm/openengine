// Package controller provide all first layer method used by either cli or REST API
package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	dslapi "openengine/dsl"
	"openengine/engine"
	"os"
	"path/filepath"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/goccy/go-yaml"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// InitLogger is the init function for logger.
// nolint: G304
func InitLogger(path string, debug, verbose bool) {
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)

	if verbose {
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	}

	log.SetLevel(log.InfoLevel)

	if debug {
		log.SetLevel(log.DebugLevel)
	}
}

// Deploy is the entrypoint of the whole process of finding solution.
func deploy(param DeployParam) (err error) {
	filename, _ := filepath.Abs(param.Path)

	yamlFile, errRead := ioutil.ReadFile(filepath.Clean(filename))
	if errRead != nil {
		err = xerrors.New(fmt.Sprintf("Unable to read DSL file:\n%v", errRead.Error()))
	}

	var dsl dslapi.DSL

	switch param.CallFrom {
	case "cli":
		errParse := yaml.UnmarshalWithOptions(yamlFile, &dsl, yaml.Strict())
		if errParse != nil {
			err = xerrors.New(fmt.Sprintf("Unable to parse DSL file:\n%v", errParse.Error()))
		}
	case "rest":
		dsl = dslapi.DSL{
			API:          param.API,
			Provisioners: param.Provisioners,
			Systems:      param.Systems,
			Tools:        param.Tools,
			Engine:       param.Engine,
			Resources:    param.Resources,
		}
	default:
		errParse := yaml.UnmarshalWithOptions(yamlFile, &dsl, yaml.Strict())
		if errParse != nil {
			err = xerrors.New(fmt.Sprintf("Unable to parse DSL file:\n%v", errParse.Error()))
		}
	}

	dsl.CreateEngine()

	if param.Noop {
		engineJSON, _ := json.MarshalIndent(dsl.Engine, "", "  ")

		fmt.Println(string(engineJSON))

		err = nil
	}

	if errCreate := dsl.Run("create"); err != nil {
		err = xerrors.New(fmt.Sprintf("Engine failed to run:\n%v", errCreate.Error()))
	}

	return err
}

// DeployParam is the parameters requested by all method which calls Deploy.
type DeployParam struct {
	CallFrom     string
	Noop         bool
	Debug        bool
	Verbose      bool
	Path         string
	Log          string
	API          []string          `yaml:"api"`
	Provisioners []string          `yaml:"provisioners"`
	Systems      []engine.System   `yaml:"systems"`
	Tools        []string          `yaml:"tools"`
	Resources    []engine.Resource `yaml:"resources"`
	Engine       engine.Engine
}

// Deploy is the method exposed to all first layer calls
//(for example, oe cli, oe rest api), it's a wrapper of deploy method.
func Deploy(c DeployParam) error {
	InitLogger(c.Log, c.Debug, c.Verbose)

	return deploy(c)
}
