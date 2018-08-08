package types

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//ProposalOperations array of values of type ProposalOperation.
type ProposalOperations []ProposalOperation

//ProposalOperation type from parameter JSON
type ProposalOperation struct {
	Operation     Operation `json:"op"`
	OperationType OpType    `json:"-"`
}

type rawProposalOperation struct {
	Operation *operationTuple `json:"op"`
}

//UnmarshalJSON unpacking the JSON parameter in the ProposalOperation type.
func (op *ProposalOperation) UnmarshalJSON(p []byte) error {
	var raw rawProposalOperation
	if err := json.Unmarshal(p, &raw); err != nil {
		return err
	}

	op.Operation = raw.Operation.Data
	op.OperationType = raw.Operation.Type
	return nil
}

//MarshalJSON function for packing the ProposalOperation type in JSON.
func (op *ProposalOperation) MarshalJSON() ([]byte, error) {
	return JSONMarshal(&rawProposalOperation{
		Operation: &operationTuple{op.Operation.Type(), op.Operation},
	})
}

//MarshalTransaction is a function of converting type ProposalOperations to bytes.
func (op ProposalOperations) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)

	enc.EncodeUVarint(uint64(len(op)))
	for _, v := range op {
		enc.Encode(v.Operation)
	}

	return enc.Err()
}
