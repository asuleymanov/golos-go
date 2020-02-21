package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//ReturnVestingDelegationOperation represents return_vesting_delegation operation data.
type ReturnVestingDelegationOperation struct {
	Account       string      `json:"account"`
	VestingShares types.Asset `json:"vesting_shares"`
}

//Type function that defines the type of operation ReturnVestingDelegationOperation.
func (op *ReturnVestingDelegationOperation) Type() OpType {
	return TypeReturnVestingDelegation
}

//Data returns the operation data ReturnVestingDelegationOperation.
func (op *ReturnVestingDelegationOperation) Data() interface{} {
	return op
}
