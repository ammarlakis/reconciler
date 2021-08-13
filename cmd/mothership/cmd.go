package cmd

import (
	startCmd "github.com/kyma-incubator/reconciler/cmd/mothership/start"
	"github.com/kyma-incubator/reconciler/internal/cli"
	"github.com/spf13/cobra"
)

func NewCmd(o *cli.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mothership",
		Short: "Manage Kyma mothership reconciler",
		Long:  "Administrative CLI tool for the Kyma mothership reconciler service",
	}

	cmd.AddCommand(startCmd.NewCmd(startCmd.NewOptions(o)))

	return cmd
}