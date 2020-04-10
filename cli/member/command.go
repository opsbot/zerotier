package member

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	var network string

	cmd := &cobra.Command{
		Use:   "member",
		Short: "manage network members",
	}
	cmd.AddCommand(
		ListCommand(),
		GetCommand(),
		PutCommand(),
	)
	cmd.PersistentFlags().StringVarP(&network, "network", "n", "", "network id")

	return cmd
}
