package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//BreakFreeReferralOperation represents break_free_referral operation data.
type BreakFreeReferralOperation struct {
	Referral   string        `json:"referral"`
	Extensions []interface{} `json:"extensions"`
}

//Type function that defines the type of operation BreakFreeReferralOperation.
func (op *BreakFreeReferralOperation) Type() OpType {
	return TypeBreakFreeReferral
}

//Data returns the operation data BreakFreeReferralOperation.
func (op *BreakFreeReferralOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type BreakFreeReferralOperation to bytes.
func (op *BreakFreeReferralOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeBreakFreeReferral.Code()))
	enc.Encode(op.Referral)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
