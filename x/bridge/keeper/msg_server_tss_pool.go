package keeper

import (
	"context"

	"bridge/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TssPool(goCtx context.Context, msg *types.MsgTssPool) (*types.MsgTssPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTssPoolResponse{}, nil
}
