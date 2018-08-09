package private_message

import (
	"encoding/json"
	"time"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
)

const apiID = "private_message"

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

//GetInbox api request get_inbox
func (api *API) GetInbox(to string, limit, offset uint16) (*[]Message, error) {
	raw, err := api.raw("get_inbox", []interface{}{to, time.Now(), limit, offset})
	if err != nil {
		return nil, err
	}
	var resp []Message
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_active_witnesses response", apiID)
	}
	return &resp, nil
}

//GetOutbox api request get_outbox
func (api *API) GetOutbox(from string, limit, offset uint16) ([]*Message, error) {
	raw, err := api.raw("get_outbox", []interface{}{from, time.Now(), limit, offset})
	if err != nil {
		return nil, err
	}
	var resp []*Message
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_active_witnesses response", apiID)
	}
	return resp, nil
}
