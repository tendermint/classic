package abci

import (
	"time"

	"github.com/tendermint/classic/crypto"
	"github.com/tendermint/classic/crypto/merkle"
)

//----------------------------------------
// Request types

type Request interface {
	AssertRequest()
}

type RequestBase struct {
}

func (_ RequestBase) AssertRequest() {}

type RequestEcho struct {
	RequestBase
	Message string
}

type RequestFlush struct {
	RequestBase
}

type RequestInfo struct {
	RequestBase
	Version      string
	BlockVersion uint64
	P2PVersion   uint64
}

// nondeterministic
type RequestSetOption struct {
	RequestBase
	Key   string
	Value string
}

type RequestInitChain struct {
	RequestBase
	Time            time.Time
	ChainID         string
	ConsensusParams *ConsensusParams
	Validators      []ValidatorUpdate
	AppStateBytes   []byte
}

type RequestQuery struct {
	RequestBase
	Data   []byte
	Path   string
	Height int64
	Prove  bool
}

type RequestBeginBlock struct {
	RequestBase
	Hash           []byte
	Header         Header
	LastCommitInfo *LastCommitInfo
	Evidences      []Evidence
}

type CheckTxType int

const (
	CheckTxTypeNew     CheckTxType = 0
	CheckTxTypeRecheck             = iota
)

type RequestCheckTx struct {
	RequestBase
	Tx   []byte
	Type CheckTxType
}

type RequestDeliverTx struct {
	RequestBase
	Tx []byte
}

type RequestEndBlock struct {
	RequestBase
	Height int64
}

type RequestCommit struct {
	RequestBase
}

//----------------------------------------
// Response types

type Response interface {
	AssertResponse()
}

type ResponseBase struct {
	Error  Error
	Data   []byte
	Events []Event

	Log  string // nondeterministic
	Info string // nondeterministic
}

func (_ ResponseBase) AssertResponse() {}

func (r ResponseBase) IsOK() bool {
	return r.Error == nil
}

func (r ResponseBase) IsErr() bool {
	return r.Error != nil
}

// nondeterministic
type ResponseException struct {
	ResponseBase
}

type ResponseEcho struct {
	ResponseBase
	Message string
}

type ResponseFlush struct {
	ResponseBase
}

type ResponseInfo struct {
	ResponseBase
	Version          string
	AppVersion       string
	LastBlockHeight  int64
	LastBlockAppHash []byte
}

// nondeterministic
type ResponseSetOption struct {
	ResponseBase
}

type ResponseInitChain struct {
	ResponseBase
	ConsensusParams *ConsensusParams
	Validators      []ValidatorUpdate
}

type ResponseQuery struct {
	ResponseBase
	Key    []byte
	Value  []byte
	Proof  *merkle.Proof
	Height int64
}

type ResponseBeginBlock struct {
	ResponseBase
}

type ResponseCheckTx struct {
	ResponseBase
	GasWanted int64 // nondeterministic
	GasUsed   int64
}

type ResponseDeliverTx struct {
	ResponseBase
	GasWanted int64
	GasUsed   int64
}

type ResponseEndBlock struct {
	ResponseBase
	ValidatorUpdates []ValidatorUpdate
	ConsensusParams  *ConsensusParams
	Events           []Event
}

type ResponseCommit struct {
	ResponseBase
}

//----------------------------------------
// Interface types

type Error interface {
	AssertABCIError()
	Error() string
}

type Event interface {
	AssertABCIEvent()
}

type Header interface {
	AssertABCIHeader()
}

type Evidence interface {
	AssertABCIEvidence()
}

//----------------------------------------
// Error types

type StringError struct {
	Message string
}

func (_ StringError) AssertABCIError() {}

func (err StringError) Error() string {
	return err.Message
}

//----------------------------------------
// Misc

type ConsensusParams struct {
	Block     BlockParams
	Evidence  EvidenceParams
	Validator ValidatorParams
}

type BlockParams struct {
	MaxBytes int64 // must be > 0
	MaxGas   int64 // must be >= -1
}

type EvidenceParams struct {
	MaxAge int64 // must be > 0
}

type ValidatorParams struct {
	PubKeyTypes []string
}

type ValidatorUpdate struct {
	PubKey crypto.PubKey
	Power  int64
}

type LastCommitInfo struct {
	Round int32
	Votes []*VoteInfo
}

type VoteInfo struct {
	Validator       // TODO consider destructuring.
	SignedLastBlock bool
}

// unstable
type Validator struct {
	Address []byte
	Power   int64
}
