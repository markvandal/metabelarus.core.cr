package cli

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func CmdAcceptInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-invite [Invite Id] [Sequence]",
		Short: "Accept an new invite",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsInviteID := string(args[0])
			sequence, err := base64.URLEncoding.DecodeString(string(args[1]))
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			uid := uuid.New().String()
			info, err := clientCtx.Keyring.NewAccount(
				uid,
				string(sequence),
				types.UnsecureNewAcctountPKPassword,
				types.DefaultWalletPath,
				hd.Secp256k1,
			)
			if err != nil {
				return err
			}

			fromAddr, fromName, err := client.GetFromFields(clientCtx.Keyring, info.GetAddress().String(), clientCtx.GenerateOnly)
			if err != nil {
				return err
			}

			clientCtx = clientCtx.WithFrom(uid).WithFromAddress(fromAddr).WithFromName(fromName)
			clientCtx.OutputFormat = "text"

			msg := types.NewMsgAcceptInvite(argsInviteID, info.GetAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			res := tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

			exportedPk, err := clientCtx.Keyring.ExportPrivKeyArmor(uid, types.UnsecureNewAcctountPKPassword)
			if err != nil {
				return err
			}

			clientCtx.Keyring.Delete(uid)

			writer := clientCtx.Output
			if writer == nil {
				writer = os.Stdout
			}

			_, err = writer.Write([]byte(
				strings.Replace(exportedPk, "\n", "\\\\n", -1),
			))
			if err != nil {
				return err
			}

			return res
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

			msg := types.NewMsgCreateInvite(
				clientCtx.GetFromAddress().String(),
				types.IdentityLevel(identityLevel),
				types.IdentityType(identityType),
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
