package crypto

import (
	"github.com/tendermint/go-amino-x"
)

func PrivKeyFromBytes(privKeyBytes []byte) (privKey PrivKey, err error) {
	err = amino.UnmarshalBinaryBare(privKeyBytes, &privKey)
	return
}

func PubKeyFromBytes(pubKeyBytes []byte) (pubKey PubKey, err error) {
	err = amino.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	return
}
