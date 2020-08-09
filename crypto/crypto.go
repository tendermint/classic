package crypto

import (
	"fmt"

	"github.com/tendermint/classic/crypto/tmhash"
	"github.com/tendermint/classic/libs/bech32"
)

//----------------------------------------
// Address

const (
	// AddressSize is the size of a pubkey address.
	AddressSize = tmhash.TruncatedSize
)

// An address is a []byte, but hex-encoded even in JSON.
// []byte leaves us the option to change the address length.
type Address []byte

func AddressHash(bz []byte) Address {
	return Address(tmhash.SumTruncated(bz))
}

func (addr Address) String() string {
	str, err := addr.MarshalAmino()
	if err != nil {
		panic(err)
	}
	return str
}

func (addr Address) MarshalAmino() (string, error) {
	// The "c" bech32 is intended to be constant,
	// and enforced upon all users of the tendermint/classic repo
	// and derivations of tendermint/classic.
	bech32Addr, err := bech32.Encode("c", addr)
	if err != nil {
		return "", err
	}
	return bech32Addr, nil
}

func (addr *Address) UnmarshalAmino(str string) error {
	pre, bz, err := bech32.Decode(str)
	if err != nil {
		return err
	}
	if pre != "c" {
		return fmt.Errorf("unexpected bech32 prefix for address. expected \"c\", got %v", pre)
	}
	*addr = cp(bz)
	return nil
}

//----------------------------------------
// PubKey

// All operations must be deterministic.
type PubKey interface {
	// Stable
	Address() Address
	Bytes() []byte
	VerifyBytes(msg []byte, sig []byte) bool
	Equals(PubKey) bool

	// Unstable
	String() string
}

//----------------------------------------
// PrivKey

// All operations must be deterministic.
type PrivKey interface {
	// Stable
	Bytes() []byte
	Sign(msg []byte) ([]byte, error)
	PubKey() PubKey
	Equals(PrivKey) bool
}

//----------------------------------------
// Symmetric

type Symmetric interface {
	Keygen() []byte
	Encrypt(plaintext []byte, secret []byte) (ciphertext []byte)
	Decrypt(ciphertext []byte, secret []byte) (plaintext []byte, err error)
}
