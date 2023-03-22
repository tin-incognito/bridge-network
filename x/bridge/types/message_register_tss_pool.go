package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterTssPool = "register_tss_pool"

var _ sdk.Msg = &MsgRegisterTssPool{}

func NewMsgRegisterTssPool(creator, poolPubKey, signature string, members []string) *MsgRegisterTssPool {
	return &MsgRegisterTssPool{
		Creator:    creator,
		Members:    members,
		PoolPubKey: poolPubKey,
		Signature:  signature,
	}
}

func (msg *MsgRegisterTssPool) Route() string {
	return RouterKey
}

func (msg *MsgRegisterTssPool) Type() string {
	return TypeMsgRegisterTssPool
}

func (msg *MsgRegisterTssPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterTssPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterTssPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
