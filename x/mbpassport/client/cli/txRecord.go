package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
)

func GetCmdCreateRecord(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-record [IdentityId] [ServiceId] [ServiceType] [Key] [UserValue] [ServiceValue] [CreationDt] [UpdateDt]",
		Short: "Creates a new record",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsIdentityId := string(args[0] )
			argsServiceId := string(args[1] )
			argsServiceType := string(args[2] )
			argsKey := string(args[3] )
			argsUserValue := string(args[4] )
			argsServiceValue := string(args[5] )
			argsCreationDt := string(args[6] )
			argsUpdateDt := string(args[7] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateRecord(cliCtx.GetFromAddress(), string(argsIdentityId), string(argsServiceId), string(argsServiceType), string(argsKey), string(argsUserValue), string(argsServiceValue), string(argsCreationDt), string(argsUpdateDt))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetRecord(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-record [id]  [IdentityId] [ServiceId] [ServiceType] [Key] [UserValue] [ServiceValue] [CreationDt] [UpdateDt]",
		Short: "Set a new record",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsIdentityId := string(args[1])
			argsServiceId := string(args[2])
			argsServiceType := string(args[3])
			argsKey := string(args[4])
			argsUserValue := string(args[5])
			argsServiceValue := string(args[6])
			argsCreationDt := string(args[7])
			argsUpdateDt := string(args[8])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetRecord(cliCtx.GetFromAddress(), id, string(argsIdentityId), string(argsServiceId), string(argsServiceType), string(argsKey), string(argsUserValue), string(argsServiceValue), string(argsCreationDt), string(argsUpdateDt))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteRecord(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-record [id]",
		Short: "Delete a new record by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteRecord(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
