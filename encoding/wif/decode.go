package wif

import (
	"fmt"

	"github.com/btcsuite/btcutil"
)

// Decode can be used to turn WIF into a raw private key (32 bytes).
func Decode(wif string) ([]byte, error) {
	w, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return nil, fmt.Errorf("failed to decode WIF : %w", err)
	}

	return w.PrivKey.Serialize(), nil
}

// GetPublicKey returns the public key associated with the given WIF
// in the 33-byte compressed format.
func GetPublicKey(wif string) ([]byte, error) {
	w, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return nil, fmt.Errorf("failed to decode WIF : %w", err)
	}

	return w.PrivKey.PubKey().SerializeCompressed(), nil
}
