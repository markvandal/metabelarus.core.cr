package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/crconsent/types"
)

func CmdCreateRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-request [initiator] [recipient] [requestType] [status] [value] [memo] [promoUrl] [creationDt] [finalDt]",
		Short: "Creates a new request",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsInitiator := string(args[0])
			argsRecipient := string(args[1])
			argsRequestType := string(args[2])
			argsStatus := string(args[3])
			argsValue, _ := strconv.ParseInt(args[4], 10, 64)
			argsMemo := string(args[5])
			argsPromoUrl := string(args[6])
			argsCreationDt := string(args[7])
			argsFinalDt := string(args[8])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			requestType, ok := types.RequestType_value[argsRequestType]
			if !ok {
				return fmt.Errorf("Unknown RequiestType: %s.", argsRequestType)
			}

			status, ok := types.Status_value[argsStatus]
			if !ok {
				return fmt.Errorf("Unknown Status: %s.", argsStatus)
			}

			msg := types.NewMsgCreateRequest(
				clientCtx.GetFromAddress().String(),
				string(argsInitiator),
				string(argsRecipient),
				types.RequestType(requestType),
				types.Status(status),
				int32(argsValue),
				string(argsMemo),
				string(argsPromoUrl),
				string(argsCreationDt),
				string(argsFinalDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-request [id] [initiator] [recipient] [requestType] [status] [value] [memo] [promoUrl] [creationDt] [finalDt]",
		Short: "Update a request",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsInitiator := string(args[1])
			argsRecipient := string(args[2])
			argsRequestType := string(args[3])
			argsStatus := string(args[4])
			argsValue, _ := strconv.ParseInt(args[5], 10, 64)
			argsMemo := string(args[6])
			argsPromoUrl := string(args[7])
			argsCreationDt := string(args[8])
			argsFinalDt := string(args[9])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			requestType, ok := types.RequestType_value[argsRequestType]
			if !ok {
				return fmt.Errorf("Unknown RequiestType: %s.", argsRequestType)
			}

			status, ok := types.Status_value[argsStatus]
			if !ok {
				return fmt.Errorf("Unknown Status: %s.", argsStatus)
			}

			msg := types.NewMsgUpdateRequest(
				clientCtx.GetFromAddress().String(),
				id,
				string(argsInitiator),
				string(argsRecipient),
				types.RequestType(requestType),
				types.Status(status),
				int32(argsValue),
				string(argsMemo),
				string(argsPromoUrl),
				string(argsCreationDt),
				string(argsFinalDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-request [id] [initiator] [recipient] [requestType] [status] [value] [memo] [promoUrl] [creationDt] [finalDt]",
		Short: "Delete a request by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRequest(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
