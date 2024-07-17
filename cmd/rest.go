package cmd

import (
	"github.com/adhityapp/go-starterkit/api/rest"
	"github.com/adhityapp/go-starterkit/bootstrap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func restCommand(cfg *bootstrap.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run a web server service",
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.WithField("component", "rest").Info("running rest")
			rest.Serve(cfg)
			return nil
		},
	}
}
