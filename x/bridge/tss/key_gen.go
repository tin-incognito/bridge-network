package tss

import (
	"bridge/x/bridge/common"
	"bridge/x/bridge/keeper"
	"bridge/x/bridge/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TssManager struct {
	k keeper.Keeper
}

func (t *TssManager) TriggerKeygen(ctx sdk.Context, nodeAccounts common.NodeAccounts) error {
	members := []string{}
	for _, v := range nodeAccounts {
		members = append(members, v.PubKeySet.Secp256k1)
	}
	keygen, err := types.NewKeygen(ctx.BlockHeight(), members, 0)
	if err != nil {
		return fmt.Errorf("fail to create a new keygen: %w", err)
	}

}
