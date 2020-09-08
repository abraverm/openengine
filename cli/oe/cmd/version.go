package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
// TODO: Replace cobra and viper with mow.cli .
func init() {
	rootCmd.AddCommand(versionCmd)
}

// nolint: gochecknoglobals
// TODO: Replace cobra and viper with mow.cli .
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Borderline",
	Long:  `All software has versions. This is Borderline's'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Borderline v0.1")
	},
}
