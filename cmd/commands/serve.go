package commands

import (
	"github.com/spf13/cobra"
)

func serveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start megatude server with echo",
		Run: func(cmd *cobra.Command, args []string) {
			megatude.Start()
		},
	}

	return cmd
}
