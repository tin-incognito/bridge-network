package keeper_test

import (
	"testing"

	keepertest "bridge/testutil/keeper"
	"bridge/testutil/nullify"
	"bridge/x/bridge/keeper"
	"bridge/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNNodeAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NodeAccount {
	items := make([]types.NodeAccount, n)
	for i := range items {
		items[i].Id = keeper.AppendNodeAccount(ctx, items[i])
	}
	return items
}

func TestNodeAccountGet(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNNodeAccount(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetNodeAccount(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestNodeAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNNodeAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNodeAccount(ctx, item.Id)
		_, found := keeper.GetNodeAccount(ctx, item.Id)
		require.False(t, found)
	}
}

func TestNodeAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNNodeAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNodeAccount(ctx)),
	)
}

func TestNodeAccountCount(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNNodeAccount(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetNodeAccountCount(ctx))
}
