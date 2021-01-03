package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

// CmdUnpackPrivKey - encryption helper command for external apps
/**
 * Allows to unpack private key with a correct blowfish algorithm.
 * This function is initally required to avoid php blowfish issue.
 */
func CmdUnpackPrivKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpack-privkey [Armored Privated Key] [Secretbox Password]",
		Short: "unpack private key",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			privateKeyArmored := args[0]
			password := args[1]

			pkUnarm, _, err := crypto.UnarmorDecryptPrivKey(privateKeyArmored, password)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(hex.EncodeToString(pkUnarm.Bytes()))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdReencryptPrivKey - encryption helper command for external apps
/**
 * The blockchain generate an account with a standard password.
 * The pk you get should be reencrypted with a more secured passowrd.
 */
func CmdReencryptPrivKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repack-privkey [Armored Private Key] [Old Secretbox Password] [New Password]",
		Short: "repack private key with a new password",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			privateKeyArmored := args[0]
			password := args[1]
			newPassword := args[2]

			pkUnarm, algo, err := crypto.UnarmorDecryptPrivKey(privateKeyArmored, password)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(crypto.EncryptArmorPrivKey(pkUnarm, newPassword, algo))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdDecryptPayload - encryption helper command for external apps
/**
 * If you unpacked a private key it's better to use your language
 * specific library rather than this hack
 */
func CmdDecryptPayload() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt-payload [Private Key] [Payload]",
		Short: "decrypt payload that was encrypted with public key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			privateKey, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}
			payload, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}

			decryptedPayload, err := mbutils.DecryptPayload(privateKey, payload)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(string(decryptedPayload))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
