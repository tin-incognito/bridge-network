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

func CmdTssPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tss-pool",
		Short: "Broadcast message TssPool",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgTssPool{}

			/*msg := types.NewMsgTssPool(*/
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
