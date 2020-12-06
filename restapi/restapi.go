// Package restapi is where all restapi go
package restapi

import (
	"openengine/controller"

	"github.com/urfave/cli/v2"
)

type RestApi struct {
}

func restDeploy(c *cli.Context) error {
	return controller.Deploy(controller.DeployParam{
		Log:      c.String("log"),
		Debug:    c.Bool("debug"),
		Verbose:  c.Bool("verbose"),
		Path:     c.Args().Get(0),
		Noop:     c.Bool("noop"),
		CallFrom: "rest",
	})
}
