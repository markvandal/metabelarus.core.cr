package cli

import (
  
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func CmdCreateAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-auth [identity] [service] [key] [status] [creationDt] [availabilityDt]",
		Short: "Creates a new auth",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsIdentity := string(args[0])
      argsService := string(args[1])
      argsKey := string(args[2])
      argsStatus := string(args[3])
      argsCreationDt := string(args[4])
      argsAvailabilityDt := string(args[5])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAuth(clientCtx.GetFromAddress().String(), string(argsIdentity), string(argsService), string(argsKey), string(argsStatus), string(argsCreationDt), string(argsAvailabilityDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdateAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-auth [id] [identity] [service] [key] [status] [creationDt] [availabilityDt]",
		Short: "Update a auth",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]
      argsIdentity := string(args[1])
      argsService := string(args[2])
      argsKey := string(args[3])
      argsStatus := string(args[4])
      argsCreationDt := string(args[5])
      argsAvailabilityDt := string(args[6])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAuth(clientCtx.GetFromAddress().String(), id, string(argsIdentity), string(argsService), string(argsKey), string(argsStatus), string(argsCreationDt), string(argsAvailabilityDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeleteAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-auth [id] [identity] [service] [key] [status] [creationDt] [availabilityDt]",
		Short: "Delete a auth by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]

        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAuth(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
