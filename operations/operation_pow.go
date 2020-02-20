package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//POWOperation represents pow operation data.
type POWOperation struct {
	WorkerAccount string                 `json:"worker_account"`
	BlockID       string                 `json:"block_id"`
	Nonce         *types.Int             `json:"nonce"`
	Work          *types.POW             `json:"work"`
	Props         *types.ChainProperties `json:"props"`
}

//Type function that defines the type of operation POWOperation.
func (op *POWOperation) Type() OpType {
	return TypePOW
}

//Data returns the operation data POWOperation.
func (op *POWOperation) Data() interface{} {
	return op
}
