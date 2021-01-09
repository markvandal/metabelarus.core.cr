package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdCreateId2Auth())
	cmd.AddCommand(CmdUpdateId2Auth())
	cmd.AddCommand(CmdDeleteId2Auth())


	cmd.AddCommand(CmdCreateAuth())
	cmd.AddCommand(CmdUpdateAuth())
	cmd.AddCommand(CmdDeleteAuth())


	cmd.AddCommand(CmdCreateId2Sign())
	cmd.AddCommand(CmdUpdateId2Sign())
	cmd.AddCommand(CmdDeleteId2Sign())


	cmd.AddCommand(CmdCreateSignature())
	cmd.AddCommand(CmdUpdateSignature())
	cmd.AddCommand(CmdDeleteSignature())


	return cmd 
}
