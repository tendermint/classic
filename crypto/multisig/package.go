package multisig

import (
	"reflect"

	"github.com/tendermint/go-amino-x"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/classic/crypto/multisig",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	amino.Type{
		Type:             reflect.TypeOf(PubKeyMultisigThreshold{}),
		Name:             "PubKeyMultisig",
		PointerPreferred: false,
	},
))
