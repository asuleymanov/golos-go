package private_message

import (
	"time"

	"github.com/asuleymanov/golos-go/transports"
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

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetInbox api request get_inbox
func (api *API) GetInbox(to string, limit, offset uint16) ([]*Message, error) {
	var resp []*Message
	err := api.call("get_inbox", []interface{}{to, time.Now(), limit, offset}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetOutbox api request get_outbox
func (api *API) GetOutbox(from string, limit, offset uint16) ([]*Message, error) {
	var resp []*Message
	err := api.call("get_outbox", []interface{}{from, time.Now(), limit, offset}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
