package operations

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//ProposalObjects array of values of type ProposalObject.
type ProposalObjects []ProposalObject

//ProposalObject type from parameter JSON
type ProposalObject struct {
	Operation     Operation `json:"op"`
	OperationType OpType    `json:"-"`
}

type rawProposalObject struct {
	Operation *operationTuple `json:"op"`
}

//UnmarshalJSON unpacking the JSON parameter in the ProposalObject type.
func (op *ProposalObject) UnmarshalJSON(p []byte) error {
	var raw rawProposalObject
	if err := json.Unmarshal(p, &raw); err != nil {
		return err
	}

	op.Operation = raw.Operation.Data
	op.OperationType = raw.Operation.Type
	return nil
}

//MarshalJSON function for packing the ProposalObject type in JSON.
func (op *ProposalObject) MarshalJSON() ([]byte, error) {
	return JSONMarshal(&rawProposalObject{
		Operation: &operationTuple{op.Operation.Type(), op.Operation},
	})
}

//MarshalTransaction is a function of converting type ProposalObjects to bytes.
func (op ProposalObjects) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)

	enc.EncodeUVarint(uint64(len(op)))
	for _, v := range op {
		enc.Encode(v.Operation)
	}

	return enc.Err()
}
