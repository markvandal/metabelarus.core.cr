package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group metabelaruscorecr queries under a subcommand
	metabelaruscorecrQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	metabelaruscorecrQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding # 1
			GetCmdListConfirmation(queryRoute, cdc),
			GetCmdGetConfirmation(queryRoute, cdc),
			GetCmdListInvitation(queryRoute, cdc),
			GetCmdGetInvitation(queryRoute, cdc),
			GetCmdListIdentity(queryRoute, cdc),
			GetCmdGetIdentity(queryRoute, cdc),
		)...,
	)

	return metabelaruscorecrQueryCmd
}
