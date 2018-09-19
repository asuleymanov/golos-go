package social_network

import (
	"github.com/asuleymanov/golos-go/transports"
)

const apiID = "social_network"

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

//GetAccountVotes api request get_account_votes
func (api *API) GetAccountVotes(author string, opts ...interface{}) ([]*Votes, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{author, opts[0]}
	} else {
		params = []interface{}{author}
	}
	var resp []*Votes
	err := api.call("get_account_votes", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetActiveVotes api request get_active_votes
func (api *API) GetActiveVotes(author, permlink string, opts ...interface{}) ([]*VoteState, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{author, permlink, opts[0]}
	} else {
		params = []interface{}{author, permlink}
	}
	var resp []*VoteState
	err := api.call("get_active_votes", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetAllContentReplies api request get_all_content_replies
func (api *API) GetAllContentReplies(parentAuthor, parentPermlink string, opts ...interface{}) ([]*Content, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{parentAuthor, parentPermlink, opts[0]}
	} else {
		params = []interface{}{parentAuthor, parentPermlink}
	}
	var resp []*Content
	err := api.call("get_all_content_replies", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetContent api request get_content
func (api *API) GetContent(author, permlink string, opts ...interface{}) (*Content, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{author, permlink, opts[0]}
	} else {
		params = []interface{}{author, permlink}
	}
	var resp Content
	err := api.call("get_content", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetContentReplies api request get_content_replies
func (api *API) GetContentReplies(parentAuthor, parentPermlink string, opts ...interface{}) ([]*Content, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{parentAuthor, parentPermlink, opts[0]}
	} else {
		params = []interface{}{parentAuthor, parentPermlink}
	}
	var resp []*Content
	err := api.call("get_content_replies", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetRepliesByLastUpdate api request get_replies_by_last_update
func (api *API) GetRepliesByLastUpdate(startAuthor, startPermlink string, limit uint32, opts ...interface{}) ([]*Content, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{startAuthor, startPermlink, limit, opts[0]}
	} else {
		params = []interface{}{startAuthor, startPermlink, limit}
	}
	var resp []*Content
	err := api.call("get_replies_by_last_update", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
