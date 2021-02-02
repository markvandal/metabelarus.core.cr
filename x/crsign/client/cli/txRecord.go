package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func CmdCreateRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-record [key] [data] [record type] [publicity] [live time] [provider - optional] [parent record id - optional]",
		Short: "Creates a new reacord",
		Args:  cobra.RangeArgs(5, 7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsKey := string(args[0])
			argsData := string(args[1])
			argsType := string(args[2])
			argsPublicity := string(args[3])
			argsLiveTime := string(args[4])
			argsProvider := ""
			argsParentId := ""
			if len(args) > 5 {
				argsProvider = args[5]
			}
			if len(args) == 7 {
				argsParentId = args[6]
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			signature, _, err := clientCtx.Keyring.Sign(clientCtx.GetFromName(), []byte(argsData))
			if err != nil {
				return err
			}

			liveTime, err := strconv.ParseInt(argsLiveTime, 10, 32)
			if err != nil {
				return err
			}

			publicity, ok := types.PublicityType_value[argsPublicity]
			if !ok {
				return fmt.Errorf("Publicity type: %s does not exist", argsPublicity)
			}

			recordType, ok := types.RecordType_value[argsType]
			if !ok {
				return fmt.Errorf("Record type: %s does not exist", argsType)
			}

			msg := types.NewMsgCreateRecord(
				clientCtx.GetFromAddress().String(),
				string(argsKey),
				string(argsData),
				hex.EncodeToString(signature),
				types.RecordType(recordType),
				types.PublicityType(publicity),
				int32(liveTime),
				string(argsProvider),
				argsParentId,
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

func CmdUpdateRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-record [id] [data] [live time] [action]",
		Short: "Update a record",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsData := string(args[1])
			argsLiveTime := string(args[2])
			argsAction := string(args[3])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			signature, _, err := clientCtx.Keyring.Sign(clientCtx.GetFromName(), []byte(argsData))
			if err != nil {
				return err
			}

			liveTime, err := strconv.ParseInt(argsLiveTime, 10, 32)
			if err != nil {
				return err
			}

			action, ok := types.RecordUpdate_value[argsAction]
			if !ok {
				return fmt.Errorf("Action type: %s does not exist", argsAction)
			}

			msg := types.NewMsgUpdateRecord(
				clientCtx.GetFromAddress().String(),
				id,
				string(argsData),
				hex.EncodeToString(signature),
				int32(liveTime),
				types.RecordUpdate(action),
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

func CmdDeleteRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-record [id]",
		Short: "Delete a record by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRecord(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
