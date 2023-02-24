package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgKeygen = "keygen"

var _ sdk.Msg = &MsgKeygen{}

func NewMsgKeygen(creator string, t int32, data string) *MsgKeygen {
	return &MsgKeygen{
		Creator: creator,
		T:       t,
		Data:    data,
	}
}

func (msg *MsgKeygen) Route() string {
	return RouterKey
}

func (msg *MsgKeygen) Type() string {
	return TypeMsgKeygen
}

func (msg *MsgKeygen) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgKeygen) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgKeygen) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
