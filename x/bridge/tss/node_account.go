package tss

import (
	"bridge/x/bridge/types"
	"math/rand"
	"time"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// GetRandomValidatorNode creates a random validator node account, used for testing
func GetRandomValidatorNode() *types.NodeAccount {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s) // #nosec G404 this is a method only used for test purpose
	accts := simtypes.RandomAccounts(r, 1)
	account := accts[0]

	/*k, _ := cosmos.Bech32ifyPubKey(cosmos.Bech32PubKeyTypeConsPub, accts[0].PubKey)*/
	/*pubKeys := types.PubKeySet{*/
	/*Secp256K1: GetRandomPubKey(),*/
	/*Ed25519:   GetRandomPubKey(),*/
	/*}*/
	/*fmt.Println(pubKeys)*/
	/*fmt.Println(k)*/
	//fmt.Println(account.PrivKey.String())

	return &types.NodeAccount{
		NodeAddress: account.Address.String(),
		PubKeySet: &types.PubKeySet{
			Secp256K1: account.PubKey.String(),
			Ed25519:   account.ConsKey.PubKey().String(),
		},
	}
}

/*func GetRandomPubKey() string {*/
/*r := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404*/
/*accts := simtypes.RandomAccounts(r, 1)*/
/*bech32PubKey, _ := legacybech32.UnmarshalPubKey(cosmos.Bech32PubKeyTypeAccPub, accts[0].PubKey.String())*/
/*pk, _ := NewPubKey(bech32PubKey.String())*/
/*return pk*/
/*}*/

/*// NewPubKey create a new instance of PubKey*/
/*// key is bech32 encoded string*/
/*func NewPubKey(key string) (string, error) {*/
/*if len(key) == 0 {*/
/*return "", nil*/
/*}*/
/*_, err := legacybech32.UnmarshalPubKey(legacybech32.AccPK, key)*/
/*if err != nil {*/
/*return "", fmt.Errorf("%s is not bech32 encoded pub key,err : %w", key, err)*/
/*}*/
/*return key, nil*/
/*}*/
