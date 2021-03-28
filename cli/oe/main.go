package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"golang.org/x/xerrors"
)

// nolint: gosec, exhaustivestruct
func initLogger(path string, debug, verbose bool) {
	format := log.TextFormatter{
		DisableTimestamp: true,
		DisableQuote:     true,
	}
	log.SetFormatter(&format)

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
		log.Debug("running oe")
	}
}

func op(path string, noop bool, opType string) error {
	filename, _ := filepath.Abs(path)

	base := filepath.Dir(filename)

	yamlFile, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return xerrors.Errorf("unable to read combination file:\n%v", err)
	}

	var c combination

	if er := yaml.UnmarshalWithOptions(yamlFile, &c, yaml.Strict()); er != nil {
		return fmt.Errorf("unable to parse combination file:\n%w", er)
	}

	e := createEngine(c, base)

	solutions, err := e.Solutions(opType)
	if err != nil {
		return fmt.Errorf("engine failure: %w", err)
	}

	if noop {
		log.Info(solutions)
		log.Debug(e.GetSpec())

		return nil
	}

	return nil
}

// nolint: funlen, exhaustivestruct
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
			Name:  "create",
			Usage: "Create resources",
			Description: "Create command will parse the DSL file, " +
				"resolve definitions and other requirements to provision requested resources",
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return xerrors.Errorf("no DSL file was provided (argument)")
				}
				initLogger(c.String("log"), c.Bool("debug"), c.Bool("verbose"))

				return op(c.Args().Get(0), c.Bool("noop"), "create")
			},
		},
		{
			Name:  "delete",
			Usage: "Delete resources",
			Description: "Delete command will parse the DSL file, " +
				"resolve definitions and other requirements to delete requested resources",
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return xerrors.Errorf("no DSL file was provided (argument)")
				}
				initLogger(c.String("log"), c.Bool("debug"), c.Bool("verbose"))

				return op(c.Args().Get(0), c.Bool("noop"), "delete")
			},
		},
		{
			Name:  "update",
			Usage: "Update resources",
			Description: "Update command will parse the DSL file, " +
				"resolve definitions and other requirements to update requested resources",
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return xerrors.Errorf("no DSL file was provided (argument)")
				}
				initLogger(c.String("log"), c.Bool("debug"), c.Bool("verbose"))

				return op(c.Args().Get(0), c.Bool("noop"), "update")
			},
		},
		{
			Name:  "read",
			Usage: "Get resources",
			Description: "Read/get command will parse the DSL file, " +
				"resolve definitions and other requirements to get requested resources",
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return xerrors.Errorf("no DSL file was provided (argument)")
				}
				initLogger(c.String("log"), c.Bool("debug"), c.Bool("verbose"))

				return op(c.Args().Get(0), c.Bool("noop"), "read")
			},
		},
	}

	app.Flags = flags

	return app.Run(args)
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
