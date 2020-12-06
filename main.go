// Package main is the real entrypoint of OpenEngine
package main

import (
	"log"
	"openengine/cli"
	"os"
)

func main() {
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
