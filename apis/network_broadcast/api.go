package network_broadcast

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
	"github.com/pkg/errors"
)

const apiID = "network_broadcast_api"

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

//BroadcastTransaction api request broadcast_transaction
func (api *API) BroadcastTransaction(tx *types.Transaction) error {
	return errors.Wrapf(api.call("broadcast_transaction", []interface{}{tx}, nil), "")
}

//BroadcastTransactionSynchronous api request broadcast_transaction_synchronous
func (api *API) BroadcastTransactionSynchronous(tx *types.Transaction) (*BroadcastResponse, error) {
	var raw json.RawMessage
	err := api.call("broadcast_transaction_synchronous", []interface{}{tx}, &raw)
	if err != nil {
		return nil, err
	}

	var resp BroadcastResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal BroadcastResponse: %v", string(raw))
	}
	return &resp, nil
}
