package version

import (
	"sort"

	"golang.org/x/mod/semver"

	abciver "github.com/tendermint/classic/abci/version"
	bcver "github.com/tendermint/classic/blockchain/version"
	p2pver "github.com/tendermint/classic/p2p/version"
)

var (
	// The major or minor versions must bump when components bump.
	// The TendermintClassic software version
	Version = "v1.0.0-rc.0"
)

func init() {
	if abciver.Version != "v1.0.0-rc.0" ||
		bcver.Version != "v1.0.0-rc.0" ||
		p2pver.Version != "v1.0.0-rc.0" {
		panic("bump Version")
	}
}

// ProtocolVersion is used to negotiate between clients.
type ProtocolVersion struct {
	Name    string // abci, p2p, app, block, etc.
	Version string
}

type ProtocolVersionSet []ProtocolVersion

func (pvs ProtocolVersionSet) Sort() {
	sort.Slice(pvs, func(i, j int) bool {
		if pvs[i].Name < pvs[j].Name {
			return true
		} else if pvs[i].Name == pvs[j].Name {
			panic("should not happen")
		} else {
			return false
		}
	})
}

func (pvs ProtocolVersionSet) GetProtocol(name string) (pv ProtocolVersion, ok bool) {
	for _, pv := range pvs {
		if pv.Name == name {
			return pv, true
		}
	}
	ok = false
	return
}
