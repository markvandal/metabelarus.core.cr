package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func CmdRequestAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-auth [service] [identity] [key]",
		Short: "Request a new auth by a service",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsIdentity := string(args[1])
			argsService := string(args[0])
			argsKey := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestAuth(
				clientCtx.GetFromAddress().String(),
				string(argsService),
				string(argsIdentity),
				string(argsKey), // @TODO Encrypt key with user's pubkey
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdConfirmAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-auth [identity] [service]",
		Short: "Request a new auth by a service",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsIdentity := string(args[0])
			argsService := string(args[1])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgConfirmAuth(
				clientCtx.GetFromAddress().String(),
				string(argsService),
				string(argsIdentity),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
