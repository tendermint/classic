package kvstore

import (
	"github.com/tendermint/classic/abci/types"
	"github.com/tendermint/classic/crypto/ed25519"
	cmn "github.com/tendermint/classic/libs/common"
)

// RandVal creates one random validator, with a key derived
// from the input value
func RandVal(i int) abci.ValidatorUpdate {
	pubkey := ed25519.GenPrivKey().PubKey()
	power := cmn.RandUint16() + 1
	v := abci.ValidatorUpdate{pubkey.Address(), pubkey, int64(power)}
	return v
}

// RandVals returns a list of cnt validators for initializing
// the application. Note that the keys are deterministically
// derived from the index in the array, while the power is
// random (Change this if not desired)
func RandVals(cnt int) []abci.ValidatorUpdate {
	res := make([]abci.ValidatorUpdate, cnt)
	for i := 0; i < cnt; i++ {
		res[i] = RandVal(i)
	}
	return res
}

// InitKVStore initializes the kvstore app with some data,
// which allows tests to pass and is fine as long as you
// don't make any tx that modify the validator state
func InitKVStore(app *PersistentKVStoreApplication) {
	app.InitChain(abci.RequestInitChain{
		Validators: RandVals(1),
	})
}
