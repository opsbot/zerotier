package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCommand returns a cobra command
func VersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "get version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version    : %s\n", Version)
			fmt.Printf("Git Hash   : %s\n", GitHash)
			fmt.Printf("Build Time : %s\n", BuildTime)
		},
	}
	return cmd
}
