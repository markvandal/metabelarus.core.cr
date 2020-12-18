package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

func GetCmdCreateConfirmation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-confirmation [idenitityID] [creationDate] [expirationDate] [confirmatorID] [centerGeo] [status] [nextTryDate]",
		Short: "Creates a new confirmation",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsIdenitityID := string(args[0] )
			argsCreationDate := string(args[1] )
			argsExpirationDate := string(args[2] )
			argsConfirmatorID := string(args[3] )
			argsCenterGeo := string(args[4] )
			argsStatus := string(args[5] )
			argsNextTryDate := string(args[6] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateConfirmation(cliCtx.GetFromAddress(), string(argsIdenitityID), string(argsCreationDate), string(argsExpirationDate), string(argsConfirmatorID), string(argsCenterGeo), string(argsStatus), string(argsNextTryDate))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetConfirmation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-confirmation [id]  [idenitityID] [creationDate] [expirationDate] [confirmatorID] [centerGeo] [status] [nextTryDate]",
		Short: "Set a new confirmation",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsIdenitityID := string(args[1])
			argsCreationDate := string(args[2])
			argsExpirationDate := string(args[3])
			argsConfirmatorID := string(args[4])
			argsCenterGeo := string(args[5])
			argsStatus := string(args[6])
			argsNextTryDate := string(args[7])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetConfirmation(cliCtx.GetFromAddress(), id, string(argsIdenitityID), string(argsCreationDate), string(argsExpirationDate), string(argsConfirmatorID), string(argsCenterGeo), string(argsStatus), string(argsNextTryDate))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteConfirmation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-confirmation [id]",
		Short: "Delete a new confirmation by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteConfirmation(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
