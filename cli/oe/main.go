package main

import (
	"fmt"

	"github.com/abraverm/openengine/cli/oe/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%v", err)
	}
}
