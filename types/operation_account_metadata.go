package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//AccountMetadataOperation represents account_metadata operation data.
type AccountMetadataOperation struct {
	Account      string           `json:"account"`
	JSONMetadata *AccountMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation AccountMetadataOperation.
func (op *AccountMetadataOperation) Type() OpType {
	return TypeAccountMetadata
}

//Data returns the operation data AccountMetadataOperation.
func (op *AccountMetadataOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountMetadataOperation to bytes.
func (op *AccountMetadataOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountMetadata.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}
