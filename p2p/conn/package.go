package conn

import (
	"github.com/tendermint/go-amino-x"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/classic/p2p/conn",
	"conn",
	amino.GetCallersDirname(),
).
	WithDependencies(
	// na
	).
	WithTypes(

		/*
			pkg.Type{ // example of overriding type name.
				Type:             reflect.TypeOf(EmbeddedSt5{}),
				Name:             "EmbeddedSt5NameOverride",
				PointerPreferred: false,
			},
		*/

		PackatPing{},
		PacketPong{},
		PacketMsg{},
	))
