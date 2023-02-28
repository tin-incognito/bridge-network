package bridge

import (
	"bridge/x/bridge/keeper"
	"bridge/x/bridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) []abci.ValidatorUpdate {
	// Set all the keygenBlock
	for _, elem := range genState.KeygenBlockList {
		k.SetKeygenBlock(ctx, elem)
	}

	// Set all the nodeAccount
	validators := []abci.ValidatorUpdate{}
	//validators := make([]abci.ValidatorUpdate, 0, len(genState.NodeAccountList))
	for _, nodeAccount := range genState.NodeAccountList {
		/*pk, err := cosmos.GetPubKeyFromBech32(cosmos.Bech32PubKeyTypeConsPub, nodeAccount.ValidatorConsPubKey)*/
		/*if err != nil {*/
		/*ctx.Logger().Error("fail to parse consensus public key", "key", nodeAccount.ValidatorConsPubKey, "error", err)*/
		/*panic(err)*/
		/*}*/
		/*validators = append(validators, abci.Ed25519ValidatorUpdate(pk.Bytes(), 100))*/

		k.SetNodeAccount(ctx, nodeAccount)
	}

	// Set nodeAccount count
	k.SetNodeAccountCount(ctx, genState.NodeAccountCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	return validators
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.KeygenBlockList = k.GetAllKeygenBlock(ctx)
	genesis.NodeAccountList = k.GetAllNodeAccount(ctx)
	genesis.NodeAccountCount = k.GetNodeAccountCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
