package transactions

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/asuleymanov/golos-go/types"
)

//RefBlockNum function returns blockNumber
func RefBlockNum(blockNumber uint32) types.UInt16 {
	return types.UInt16(blockNumber)
}

//RefBlockPrefix function returns block prefix
func RefBlockPrefix(blockID string) (types.UInt32, error) {
	// Block ID is hex-encoded.
	rawBlockID, err := hex.DecodeString(blockID)
	if err != nil {
		return 0, fmt.Errorf("networkbroadcast: failed to decode block ID: %v \n Error : %w", blockID, err)
	}

	// Raw prefix = raw block ID [4:8].
	// Make sure we don't trigger a slice bounds out of range panic.
	if len(rawBlockID) < 8 {
		return 0, fmt.Errorf("networkbroadcast: invalid block ID: %v", blockID)
	}
	rawPrefix := rawBlockID[4:8]

	// Decode the prefix.
	var prefix uint32
	if err := binary.Read(bytes.NewReader(rawPrefix), binary.LittleEndian, &prefix); err != nil {
		return 0, fmt.Errorf("networkbroadcast: failed to read block prefix: %v \n Error : %w", rawPrefix, err)
	}

	// Done, return the prefix.
	return types.UInt32(prefix), nil
}
