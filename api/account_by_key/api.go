package account_by_key

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
)

const apiID = "account_by_key"

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

//GetKeyReferences api request get_key_references
func (api *API) GetKeyReferences(pubkey string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_key_references", []interface{}{pubkey}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
