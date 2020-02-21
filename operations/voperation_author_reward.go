package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//AuthorRewardOperation represents author_reward operation data.
type AuthorRewardOperation struct {
	Author        string      `json:"author"`
	Permlink      string      `json:"permlink"`
	SbdPayout     types.Asset `json:"sbd_payout"`
	SteemPayout   types.Asset `json:"steem_payout"`
	VestingPayout types.Asset `json:"vesting_payout"`
}

//Type function that defines the type of operation AuthorRewardOperation.
func (op *AuthorRewardOperation) Type() OpType {
	return TypeAuthorReward
}

//Data returns the operation data AuthorRewardOperation.
func (op *AuthorRewardOperation) Data() interface{} {
	return op
}
