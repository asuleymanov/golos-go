package api

import (
	"github.com/asuleymanov/golos-go/types"
)

//operation_history

//GetOpsInBlock api request get_ops_in_block
func (api *API) GetOpsInBlock(blockNum uint32, onlyVirtual bool) ([]*types.OperationObject, error) {
	var resp []*types.OperationObject
	err := api.call("operation_history", "get_ops_in_block", []interface{}{blockNum, onlyVirtual}, &resp)
	return resp, err
}

//GetTransaction api request get_transaction
func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	var resp types.Transaction
	err := api.call("operation_history", "get_transaction", []string{id}, &resp)
	return &resp, err
}
