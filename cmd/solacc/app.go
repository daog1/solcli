package main

import (
	"encoding/hex"
	"os"
	"path"
	"path/filepath"

	"github.com/gagliardetto/solana-go"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var LocalPrivateKey = "/.config/solana/id.json"

func getLocalPrivateKey() (solana.PrivateKey, error) {
	// Load the account that you will send funds FROM:
	abspath := path.Join(os.Getenv("HOME"), LocalPrivateKey)
	accountFrom, err := getFileKey(abspath)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}

	return accountFrom, nil
}
func getFileKey(abspath string) (solana.PrivateKey, error) {
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile(abspath)

	if err != nil {
		return nil, err
	}
	log.Debug("account load public key:", accountFrom.PublicKey().String())
	return accountFrom, nil
}
func main() {
	log.SetLevel(logrus.DebugLevel)
	args := os.Args
	var wallet solana.PrivateKey
	if len(args) > 1 {
		if filepath.IsAbs(args[1]) {
			wallet, _ = getFileKey(args[1])
		}
	} else {
		wallet, _ = getLocalPrivateKey()
	}
	log.Debug("wallet hex key:", hex.EncodeToString(wallet))
	log.Debug("wallet key:", wallet.String())
}
