package main

import (
	"github.com/tendermint/classic/crypto/ed25519"
	"github.com/tendermint/classic/crypto/merkle"
	"github.com/tendermint/classic/crypto/multisig"
	"github.com/tendermint/classic/crypto/secp256k1"
	"github.com/tendermint/go-amino-x"
	"github.com/tendermint/go-amino-x/genproto"
)

func main() {
	pkgs := []*amino.Package{
		ed25519.Package,
		secp256k1.Package,
		multisig.Package,
		merkle.Package,
	}
	for _, pkg := range pkgs {
		genproto.WriteProto3Schema(pkg)
		genproto.WriteProtoBindings(pkg)
		genproto.MakeProtoFolder(pkg, "proto")
	}
	for _, pkg := range pkgs {
		genproto.RunProtoc(pkg, "proto")
	}
}
