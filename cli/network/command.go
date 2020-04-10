package network

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "network",
		Short: "manage networks",
	}
	cmd.AddCommand(
		ListCommand(),
		GetCommand(),
		PutCommand(),
		DeleteCommand(),
	)

	return cmd
}
