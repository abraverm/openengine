package cmd

import (
	"fmt"
	"github.com/abraverm/engine/cli/common"
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
		log.Fatal("Unable to read DSL file:", err)
	}

	var dsl common.DSL
	err = yaml.UnmarshalWithOptions(yamlFile, &dsl, yaml.Strict())
	if err != nil {
		log.Fatalf("Unable to parse DSL file: %v", fmt.Sprintf(err.Error()))
	}
	if err = dsl.CreateEngine(); err != nil {
		log.Fatal(err)
	}
	if err := dsl.Run("create"); err != nil {
		log.Fatal(err)
	}
}
