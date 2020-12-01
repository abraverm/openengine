package controller

import (
	"os"
	"testing"
)

var deployParam *DeployParam

func resetDeployParamDefault() {
	deployParam = &DeployParam{
		Log:     "oe.log",
		Debug:   true,
		Verbose: true,
		Path:    "testdata/empty",
		Noop:    true,
	}
}

func resetDeployParamREST() {
	deployParam = &DeployParam{
		Log:      "oe.log",
		CallFrom: "rest",
		Debug:    true,
		Verbose:  true,
		Noop:     true,
	}
}
func resetDeployParamCLI() {
	deployParam = &DeployParam{
		Log:      "oe.log",
		CallFrom: "cli",
		Debug:    true,
		Verbose:  true,
		Path:     "testdata/bdsl.yaml",
		Noop:     true,
	}
}

func testDeploy(t *testing.T) {
	err := Deploy(*deployParam)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDeployWrapperREST(t *testing.T) {
	resetDeployParamREST()
	testDeploy(t)
}

func TestDeployWrapperCLI(t *testing.T) {
	resetDeployParamCLI()
	testDeploy(t)
}

func TestDeployWrapperDefault(t *testing.T) {
	resetDeployParamDefault()
	testDeploy(t)
}

func TestMain(m *testing.M) {
	exitcode := m.Run()
	os.RemoveAll("oe.log") // remove the directory and its contents.
	os.Exit(exitcode)
}
