package operation_history

import (
	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
)

const apiID = "operation_history"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetOpsInBlock api request get_ops_in_block
func (api *API) GetOpsInBlock(blockNum uint32, onlyVirtual bool) ([]*types.OperationObject, error) {
	var resp []*types.OperationObject
	err := api.call("get_ops_in_block", []interface{}{blockNum, onlyVirtual}, &resp)
	return resp, err
}

//GetTransaction api request get_transaction
func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	var resp types.Transaction
	err := api.call("get_transaction", []string{id}, &resp)
	return &resp, err
}
