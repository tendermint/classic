package types

import (
	amino "github.com/tendermint/go-amino-x"
	"github.com/tendermint/classic/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}
