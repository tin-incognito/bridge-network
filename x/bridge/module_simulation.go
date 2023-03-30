package bridge

import (
	"math/rand"

	"bridge/testutil/sample"
	bridgesimulation "bridge/x/bridge/simulation"
	"bridge/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = bridgesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgKeygen = "op_weight_msg_keygen"
	// TODO: Determine the simulation weight value
	defaultWeightMsgKeygen int = 100

	opWeightMsgTssPool = "op_weight_msg_tss_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTssPool int = 100

	opWeightMsgTssKeySign = "op_weight_msg_tss_key_sign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTssKeySign int = 100

	opWeightMsgRegisterTssPool = "op_weight_msg_register_tss_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterTssPool int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bridgeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bridgeGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgKeygen int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgKeygen, &weightMsgKeygen, nil,
		func(_ *rand.Rand) {
			weightMsgKeygen = defaultWeightMsgKeygen
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgKeygen,
		bridgesimulation.SimulateMsgKeygen(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTssPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTssPool, &weightMsgTssPool, nil,
		func(_ *rand.Rand) {
			weightMsgTssPool = defaultWeightMsgTssPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTssPool,
		bridgesimulation.SimulateMsgTssPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTssKeySign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTssKeySign, &weightMsgTssKeySign, nil,
		func(_ *rand.Rand) {
			weightMsgTssKeySign = defaultWeightMsgTssKeySign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTssKeySign,
		bridgesimulation.SimulateMsgTssKeySign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterTssPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterTssPool, &weightMsgRegisterTssPool, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterTssPool = defaultWeightMsgRegisterTssPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterTssPool,
		bridgesimulation.SimulateMsgRegisterTssPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
