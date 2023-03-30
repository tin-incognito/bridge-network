package keeper

import (
	"context"

	"bridge/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TssKeySign(goCtx context.Context, msg *types.MsgTssKeySign) (*types.MsgTssKeySignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTssKeySignResponse{}, nil
}
