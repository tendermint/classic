package main

import (
	"fmt"
	"os"

	amino "github.com/tendermint/go-amino-x"
	cryptoAmino "github.com/tendermint/classic/crypto/encoding/amino"
)

func main() {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
