package cli

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

// GetCmdCreateIdentity - create identity via command line
/**
 * @TODO this function is for debug and genesis purposes only.
 * It shouldn't be supported on production.
 * @TODO all arguments should be set via optional flags
 */
func GetCmdCreateIdentity(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-identity [details] [idenitityType] [authPubKey]",
		Short: "Creates a new Identity",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			argsDetails := string(args[0])
			argsIdenitityType, err := strconv.Atoi(string(args[1]))
			if err != nil {
				return err
			}
			argsAuthPubKey := string(args[2])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgCreateIdentity(
				cliCtx.GetFromAddress(),
				types.IdentityType(argsIdenitityType),
				argsDetails,
				argsAuthPubKey,
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSetIdentity - change identity via command line
/**
 * @TODO this function is for debug and genesis purposes only.
 * It shouldn't be supported on production.
 * @TODO all arguments should be set via optional flags
 */
func GetCmdSetIdentity(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-identity [ID] [Details] [IdentityType] [AuthPubKey]",
		Short: "Set a new Identity",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			argsID := string(args[0])
			argsDetails := string(args[1])
			argsIdenitityType, err := strconv.Atoi(string(args[2]))
			if err != nil {
				return err
			}
			argsAuthPubKey := string(args[3])

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QueryGetIdentity, argsID), nil)
			if err != nil {
				return err
			}

			var identity types.Identity
			cdc.MustUnmarshalJSON(res, &identity)
			cliCtx.PrintOutput("Previous state: ")
			cliCtx.PrintOutput(identity)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgSetIdentity(
				argsID,
				identity.AccountID,
				types.IdentityType(argsIdenitityType),
				string(argsDetails),
				argsAuthPubKey,
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
