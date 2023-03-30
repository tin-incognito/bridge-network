package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTssKeySign = "tss_key_sign"

var _ sdk.Msg = &MsgTssKeySign{}

func NewMsgTssKeySign(creator string) *MsgTssKeySign {
	return &MsgTssKeySign{
		Creator: creator,
	}
}

func (msg *MsgTssKeySign) Route() string {
	return RouterKey
}

func (msg *MsgTssKeySign) Type() string {
	return TypeMsgTssKeySign
}

func (msg *MsgTssKeySign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTssKeySign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTssKeySign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
