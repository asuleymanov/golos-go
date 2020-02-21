package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//AccountReferral is an additional structure used by other structures.
type AccountReferral struct {
	Referrer     string `json:"referrer"`
	InterestRate uint16 `json:"interest_rate"`
	EndDate      Time   `json:"end_date"`
	BreakFee     Asset  `json:"break_fee"`
}

//MarshalTransaction is a function of converting type AccountReferral to bytes.
func (cp *AccountReferral) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.Referrer)
	enc.Encode(cp.InterestRate)
	enc.Encode(cp.EndDate)
	enc.Encode(cp.BreakFee)
	return enc.Err()
}
