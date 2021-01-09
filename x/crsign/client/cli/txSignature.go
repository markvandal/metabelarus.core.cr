package cli

import (
  
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func CmdCreateSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-signature [identity] [service] [key] [secret] [creationDt] [availabilityDt]",
		Short: "Creates a new signature",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsIdentity := string(args[0])
      argsService := string(args[1])
      argsKey := string(args[2])
      argsSecret := string(args[3])
      argsCreationDt := string(args[4])
      argsAvailabilityDt := string(args[5])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSignature(clientCtx.GetFromAddress().String(), string(argsIdentity), string(argsService), string(argsKey), string(argsSecret), string(argsCreationDt), string(argsAvailabilityDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdateSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-signature [id] [identity] [service] [key] [secret] [creationDt] [availabilityDt]",
		Short: "Update a signature",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]
      argsIdentity := string(args[1])
      argsService := string(args[2])
      argsKey := string(args[3])
      argsSecret := string(args[4])
      argsCreationDt := string(args[5])
      argsAvailabilityDt := string(args[6])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSignature(clientCtx.GetFromAddress().String(), id, string(argsIdentity), string(argsService), string(argsKey), string(argsSecret), string(argsCreationDt), string(argsAvailabilityDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeleteSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-signature [id] [identity] [service] [key] [secret] [creationDt] [availabilityDt]",
		Short: "Delete a signature by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]

        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSignature(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
