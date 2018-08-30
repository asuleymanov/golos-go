package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//ProposalDeleteOperation represents proposal_delete operation data.
type ProposalDeleteOperation struct {
	Author     string        `json:"author"`
	Title      string        `json:"title"`
	Requester  string        `json:"requester"`
	Extensions []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ProposalDeleteOperation.
func (op *ProposalDeleteOperation) Type() OpType {
	return TypeProposalDelete
}

//Data returns the operation data ProposalDeleteOperation.
func (op *ProposalDeleteOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ProposalDeleteOperation to bytes.
func (op *ProposalDeleteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProposalDelete.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Title)
	enc.Encode(op.Requester)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
