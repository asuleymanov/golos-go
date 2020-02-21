package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//TransitToCyberwayOperation represents transit_to_cyberway operation data.
type TransitToCyberwayOperation struct {
	Owner         string `json:"owner"`
	VoteToTransit bool   `json:"vote_to_transit"`
}

//Type function that defines the type of operation TransitToCyberwayOperation.
func (op *TransitToCyberwayOperation) Type() OpType {
	return TypeTransitToCyberway
}

//Data returns the operation data TransitToCyberwayOperation.
func (op *TransitToCyberwayOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type TransitToCyberwayOperation to bytes.
func (op *TransitToCyberwayOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransitToCyberway.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.VoteToTransit)
	return enc.Err()
}
