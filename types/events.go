package types

import (
	abci "github.com/tendermint/classic/abci/types"
	"github.com/tendermint/classic/libs/events"
)

// TMEvent implements events.Event.
type TMEvent interface {
	abci.Event
	events.Event
}

func (_ EventNewBlock) AssertABCIEvent()            {}
func (_ EventNewBlock) AssertEvent()                {}
func (_ EventNewBlockHeader) AssertABCIEvent()      {}
func (_ EventNewBlockHeader) AssertEvent()          {}
func (_ EventTx) AssertABCIEvent()                  {}
func (_ EventTx) AssertEvent()                      {}
func (_ EventRoundState) AssertABCIEvent()          {}
func (_ EventRoundState) AssertEvent()              {}
func (_ EventNewRound) AssertABCIEvent()            {}
func (_ EventNewRound) AssertEvent()                {}
func (_ EventCompleteProposal) AssertABCIEvent()    {}
func (_ EventCompleteProposal) AssertEvent()        {}
func (_ EventVote) AssertABCIEvent()                {}
func (_ EventVote) AssertEvent()                    {}
func (_ EventString) AssertABCIEvent()              {}
func (_ EventString) AssertEvent()                  {}
func (_ EventValidatorSetUpdates) AssertABCIEvent() {}
func (_ EventValidatorSetUpdates) AssertEvent()     {}

// Most event messages are basic types (a block, a transaction)
// but some (an input to a call tx or a receive) are more exotic

type EventNewBlock struct {
	Block *Block `json:"block"`

	ResultBeginBlock abci.ResponseBeginBlock `json:"result_begin_block"`
	ResultEndBlock   abci.ResponseEndBlock   `json:"result_end_block"`
}

// light weight event for benchmarking
type EventNewBlockHeader struct {
	Header Header `json:"header"`

	ResultBeginBlock abci.ResponseBeginBlock `json:"result_begin_block"`
	ResultEndBlock   abci.ResponseEndBlock   `json:"result_end_block"`
}

// All txs fire EventTx
type EventTx struct {
	TxResult
}

// NOTE: This goes into the replay WAL
type EventRoundState struct {
	Height int64  `json:"height"`
	Round  int    `json:"round"`
	Step   string `json:"step"`
}

type EventNewRound struct {
	Height int64  `json:"height"`
	Round  int    `json:"round"`
	Step   string `json:"step"`

	Proposer Validator `json:"proposer"`
}

type EventCompleteProposal struct {
	Height int64  `json:"height"`
	Round  int    `json:"round"`
	Step   string `json:"step"`

	BlockID BlockID `json:"block_id"`
}

type EventVote struct {
	Vote *Vote
}

type EventString string

type EventValidatorSetUpdates struct {
	ValidatorUpdates []abci.ValidatorUpdate `json:"validator_updates"`
}
