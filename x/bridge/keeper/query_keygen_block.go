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

func (k Keeper) KeygenBlockAll(goCtx context.Context, req *types.QueryAllKeygenBlockRequest) (*types.QueryAllKeygenBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var keygenBlocks []types.KeygenBlock
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	keygenBlockStore := prefix.NewStore(store, types.KeyPrefix(types.KeygenBlockKeyPrefix))

	pageRes, err := query.Paginate(keygenBlockStore, req.Pagination, func(key []byte, value []byte) error {
		var keygenBlock types.KeygenBlock
		if err := k.cdc.Unmarshal(value, &keygenBlock); err != nil {
			return err
		}

		keygenBlocks = append(keygenBlocks, keygenBlock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKeygenBlockResponse{KeygenBlock: keygenBlocks, Pagination: pageRes}, nil
}

func (k Keeper) KeygenBlock(goCtx context.Context, req *types.QueryGetKeygenBlockRequest) (*types.QueryGetKeygenBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetKeygenBlock(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKeygenBlockResponse{KeygenBlock: val}, nil
}
