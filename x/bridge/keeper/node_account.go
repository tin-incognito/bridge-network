package keeper

import (
	"encoding/binary"

	"bridge/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetNodeAccountCount get the total number of nodeAccount
func (k Keeper) GetNodeAccountCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.NodeAccountCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetNodeAccountCount set the total number of nodeAccount
func (k Keeper) SetNodeAccountCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.NodeAccountCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendNodeAccount appends a nodeAccount in the store with a new id and update the count
func (k Keeper) AppendNodeAccount(
	ctx sdk.Context,
	nodeAccount types.NodeAccount,
) uint64 {
	// Create the nodeAccount
	count := k.GetNodeAccountCount(ctx)

	// Set the ID of the appended value
	nodeAccount.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeAccountKey))
	appendedValue := k.cdc.MustMarshal(&nodeAccount)
	store.Set(GetNodeAccountIDBytes(nodeAccount.Id), appendedValue)

	// Update nodeAccount count
	k.SetNodeAccountCount(ctx, count+1)

	return count
}

// SetNodeAccount set a specific nodeAccount in the store
func (k Keeper) SetNodeAccount(ctx sdk.Context, nodeAccount types.NodeAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeAccountKey))
	b := k.cdc.MustMarshal(&nodeAccount)
	store.Set(GetNodeAccountIDBytes(nodeAccount.Id), b)
}

// GetNodeAccount returns a nodeAccount from its id
func (k Keeper) GetNodeAccount(ctx sdk.Context, id uint64) (val types.NodeAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeAccountKey))
	b := store.Get(GetNodeAccountIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNodeAccount removes a nodeAccount from the store
func (k Keeper) RemoveNodeAccount(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeAccountKey))
	store.Delete(GetNodeAccountIDBytes(id))
}

// GetAllNodeAccount returns all nodeAccount
func (k Keeper) GetAllNodeAccount(ctx sdk.Context) (list []types.NodeAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeAccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NodeAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetNodeAccountIDBytes returns the byte representation of the ID
func GetNodeAccountIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetNodeAccountIDFromBytes returns ID in uint64 format from a byte array
func GetNodeAccountIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
