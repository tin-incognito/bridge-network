package keeper

import (
	"context"

	"bridge/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NodeAccountAll(goCtx context.Context, req *types.QueryAllNodeAccountRequest) (*types.QueryAllNodeAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nodeAccounts []types.NodeAccount
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	nodeAccountStore := prefix.NewStore(store, types.KeyPrefix(types.NodeAccountKey))

	pageRes, err := query.Paginate(nodeAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var nodeAccount types.NodeAccount
		if err := k.cdc.Unmarshal(value, &nodeAccount); err != nil {
			return err
		}

		nodeAccounts = append(nodeAccounts, nodeAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNodeAccountResponse{NodeAccount: nodeAccounts, Pagination: pageRes}, nil
}

func (k Keeper) NodeAccount(goCtx context.Context, req *types.QueryGetNodeAccountRequest) (*types.QueryGetNodeAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	nodeAccount, found := k.GetNodeAccount(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetNodeAccountResponse{NodeAccount: nodeAccount}, nil
}
