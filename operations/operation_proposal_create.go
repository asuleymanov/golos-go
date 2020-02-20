package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"
)

//ProposalCreateOperation represents proposal_create operation data.
type ProposalCreateOperation struct {
	Author             string          `json:"author"`
	Title              string          `json:"title"`
	Memo               string          `json:"memo"`
	ExpirationTime     *types.Time     `json:"expiration_time"`
	ProposedOperations ProposalObjects `json:"proposed_operations"`
	ReviewPeriodTime   *types.Time     `json:"review_period_time,omnitempty"`
	Extensions         []interface{}   `json:"extensions"`
}

//Type function that defines the type of operation ProposalCreateOperation.
func (op *ProposalCreateOperation) Type() OpType {
	return TypeProposalCreate
}

//Data returns the operation data ProposalCreateOperation.
func (op *ProposalCreateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ProposalCreateOperation to bytes.
func (op *ProposalCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProposalCreate.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Title)
	enc.Encode(op.Memo)
	enc.Encode(op.ExpirationTime)
	enc.Encode(op.ProposedOperations)
	enc.Encode(byte(1))
	enc.Encode(op.ReviewPeriodTime)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
