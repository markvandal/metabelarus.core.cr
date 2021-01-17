package cli

import (
    "context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
    "github.com/metabelarus/mbcorecr/x/crsign/types"
)

func CmdListSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-signature",
		Short: "list all signature",
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

            params := &types.QueryAllSignatureRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.SignatureAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-signature [id]",
		Short: "shows a signature",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)
            clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryGetSignatureRequest{
                Id: args[0],
            }

            res, err := queryClient.Signature(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
