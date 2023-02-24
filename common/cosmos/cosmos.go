package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"
)

var (
	GetConfig               = sdk.GetConfig
	NewContext              = sdk.NewContext
	GetPubKeyFromBech32     = legacybech32.UnmarshalPubKey // nolint SA1019 deprecated
	Bech32ifyPubKey         = legacybech32.MarshalPubKey   // nolint SA1019 deprecated
	Bech32PubKeyTypeConsPub = legacybech32.ConsPK
	Bech32PubKeyTypeAccPub  = legacybech32.AccPK
)
