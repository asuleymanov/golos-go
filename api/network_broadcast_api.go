package api

import (
	"github.com/asuleymanov/golos-go/operations"
)

//network_broadcast_api

//BroadcastTransaction api request broadcast_transaction
func (api *API) BroadcastTransaction(tx *operations.Transaction) error {
	return api.call("network_broadcast_api", "broadcast_transaction", []interface{}{tx}, nil)
}

//BroadcastTransactionSynchronous api request broadcast_transaction_synchronous
func (api *API) BroadcastTransactionSynchronous(tx *operations.Transaction, maxBlockAge ...uint32) (*BroadcastResponse, error) {
	var params []interface{}
	if len(maxBlockAge) > 0 {
		params = []interface{}{tx, maxBlockAge[0]}
	} else {
		params = []interface{}{tx}
	}
	var resp BroadcastResponse
	err := api.call("network_broadcast_api", "broadcast_transaction_synchronous", []interface{}{params}, &resp)
	return &resp, err
}
