package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/spf13/cobra"
)

func GetCmdListIdentity(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-identity",
		Short: "list all Identity",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListIdentity, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Identity\n%s\n", err.Error())
				return nil
			}
			var out []types.Identity
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetIdentity(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-identity [key]",
		Short: "Query a Identity by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetIdentity, key), nil)
			if err != nil {
				fmt.Printf("could not resolve Identity %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Identity
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
