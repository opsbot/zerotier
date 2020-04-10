package user

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "manage users",
	}
	cmd.AddCommand(
		ListCommand(),
		GetCommand(),
		UpdateCommand(),
		DeleteCommand(),
	)

	return cmd
}
