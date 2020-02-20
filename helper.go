package golos

import (
	"github.com/asuleymanov/golos-go/operations"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/types"
)

//SetAsyncProtocol enables or disables the asynchronous operation protocol
func (client *Client) SetAsyncProtocol(value bool) {
	client.asyncProtocol = value
}

//SetKeys you can specify keys for signing transactions.
func (client *Client) SetKeys(keys *Keys) {
	client.CurrentKeys = keys
}

//SetAsset returns data of type Asset
func SetAsset(amount float64, symbol string) *types.Asset {
	return &types.Asset{Amount: amount, Symbol: symbol}
}

//JSONTrxString generate Trx to String
func JSONTrxString(v *transactions.SignedTransaction) (string, error) {
	ans, err := types.JSONMarshal(v)
	if err != nil {
		return "", err
	}
	return string(ans), nil
}

//JSONOpString generate Operations to String
func JSONOpString(v []operations.Operation) (string, error) {
	var tx operations.Operations

	tx = append(tx, v...)

	ans, err := types.JSONMarshal(tx)
	if err != nil {
		return "", err
	}
	return string(ans), nil
}

//GenerateProposalOperation generate []Operation to ProposalOperations
func GenerateProposalOperation(ops []operations.Operation) operations.ProposalObjects {
	var ans operations.ProposalObjects

	for _, val := range ops {
		ans = append(ans, operations.ProposalObject{Operation: val, OperationType: val.Type()})
	}

	return ans
}
