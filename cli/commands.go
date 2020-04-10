package cli

import (
	"github.com/spf13/cobra"

	"github.com/opsbot/zerotier/cli/member"
	"github.com/opsbot/zerotier/cli/network"
	"github.com/opsbot/zerotier/cli/service"
	"github.com/opsbot/zerotier/cli/user"
)

// AddCommands add the commands from subcommand groups to the root command
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		member.Command(),
		network.Command(),
		service.Command(),
		user.Command(),
		VersionCommand(),
	)
}
