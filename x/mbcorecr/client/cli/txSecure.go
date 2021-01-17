package cli

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/crypto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func CmdDecrypt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt [Payload]",
		Short: "Decrypt payload encrypted with public key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			payload := string(args[0])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			armor, err := clientCtx.Keyring.ExportPrivKeyArmorByAddress(
				clientCtx.GetFromAddress(),
				types.UnsecureNewAcctountPKPassword,
			)
			if err != nil {
				return err
			}

			pk, _, err := crypto.UnarmorDecryptPrivKey(armor, types.UnsecureNewAcctountPKPassword)
			if err != nil {
				return err
			}

			encrypted, err := base64.URLEncoding.DecodeString(payload)
			if err != nil {
				return err
			}

			data, err := mbutils.DecryptPayload(pk.Bytes(), encrypted)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", data))
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
