package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"github.com/spf13/cobra"
)

func CmdShowAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-auth [id]",
		Short: "shows a auth",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAuthRequest{
				Id: args[0],
			}

			res, err := queryClient.Auth(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowId2Service() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-id2services [id]",
		Short: "shows a id to services map",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetId2ServiceRequest{
				Id: args[0],
			}

			res, err := queryClient.Id2Service(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
