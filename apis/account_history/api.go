package account_history

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
	"github.com/pkg/errors"
)

const apiID = "account_history"

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

//GetAccountHistory api request get_account_history
func (api *API) GetAccountHistory(account string, from int64, limit uint32) ([]*types.OperationObject, error) {
	raw, err := api.raw("get_account_history", []interface{}{account, from, limit})
	if err != nil {
		return nil, err
	}
	var tmp1 [][]interface{}
	if err := json.Unmarshal([]byte(*raw), &tmp1); err != nil {
		return nil, errors.Wrapf(err, "")
	}
	var resp []*types.OperationObject
	for _, v := range tmp1 {
		byteData, errm := json.Marshal(&v[1])
		if errm != nil {
			return nil, errors.Wrapf(errm, "")
		}
		var tmp *types.OperationObject
		if err := json.Unmarshal(byteData, &tmp); err != nil {
			return nil, errors.Wrapf(err, "")
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}
