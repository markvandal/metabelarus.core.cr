package cli

import (
  
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

func CmdCreateInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-invite [inviter] [invitee] [level] [key] [creationDt]",
		Short: "Creates a new invite",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsInviter := string(args[0])
      argsInvitee := string(args[1])
      argsLevel := string(args[2])
      argsKey := string(args[3])
      argsCreationDt := string(args[4])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateInvite(clientCtx.GetFromAddress().String(), string(argsInviter), string(argsInvitee), string(argsLevel), string(argsKey), string(argsCreationDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdateInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-invite [id] [inviter] [invitee] [level] [key] [creationDt]",
		Short: "Update a invite",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]
      argsInviter := string(args[1])
      argsInvitee := string(args[2])
      argsLevel := string(args[3])
      argsKey := string(args[4])
      argsCreationDt := string(args[5])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateInvite(clientCtx.GetFromAddress().String(), id, string(argsInviter), string(argsInvitee), string(argsLevel), string(argsKey), string(argsCreationDt))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeleteInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-invite [id] [inviter] [invitee] [level] [key] [creationDt]",
		Short: "Delete a invite by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            id := args[0]

        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteInvite(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
