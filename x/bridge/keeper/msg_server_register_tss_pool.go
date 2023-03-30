package keeper

import (
	"context"

	"bridge/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterTssPool(goCtx context.Context, msg *types.MsgRegisterTssPool) (*types.MsgRegisterTssPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterTssPoolResponse{}, nil
}
