package node

import (
	amino "github.com/tendermint/go-amino-x"
	cryptoAmino "github.com/tendermint/classic/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
