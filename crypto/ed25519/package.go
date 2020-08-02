package ed25519

import (
	"reflect"

	"github.com/tendermint/go-amino-x"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/classic/crypto/ed25519",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	amino.Type{
		Type:             reflect.TypeOf(PubKeyEd25519{}),
		Name:             "PubKeyEd25519",
		PointerPreferred: false,
	},
	amino.Type{
		Type:             reflect.TypeOf(PrivKeyEd25519{}),
		Name:             "PrivKeyEd25519",
		PointerPreferred: false,
	},
))
