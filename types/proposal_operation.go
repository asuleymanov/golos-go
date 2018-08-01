package types

import (
	"encoding/json"
)

type ProposalOperation struct {
    Operation Operation `json:"op"`
    OperationType          OpType    `json:"-"`
}

type rawProposalOperation struct {
	Operation              *operationTuple `json:"op"`
}

func (op *ProposalOperation) UnmarshalJSON(p []byte) error {
	var raw rawOperationObject
	if err := json.Unmarshal(p, &raw); err != nil {
		return err
	}

	op.Operation = raw.Operation.Data
	op.OperationType = raw.Operation.Type
	return nil
}


func (op *ProposalOperation) MarshalJSON() ([]byte, error) {
	return JSONMarshal(&rawOperationObject{
		Operation:              &operationTuple{op.Operation.Type(), op.Operation},
	})
}
