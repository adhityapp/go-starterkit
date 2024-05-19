package cmd

import (
	"os"

	"github.com/adhityapp/go-starterkit/bootstrap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Execute() error {
	// todo configurable logging level
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	var (
		root = &cobra.Command{}
		cfg  = bootstrap.Init()
	)

	root.AddCommand(
		restCommand(cfg),
	)

	return root.Execute()
}
