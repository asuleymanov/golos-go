package operations

import (
	"github.com/asuleymanov/golos-go/types"
)

//ConvertSbdDebtOperation represents convert_sbd_debt operation data.
type ConvertSbdDebtOperation struct {
	Owner              string      `json:"owner"`
	SbdAmount          types.Asset `json:"sbd_amount"`
	SteemAmount        types.Asset `json:"steem_amount"`
	SavingsSbdAmount   types.Asset `json:"savings_sbd_amount"`
	SavingsSteemAmount types.Asset `json:"savings_steem_amount"`
}

//Type function that defines the type of operation ConvertSbdDebtOperation.
func (op *ConvertSbdDebtOperation) Type() OpType {
	return TypeConvertSbdDebt
}

//Data returns the operation data ConvertSbdDebtOperation.
func (op *ConvertSbdDebtOperation) Data() interface{} {
	return op
}
