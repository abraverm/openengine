package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/bmizerany/assert"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func Test_deploy(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	c := &cobra.Command{Use: "deploy", Run: deploy}

	defer func() { log.StandardLogger().ExitFunc = nil }()
	var fatal bool
	log.StandardLogger().ExitFunc = func(int) { fatal = true }
	log.SetOutput(ioutil.Discard)
	tests := []struct {
		name      string
		fixture   string
		returnErr bool
		args      args
	}{
		{"empty", "testdata/empty", false, args{c, nil}},
		{"not_found", "testdata/nill", true, args{c, nil}},
		{"bad_yaml", "testdata/bad_yaml", true, args{c, nil}},
		{"bad_dsl", "testdata/bad_dsl", true, args{c, nil}},
	}
	for _, tt := range tests {
		viper.SetDefault("dsl", tt.fixture)
		fatal = false
		t.Run(tt.name, func(t *testing.T) {
			deploy(tt.args.cmd, tt.args.args)
		})
		assert.Equal(t, tt.returnErr, fatal)
	}
}
