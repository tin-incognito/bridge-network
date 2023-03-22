package keeper

import (
	"bridge/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRegisterKeygen set a specific registerKeygen in the store from its index
func (k Keeper) SetRegisterKeygen(ctx sdk.Context, registerKeygen types.RegisterKeygen) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterKeygenKeyPrefix))
	b := k.cdc.MustMarshal(&registerKeygen)
	store.Set(types.RegisterKeygenKey(
		registerKeygen.Index,
	), b)
}

// GetRegisterKeygen returns a registerKeygen from its index
func (k Keeper) GetRegisterKeygen(
	ctx sdk.Context,
	index string,

) (val types.RegisterKeygen, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterKeygenKeyPrefix))

	b := store.Get(types.RegisterKeygenKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegisterKeygen removes a registerKeygen from the store
func (k Keeper) RemoveRegisterKeygen(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterKeygenKeyPrefix))
	store.Delete(types.RegisterKeygenKey(
		index,
	))
}

// GetAllRegisterKeygen returns all registerKeygen
func (k Keeper) GetAllRegisterKeygen(ctx sdk.Context) (list []types.RegisterKeygen) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegisterKeygenKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RegisterKeygen
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
