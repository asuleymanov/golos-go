package types

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

type Asset struct {
	Amount float64
	Symbol string
}

func (op *Asset) UnmarshalJSON(data []byte) error {
	str, errunq := strconv.Unquote(string(data))
	if errunq != nil {
		return errunq
	}
	param := strings.Split(str, " ")
	if s, errpf := strconv.ParseFloat(param[0], 64); errpf != nil {
		return errpf
	} else {
		op.Amount = s
	}
	op.Symbol = param[1]
	return nil
}

func (op *Asset) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.String())
}

func (op *Asset) MarshalTransaction(encoder *transaction.Encoder) error {
	ans, err := json.Marshal(op)
	if err != nil {
		return err
	}

	str, err := strconv.Unquote(string(ans))
	if err != nil {
		return err
	}
	return encoder.EncodeMoney(str)
}

func (op *Asset) String() string {
	ammf := strconv.FormatFloat(op.Amount, 'f', 3, 64)
	return ammf + " " + op.Symbol
}

func (op *Asset) StringAmount() string {
	return strconv.FormatFloat(op.Amount, 'f', 3, 64)
}
