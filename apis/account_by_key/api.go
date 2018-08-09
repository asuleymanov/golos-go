package account_by_key

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
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

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//GetKeyReferences api request get_key_references
func (api *API) GetKeyReferences(pubkey string) (*json.RawMessage, error) {
	return api.raw("get_key_references", []interface{}{pubkey})
}
