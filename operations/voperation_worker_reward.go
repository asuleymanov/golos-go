package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//WorkerRewardOperation represents worker_reward operation data.
type WorkerRewardOperation struct {
	Worker                string      `json:"worker"`
	WorkerRequestAuthor   string      `json:"worker_request_author"`
	WorkerRequestPermlink string      `json:"worker_request_permlink"`
	Reward                types.Asset `json:"reward"`
	RewardInVestsIfVest   types.Asset `json:"reward_in_vests_if_vest"`
}

//Type function that defines the type of operation WorkerRewardOperation.
func (op *WorkerRewardOperation) Type() OpType {
	return TypeWorkerReward
}

//Data returns the operation data WorkerRewardOperation.
func (op *WorkerRewardOperation) Data() interface{} {
	return op
}
