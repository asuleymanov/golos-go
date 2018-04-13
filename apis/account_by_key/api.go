package account_by_key

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
)

const apiID = "account_by_key"

type API struct {
	caller transports.Caller
}

func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = struct{}{}

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//GetKeyReferences
func (api *API) GetKeyReferences(pubkey string) (*json.RawMessage, error) {
	return api.raw("get_key_references", []interface{}{pubkey})
}
