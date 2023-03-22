package bridge_test

import (
	"testing"

	keepertest "bridge/testutil/keeper"
	"bridge/testutil/nullify"
	"bridge/x/bridge"
	"bridge/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KeygenBlockList: []types.KeygenBlock{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		NodeAccountList: []types.NodeAccount{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		NodeAccountCount: 2,
		RegisterKeygenList: []types.RegisterKeygen{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BridgeKeeper(t)
	bridge.InitGenesis(ctx, *k, genesisState)
	got := bridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.KeygenBlockList, got.KeygenBlockList)
	require.ElementsMatch(t, genesisState.NodeAccountList, got.NodeAccountList)
	require.Equal(t, genesisState.NodeAccountCount, got.NodeAccountCount)
	require.ElementsMatch(t, genesisState.RegisterKeygenList, got.RegisterKeygenList)
	// this line is used by starport scaffolding # genesis/test/assert
}
