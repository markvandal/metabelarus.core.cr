package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	metabelaruscorecrTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	metabelaruscorecrTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding # 1
		GetCmdCreateConfirmation(cdc),
		GetCmdSetConfirmation(cdc),
		GetCmdDeleteConfirmation(cdc),
		GetCmdCreateInvitation(cdc),
		GetCmdSetInvitation(cdc),
		GetCmdDeleteInvitation(cdc),
		GetCmdCreateIdentity(cdc),
		GetCmdSetIdentity(cdc),
		GetCmdDeleteIdentity(cdc),
	)...)

	return metabelaruscorecrTxCmd
}
