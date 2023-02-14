package keeper_test

import (
	"context"
	"testing"

	keepertest "bridge/testutil/keeper"
	"bridge/x/bridge/keeper"
	"bridge/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BridgeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
