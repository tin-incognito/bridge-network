package keeper

import (
	"context"

	"bridge/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RegisterKeygenAll(goCtx context.Context, req *types.QueryAllRegisterKeygenRequest) (*types.QueryAllRegisterKeygenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registerKeygens []types.RegisterKeygen
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	registerKeygenStore := prefix.NewStore(store, types.KeyPrefix(types.RegisterKeygenKeyPrefix))

	pageRes, err := query.Paginate(registerKeygenStore, req.Pagination, func(key []byte, value []byte) error {
		var registerKeygen types.RegisterKeygen
		if err := k.cdc.Unmarshal(value, &registerKeygen); err != nil {
			return err
		}

		registerKeygens = append(registerKeygens, registerKeygen)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegisterKeygenResponse{RegisterKeygen: registerKeygens, Pagination: pageRes}, nil
}

func (k Keeper) RegisterKeygen(goCtx context.Context, req *types.QueryGetRegisterKeygenRequest) (*types.QueryGetRegisterKeygenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetRegisterKeygen(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRegisterKeygenResponse{RegisterKeygen: val}, nil
}
