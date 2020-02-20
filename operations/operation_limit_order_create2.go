package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"
)

//LimitOrderCreate2Operation represents limit_order_create2 operation data.
type LimitOrderCreate2Operation struct {
	Owner        string          `json:"owner"`
	OrderID      uint32          `json:"orderid"`
	AmountToSell *types.Asset    `json:"amount_to_sell"`
	ExchangeRate *types.ExchRate `json:"exchange_rate"`
	FillOrKill   bool            `json:"fill_or_kill"`
	Expiration   *types.Time     `json:"expiration"`
}

//Type function that defines the type of operation LimitOrderCreate2Operation.
func (op *LimitOrderCreate2Operation) Type() OpType {
	return TypeLimitOrderCreate2
}

//Data returns the operation data LimitOrderCreate2Operation.
func (op *LimitOrderCreate2Operation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type LimitOrderCreate2Operation to bytes.
func (op *LimitOrderCreate2Operation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate2.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	enc.Encode(op.AmountToSell)
	enc.EncodeBool(op.FillOrKill)
	enc.Encode(op.ExchangeRate)
	enc.Encode(op.Expiration)
	return enc.Err()
}
