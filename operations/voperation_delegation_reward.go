package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//DelegationRewardOperation represents delegation_reward operation data.
type DelegationRewardOperation struct {
	Delegator      string      `json:"delegator"`
	Delegatee      string      `json:"delegatee"`
	PayoutStrategy int         `json:"payout_strategy"`
	VestingShares  types.Asset `json:"vesting_shares"`
}

//Type function that defines the type of operation DelegationRewardOperation.
func (op *DelegationRewardOperation) Type() OpType {
	return TypeDelegationReward
}

//Data returns the operation data DelegationRewardOperation.
func (op *DelegationRewardOperation) Data() interface{} {
	return op
}
