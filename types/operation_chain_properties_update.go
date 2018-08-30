package types

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//ChainPropertiesUpdateOperation represents chain_properties_update operation data.
type ChainPropertiesUpdateOperation struct {
	Owner string        `json:"owner"`
	Props []interface{} `json:"props"`
}

//Type function that defines the type of operation ChainPropertiesUpdateOperation.
func (op *ChainPropertiesUpdateOperation) Type() OpType {
	return TypeChainPropertiesUpdate
}

//Data returns the operation data ChainPropertiesUpdateOperation.
func (op *ChainPropertiesUpdateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ChainPropertiesUpdateOperation to bytes.
func (op *ChainPropertiesUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChainPropertiesUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(byte(1))
	z, _ := json.Marshal(op.Props[1])
	var d ChainProperties
	_ = json.Unmarshal(z, &d)
	enc.Encode(&d)
	return enc.Err()
}
