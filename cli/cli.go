// Package cli is the Command Line Interface of Openengine, it's also the entrypoint of the program
package cli

import (
	"openengine/controller"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

// nolint:G304
func Run(args []string) error {
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
			Action: cliDeploy,
		},
	}

	app.Flags = flags

	return app.Run(args)
}

func cliDeploy(c *cli.Context) error {
	return controller.Deploy(controller.DeployParam{
		Log:     c.String("log"),
		Debug:   c.Bool("debug"),
		Verbose: c.Bool("verbose"),
		Path:    c.Args().Get(0),
		Noop:    c.Bool("noop"),
	})
}
