package tss

import (
	"bridge/x/bridge/keeper"
	"bridge/x/bridge/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TssManager struct {
	k keeper.Keeper
}

func NewTssManager(k keeper.Keeper) *TssManager {
	return &TssManager{k: k}
}

func (t *TssManager) TriggerKeygen(ctx sdk.Context, nodeAccounts []types.NodeAccount) error {
	members := []string{}
	for _, v := range nodeAccounts {
		members = append(members, v.PubKeySet.Secp256K1)
	}
	_, err := types.NewKeygen(ctx.BlockHeight(), members, 0)
	if err != nil {
		return fmt.Errorf("fail to create a new keygen: %w", err)
	}
	return nil
}
