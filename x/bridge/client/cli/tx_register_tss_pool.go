package cli

import (
	"strconv"

	"bridge/x/bridge/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRegisterTssPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-tss-pool",
		Short: "Broadcast message RegisterTssPool",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRegisterTssPool{}

			/*msg := types.NewMsgRegisterTssPool(*/
			/*clientCtx.GetFromAddress().String(),*/
			/*)*/
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
