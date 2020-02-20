package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//RejectVestingSharesDelegationOperation represents reject_vesting_shares_delegation operation data.
type RejectVestingSharesDelegationOperation struct {
	Delegator  string        `json:"delegator"`
	Delegatee  string        `json:"delegatee"`
	Extensions []interface{} `json:"extensions"`
}

//Type function that defines the type of operation RejectVestingSharesDelegationOperation.
func (op *RejectVestingSharesDelegationOperation) Type() OpType {
	return TypeRejectVestingSharesDelegation
}

//Data returns the operation data RejectVestingSharesDelegationOperation.
func (op *RejectVestingSharesDelegationOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type RejectVestingSharesDelegationOperation to bytes.
func (op *RejectVestingSharesDelegationOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeRejectVestingSharesDelegation.Code()))
	enc.Encode(op.Delegator)
	enc.Encode(op.Delegatee)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
