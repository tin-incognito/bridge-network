package types

import (
	"crypto/sha256"
	"encoding/hex"
	fmt "fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTssPool = "tss_pool"

var _ sdk.Msg = &MsgTssPool{}

// NewMsgTssPool is a constructor function for MsgTssPool
func NewMsgTssPool(
	creator string, pks []string,
	poolpk string, keygenType int32,
	height int64, bl *Blame, chains []string,
	keygenTime int64,
) (*MsgTssPool, error) {
	id, err := getTssID(pks, poolpk, height, bl)
	if err != nil {
		return nil, fmt.Errorf("fail to get tss id: %w", err)
	}
	return &MsgTssPool{
		Creator:    creator,
		Id:         id,
		PubKeys:    pks,
		PoolPubKey: poolpk,
		Height:     height,
		KeygenType: keygenType,
		Blame:      bl,
		Chains:     chains,
		KeygenTime: keygenTime,
	}, nil
}

// getTssID
func getTssID(members []string, poolPk string, height int64, bl *Blame) (string, error) {
	// ensure input pubkeys list is deterministically sorted
	sort.SliceStable(members, func(i, j int) bool {
		return members[i] < members[j]
	})

	pubkeys := make([]string, len(bl.BlameNodes))
	for i, node := range bl.BlameNodes {
		pubkeys[i] = node.Pubkey
	}
	sort.SliceStable(pubkeys, func(i, j int) bool {
		return pubkeys[i] < pubkeys[j]
	})

	sb := strings.Builder{}
	for _, item := range members {
		sb.WriteString("m:" + item)
	}
	for _, item := range pubkeys {
		sb.WriteString("p:" + item)
	}
	sb.WriteString(poolPk)
	sb.WriteString(fmt.Sprintf("%d", height))
	hash := sha256.New()
	_, err := hash.Write([]byte(sb.String()))
	if err != nil {
		return "", fmt.Errorf("fail to get tss id: %w", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (msg *MsgTssPool) Route() string {
	return RouterKey
}

func (msg *MsgTssPool) Type() string {
	return TypeMsgTssPool
}

func (msg *MsgTssPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTssPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTssPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
