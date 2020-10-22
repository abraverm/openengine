package main

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
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"golang.org/x/xerrors"
)

// nolint: G304
func initLogger(path string, debug, verbose bool) {
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

func delete(path string, noop bool) (results []string, err error) {
	filename, _ := filepath.Abs(path)

	yamlFile, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, xerrors.Errorf("Unable to read DSL file:\n%v", err)
	}

	var dsl DSL

	err = yaml.UnmarshalWithOptions(yamlFile, &dsl, yaml.Strict())
	if err != nil {
		return nil, xerrors.Errorf("Unable to parse DSL file:\n%v", err.Error())
	}

	dsl.CreateEngine()

	if noop {
		engineJSON, _ := json.MarshalIndent(dsl.Engine, "", "  ")

		fmt.Println(string(engineJSON))

		return nil, nil
	}

	if results, err := dsl.Delete("delete"); err != nil {
		return nil, xerrors.Errorf("Engine failed to run:\n%v", err)
	} else {
		return results, nil
	}

	// return results, nil
}

// nolint: funlen
func run(args []string) error {
	app := &cli.App{
		Name:        "oe",
		Description: "OpenEgnine command line tool",
		Authors: []*cli.Author{
			{
				Name:  "Alexander Braverman Masis",
				Email: "abraverm@redhat.com",
			},
		},
	}

	flags := []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:    "noop",
			Aliases: []string{"n"},
			Usage:   "No operation - complete the command without its actual execution",
			Value:   false,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Debug level of logging",
			Value:   false,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Print log to console",
			Value:   false,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "log",
			Aliases: []string{"l"},
			Usage:   "Log file path",
			Value:   "oe.log",
		}),
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Config file path",
			Value:   "oe.yaml",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "deploy",
			Usage: "Create resources",
			Description: "Deploy command will parse the DSL file, " +
				"resolve the APIs and other requirements to provision requested resources",
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return xerrors.Errorf("no DSL file was provider (argument)")
				}
				initLogger(c.String("log"), c.Bool("debug"), c.Bool("verbose"))

				return deploy(c.Args().Get(0), c.Bool("noop"))
			},
		},
		{
			Name:  "delete",
			Usage: "Delete resources",
			Description: "Delete command will parse the DSL file, " +
				"resolve the APIs and other requirements to delete requested resources",
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return xerrors.Errorf("no DSL file was provider (argument)")
				}
				initLogger(c.String("log"), c.Bool("debug"), c.Bool("verbose"))
				results, err := delete(c.Args().Get(0), c.Bool("noop"))
				fmt.Println("This is the result of deleting", results)
				return err
			},
		},
	}

	app.Flags = flags

	return app.Run(args)
}

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
