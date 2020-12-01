// Package controller provide all first layer method used by either cli or REST API
package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/goccy/go-yaml"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

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

// Deploy is the entrypoint of the whole process of finding solution
func deploy(path string, noop bool) error {
	filename, _ := filepath.Abs(path)

	yamlFile, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return xerrors.Errorf("Unable to read DSL file:\n%v", err)
	}

	var dsl DSL

	err = yaml.UnmarshalWithOptions(yamlFile, &dsl, yaml.Strict())
	if err != nil {
		return xerrors.Errorf("Unable to parse DSL file:\n%v", err.Error())
	}

	dsl.CreateEngine()

	if noop {
		engineJSON, _ := json.MarshalIndent(dsl.Engine, "", "  ")

		fmt.Println(string(engineJSON))

		return nil
	}

	if err := dsl.Run("create"); err != nil {
		return xerrors.Errorf("Engine failed to run:\n%v", err)
	}

	return nil
}

// DeployParam is the parameters requested by all method which calls Deploy
type DeployParam struct {
	Log     string
	Debug   bool
	Verbose bool
	Path    string
	Noop    bool
}

// Deploy is the method exposed to all first layer calls (for example, oe cli, oe rest api), it's a wrapper of deploy method
func Deploy(c DeployParam) error {
	InitLogger(c.Log, c.Debug, c.Verbose)

	return deploy(c.Path, c.Noop)
}
