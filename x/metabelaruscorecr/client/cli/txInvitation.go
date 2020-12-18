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

func GetCmdCreateInvitation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-invitation [inviterId] [identityId] [creationDate] [activationPubKey]",
		Short: "Creates a new invitation",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsInviterId := string(args[0] )
			argsIdentityId := string(args[1] )
			argsCreationDate := string(args[2] )
			argsActivationPubKey := string(args[3] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateInvitation(cliCtx.GetFromAddress(), string(argsInviterId), string(argsIdentityId), string(argsCreationDate), string(argsActivationPubKey))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetInvitation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-invitation [id]  [inviterId] [identityId] [creationDate] [activationPubKey]",
		Short: "Set a new invitation",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsInviterId := string(args[1])
			argsIdentityId := string(args[2])
			argsCreationDate := string(args[3])
			argsActivationPubKey := string(args[4])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetInvitation(cliCtx.GetFromAddress(), id, string(argsInviterId), string(argsIdentityId), string(argsCreationDate), string(argsActivationPubKey))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteInvitation(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-invitation [id]",
		Short: "Delete a new invitation by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteInvitation(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
