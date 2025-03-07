package cli

import (
	"fmt"

	"omni/x/omni/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group omni queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListBalance())
	cmd.AddCommand(CmdListWhitelist())
	cmd.AddCommand(CmdShowWhitelist())
	cmd.AddCommand(CmdListObserveVote())
	cmd.AddCommand(CmdShowObserveVote())
	// this line is used by starport scaffolding # 1

	return cmd
}
