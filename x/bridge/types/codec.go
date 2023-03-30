package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgKeygen{}, "bridge/Keygen", nil)
	cdc.RegisterConcrete(&MsgTssPool{}, "bridge/TssPool", nil)
	cdc.RegisterConcrete(&MsgTssKeySign{}, "bridge/TssKeySign", nil)
	cdc.RegisterConcrete(&MsgRegisterTssPool{}, "bridge/RegisterTssPool", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgKeygen{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTssPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTssKeySign{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterTssPool{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
