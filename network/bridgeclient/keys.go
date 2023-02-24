package bridgeclient

import (
	"bridge/app"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/crypto"
	ckeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// folder name for bridge network bridgeNetwork cli
	bridgeNetworkCliFolderName = `.bridge`
)

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
	_, err = kb.Key(signerName)
	if err != nil {
		return nil, fmt.Errorf("fail to get signer info(%s): %w", signerName, err)
	}
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
		cliDir = filepath.Join(usr.HomeDir, bridgeNetworkCliFolderName)
	}

	encodingConfig := app.MakeEncodingConfig()
	return ckeys.New(sdk.KeyringServiceName(), ckeys.BackendFile, cliDir, reader, encodingConfig.Marshaler)
}

// Keys manages all the keys used by thorchain
type Keys struct {
	signerName string
	password   string // TODO this is a bad way , need to fix it
	kb         ckeys.Keyring
}

// NewKeysWithKeybase create a new instance of Keys
func NewKeysWithKeybase(kb ckeys.Keyring, name, password string) *Keys {
	return &Keys{
		signerName: name,
		password:   password,
		kb:         kb,
	}
}

// GetPrivateKey return the private key
func (k *Keys) GetPrivateKey() (cryptotypes.PrivKey, error) {
	// return k.kb.ExportPrivateKeyObject(k.signerName)
	privKeyArmor, err := k.kb.ExportPrivKeyArmor(k.signerName, k.password)
	if err != nil {
		return nil, err
	}
	priKey, _, err := crypto.UnarmorDecryptPrivKey(privKeyArmor, k.password)
	if err != nil {
		return nil, fmt.Errorf("fail to unarmor private key: %w", err)
	}
	return priKey, nil
}
