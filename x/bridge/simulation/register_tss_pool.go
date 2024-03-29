package simulation

import (
	"math/rand"

	"bridge/x/bridge/keeper"
	"bridge/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRegisterTssPool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterTssPool{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RegisterTssPool simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterTssPool simulation not implemented"), nil, nil
	}
}
