package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"
)

//DelegateVestingSharesWithInterestOperation represents delegate_vesting_shares_with_interest operation data.
/*
PayoutStrategy:
to_delegator - 0
to_delegated_vesting - 1
*/
type DelegateVestingSharesWithInterestOperation struct {
	Delegator      string        `json:"delegator"`
	Delegatee      string        `json:"delegatee"`
	VestingShares  *types.Asset  `json:"vesting_shares"`
	InterestRate   uint16        `json:"interest_rate"`
	PayoutStrategy int           `json:"payout_strategy"`
	Extensions     []interface{} `json:"extensions"`
}

//Type function that defines the type of operation DelegateVestingSharesWithInterestOperation.
func (op *DelegateVestingSharesWithInterestOperation) Type() OpType {
	return TypeDelegateVestingSharesWithInterest
}

//Data returns the operation data DelegateVestingSharesWithInterestOperation.
func (op *DelegateVestingSharesWithInterestOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type DelegateVestingSharesWithInterestOperation to bytes.
func (op *DelegateVestingSharesWithInterestOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDelegateVestingSharesWithInterest.Code()))
	enc.Encode(op.Delegator)
	enc.Encode(op.Delegatee)
	enc.Encode(op.VestingShares)
	enc.Encode(op.InterestRate)
	enc.Encode(op.PayoutStrategy)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
