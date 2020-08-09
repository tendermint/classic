package version

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/mod/semver"
)

// ProtocolVersion is used to negotiate between clients.
type ProtocolVersion struct {
	Name     string // abci, p2p, app, block, etc.
	Version  string
	Optional bool // default required.
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

// Returns an error if not compatible.
// Otherwise, returns the set of compatible interfaces.
// Only the Major and Minor versions are returned; Patch, Prerelease, and Build
// portions of Semver2.0 are discarded in the resulting intersection
// ProtocolVersionSet.
// TODO: test
func (pvs ProtocolVersionSet) CompatibleWith(other ProtocolVersionSet) (res ProtocolVersionSet, err error) {
	var errs []string
	type pvpair [2]*ProtocolVersion
	name2Pair := map[string]*pvpair{}
	for _, pv := range pvs {
		name2Pair[pv.Name] = &pvpair{&pv, nil}
	}
	for _, pv := range other {
		item, ok := name2Pair[pv.Name]
		if ok {
			item[1] = &pv
		} else {
			name2Pair[pv.Name] = &pvpair{nil, &pv}
		}
	}
	for _, pair := range name2Pair {
		pv1, pv2 := pair[0], pair[1]
		if pv1 == nil {
			if pv2.Optional {
				// fine
			} else {
				errs = append(errs, fmt.Sprintf("Other ProtocolVersionSet requires %v", pv2))
			}
		} else if pv2 == nil {
			if pv1.Optional {
				// fine
			} else {
				errs = append(errs, fmt.Sprintf("Our ProtocolVersionSet requires %v", pv1))
			}
		} else {
			pv1mm := semver.MajorMinor(pv1.Version)
			pv2mm := semver.MajorMinor(pv2.Version)
			if semver.Major(pv1mm) == semver.Major(pv2mm) {
				if semver.Compare(semver.Major(pv1mm), semver.Major(pv2mm)) > 0 {
					res = append(res, ProtocolVersion{Name: pv1.Name, Version: pv2mm, Optional: pv1.Optional && pv2.Optional})
				} else {
					res = append(res, ProtocolVersion{Name: pv1.Name, Version: pv1mm, Optional: pv1.Optional && pv2.Optional})
				}
			} else {
				errs = append(errs, fmt.Sprintf("ProtocolVersions not compatible: %v vs %v", pv1, pv2))
			}
		}
	}
	if errs != nil {
		return res, fmt.Errorf("ProtocolVersionSet not compatible...\n%s", strings.Join(errs, "\n"))
	}
	return res, nil
}
