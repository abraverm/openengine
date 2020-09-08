package cmd

import (
	"fmt"
	"io"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// nolint: gochecknoglobals
// TODO: Replace cobra and viper with mow.cli .
var (
	// Used for flags.
	cfgFile string
	logFile os.File
	rootCmd = &cobra.Command{
		Use:   "openengine",
		Short: "Borderline command line tool",
		Long:  `CLI for processing DSL Borderline engine `,
	}
)

// Execute executes the root command.
func Execute() (err error) {
	defer func() {
		if logErr := logFile.Close(); logErr != nil {
			err = logErr
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

// nolint: gochecknoinits
// TODO: Replace cobra and viper with mow.cli .
func init() {
	cobra.OnInitialize(initConfig, initLogger)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "openengine.yaml", "config file")
	rootCmd.PersistentFlags().StringP("dsl", "b", "bdsl.yaml", "DSL file")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug log level")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Print log to stdout")
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix("bl")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	_ = viper.BindPFlag("dsl", rootCmd.PersistentFlags().Lookup("dsl"))
	_ = viper.BindPFlag("log.debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("log.verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.SetDefault("log.file", "openengine.log")
	viper.SetDefault("cache_path", ".openengine")
}

func initLogger() {
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	logFile, err := os.OpenFile(viper.GetString("log.file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)

	if viper.GetBool(("log.verbose")) {
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	}

	log.SetLevel(log.InfoLevel)

	if viper.GetBool("log.debug") {
		log.SetLevel(log.DebugLevel)
	}
}
