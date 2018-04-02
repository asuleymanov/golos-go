package account_by_key

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
)

const APIID = "account_by_key"

type API struct {
	caller transports.Caller
}

func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var EmptyParams = []string{}

func (api *API) Raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{APIID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to call %v\n", APIID, method)
	}
	return &resp, nil
}

//get_key_references
func (api *API) GetKeyReferences(pubkey string) (*json.RawMessage, error) {
	return api.Raw("get_key_references", []interface{}{pubkey})
}
