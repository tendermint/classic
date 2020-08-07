package abci

import (
	proto "google.golang.org/protobuf/proto"
	amino "github.com/tendermint/go-amino-x"
	abcipb "github.com/tendermint/classic/abci/types/pb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	time "time"
	anypb "google.golang.org/protobuf/types/known/anypb"
	merklepb "github.com/tendermint/classic/crypto/merkle/pb"
	merkle "github.com/tendermint/classic/crypto/merkle"
	crypto "github.com/tendermint/classic/crypto"
)

func (goo RequestBase) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestBase
	{
		if isRequestBaseEmptyRepr(goo) {
			var pbov *abcipb.RequestBase
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestBase)
	}
	msg = pbo
	return
}
func (goo RequestBase) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestBase)
	msg = pbo
	return
}
func (goo *RequestBase) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestBase = msg.(*abcipb.RequestBase)
	{
		if pbo != nil {
		}
	}
	return
}
func (_ RequestBase) GetTypeURL() (typeURL string) {
	return "/abci.RequestBase"
}
func isRequestBaseEmptyRepr(goor RequestBase) (empty bool) {
	{
		empty = true
	}
	return
}
func (goo RequestEcho) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestEcho
	{
		if isRequestEchoEmptyRepr(goo) {
			var pbov *abcipb.RequestEcho
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestEcho)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			pbo.Message = goo.Message
		}
	}
	msg = pbo
	return
}
func (goo RequestEcho) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestEcho)
	msg = pbo
	return
}
func (goo *RequestEcho) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestEcho = msg.(*abcipb.RequestEcho)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).Message = string(pbo.Message)
			}
		}
	}
	return
}
func (_ RequestEcho) GetTypeURL() (typeURL string) {
	return "/abci.RequestEcho"
}
func isRequestEchoEmptyRepr(goor RequestEcho) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if goor.Message != "" {
				return false
			}
		}
	}
	return
}
func (goo RequestFlush) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestFlush
	{
		if isRequestFlushEmptyRepr(goo) {
			var pbov *abcipb.RequestFlush
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestFlush)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
	}
	msg = pbo
	return
}
func (goo RequestFlush) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestFlush)
	msg = pbo
	return
}
func (goo *RequestFlush) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestFlush = msg.(*abcipb.RequestFlush)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ RequestFlush) GetTypeURL() (typeURL string) {
	return "/abci.RequestFlush"
}
func isRequestFlushEmptyRepr(goor RequestFlush) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo RequestInfo) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestInfo
	{
		if isRequestInfoEmptyRepr(goo) {
			var pbov *abcipb.RequestInfo
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestInfo)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			pbo.Version = goo.Version
		}
		{
			pbo.BlockVersion = goo.BlockVersion
		}
		{
			pbo.P2PVersion = goo.P2PVersion
		}
	}
	msg = pbo
	return
}
func (goo RequestInfo) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestInfo)
	msg = pbo
	return
}
func (goo *RequestInfo) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestInfo = msg.(*abcipb.RequestInfo)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).Version = string(pbo.Version)
			}
			{
				(*goo).BlockVersion = uint64(pbo.BlockVersion)
			}
			{
				(*goo).P2PVersion = uint64(pbo.P2PVersion)
			}
		}
	}
	return
}
func (_ RequestInfo) GetTypeURL() (typeURL string) {
	return "/abci.RequestInfo"
}
func isRequestInfoEmptyRepr(goor RequestInfo) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if goor.Version != "" {
				return false
			}
		}
		{
			if goor.BlockVersion != 0 {
				return false
			}
		}
		{
			if goor.P2PVersion != 0 {
				return false
			}
		}
	}
	return
}
func (goo RequestSetOption) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestSetOption
	{
		if isRequestSetOptionEmptyRepr(goo) {
			var pbov *abcipb.RequestSetOption
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestSetOption)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			pbo.Key = goo.Key
		}
		{
			pbo.Value = goo.Value
		}
	}
	msg = pbo
	return
}
func (goo RequestSetOption) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestSetOption)
	msg = pbo
	return
}
func (goo *RequestSetOption) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestSetOption = msg.(*abcipb.RequestSetOption)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).Key = string(pbo.Key)
			}
			{
				(*goo).Value = string(pbo.Value)
			}
		}
	}
	return
}
func (_ RequestSetOption) GetTypeURL() (typeURL string) {
	return "/abci.RequestSetOption"
}
func isRequestSetOptionEmptyRepr(goor RequestSetOption) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if goor.Key != "" {
				return false
			}
		}
		{
			if goor.Value != "" {
				return false
			}
		}
	}
	return
}
func (goo RequestInitChain) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestInitChain
	{
		if isRequestInitChainEmptyRepr(goo) {
			var pbov *abcipb.RequestInitChain
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestInitChain)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			if !amino.IsEmptyTime(goo.Time) {
				pbo.Time = timestamppb.New(goo.Time)
			}
		}
		{
			pbo.ChainID = goo.ChainID
		}
		{
			if goo.ConsensusParams != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.ConsensusParams.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.ConsensusParams = pbom.(*abcipb.ConsensusParams)
				if pbo.ConsensusParams == nil {
					pbo.ConsensusParams = new(abcipb.ConsensusParams)
				}
			}
		}
		{
			goorl := len(goo.Validators)
			if goorl == 0 {
				pbo.Validators = nil
			} else {
				var pbos = make([]*abcipb.ValidatorUpdate, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Validators[i]
						{
							pbom := proto.Message(nil)
							pbom, err = goore.ToPBMessage(cdc)
							if err != nil {
								return
							}
							pbos[i] = pbom.(*abcipb.ValidatorUpdate)
						}
					}
				}
				pbo.Validators = pbos
			}
		}
		{
			goorl := len(goo.AppStateBytes)
			if goorl == 0 {
				pbo.AppStateBytes = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.AppStateBytes[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.AppStateBytes = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo RequestInitChain) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestInitChain)
	msg = pbo
	return
}
func (goo *RequestInitChain) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestInitChain = msg.(*abcipb.RequestInitChain)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).Time = pbo.Time.AsTime()
			}
			{
				(*goo).ChainID = string(pbo.ChainID)
			}
			{
				if pbo.ConsensusParams != nil {
					(*goo).ConsensusParams = new(ConsensusParams)
					err = (*goo).ConsensusParams.FromPBMessage(cdc, pbo.ConsensusParams)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Validators != nil {
					pbol = len(pbo.Validators)
				}
				if pbol == 0 {
					(*goo).Validators = nil
				} else {
					var goors = make([]ValidatorUpdate, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Validators[i]
							{
								pboev := pboe
								if pboev != nil {
									err = goors[i].FromPBMessage(cdc, pboev)
									if err != nil {
										return
									}
								}
							}
						}
					}
					(*goo).Validators = goors
				}
			}
			{
				var pbol int = 0
				if pbo.AppStateBytes != nil {
					pbol = len(pbo.AppStateBytes)
				}
				if pbol == 0 {
					(*goo).AppStateBytes = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.AppStateBytes[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).AppStateBytes = goors
				}
			}
		}
	}
	return
}
func (_ RequestInitChain) GetTypeURL() (typeURL string) {
	return "/abci.RequestInitChain"
}
func isRequestInitChainEmptyRepr(goor RequestInitChain) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if !amino.IsEmptyTime(goor.Time) {
				return false
			}
		}
		{
			if goor.ChainID != "" {
				return false
			}
		}
		{
			if goor.ConsensusParams != nil {
				return false
			}
		}
		{
			if len(goor.Validators) != 0 {
				return false
			}
		}
		{
			if len(goor.AppStateBytes) != 0 {
				return false
			}
		}
	}
	return
}
func (goo RequestQuery) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestQuery
	{
		if isRequestQueryEmptyRepr(goo) {
			var pbov *abcipb.RequestQuery
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestQuery)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			goorl := len(goo.Data)
			if goorl == 0 {
				pbo.Data = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Data[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Data = pbos
			}
		}
		{
			pbo.Path = goo.Path
		}
		{
			pbo.Height = goo.Height
		}
		{
			pbo.Prove = goo.Prove
		}
	}
	msg = pbo
	return
}
func (goo RequestQuery) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestQuery)
	msg = pbo
	return
}
func (goo *RequestQuery) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestQuery = msg.(*abcipb.RequestQuery)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Data != nil {
					pbol = len(pbo.Data)
				}
				if pbol == 0 {
					(*goo).Data = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Data[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Data = goors
				}
			}
			{
				(*goo).Path = string(pbo.Path)
			}
			{
				(*goo).Height = int64(pbo.Height)
			}
			{
				(*goo).Prove = bool(pbo.Prove)
			}
		}
	}
	return
}
func (_ RequestQuery) GetTypeURL() (typeURL string) {
	return "/abci.RequestQuery"
}
func isRequestQueryEmptyRepr(goor RequestQuery) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if len(goor.Data) != 0 {
				return false
			}
		}
		{
			if goor.Path != "" {
				return false
			}
		}
		{
			if goor.Height != 0 {
				return false
			}
		}
		{
			if goor.Prove != false {
				return false
			}
		}
	}
	return
}
func (goo RequestBeginBlock) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestBeginBlock
	{
		if isRequestBeginBlockEmptyRepr(goo) {
			var pbov *abcipb.RequestBeginBlock
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestBeginBlock)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			goorl := len(goo.Hash)
			if goorl == 0 {
				pbo.Hash = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Hash[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Hash = pbos
			}
		}
		{
			if goo.Header != nil {
				typeUrl := cdc.GetTypeURL(goo.Header)
				bz := []byte(nil)
				bz, err = cdc.Marshal(goo.Header)
				if err != nil {
					return
				}
				pbo.Header = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			if goo.LastCommitInfo != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.LastCommitInfo.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.LastCommitInfo = pbom.(*abcipb.LastCommitInfo)
				if pbo.LastCommitInfo == nil {
					pbo.LastCommitInfo = new(abcipb.LastCommitInfo)
				}
			}
		}
		{
			goorl := len(goo.Violations)
			if goorl == 0 {
				pbo.Violations = nil
			} else {
				var pbos = make([]*abcipb.Violation, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Violations[i]
						{
							pbom := proto.Message(nil)
							pbom, err = goore.ToPBMessage(cdc)
							if err != nil {
								return
							}
							pbos[i] = pbom.(*abcipb.Violation)
						}
					}
				}
				pbo.Violations = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo RequestBeginBlock) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestBeginBlock)
	msg = pbo
	return
}
func (goo *RequestBeginBlock) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestBeginBlock = msg.(*abcipb.RequestBeginBlock)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Hash != nil {
					pbol = len(pbo.Hash)
				}
				if pbol == 0 {
					(*goo).Hash = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Hash[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Hash = goors
				}
			}
			{
				typeUrl := pbo.Header.TypeUrl
				bz := pbo.Header.Value
				goorp := &(*goo).Header
				err = cdc.UnmarshalAny(typeUrl, bz, goorp)
				if err != nil {
					return
				}
			}
			{
				if pbo.LastCommitInfo != nil {
					(*goo).LastCommitInfo = new(LastCommitInfo)
					err = (*goo).LastCommitInfo.FromPBMessage(cdc, pbo.LastCommitInfo)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Violations != nil {
					pbol = len(pbo.Violations)
				}
				if pbol == 0 {
					(*goo).Violations = nil
				} else {
					var goors = make([]Violation, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Violations[i]
							{
								pboev := pboe
								if pboev != nil {
									err = goors[i].FromPBMessage(cdc, pboev)
									if err != nil {
										return
									}
								}
							}
						}
					}
					(*goo).Violations = goors
				}
			}
		}
	}
	return
}
func (_ RequestBeginBlock) GetTypeURL() (typeURL string) {
	return "/abci.RequestBeginBlock"
}
func isRequestBeginBlockEmptyRepr(goor RequestBeginBlock) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if len(goor.Hash) != 0 {
				return false
			}
		}
		{
			if goor.Header != nil {
				return false
			}
		}
		{
			if goor.LastCommitInfo != nil {
				return false
			}
		}
		{
			if len(goor.Violations) != 0 {
				return false
			}
		}
	}
	return
}
func (goo RequestCheckTx) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestCheckTx
	{
		if isRequestCheckTxEmptyRepr(goo) {
			var pbov *abcipb.RequestCheckTx
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestCheckTx)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			goorl := len(goo.Tx)
			if goorl == 0 {
				pbo.Tx = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Tx[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Tx = pbos
			}
		}
		{
			pbo.Type = int64(goo.Type)
		}
	}
	msg = pbo
	return
}
func (goo RequestCheckTx) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestCheckTx)
	msg = pbo
	return
}
func (goo *RequestCheckTx) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestCheckTx = msg.(*abcipb.RequestCheckTx)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Tx != nil {
					pbol = len(pbo.Tx)
				}
				if pbol == 0 {
					(*goo).Tx = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Tx[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Tx = goors
				}
			}
			{
				(*goo).Type = CheckTxType(int(pbo.Type))
			}
		}
	}
	return
}
func (_ RequestCheckTx) GetTypeURL() (typeURL string) {
	return "/abci.RequestCheckTx"
}
func isRequestCheckTxEmptyRepr(goor RequestCheckTx) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if len(goor.Tx) != 0 {
				return false
			}
		}
		{
			if goor.Type != 0 {
				return false
			}
		}
	}
	return
}
func (goo RequestDeliverTx) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestDeliverTx
	{
		if isRequestDeliverTxEmptyRepr(goo) {
			var pbov *abcipb.RequestDeliverTx
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestDeliverTx)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			goorl := len(goo.Tx)
			if goorl == 0 {
				pbo.Tx = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Tx[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Tx = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo RequestDeliverTx) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestDeliverTx)
	msg = pbo
	return
}
func (goo *RequestDeliverTx) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestDeliverTx = msg.(*abcipb.RequestDeliverTx)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Tx != nil {
					pbol = len(pbo.Tx)
				}
				if pbol == 0 {
					(*goo).Tx = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Tx[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Tx = goors
				}
			}
		}
	}
	return
}
func (_ RequestDeliverTx) GetTypeURL() (typeURL string) {
	return "/abci.RequestDeliverTx"
}
func isRequestDeliverTxEmptyRepr(goor RequestDeliverTx) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if len(goor.Tx) != 0 {
				return false
			}
		}
	}
	return
}
func (goo RequestEndBlock) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestEndBlock
	{
		if isRequestEndBlockEmptyRepr(goo) {
			var pbov *abcipb.RequestEndBlock
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestEndBlock)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
		{
			pbo.Height = goo.Height
		}
	}
	msg = pbo
	return
}
func (goo RequestEndBlock) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestEndBlock)
	msg = pbo
	return
}
func (goo *RequestEndBlock) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestEndBlock = msg.(*abcipb.RequestEndBlock)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).Height = int64(pbo.Height)
			}
		}
	}
	return
}
func (_ RequestEndBlock) GetTypeURL() (typeURL string) {
	return "/abci.RequestEndBlock"
}
func isRequestEndBlockEmptyRepr(goor RequestEndBlock) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
		{
			if goor.Height != 0 {
				return false
			}
		}
	}
	return
}
func (goo RequestCommit) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.RequestCommit
	{
		if isRequestCommitEmptyRepr(goo) {
			var pbov *abcipb.RequestCommit
			msg = pbov
			return
		}
		pbo = new(abcipb.RequestCommit)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.RequestBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.RequestBase = pbom.(*abcipb.RequestBase)
		}
	}
	msg = pbo
	return
}
func (goo RequestCommit) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.RequestCommit)
	msg = pbo
	return
}
func (goo *RequestCommit) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.RequestCommit = msg.(*abcipb.RequestCommit)
	{
		if pbo != nil {
			{
				if pbo.RequestBase != nil {
					err = (*goo).RequestBase.FromPBMessage(cdc, pbo.RequestBase)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ RequestCommit) GetTypeURL() (typeURL string) {
	return "/abci.RequestCommit"
}
func isRequestCommitEmptyRepr(goor RequestCommit) (empty bool) {
	{
		empty = true
		{
			e := isRequestBaseEmptyRepr(goor.RequestBase)
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo ResponseBase) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseBase
	{
		if isResponseBaseEmptyRepr(goo) {
			var pbov *abcipb.ResponseBase
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseBase)
		{
			if goo.Error != nil {
				typeUrl := cdc.GetTypeURL(goo.Error)
				bz := []byte(nil)
				bz, err = cdc.Marshal(goo.Error)
				if err != nil {
					return
				}
				pbo.Error = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			goorl := len(goo.Data)
			if goorl == 0 {
				pbo.Data = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Data[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Data = pbos
			}
		}
		{
			goorl := len(goo.Events)
			if goorl == 0 {
				pbo.Events = nil
			} else {
				var pbos = make([]*anypb.Any, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Events[i]
						{
							if goore != nil {
								typeUrl := cdc.GetTypeURL(goore)
								bz := []byte(nil)
								bz, err = cdc.Marshal(goore)
								if err != nil {
									return
								}
								pbos[i] = &anypb.Any{TypeUrl: typeUrl, Value: bz}
							}
						}
					}
				}
				pbo.Events = pbos
			}
		}
		{
			pbo.Log = goo.Log
		}
		{
			pbo.Info = goo.Info
		}
	}
	msg = pbo
	return
}
func (goo ResponseBase) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseBase)
	msg = pbo
	return
}
func (goo *ResponseBase) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseBase = msg.(*abcipb.ResponseBase)
	{
		if pbo != nil {
			{
				typeUrl := pbo.Error.TypeUrl
				bz := pbo.Error.Value
				goorp := &(*goo).Error
				err = cdc.UnmarshalAny(typeUrl, bz, goorp)
				if err != nil {
					return
				}
			}
			{
				var pbol int = 0
				if pbo.Data != nil {
					pbol = len(pbo.Data)
				}
				if pbol == 0 {
					(*goo).Data = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Data[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Data = goors
				}
			}
			{
				var pbol int = 0
				if pbo.Events != nil {
					pbol = len(pbo.Events)
				}
				if pbol == 0 {
					(*goo).Events = nil
				} else {
					var goors = make([]Event, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Events[i]
							{
								pboev := pboe
								typeUrl := pboev.TypeUrl
								bz := pboev.Value
								goorp := &goors[i]
								err = cdc.UnmarshalAny(typeUrl, bz, goorp)
								if err != nil {
									return
								}
							}
						}
					}
					(*goo).Events = goors
				}
			}
			{
				(*goo).Log = string(pbo.Log)
			}
			{
				(*goo).Info = string(pbo.Info)
			}
		}
	}
	return
}
func (_ ResponseBase) GetTypeURL() (typeURL string) {
	return "/abci.ResponseBase"
}
func isResponseBaseEmptyRepr(goor ResponseBase) (empty bool) {
	{
		empty = true
		{
			if goor.Error != nil {
				return false
			}
		}
		{
			if len(goor.Data) != 0 {
				return false
			}
		}
		{
			if len(goor.Events) != 0 {
				return false
			}
		}
		{
			if goor.Log != "" {
				return false
			}
		}
		{
			if goor.Info != "" {
				return false
			}
		}
	}
	return
}
func (goo ResponseException) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseException
	{
		if isResponseExceptionEmptyRepr(goo) {
			var pbov *abcipb.ResponseException
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseException)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
	}
	msg = pbo
	return
}
func (goo ResponseException) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseException)
	msg = pbo
	return
}
func (goo *ResponseException) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseException = msg.(*abcipb.ResponseException)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ ResponseException) GetTypeURL() (typeURL string) {
	return "/abci.ResponseException"
}
func isResponseExceptionEmptyRepr(goor ResponseException) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo ResponseEcho) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseEcho
	{
		if isResponseEchoEmptyRepr(goo) {
			var pbov *abcipb.ResponseEcho
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseEcho)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			pbo.Message = goo.Message
		}
	}
	msg = pbo
	return
}
func (goo ResponseEcho) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseEcho)
	msg = pbo
	return
}
func (goo *ResponseEcho) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseEcho = msg.(*abcipb.ResponseEcho)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).Message = string(pbo.Message)
			}
		}
	}
	return
}
func (_ ResponseEcho) GetTypeURL() (typeURL string) {
	return "/abci.ResponseEcho"
}
func isResponseEchoEmptyRepr(goor ResponseEcho) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if goor.Message != "" {
				return false
			}
		}
	}
	return
}
func (goo ResponseFlush) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseFlush
	{
		if isResponseFlushEmptyRepr(goo) {
			var pbov *abcipb.ResponseFlush
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseFlush)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
	}
	msg = pbo
	return
}
func (goo ResponseFlush) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseFlush)
	msg = pbo
	return
}
func (goo *ResponseFlush) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseFlush = msg.(*abcipb.ResponseFlush)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ ResponseFlush) GetTypeURL() (typeURL string) {
	return "/abci.ResponseFlush"
}
func isResponseFlushEmptyRepr(goor ResponseFlush) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo ResponseInfo) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseInfo
	{
		if isResponseInfoEmptyRepr(goo) {
			var pbov *abcipb.ResponseInfo
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseInfo)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			pbo.ABCIVersion = goo.ABCIVersion
		}
		{
			pbo.AppVersion = goo.AppVersion
		}
		{
			pbo.LastBlockHeight = goo.LastBlockHeight
		}
		{
			goorl := len(goo.LastBlockAppHash)
			if goorl == 0 {
				pbo.LastBlockAppHash = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.LastBlockAppHash[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.LastBlockAppHash = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo ResponseInfo) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseInfo)
	msg = pbo
	return
}
func (goo *ResponseInfo) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseInfo = msg.(*abcipb.ResponseInfo)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).ABCIVersion = string(pbo.ABCIVersion)
			}
			{
				(*goo).AppVersion = string(pbo.AppVersion)
			}
			{
				(*goo).LastBlockHeight = int64(pbo.LastBlockHeight)
			}
			{
				var pbol int = 0
				if pbo.LastBlockAppHash != nil {
					pbol = len(pbo.LastBlockAppHash)
				}
				if pbol == 0 {
					(*goo).LastBlockAppHash = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.LastBlockAppHash[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).LastBlockAppHash = goors
				}
			}
		}
	}
	return
}
func (_ ResponseInfo) GetTypeURL() (typeURL string) {
	return "/abci.ResponseInfo"
}
func isResponseInfoEmptyRepr(goor ResponseInfo) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if goor.ABCIVersion != "" {
				return false
			}
		}
		{
			if goor.AppVersion != "" {
				return false
			}
		}
		{
			if goor.LastBlockHeight != 0 {
				return false
			}
		}
		{
			if len(goor.LastBlockAppHash) != 0 {
				return false
			}
		}
	}
	return
}
func (goo ResponseSetOption) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseSetOption
	{
		if isResponseSetOptionEmptyRepr(goo) {
			var pbov *abcipb.ResponseSetOption
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseSetOption)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
	}
	msg = pbo
	return
}
func (goo ResponseSetOption) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseSetOption)
	msg = pbo
	return
}
func (goo *ResponseSetOption) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseSetOption = msg.(*abcipb.ResponseSetOption)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ ResponseSetOption) GetTypeURL() (typeURL string) {
	return "/abci.ResponseSetOption"
}
func isResponseSetOptionEmptyRepr(goor ResponseSetOption) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo ResponseInitChain) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseInitChain
	{
		if isResponseInitChainEmptyRepr(goo) {
			var pbov *abcipb.ResponseInitChain
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseInitChain)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			if goo.ConsensusParams != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.ConsensusParams.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.ConsensusParams = pbom.(*abcipb.ConsensusParams)
				if pbo.ConsensusParams == nil {
					pbo.ConsensusParams = new(abcipb.ConsensusParams)
				}
			}
		}
		{
			goorl := len(goo.Validators)
			if goorl == 0 {
				pbo.Validators = nil
			} else {
				var pbos = make([]*abcipb.ValidatorUpdate, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Validators[i]
						{
							pbom := proto.Message(nil)
							pbom, err = goore.ToPBMessage(cdc)
							if err != nil {
								return
							}
							pbos[i] = pbom.(*abcipb.ValidatorUpdate)
						}
					}
				}
				pbo.Validators = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo ResponseInitChain) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseInitChain)
	msg = pbo
	return
}
func (goo *ResponseInitChain) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseInitChain = msg.(*abcipb.ResponseInitChain)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.ConsensusParams != nil {
					(*goo).ConsensusParams = new(ConsensusParams)
					err = (*goo).ConsensusParams.FromPBMessage(cdc, pbo.ConsensusParams)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Validators != nil {
					pbol = len(pbo.Validators)
				}
				if pbol == 0 {
					(*goo).Validators = nil
				} else {
					var goors = make([]ValidatorUpdate, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Validators[i]
							{
								pboev := pboe
								if pboev != nil {
									err = goors[i].FromPBMessage(cdc, pboev)
									if err != nil {
										return
									}
								}
							}
						}
					}
					(*goo).Validators = goors
				}
			}
		}
	}
	return
}
func (_ ResponseInitChain) GetTypeURL() (typeURL string) {
	return "/abci.ResponseInitChain"
}
func isResponseInitChainEmptyRepr(goor ResponseInitChain) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if goor.ConsensusParams != nil {
				return false
			}
		}
		{
			if len(goor.Validators) != 0 {
				return false
			}
		}
	}
	return
}
func (goo ResponseQuery) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseQuery
	{
		if isResponseQueryEmptyRepr(goo) {
			var pbov *abcipb.ResponseQuery
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseQuery)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			goorl := len(goo.Key)
			if goorl == 0 {
				pbo.Key = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Key[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Key = pbos
			}
		}
		{
			goorl := len(goo.Value)
			if goorl == 0 {
				pbo.Value = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Value[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Value = pbos
			}
		}
		{
			if goo.Proof != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.Proof.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.Proof = pbom.(*merklepb.Proof)
				if pbo.Proof == nil {
					pbo.Proof = new(merklepb.Proof)
				}
			}
		}
		{
			pbo.Height = goo.Height
		}
	}
	msg = pbo
	return
}
func (goo ResponseQuery) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseQuery)
	msg = pbo
	return
}
func (goo *ResponseQuery) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseQuery = msg.(*abcipb.ResponseQuery)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Key != nil {
					pbol = len(pbo.Key)
				}
				if pbol == 0 {
					(*goo).Key = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Key[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Key = goors
				}
			}
			{
				var pbol int = 0
				if pbo.Value != nil {
					pbol = len(pbo.Value)
				}
				if pbol == 0 {
					(*goo).Value = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Value[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Value = goors
				}
			}
			{
				if pbo.Proof != nil {
					(*goo).Proof = new(merkle.Proof)
					err = (*goo).Proof.FromPBMessage(cdc, pbo.Proof)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).Height = int64(pbo.Height)
			}
		}
	}
	return
}
func (_ ResponseQuery) GetTypeURL() (typeURL string) {
	return "/abci.ResponseQuery"
}
func isResponseQueryEmptyRepr(goor ResponseQuery) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if len(goor.Key) != 0 {
				return false
			}
		}
		{
			if len(goor.Value) != 0 {
				return false
			}
		}
		{
			if goor.Proof != nil {
				return false
			}
		}
		{
			if goor.Height != 0 {
				return false
			}
		}
	}
	return
}
func (goo ResponseBeginBlock) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseBeginBlock
	{
		if isResponseBeginBlockEmptyRepr(goo) {
			var pbov *abcipb.ResponseBeginBlock
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseBeginBlock)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
	}
	msg = pbo
	return
}
func (goo ResponseBeginBlock) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseBeginBlock)
	msg = pbo
	return
}
func (goo *ResponseBeginBlock) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseBeginBlock = msg.(*abcipb.ResponseBeginBlock)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ ResponseBeginBlock) GetTypeURL() (typeURL string) {
	return "/abci.ResponseBeginBlock"
}
func isResponseBeginBlockEmptyRepr(goor ResponseBeginBlock) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo ResponseCheckTx) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseCheckTx
	{
		if isResponseCheckTxEmptyRepr(goo) {
			var pbov *abcipb.ResponseCheckTx
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseCheckTx)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			pbo.GasWanted = goo.GasWanted
		}
		{
			pbo.GasUsed = goo.GasUsed
		}
	}
	msg = pbo
	return
}
func (goo ResponseCheckTx) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseCheckTx)
	msg = pbo
	return
}
func (goo *ResponseCheckTx) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseCheckTx = msg.(*abcipb.ResponseCheckTx)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).GasWanted = int64(pbo.GasWanted)
			}
			{
				(*goo).GasUsed = int64(pbo.GasUsed)
			}
		}
	}
	return
}
func (_ ResponseCheckTx) GetTypeURL() (typeURL string) {
	return "/abci.ResponseCheckTx"
}
func isResponseCheckTxEmptyRepr(goor ResponseCheckTx) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if goor.GasWanted != 0 {
				return false
			}
		}
		{
			if goor.GasUsed != 0 {
				return false
			}
		}
	}
	return
}
func (goo ResponseDeliverTx) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseDeliverTx
	{
		if isResponseDeliverTxEmptyRepr(goo) {
			var pbov *abcipb.ResponseDeliverTx
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseDeliverTx)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			pbo.GasWanted = goo.GasWanted
		}
		{
			pbo.GasUsed = goo.GasUsed
		}
	}
	msg = pbo
	return
}
func (goo ResponseDeliverTx) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseDeliverTx)
	msg = pbo
	return
}
func (goo *ResponseDeliverTx) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseDeliverTx = msg.(*abcipb.ResponseDeliverTx)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).GasWanted = int64(pbo.GasWanted)
			}
			{
				(*goo).GasUsed = int64(pbo.GasUsed)
			}
		}
	}
	return
}
func (_ ResponseDeliverTx) GetTypeURL() (typeURL string) {
	return "/abci.ResponseDeliverTx"
}
func isResponseDeliverTxEmptyRepr(goor ResponseDeliverTx) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if goor.GasWanted != 0 {
				return false
			}
		}
		{
			if goor.GasUsed != 0 {
				return false
			}
		}
	}
	return
}
func (goo ResponseEndBlock) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseEndBlock
	{
		if isResponseEndBlockEmptyRepr(goo) {
			var pbov *abcipb.ResponseEndBlock
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseEndBlock)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			goorl := len(goo.ValidatorUpdates)
			if goorl == 0 {
				pbo.ValidatorUpdates = nil
			} else {
				var pbos = make([]*abcipb.ValidatorUpdate, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.ValidatorUpdates[i]
						{
							pbom := proto.Message(nil)
							pbom, err = goore.ToPBMessage(cdc)
							if err != nil {
								return
							}
							pbos[i] = pbom.(*abcipb.ValidatorUpdate)
						}
					}
				}
				pbo.ValidatorUpdates = pbos
			}
		}
		{
			if goo.ConsensusParams != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.ConsensusParams.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.ConsensusParams = pbom.(*abcipb.ConsensusParams)
				if pbo.ConsensusParams == nil {
					pbo.ConsensusParams = new(abcipb.ConsensusParams)
				}
			}
		}
		{
			goorl := len(goo.Events)
			if goorl == 0 {
				pbo.Events = nil
			} else {
				var pbos = make([]*anypb.Any, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Events[i]
						{
							if goore != nil {
								typeUrl := cdc.GetTypeURL(goore)
								bz := []byte(nil)
								bz, err = cdc.Marshal(goore)
								if err != nil {
									return
								}
								pbos[i] = &anypb.Any{TypeUrl: typeUrl, Value: bz}
							}
						}
					}
				}
				pbo.Events = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo ResponseEndBlock) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseEndBlock)
	msg = pbo
	return
}
func (goo *ResponseEndBlock) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseEndBlock = msg.(*abcipb.ResponseEndBlock)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.ValidatorUpdates != nil {
					pbol = len(pbo.ValidatorUpdates)
				}
				if pbol == 0 {
					(*goo).ValidatorUpdates = nil
				} else {
					var goors = make([]ValidatorUpdate, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.ValidatorUpdates[i]
							{
								pboev := pboe
								if pboev != nil {
									err = goors[i].FromPBMessage(cdc, pboev)
									if err != nil {
										return
									}
								}
							}
						}
					}
					(*goo).ValidatorUpdates = goors
				}
			}
			{
				if pbo.ConsensusParams != nil {
					(*goo).ConsensusParams = new(ConsensusParams)
					err = (*goo).ConsensusParams.FromPBMessage(cdc, pbo.ConsensusParams)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Events != nil {
					pbol = len(pbo.Events)
				}
				if pbol == 0 {
					(*goo).Events = nil
				} else {
					var goors = make([]Event, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Events[i]
							{
								pboev := pboe
								typeUrl := pboev.TypeUrl
								bz := pboev.Value
								goorp := &goors[i]
								err = cdc.UnmarshalAny(typeUrl, bz, goorp)
								if err != nil {
									return
								}
							}
						}
					}
					(*goo).Events = goors
				}
			}
		}
	}
	return
}
func (_ ResponseEndBlock) GetTypeURL() (typeURL string) {
	return "/abci.ResponseEndBlock"
}
func isResponseEndBlockEmptyRepr(goor ResponseEndBlock) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if len(goor.ValidatorUpdates) != 0 {
				return false
			}
		}
		{
			if goor.ConsensusParams != nil {
				return false
			}
		}
		{
			if len(goor.Events) != 0 {
				return false
			}
		}
	}
	return
}
func (goo ResponseCommit) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ResponseCommit
	{
		if isResponseCommitEmptyRepr(goo) {
			var pbov *abcipb.ResponseCommit
			msg = pbov
			return
		}
		pbo = new(abcipb.ResponseCommit)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
	}
	msg = pbo
	return
}
func (goo ResponseCommit) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ResponseCommit)
	msg = pbo
	return
}
func (goo *ResponseCommit) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ResponseCommit = msg.(*abcipb.ResponseCommit)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ ResponseCommit) GetTypeURL() (typeURL string) {
	return "/abci.ResponseCommit"
}
func isResponseCommitEmptyRepr(goor ResponseCommit) (empty bool) {
	{
		empty = true
		{
			e := isResponseBaseEmptyRepr(goor.ResponseBase)
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo ConsensusParams) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ConsensusParams
	{
		if isConsensusParamsEmptyRepr(goo) {
			var pbov *abcipb.ConsensusParams
			msg = pbov
			return
		}
		pbo = new(abcipb.ConsensusParams)
		{
			if goo.Block != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.Block.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.Block = pbom.(*abcipb.BlockParams)
				if pbo.Block == nil {
					pbo.Block = new(abcipb.BlockParams)
				}
			}
		}
		{
			if goo.Evidence != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.Evidence.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.Evidence = pbom.(*abcipb.EvidenceParams)
				if pbo.Evidence == nil {
					pbo.Evidence = new(abcipb.EvidenceParams)
				}
			}
		}
		{
			if goo.Validator != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.Validator.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.Validator = pbom.(*abcipb.ValidatorParams)
				if pbo.Validator == nil {
					pbo.Validator = new(abcipb.ValidatorParams)
				}
			}
		}
	}
	msg = pbo
	return
}
func (goo ConsensusParams) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ConsensusParams)
	msg = pbo
	return
}
func (goo *ConsensusParams) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ConsensusParams = msg.(*abcipb.ConsensusParams)
	{
		if pbo != nil {
			{
				if pbo.Block != nil {
					(*goo).Block = new(BlockParams)
					err = (*goo).Block.FromPBMessage(cdc, pbo.Block)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.Evidence != nil {
					(*goo).Evidence = new(EvidenceParams)
					err = (*goo).Evidence.FromPBMessage(cdc, pbo.Evidence)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.Validator != nil {
					(*goo).Validator = new(ValidatorParams)
					err = (*goo).Validator.FromPBMessage(cdc, pbo.Validator)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ ConsensusParams) GetTypeURL() (typeURL string) {
	return "/abci.ConsensusParams"
}
func isConsensusParamsEmptyRepr(goor ConsensusParams) (empty bool) {
	{
		empty = true
		{
			if goor.Block != nil {
				return false
			}
		}
		{
			if goor.Evidence != nil {
				return false
			}
		}
		{
			if goor.Validator != nil {
				return false
			}
		}
	}
	return
}
func (goo BlockParams) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.BlockParams
	{
		if isBlockParamsEmptyRepr(goo) {
			var pbov *abcipb.BlockParams
			msg = pbov
			return
		}
		pbo = new(abcipb.BlockParams)
		{
			pbo.MaxTxBytes = goo.MaxTxBytes
		}
		{
			pbo.MaxGas = goo.MaxGas
		}
		{
			pbo.TimeIotaMS = goo.TimeIotaMS
		}
	}
	msg = pbo
	return
}
func (goo BlockParams) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.BlockParams)
	msg = pbo
	return
}
func (goo *BlockParams) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.BlockParams = msg.(*abcipb.BlockParams)
	{
		if pbo != nil {
			{
				(*goo).MaxTxBytes = int64(pbo.MaxTxBytes)
			}
			{
				(*goo).MaxGas = int64(pbo.MaxGas)
			}
			{
				(*goo).TimeIotaMS = int64(pbo.TimeIotaMS)
			}
		}
	}
	return
}
func (_ BlockParams) GetTypeURL() (typeURL string) {
	return "/abci.BlockParams"
}
func isBlockParamsEmptyRepr(goor BlockParams) (empty bool) {
	{
		empty = true
		{
			if goor.MaxTxBytes != 0 {
				return false
			}
		}
		{
			if goor.MaxGas != 0 {
				return false
			}
		}
		{
			if goor.TimeIotaMS != 0 {
				return false
			}
		}
	}
	return
}
func (goo EvidenceParams) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.EvidenceParams
	{
		if isEvidenceParamsEmptyRepr(goo) {
			var pbov *abcipb.EvidenceParams
			msg = pbov
			return
		}
		pbo = new(abcipb.EvidenceParams)
		{
			pbo.MaxAge = goo.MaxAge
		}
	}
	msg = pbo
	return
}
func (goo EvidenceParams) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.EvidenceParams)
	msg = pbo
	return
}
func (goo *EvidenceParams) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.EvidenceParams = msg.(*abcipb.EvidenceParams)
	{
		if pbo != nil {
			{
				(*goo).MaxAge = int64(pbo.MaxAge)
			}
		}
	}
	return
}
func (_ EvidenceParams) GetTypeURL() (typeURL string) {
	return "/abci.EvidenceParams"
}
func isEvidenceParamsEmptyRepr(goor EvidenceParams) (empty bool) {
	{
		empty = true
		{
			if goor.MaxAge != 0 {
				return false
			}
		}
	}
	return
}
func (goo ValidatorParams) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ValidatorParams
	{
		if isValidatorParamsEmptyRepr(goo) {
			var pbov *abcipb.ValidatorParams
			msg = pbov
			return
		}
		pbo = new(abcipb.ValidatorParams)
		{
			goorl := len(goo.PubKeyTypeURLs)
			if goorl == 0 {
				pbo.PubKeyTypeURLs = nil
			} else {
				var pbos = make([]string, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.PubKeyTypeURLs[i]
						{
							pbos[i] = goore
						}
					}
				}
				pbo.PubKeyTypeURLs = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo ValidatorParams) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ValidatorParams)
	msg = pbo
	return
}
func (goo *ValidatorParams) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ValidatorParams = msg.(*abcipb.ValidatorParams)
	{
		if pbo != nil {
			{
				var pbol int = 0
				if pbo.PubKeyTypeURLs != nil {
					pbol = len(pbo.PubKeyTypeURLs)
				}
				if pbol == 0 {
					(*goo).PubKeyTypeURLs = nil
				} else {
					var goors = make([]string, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.PubKeyTypeURLs[i]
							{
								pboev := pboe
								goors[i] = string(pboev)
							}
						}
					}
					(*goo).PubKeyTypeURLs = goors
				}
			}
		}
	}
	return
}
func (_ ValidatorParams) GetTypeURL() (typeURL string) {
	return "/abci.ValidatorParams"
}
func isValidatorParamsEmptyRepr(goor ValidatorParams) (empty bool) {
	{
		empty = true
		{
			if len(goor.PubKeyTypeURLs) != 0 {
				return false
			}
		}
	}
	return
}
func (goo ValidatorUpdate) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.ValidatorUpdate
	{
		if isValidatorUpdateEmptyRepr(goo) {
			var pbov *abcipb.ValidatorUpdate
			msg = pbov
			return
		}
		pbo = new(abcipb.ValidatorUpdate)
		{
			if goo.PubKey != nil {
				typeUrl := cdc.GetTypeURL(goo.PubKey)
				bz := []byte(nil)
				bz, err = cdc.Marshal(goo.PubKey)
				if err != nil {
					return
				}
				pbo.PubKey = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			pbo.Power = goo.Power
		}
	}
	msg = pbo
	return
}
func (goo ValidatorUpdate) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.ValidatorUpdate)
	msg = pbo
	return
}
func (goo *ValidatorUpdate) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.ValidatorUpdate = msg.(*abcipb.ValidatorUpdate)
	{
		if pbo != nil {
			{
				typeUrl := pbo.PubKey.TypeUrl
				bz := pbo.PubKey.Value
				goorp := &(*goo).PubKey
				err = cdc.UnmarshalAny(typeUrl, bz, goorp)
				if err != nil {
					return
				}
			}
			{
				(*goo).Power = int64(pbo.Power)
			}
		}
	}
	return
}
func (_ ValidatorUpdate) GetTypeURL() (typeURL string) {
	return "/abci.ValidatorUpdate"
}
func isValidatorUpdateEmptyRepr(goor ValidatorUpdate) (empty bool) {
	{
		empty = true
		{
			if goor.PubKey != nil {
				return false
			}
		}
		{
			if goor.Power != 0 {
				return false
			}
		}
	}
	return
}
func (goo LastCommitInfo) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.LastCommitInfo
	{
		if isLastCommitInfoEmptyRepr(goo) {
			var pbov *abcipb.LastCommitInfo
			msg = pbov
			return
		}
		pbo = new(abcipb.LastCommitInfo)
		{
			pbo.Round = goo.Round
		}
		{
			goorl := len(goo.Votes)
			if goorl == 0 {
				pbo.Votes = nil
			} else {
				var pbos = make([]*abcipb.VoteInfo, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Votes[i]
						{
							if goore != nil {
								pbom := proto.Message(nil)
								pbom, err = goore.ToPBMessage(cdc)
								if err != nil {
									return
								}
								pbos[i] = pbom.(*abcipb.VoteInfo)
								if pbos[i] == nil {
									pbos[i] = new(abcipb.VoteInfo)
								}
							}
						}
					}
				}
				pbo.Votes = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo LastCommitInfo) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.LastCommitInfo)
	msg = pbo
	return
}
func (goo *LastCommitInfo) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.LastCommitInfo = msg.(*abcipb.LastCommitInfo)
	{
		if pbo != nil {
			{
				(*goo).Round = int32(pbo.Round)
			}
			{
				var pbol int = 0
				if pbo.Votes != nil {
					pbol = len(pbo.Votes)
				}
				if pbol == 0 {
					(*goo).Votes = nil
				} else {
					var goors = make([]*VoteInfo, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Votes[i]
							{
								pboev := pboe
								if pboev != nil {
									goors[i] = new(VoteInfo)
									err = goors[i].FromPBMessage(cdc, pboev)
									if err != nil {
										return
									}
								}
							}
						}
					}
					(*goo).Votes = goors
				}
			}
		}
	}
	return
}
func (_ LastCommitInfo) GetTypeURL() (typeURL string) {
	return "/abci.LastCommitInfo"
}
func isLastCommitInfoEmptyRepr(goor LastCommitInfo) (empty bool) {
	{
		empty = true
		{
			if goor.Round != 0 {
				return false
			}
		}
		{
			if len(goor.Votes) != 0 {
				return false
			}
		}
	}
	return
}
func (goo VoteInfo) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.VoteInfo
	{
		if isVoteInfoEmptyRepr(goo) {
			var pbov *abcipb.VoteInfo
			msg = pbov
			return
		}
		pbo = new(abcipb.VoteInfo)
		{
			goorl := len(goo.Address)
			if goorl == 0 {
				pbo.Address = nil
			} else {
				var pbos = make([]uint8, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Address[i]
						{
							pbos[i] = byte(goore)
						}
					}
				}
				pbo.Address = pbos
			}
		}
		{
			pbo.Power = goo.Power
		}
		{
			pbo.SignedLastBlock = goo.SignedLastBlock
		}
	}
	msg = pbo
	return
}
func (goo VoteInfo) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.VoteInfo)
	msg = pbo
	return
}
func (goo *VoteInfo) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.VoteInfo = msg.(*abcipb.VoteInfo)
	{
		if pbo != nil {
			{
				var pbol int = 0
				if pbo.Address != nil {
					pbol = len(pbo.Address)
				}
				if pbol == 0 {
					(*goo).Address = nil
				} else {
					var goors = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Address[i]
							{
								pboev := pboe
								goors[i] = uint8(uint8(pboev))
							}
						}
					}
					(*goo).Address = goors
				}
			}
			{
				(*goo).Power = int64(pbo.Power)
			}
			{
				(*goo).SignedLastBlock = bool(pbo.SignedLastBlock)
			}
		}
	}
	return
}
func (_ VoteInfo) GetTypeURL() (typeURL string) {
	return "/abci.VoteInfo"
}
func isVoteInfoEmptyRepr(goor VoteInfo) (empty bool) {
	{
		empty = true
		{
			if len(goor.Address) != 0 {
				return false
			}
		}
		{
			if goor.Power != 0 {
				return false
			}
		}
		{
			if goor.SignedLastBlock != false {
				return false
			}
		}
	}
	return
}
func (goo Validator) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.Validator
	{
		if isValidatorEmptyRepr(goo) {
			var pbov *abcipb.Validator
			msg = pbov
			return
		}
		pbo = new(abcipb.Validator)
		{
			if goo.PubKey != nil {
				typeUrl := cdc.GetTypeURL(goo.PubKey)
				bz := []byte(nil)
				bz, err = cdc.Marshal(goo.PubKey)
				if err != nil {
					return
				}
				pbo.PubKey = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			pbo.Power = goo.Power
		}
	}
	msg = pbo
	return
}
func (goo Validator) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.Validator)
	msg = pbo
	return
}
func (goo *Validator) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.Validator = msg.(*abcipb.Validator)
	{
		if pbo != nil {
			{
				typeUrl := pbo.PubKey.TypeUrl
				bz := pbo.PubKey.Value
				goorp := &(*goo).PubKey
				err = cdc.UnmarshalAny(typeUrl, bz, goorp)
				if err != nil {
					return
				}
			}
			{
				(*goo).Power = int64(pbo.Power)
			}
		}
	}
	return
}
func (_ Validator) GetTypeURL() (typeURL string) {
	return "/abci.Validator"
}
func isValidatorEmptyRepr(goor Validator) (empty bool) {
	{
		empty = true
		{
			if goor.PubKey != nil {
				return false
			}
		}
		{
			if goor.Power != 0 {
				return false
			}
		}
	}
	return
}
func (goo Violation) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.Violation
	{
		if isViolationEmptyRepr(goo) {
			var pbov *abcipb.Violation
			msg = pbov
			return
		}
		pbo = new(abcipb.Violation)
		{
			if goo.Evidence != nil {
				typeUrl := cdc.GetTypeURL(goo.Evidence)
				bz := []byte(nil)
				bz, err = cdc.Marshal(goo.Evidence)
				if err != nil {
					return
				}
				pbo.Evidence = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			goorl := len(goo.Validators)
			if goorl == 0 {
				pbo.Validators = nil
			} else {
				var pbos = make([]*abcipb.Validator, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Validators[i]
						{
							pbom := proto.Message(nil)
							pbom, err = goore.ToPBMessage(cdc)
							if err != nil {
								return
							}
							pbos[i] = pbom.(*abcipb.Validator)
						}
					}
				}
				pbo.Validators = pbos
			}
		}
		{
			pbo.Height = goo.Height
		}
		{
			if !amino.IsEmptyTime(goo.Time) {
				pbo.Time = timestamppb.New(goo.Time)
			}
		}
		{
			pbo.TotalVotingPower = goo.TotalVotingPower
		}
	}
	msg = pbo
	return
}
func (goo Violation) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.Violation)
	msg = pbo
	return
}
func (goo *Violation) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.Violation = msg.(*abcipb.Violation)
	{
		if pbo != nil {
			{
				typeUrl := pbo.Evidence.TypeUrl
				bz := pbo.Evidence.Value
				goorp := &(*goo).Evidence
				err = cdc.UnmarshalAny(typeUrl, bz, goorp)
				if err != nil {
					return
				}
			}
			{
				var pbol int = 0
				if pbo.Validators != nil {
					pbol = len(pbo.Validators)
				}
				if pbol == 0 {
					(*goo).Validators = nil
				} else {
					var goors = make([]Validator, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Validators[i]
							{
								pboev := pboe
								if pboev != nil {
									err = goors[i].FromPBMessage(cdc, pboev)
									if err != nil {
										return
									}
								}
							}
						}
					}
					(*goo).Validators = goors
				}
			}
			{
				(*goo).Height = int64(pbo.Height)
			}
			{
				(*goo).Time = pbo.Time.AsTime()
			}
			{
				(*goo).TotalVotingPower = int64(pbo.TotalVotingPower)
			}
		}
	}
	return
}
func (_ Violation) GetTypeURL() (typeURL string) {
	return "/abci.Violation"
}
func isViolationEmptyRepr(goor Violation) (empty bool) {
	{
		empty = true
		{
			if goor.Evidence != nil {
				return false
			}
		}
		{
			if len(goor.Validators) != 0 {
				return false
			}
		}
		{
			if goor.Height != 0 {
				return false
			}
		}
		{
			if !amino.IsEmptyTime(goor.Time) {
				return false
			}
		}
		{
			if goor.TotalVotingPower != 0 {
				return false
			}
		}
	}
	return
}
func (goo SimpleEvent) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.SimpleEvent
	{
		if isSimpleEventEmptyRepr(goo) {
			var pbov *abcipb.SimpleEvent
			msg = pbov
			return
		}
		pbo = &abcipb.SimpleEvent{Value: goo}
	}
	msg = pbo
	return
}
func (goo SimpleEvent) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.SimpleEvent)
	msg = pbo
	return
}
func (goo *SimpleEvent) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.SimpleEvent = msg.(*abcipb.SimpleEvent)
	{
		*goo = SimpleEvent(pbo.Value)
	}
	return
}
func (_ SimpleEvent) GetTypeURL() (typeURL string) {
	return "/abci.SimpleEvent"
}
func isSimpleEventEmptyRepr(goor SimpleEvent) (empty bool) {
	{
		empty = true
		if goor != "" {
			return false
		}
	}
	return
}
func (goo MockHeader) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *abcipb.MockHeader
	{
		if isMockHeaderEmptyRepr(goo) {
			var pbov *abcipb.MockHeader
			msg = pbov
			return
		}
		pbo = new(abcipb.MockHeader)
		{
			pbo.Version = goo.Version
		}
		{
			pbo.ChainID = goo.ChainID
		}
		{
			pbo.Height = goo.Height
		}
		{
			if !amino.IsEmptyTime(goo.Time) {
				pbo.Time = timestamppb.New(goo.Time)
			}
		}
		{
			pbo.NumTxs = goo.NumTxs
		}
		{
			pbo.TotalTxs = goo.TotalTxs
		}
	}
	msg = pbo
	return
}
func (goo MockHeader) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(abcipb.MockHeader)
	msg = pbo
	return
}
func (goo *MockHeader) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *abcipb.MockHeader = msg.(*abcipb.MockHeader)
	{
		if pbo != nil {
			{
				(*goo).Version = string(pbo.Version)
			}
			{
				(*goo).ChainID = string(pbo.ChainID)
			}
			{
				(*goo).Height = int64(pbo.Height)
			}
			{
				(*goo).Time = pbo.Time.AsTime()
			}
			{
				(*goo).NumTxs = int64(pbo.NumTxs)
			}
			{
				(*goo).TotalTxs = int64(pbo.TotalTxs)
			}
		}
	}
	return
}
func (_ MockHeader) GetTypeURL() (typeURL string) {
	return "/abci.MockHeader"
}
func isMockHeaderEmptyRepr(goor MockHeader) (empty bool) {
	{
		empty = true
		{
			if goor.Version != "" {
				return false
			}
		}
		{
			if goor.ChainID != "" {
				return false
			}
		}
		{
			if goor.Height != 0 {
				return false
			}
		}
		{
			if !amino.IsEmptyTime(goor.Time) {
				return false
			}
		}
		{
			if goor.NumTxs != 0 {
				return false
			}
		}
		{
			if goor.TotalTxs != 0 {
				return false
			}
		}
	}
	return
}
