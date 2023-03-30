package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/types"

	"bridge/app"
	prefix "bridge/cmd"
	"bridge/cmd/bridged/cmd"
)

func main() {
	/*config := cosmos.GetConfig()*/
	/*config.SetBech32PrefixForAccount(prefix.Bech32PrefixAccAddr, prefix.Bech32PrefixAccPub)*/
	/*config.SetBech32PrefixForValidator(prefix.Bech32PrefixValAddr, prefix.Bech32PrefixValPub)*/
	/*config.SetBech32PrefixForConsensusNode(prefix.Bech32PrefixConsAddr, prefix.Bech32PrefixConsPub)*/
	/*config.SetCoinType(prefix.BridgeChainCoinType)*/
	/*config.SetPurpose(prefix.BridgeChainCoinPurpose)*/
	/*config.Seal()*/

	types.SetCoinDenomRegex(func() string {
		return prefix.DenomRegex
	})

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
