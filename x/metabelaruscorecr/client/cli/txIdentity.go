package cli

import (
	"bufio"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

func GetCmdCreateIdentity(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-Identity [accountID] [details] [creationDt] [idenitityType] [authPubKey]",
		Short: "Creates a new Identity",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAccountID := string(args[0])
			argsDetails := string(args[1])
			argsCreationDt := string(args[2])
			argsIdenitityType := string(args[3])
			argsAuthPubKey := string(args[4])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateIdentity(cliCtx.GetFromAddress(), string(argsAccountID), string(argsDetails), string(argsCreationDt), string(argsIdenitityType), string(argsAuthPubKey))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdSetIdentity(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-Identity [id]  [accountID] [details] [creationDt] [idenitityType] [authPubKey]",
		Short: "Set a new Identity",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {

			argsDetails := string(args[1])
			argsCreationDt := string(args[2])
			argsIdenitityType, err := strconv.Atoi(string(args[3]))
			if err != nil {
				return err
			}
			argsAuthPubKey := string(args[4])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			creationDt, err := time.ParseInLocation("2006-Jan-02", argsCreationDt, types.BelarusLocation)
			if err != nil {
				return err
			}

			msg, err := types.NewMsgSetIdentity(
				cliCtx.GetFromAddress(),
				argsIdenitityType,
				string(argsDetails),
				creationDt,
			)
			if err != nil {
				return err
			}

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteIdentity(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-Identity [id]",
		Short: "Delete a new Identity by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteIdentity(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
