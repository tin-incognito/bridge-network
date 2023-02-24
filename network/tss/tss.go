package tss

import (
	"bridge/network/p2p"
	"bridge/network/tss/storage"
	"fmt"

	coskey "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"

	//bkeygen "github.com/bnb-chain/tss-lib/ecdsa/keygen"

	tcrypto "github.com/tendermint/tendermint/crypto"
)

type Tss struct {
}

func NewTss(
	cmdBootstrapPeers p2p.AddrList,
	p2pPort int,
	priKey tcrypto.PrivKey,
	rendezvous,
	baseFolder string,
	conf Config,
	//preParams *bkeygen.LocalPreParams,
	externalIP string,
) (*Tss, error) {
	pk := coskey.PubKey{
		Key: priKey.PubKey().Bytes()[:],
	}

	pubKey, err := sdk.MarshalPubKey(sdk.AccPK, &pk)
	if err != nil {
		return nil, fmt.Errorf("fail to genearte the key: %w", err)
	}

	stateManager, err := storage.NewFileStateMgr(baseFolder)
	if err != nil {
		return nil, fmt.Errorf("fail to create file state manager")
	}

	var bootstrapPeers p2p.AddrList
	savedPeers, err := stateManager.RetrieveP2PAddresses()
	if err != nil {
		bootstrapPeers = cmdBootstrapPeers
	} else {
		bootstrapPeers = savedPeers
		bootstrapPeers = append(bootstrapPeers, cmdBootstrapPeers...)
	}

	return &Tss{}, nil
}

func (t *Tss) Start() error {
	return nil
}
