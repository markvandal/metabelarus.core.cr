package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func CmdCreateSignatureList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-signatureList [rootSignatureId]",
		Short: "Creates a new signatureList",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRootSignatureId := string(args[0])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSignatureList(
				clientCtx.GetFromAddress().String(),
				string(argsRootSignatureId),
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

func CmdUpdateSignatureList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-signatureList [id] [rootSignatureId] [lastSignatureId] [nextSignatureId] [metadata]",
		Short: "Update a signatureList",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsRootSignatureId := string(args[1])
			argsLastSignatureId := string(args[2])
			argsNextSignatureId := string(args[3])
			argsMetadata := string(args[4])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSignatureList(clientCtx.GetFromAddress().String(), id, string(argsRootSignatureId), string(argsLastSignatureId), string(argsNextSignatureId), string(argsMetadata))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
