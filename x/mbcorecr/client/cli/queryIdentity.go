package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
	"github.com/spf13/cobra"
)

func CmdListIdentity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-identity",
		Short: "list all identity",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIdentityRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IdentityAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListId2Addr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-id2addr",
		Short: "list all maps of identities to address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllId2AddrRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Id2AddrAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowIdentity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-identity [id]",
		Short: "shows a identity",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetIdentityRequest{
				Id: args[0],
			}

			res, err := queryClient.Identity(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
