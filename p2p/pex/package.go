package pex

import (
	"reflect"

	"github.com/tendermint/go-amino-x"
	"github.com/tendermint/go-amino-x/pkg"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/classic/p2p/conn",
	"p2p", // keep short, do not change.
	amino.GetCallersDirname(),
).
	WithDependencies(
	// NA
	).
	WithTypes(

		// NOTE: Keep the names short.
		pkg.Type{
			Type:             reflect.TypeOf(pexRequestMessage{}),
			Name:             "PexRequest",
			PointerPreferred: true,
		},
		pkg.Type{
			Type:             reflect.TypeOf(pexAddrsMessage{}),
			Name:             "PexAddrs",
			PointerPreferred: true,
		},
	))
