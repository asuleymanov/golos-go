// +build !nosigning

package transactions

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/operations"
	"github.com/asuleymanov/golos-go/types"
)

//SignedTransaction structure of a signed transaction
type SignedTransaction struct {
	*operations.Transaction
}

//NewSignedTransaction initialization of a new signed transaction
func NewSignedTransaction(tx *operations.Transaction) *SignedTransaction {
	if tx.Expiration == nil {
		expiration := time.Now().Add(30 * time.Second).UTC()
		tx.Expiration = &types.Time{Time: &expiration}
	}

	return &SignedTransaction{tx}
}

//Serialize function serializes a transaction
func (tx *SignedTransaction) Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	if err := encoder.Encode(tx.Transaction); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

//Digest function that returns a digest from a serialized transaction
func (tx *SignedTransaction) Digest(chain string) ([]byte, error) {
	var msgBuffer bytes.Buffer

	// Write the chain ID.
	rawChainID, err := hex.DecodeString(chain)
	if err != nil {
		return nil, fmt.Errorf("failed to decode chain ID: %v \n Error : %s", chain, err)
	}

	if _, err := msgBuffer.Write(rawChainID); err != nil {
		return nil, fmt.Errorf("failed to write chain ID : %s", err)
	}

	// Write the serialized transaction.
	rawTx, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	if _, err := msgBuffer.Write(rawTx); err != nil {
		return nil, fmt.Errorf("failed to write serialized transaction : %s", err)
	}

	// Compute the digest.
	digest := sha256.Sum256(msgBuffer.Bytes())
	return digest[:], nil
}

//Sign function directly generating transaction signature
func (tx *SignedTransaction) Sign(privKeys [][]byte, chain string) error {
	var buf bytes.Buffer
	chainid, errdec := hex.DecodeString(chain)
	if errdec != nil {
		return errdec
	}

	txRaw, err := tx.Serialize()
	if err != nil {
		return err
	}

	buf.Write(chainid)
	buf.Write(txRaw)
	data := buf.Bytes()
	//msg_sha := crypto.Sha256(buf.Bytes())

	var sigsHex []string

	for _, privB := range privKeys {
		sigBytes, err := tx.SignSingle(privB, data)
		if err != nil {
			return err
		}
		sigsHex = append(sigsHex, hex.EncodeToString(sigBytes))
	}

	tx.Transaction.Signatures = sigsHex
	return nil
}
