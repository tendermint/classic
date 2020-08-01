package mempool

import (
	amino "github.com/tendermint/go-amino-x"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMempoolMessages(cdc)
}
