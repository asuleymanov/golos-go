package account_history

import (
	"encoding/json"
	"fmt"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
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

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetAccountHistory api request get_account_history
func (api *API) GetAccountHistory(account string, from int64, limit uint32) ([]*types.OperationObject, error) {
	if limit > 10000 {
		return nil, fmt.Errorf("%v: get_account_history -> limit must not exceed 10000", apiID)
	}
	if from == 0 {
		return nil, fmt.Errorf("%v: get_account_history -> from can not have the value 0", apiID)
	}
	if from < int64(limit) && !(from < 0) {
		return nil, fmt.Errorf("%v: get_account_history -> from must be greater than or equal to the limit", apiID)
	}
	var raw json.RawMessage
	err := api.call("get_account_history", []interface{}{account, from, limit}, &raw)
	if err != nil {
		return nil, err
	}
	var tmp1 [][]interface{}
	if err := json.Unmarshal([]byte(raw), &tmp1); err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	for _, v := range tmp1 {
		byteData, errm := json.Marshal(&v[1])
		if errm != nil {
			return nil, errm
		}
		var tmp *types.OperationObject
		if err := json.Unmarshal(byteData, &tmp); err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}
