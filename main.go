// Package main is the real entrypoint of OpenEngine
package main

import (
	"log"
	"openengine/cli"
	"os"
)

// nolint:G304
func main() {
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
