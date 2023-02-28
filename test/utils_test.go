package test

import (
	"bridge/x/bridge/tss"
	"fmt"
	"testing"
)

func TestTemp(t *testing.T) {
	fmt.Println(tss.GetRandomValidatorNode())
	/*a := simulation.Account{}*/
	/*a.PrivKey*/
	/*a.PrivKey = secp256k1.GenPrivKeyFromSecret(privkeySeed)*/
	/*a.PubKey = a.PrivKey.PubKey()*/
	/*a.Address = sdk.AccAddress(a.PubKey.Address())*/

	/*a.ConsKey = ed25519.GenPrivKeyFromSecret(privkeySeed)*/

}
