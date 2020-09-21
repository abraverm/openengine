package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"default", false},
	}
	cfgFile = "testdata/empty"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs([]string{"help"})
			rootCmd.AddCommand(&cobra.Command{Use: "^\\QTestExecute\\E$", Run: func(cmd *cobra.Command, args []string) {}})
			if err := Execute(); (err != nil) != tt.wantErr {
				// t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_initLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
