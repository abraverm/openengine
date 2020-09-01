package cmd

import (
	"fmt"
	"github.com/abraverm/openengine/cli/common"
	yaml "github.com/goccy/go-yaml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(deployCommand)
}

var deployCommand = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy resources",
	Long:  `Deploy command will parse the DSL file, resolve the APIs and other requirements to provision requested resources`,
	Run:   deploy,
}

func deploy(cmd *cobra.Command, args []string) {
	filename, _ := filepath.Abs(viper.GetString("dsl"))
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read DSL file:\n%v", err)
	}

	var dsl common.DSL
	err = yaml.UnmarshalWithOptions(yamlFile, &dsl, yaml.Strict())
	if err != nil {
		log.Fatalf("Unable to parse DSL file:\n%v", fmt.Sprintf(err.Error()))
	}
	dsl.CreateEngine()
	if err := dsl.Run("create"); err != nil {
		log.Fatalf("Engine failed to run:\n%v", err)
	}
}
