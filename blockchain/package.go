package blockchain

import (
	"reflect"

	"github.com/tendermint/go-amino-x/pkg"
)

var Package = pkg.NewPackage(
	"github.com/tendermint/classic/blockchain",
	"tm",
	pkg.GetCallersDirName(),
).WithDependencies().WithTypes(
	pkg.Type{
		Type:             reflect.TypeOf(bcBlockRequestMessage{}),
		Name:             "BlockRequest",
		PointerPreferred: true,
	},
	pkg.Type{
		Type:             reflect.TypeOf(bcBlockResponseMessage{}),
		Name:             "BlockResponse",
		PointerPreferred: true,
	},
	pkg.Type{
		Type:             reflect.TypeOf(bcNoBlockResponseMessage{}),
		Name:             "NoBlockResponse",
		PointerPreferred: true,
	},
	pkg.Type{
		Type:             reflect.TypeOf(bcStausRequestMessage{}),
		Name:             "StatusRequest",
		PointerPreferred: true,
	},
	pkg.Type{
		Type:             reflect.TypeOf(bcStausResponseMessage{}),
		Name:             "StatusResponse",
		PointerPreferred: true,
	},
)
