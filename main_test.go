package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEntryPoint(t *testing.T) {
	fmt.Println(os.Args)

	os.Args = []string{"oe"}
	main()
}

func TestMain(m *testing.M) {
	exitcode := m.Run()
	os.RemoveAll("oe.log")
	os.Exit(exitcode)
}
