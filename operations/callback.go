package operations

import (
	"encoding/json"
)

type CallbackBlockOperations struct {
	TrxInBlock    int       `json:"trx_in_block"`
	OpInTrx       int       `json:"op_in_trx"`
	VirtualOp     int       `json:"virtual_op"`
	Operation     Operation `json:"op"`
	OperationType OpType    `json:"-"`
}

type rawCallbackBlockOperations struct {
	TrxInBlock int            `json:"trx_in_block"`
	OpInTrx    int            `json:"op_in_trx"`
	VirtualOp  int            `json:"virtual_op"`
	Operation  operationTuple `json:"op"`
}

//UnmarshalJSON unpacking the JSON parameter in the CallbackBlockOperations type.
func (op *CallbackBlockOperations) UnmarshalJSON(p []byte) error {
	var raw rawCallbackBlockOperations
	if err := json.Unmarshal(p, &raw); err != nil {
		return err
	}

	op.TrxInBlock = raw.TrxInBlock
	op.OpInTrx = raw.OpInTrx
	op.VirtualOp = raw.VirtualOp
	op.Operation = raw.Operation.Data
	op.OperationType = raw.Operation.Type
	return nil
}

//MarshalJSON function for packing the CallbackBlockOperations type in JSON.
func (op *CallbackBlockOperations) MarshalJSON() ([]byte, error) {
	return JSONMarshal(&rawCallbackBlockOperations{
		TrxInBlock: op.TrxInBlock,
		OpInTrx:    op.OpInTrx,
		VirtualOp:  op.VirtualOp,
		Operation:  operationTuple{op.Operation.Type(), op.Operation},
	})
}
