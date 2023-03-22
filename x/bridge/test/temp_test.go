package test

import (
	"bridge/app"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	ckeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"
)

type PubKey string

// NewPubKey create a new instance of PubKey
// key is bech32 encoded string
func NewPubKey(key string) (PubKey, error) {
	if len(key) == 0 {
		return "", nil
	}
	_, err := legacybech32.UnmarshalPubKey(legacybech32.AccPK, key)
	if err != nil {
		return "", fmt.Errorf("%s is not bech32 encoded pub key,err : %w", key, err)
	}
	return PubKey(key), nil
}

// GetKeyringKeybase return keyring and key info
func GetKeyringKeybase(chainHomeFolder, signerName, password string) (ckeys.Keyring, error) {
	if len(signerName) == 0 {
		return nil, fmt.Errorf("signer name is empty")
	}
	if len(password) == 0 {
		return nil, fmt.Errorf("password is empty")
	}

	buf := bytes.NewBufferString(password)
	// the library used by keyring is using ReadLine , which expect a new line
	buf.WriteByte('\n')
	kb, err := getKeybase(chainHomeFolder, buf)
	if err != nil {
		return nil, fmt.Errorf("fail to get keybase,err:%w", err)
	}
	// the keyring library which used by cosmos sdk , will use interactive terminal if it detect it has one
	// this will temporary trick it think there is no interactive terminal, thus will read the password from the buffer provided
	oldStdIn := os.Stdin
	defer func() {
		os.Stdin = oldStdIn
	}()
	os.Stdin = nil
	return kb, nil
}

// getKeybase will create an instance of Keybase
func getKeybase(home string, reader io.Reader) (ckeys.Keyring, error) {
	cliDir := home
	if len(home) == 0 {
		usr, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("fail to get current user,err:%w", err)
		}
		cliDir = filepath.Join(usr.HomeDir, ".bridge")
	}

	encodingConfig := app.MakeEncodingConfig()
	return ckeys.New(sdk.KeyringServiceName(), ckeys.BackendOS, cliDir, reader, encodingConfig.Marshaler)
}

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}

func TestTemp(t *testing.T) {
	initSDKConfig()

	kb, err := GetKeyringKeybase("", "validator2", "11234566")
	if err != nil {
		panic(err)
	}

	myValidator, err := kb.Key("validator2")
	if err != nil {
		panic(err)
	}
	pubKey, err := myValidator.GetPubKey()
	if err != nil {
		panic(err)
	}
	bech32PubKey, _ := legacybech32.MarshalPubKey(legacybech32.AccPK, pubKey)
	pk, _ := NewPubKey(bech32PubKey)
	fmt.Println("pk:", pk)

	//accs[i].Address = sdk.AccAddress(accs[i].PubKey.Address())

	//accs[i].ConsKey = ed25519.GenPrivKeyFromSecret(privkeySeed)
}
