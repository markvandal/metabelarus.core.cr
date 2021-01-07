package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

func CmdCreateIdentity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-identity [AccountID] [IdentityType] [Details]",
		Short: "Creates a new identity",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAccountID := string(args[0])

			argsIdentityType, err := strconv.Atoi(string(args[1]))
			if err != nil {
				return err
			}

			argsDetails := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateIdentity(
				clientCtx.GetFromAddress().String(),
				string(argsAccountID),
				types.IdentityType(argsIdentityType),
				string(argsDetails),
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

func CmdUpdateIdentity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-identity [id] [AccountID] [IdentityType] [Details]",
		Short: "Update a identity",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsAccountID := string(args[1])
			argsIdentityType, err := strconv.Atoi(string(args[2]))
			if err != nil {
				return err
			}
			argsDetails := string(args[3])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateIdentity(
				clientCtx.GetFromAddress().String(),
				id,
				string(argsAccountID),
				types.IdentityType(argsIdentityType),
				string(argsDetails),
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

func CmdDeleteIdentity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-identity [id] [AccountID] [IdentityType] [Details] [CreationDt]",
		Short: "Delete a identity by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteIdentity(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
