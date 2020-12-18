package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
    "github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

func GetCmdListInvitation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-invitation",
		Short: "list all invitation",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListInvitation, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Invitation\n%s\n", err.Error())
				return nil
			}
			var out []types.Invitation
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetInvitation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-invitation [key]",
		Short: "Query a invitation by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetInvitation, key), nil)
			if err != nil {
				fmt.Printf("could not resolve invitation %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Invitation
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
