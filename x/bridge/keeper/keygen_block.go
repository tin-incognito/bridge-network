package keeper

import (
	"bridge/x/bridge/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKeygenBlock set a specific keygenBlock in the store from its index
func (k Keeper) SetKeygenBlock(ctx sdk.Context, keygenBlock types.KeygenBlock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeygenBlockKeyPrefix))
	b := k.cdc.MustMarshal(&keygenBlock)
	store.Set(types.KeygenBlockKey(
		keygenBlock.Index,
	), b)
}

// GetKeygenBlock returns a keygenBlock from its index
func (k Keeper) GetKeygenBlock(
	ctx sdk.Context,
	index string,

) (val types.KeygenBlock, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeygenBlockKeyPrefix))

	b := store.Get(types.KeygenBlockKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKeygenBlock removes a keygenBlock from the store
func (k Keeper) RemoveKeygenBlock(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeygenBlockKeyPrefix))
	store.Delete(types.KeygenBlockKey(
		index,
	))
}

// GetAllKeygenBlock returns all keygenBlock
func (k Keeper) GetAllKeygenBlock(ctx sdk.Context) (list []types.KeygenBlock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeygenBlockKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KeygenBlock
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
