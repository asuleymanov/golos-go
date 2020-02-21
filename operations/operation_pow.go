package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//POWOperation represents pow operation data.
type POWOperation struct {
	WorkerAccount string                   `json:"worker_account"`
	BlockID       string                   `json:"block_id"`
	Nonce         uint64                   `json:"nonce"`
	Work          types.POW                `json:"work"`
	Props         types.ChainPropertiesOLD `json:"props"`
}

//Type function that defines the type of operation POWOperation.
func (op *POWOperation) Type() OpType {
	return TypePOW
}

//Data returns the operation data POWOperation.
func (op *POWOperation) Data() interface{} {
	return op
}
