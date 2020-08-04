package types

import (
	cmn "github.com/tendermint/classic/libs/common"
)

// cdcEncode returns nil if the input is nil, otherwise returns
// cdc.MustMarshal(item)
func cdcEncode(item interface{}) []byte {
	if item != nil && !cmn.IsTypedNil(item) && !cmn.IsEmpty(item) {
		return cdc.MustMarshal(item)
	}
	return nil
}
