package keeper

import (
	"context"

	"bridge/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Keygen(goCtx context.Context, msg *types.MsgKeygen) (*types.MsgKeygenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgKeygenResponse{}, nil
}
