package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group mbcorecr queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListInvite())
	cmd.AddCommand(CmdShowInvite())


	cmd.AddCommand(CmdListIdentity())
	cmd.AddCommand(CmdShowIdentity())
	cmd.AddCommand(CmdUnpackPrivKey())
	cmd.AddCommand(CmdDecryptPayload())
	cmd.AddCommand(CmdReencryptPrivKey())

	return cmd
}