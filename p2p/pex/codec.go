package pex

import (
	amino "github.com/tendermint/go-amino-x"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	RegisterPexMessage(cdc)
}
