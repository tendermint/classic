package secp256k1

import (
	"reflect"

	"github.com/tendermint/go-amino-x"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/classic/crypto/secp256k1",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	amino.Type{
		Type:             reflect.TypeOf(PubKeySecp256k1{}),
		Name:             "PubKeySecp256k1",
		PointerPreferred: false,
	},
	amino.Type{
		Type:             reflect.TypeOf(PrivKeySecp256k1{}),
		Name:             "PrivKeySecp256k1",
		PointerPreferred: false,
	},
))
