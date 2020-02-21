package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//TotalCommentRewardOperation represents total_comment_reward operation data.
type TotalCommentRewardOperation struct {
	Author           string      `json:"author"`
	Permlink         string      `json:"permlink"`
	AuthorReward     types.Asset `json:"author_reward"`
	BenefactorReward types.Asset `json:"benefactor_reward"`
	CuratorReward    types.Asset `json:"curator_reward"`
	NetRshares       int64       `json:"net_rshares"`
}

//Type function that defines the type of operation TotalCommentRewardOperation.
func (op *TotalCommentRewardOperation) Type() OpType {
	return TypeTotalCommentReward
}

//Data returns the operation data TotalCommentRewardOperation.
func (op *TotalCommentRewardOperation) Data() interface{} {
	return op
}
