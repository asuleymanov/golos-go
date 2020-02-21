package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//WorkerRequestVoteOperation represents worker_request_vote operation data.
type WorkerRequestVoteOperation struct {
	Voter       string        `json:"voter"`
	Author      string        `json:"author"`
	Permlink    string        `json:"permlink"`
	VotePercent int16         `json:"vote_percent"`
	Extensions  []interface{} `json:"extensions"`
}

//Type function that defines the type of operation WorkerRequestVoteOperation.
func (op *WorkerRequestVoteOperation) Type() OpType {
	return TypeWorkerRequestVote
}

//Data returns the operation data WorkerRequestVoteOperation.
func (op *WorkerRequestVoteOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type WorkerRequestVoteOperation to bytes.
func (op *WorkerRequestVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWorkerRequestVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.VotePercent)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
