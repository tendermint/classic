package types

import (
	"github.com/tendermint/go-amino-x"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/classic/types",
	"tm",
	amino.GetCallersDirname(),
).
	WithGoPkgName("tm").
	WithDependencies().
	WithTypes(

		// Event types
		EventDataNewBlock{},
		EventDataNewBlockHeader{},
		EventDataTx{},
		EventDataRoundState{},
		EventDataNewRound{},
		EventDataCompleteProposal{},
		EventDataVote{},
		EventDataValidatorSetUpdates{},
		EventDataString(""),

		// Evidence types
		DuplicateVoteEvidence{},
		MockGoodEvidence{},
		MockRandomGoodEvidence{},
		MockBadEvidence{},
	))
