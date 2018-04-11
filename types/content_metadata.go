package types

import (
	"encoding/json"
	"strconv"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

type ContentMetadata map[string]interface{}

func (op *ContentMetadata) UnmarshalJSON(p []byte) error {
	str, _ := strconv.Unquote(string(p))
	if str == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(str), *op); err != nil {
		return err
	}

	return nil
}

func (op *ContentMetadata) MarshalJSON() ([]byte, error) {
	ans, err := json.Marshal(*op)
	if err != nil {
		return []byte{}, err
	}
	return []byte(strconv.Quote(string(ans))), nil
}

func (op *ContentMetadata) MarshalTransaction(encoder *transaction.Encoder) error {
	ans, err := json.Marshal(*op)
	if err != nil {
		return err
	}

	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeString(string(ans))
	return enc.Err()
}
