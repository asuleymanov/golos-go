package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//WorkerRequestDeleteOperation represents worker_request_delete operation data.
type WorkerRequestDeleteOperation struct {
	Author     string        `json:"author"`
	Permlink   string        `json:"permlink"`
	Extensions []interface{} `json:"extensions"`
}

//Type function that defines the type of operation WorkerRequestDeleteOperation.
func (op *WorkerRequestDeleteOperation) Type() OpType {
	return TypeWorkerRequestDelete
}

//Data returns the operation data WorkerRequestDeleteOperation.
func (op *WorkerRequestDeleteOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type WorkerRequestDeleteOperation to bytes.
func (op *WorkerRequestDeleteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWorkerRequestDelete.Code()))
	enc.EncodeString(op.Author)
	enc.EncodeString(op.Permlink)
	//enc.Encode(Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
