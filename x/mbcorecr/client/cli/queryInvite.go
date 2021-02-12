package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
	"github.com/spf13/cobra"
)

func CmdShowInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-invite [id]",
		Short: "shows a invite",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetInviteRequest{
				Id: args[0],
			}

			res, err := queryClient.Invite(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
