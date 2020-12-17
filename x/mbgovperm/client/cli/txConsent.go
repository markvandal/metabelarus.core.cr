package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
)

func GetCmdCreateConsent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-consent [extserviceId] [passportId] [resolution]",
		Short: "Creates a new consent",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsExtserviceId := string(args[0] )
			argsPassportId := string(args[1] )
			argsResolution := string(args[2] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateConsent(cliCtx.GetFromAddress(), string(argsExtserviceId), string(argsPassportId), string(argsResolution))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetConsent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-consent [id]  [extserviceId] [passportId] [resolution]",
		Short: "Set a new consent",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsExtserviceId := string(args[1])
			argsPassportId := string(args[2])
			argsResolution := string(args[3])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetConsent(cliCtx.GetFromAddress(), id, string(argsExtserviceId), string(argsPassportId), string(argsResolution))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteConsent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-consent [id]",
		Short: "Delete a new consent by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteConsent(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
