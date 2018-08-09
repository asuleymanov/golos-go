package operation_history

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
	"github.com/pkg/errors"
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

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//GetOpsInBlock api request get_ops_in_block
func (api *API) GetOpsInBlock(blockNum uint32, onlyVirtual bool) ([]*types.OperationObject, error) {
	raw, err := api.raw("get_ops_in_block", []interface{}{blockNum, onlyVirtual})
	if err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_ops_in_block response", apiID)
	}
	return resp, nil
}

//GetTransaction api request get_transaction
func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	raw, err := api.raw("get_transaction", []string{id})
	if err != nil {
		return nil, err
	}
	var resp types.Transaction
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_transaction response", apiID)
	}
	return &resp, nil
}
