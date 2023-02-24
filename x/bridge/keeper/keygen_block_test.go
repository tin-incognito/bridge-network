package keeper_test

import (
	"strconv"
	"testing"

	keepertest "bridge/testutil/keeper"
	"bridge/testutil/nullify"
	"bridge/x/bridge/keeper"
	"bridge/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNKeygenBlock(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.KeygenBlock {
	items := make([]types.KeygenBlock, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetKeygenBlock(ctx, items[i])
	}
	return items
}

func TestKeygenBlockGet(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNKeygenBlock(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetKeygenBlock(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestKeygenBlockRemove(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNKeygenBlock(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKeygenBlock(ctx,
			item.Index,
		)
		_, found := keeper.GetKeygenBlock(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestKeygenBlockGetAll(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNKeygenBlock(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKeygenBlock(ctx)),
	)
}
