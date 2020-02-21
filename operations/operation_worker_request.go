package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"
)

//WorkerRequestOperation represents worker_request operation data.
type WorkerRequestOperation struct {
	Author            string        `json:"author"`
	Permlink          string        `json:"permlink"`
	Worker            string        `json:"worker"`
	RequiredAmountMin types.Asset   `json:"required_amount_min"`
	RequiredAmountMax types.Asset   `json:"required_amount_max"`
	VestReward        bool          `json:"vest_reward"`
	Duration          uint32        `json:"duration"`
	Extensions        []interface{} `json:"extensions"`
}

//Type function that defines the type of operation WorkerRequestOperation.
func (op *WorkerRequestOperation) Type() OpType {
	return TypeWorkerRequest
}

//Data returns the operation data WorkerRequestOperation.
func (op *WorkerRequestOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type WorkerRequestOperation to bytes.
func (op *WorkerRequestOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWorkerRequest.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Worker)
	enc.Encode(op.RequiredAmountMin)
	enc.Encode(op.RequiredAmountMax)
	enc.Encode(op.VestReward)
	enc.Encode(op.Duration)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
