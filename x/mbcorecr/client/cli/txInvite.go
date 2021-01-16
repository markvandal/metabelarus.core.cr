package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/helper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func CmdAcceptInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-invite [Invite Id]",
		Short: "Accept an new invite",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsInviteID := string(args[0])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			inviteHelper := helper.NewClientInviteHelper(&clientCtx)
			err = inviteHelper.ProduceAccount()
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptInvite(
				argsInviteID,
				clientCtx.GetFromAddress().String(),
				inviteHelper.GetAddressString(),
				inviteHelper.GetPubKeyString(),
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

func CmdCreateInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use: fmt.Sprintf(
			"create-invite [Level: %s] [Identity Type: %s]",
			strings.Join(mbutils.EnumMapToList(types.IdentityLevel_name), "/"),
			strings.Join(mbutils.EnumMapToList(types.IdentityType_name), "/"),
		),
		Short: "Creates a new invite",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsLevel := string(args[0])
			argsType := string(args[1])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			identityLevel, ok := types.IdentityLevel_value[argsLevel]
			if !ok {
				return fmt.Errorf("Identity level: %s does not exist", argsLevel)
			}

			identityType, ok := types.IdentityType_value[argsType]
			if !ok {
				return fmt.Errorf("Identity type: %s does not exist", argsType)
			}

			inviteHelper := helper.NewClientInviteHelper(&clientCtx)
			err = inviteHelper.ProduceAccount()
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateInvite(
				clientCtx.GetFromAddress().String(),
				types.IdentityLevel(identityLevel),
				types.IdentityType(identityType),
				inviteHelper.GetAddressString(),
				inviteHelper.GetPubKeyString(),
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
