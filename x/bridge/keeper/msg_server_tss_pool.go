package keeper

import (
	"context"
	"strconv"

	"bridge/x/bridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TssPool(goCtx context.Context, msg *types.MsgTssPool) (*types.MsgTssPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	resgisterKeygen := &types.RegisterKeygen{
		Index:      strconv.FormatInt(msg.Height, 10),
		Height:     msg.Height,
		PoolPubKey: msg.PoolPubKey,
		Members:    msg.PubKeys,
	}
	k.SetRegisterKeygen(ctx, *resgisterKeygen)

	return &types.MsgTssPoolResponse{}, nil
}
