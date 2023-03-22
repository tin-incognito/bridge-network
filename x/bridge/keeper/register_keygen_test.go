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

func createNRegisterKeygen(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RegisterKeygen {
	items := make([]types.RegisterKeygen, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRegisterKeygen(ctx, items[i])
	}
	return items
}

func TestRegisterKeygenGet(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNRegisterKeygen(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRegisterKeygen(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRegisterKeygenRemove(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNRegisterKeygen(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegisterKeygen(ctx,
			item.Index,
		)
		_, found := keeper.GetRegisterKeygen(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRegisterKeygenGetAll(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNRegisterKeygen(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegisterKeygen(ctx)),
	)
}
