package service

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "service",
		Short: "api service",
	}
	cmd.AddCommand(
		RandomTokenCommand(),
		StatusCommand(),
	)

	return cmd
}
