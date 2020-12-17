package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/types"
)

func GetCmdCreateAllowance(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-allowance [passportId] [allowanceType] [resolution]",
		Short: "Creates a new allowance",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPassportId := string(args[0] )
			argsAllowanceType := string(args[1] )
			argsResolution := string(args[2] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateAllowance(cliCtx.GetFromAddress(), string(argsPassportId), string(argsAllowanceType), string(argsResolution))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetAllowance(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-allowance [id]  [passportId] [allowanceType] [resolution]",
		Short: "Set a new allowance",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsPassportId := string(args[1])
			argsAllowanceType := string(args[2])
			argsResolution := string(args[3])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetAllowance(cliCtx.GetFromAddress(), id, string(argsPassportId), string(argsAllowanceType), string(argsResolution))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteAllowance(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-allowance [id]",
		Short: "Delete a new allowance by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteAllowance(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
