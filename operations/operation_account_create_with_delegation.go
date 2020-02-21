package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"
)

//AccountCreateWithDelegationOperation represents account_create_with_delegation operation data.
type AccountCreateWithDelegationOperation struct {
	Fee            types.Asset     `json:"fee"`
	Delegation     types.Asset     `json:"delegation"`
	Creator        string          `json:"creator"`
	NewAccountName string          `json:"new_account_name"`
	Owner          types.Authority `json:"owner"`
	Active         types.Authority `json:"active"`
	Posting        types.Authority `json:"posting"`
	MemoKey        string          `json:"memo_key"`
	JSONMetadata   string          `json:"json_metadata"`
	Extensions     []interface{}   `json:"extensions"`
}

//Type function that defines the type of operation AccountCreateWithDelegationOperation.
func (op *AccountCreateWithDelegationOperation) Type() OpType {
	return TypeAccountCreateWithDelegation
}

//Data returns the operation data AccountCreateWithDelegationOperation.
func (op *AccountCreateWithDelegationOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountCreateWithDelegationOperation to bytes.
func (op *AccountCreateWithDelegationOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountCreateWithDelegation.Code()))
	enc.Encode(op.Fee)
	enc.Encode(op.Delegation)
	enc.Encode(op.Creator)
	enc.Encode(op.NewAccountName)
	enc.Encode(op.Owner)
	enc.Encode(op.Active)
	enc.Encode(op.Posting)
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	/*if len(op.Extensions)>0{

	}else{*/
	enc.Encode(byte(0))
	//}
	return enc.Err()
}
