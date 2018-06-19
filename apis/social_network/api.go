package social_network

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
)

const apiID = "social_network"

type API struct {
	caller transports.Caller
}

func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var EmptyParams = []string{}

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//GetRepliesByLastUpdate api request get_replies_by_last_update
func (api *API) GetRepliesByLastUpdate(startAuthor, startPermlink string, limit uint32, opts ...interface{}) ([]*Content, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{startAuthor, startPermlink, limit, opts[0]}
	} else {
		params = []interface{}{startAuthor, startPermlink, limit}
	}
	raw, err := api.raw("get_replies_by_last_update", params)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_replies_by_last_update response", apiID)
	}
	return resp, nil
}

//GetTrendingCategories api request get_trending_categories
func (api *API) GetTrendingCategories(after string, limit uint32) ([]*Categories, error) {
	raw, err := api.raw("get_trending_categories", []interface{}{after, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Categories
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_trending_categories response", apiID)
	}
	return resp, nil
}

//GetBestCategories api request get_best_categories
func (api *API) GetBestCategories(after string, limit uint32) (*json.RawMessage, error) {
	return api.raw("get_best_categories", []interface{}{after, limit})
}

//GetActiveCategories api request get_active_categories
func (api *API) GetActiveCategories(after string, limit uint32) (*json.RawMessage, error) {
	return api.raw("get_active_categories", []interface{}{after, limit})
}

//GetRecentCategories api request get_recent_categories
func (api *API) GetRecentCategories(after string, limit uint32) (*json.RawMessage, error) {
	return api.raw("get_recent_categories", []interface{}{after, limit})
}

//GetActiveVotes api request get_active_votes
func (api *API) GetActiveVotes(author, permlink string, opts ...interface{}) ([]*VoteState, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{author, permlink, opts[0]}
	} else {
		params = []interface{}{author, permlink}
	}
	raw, err := api.raw("get_active_votes", params)
	if err != nil {
		return nil, err
	}
	var resp []*VoteState
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_active_votes response", apiID)
	}
	return resp, nil
}

//GetAccountVotes api request get_account_votes
func (api *API) GetAccountVotes(author string, opts ...interface{}) ([]*Votes, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{author, opts[0]}
	} else {
		params = []interface{}{author}
	}
	raw, err := api.raw("get_account_votes", params)
	if err != nil {
		return nil, err
	}
	var resp []*Votes
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_account_votes response", apiID)
	}
	return resp, nil
}

//GetContentReplies api request get_content_replies
func (api *API) GetContentReplies(parentAuthor, parentPermlink string, opts ...interface{}) ([]*Content, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{parentAuthor, parentPermlink, opts[0]}
	} else {
		params = []interface{}{parentAuthor, parentPermlink}
	}
	raw, err := api.raw("get_content_replies", params)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_content_replies response", apiID)
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
	raw, err := api.raw("get_content", params)
	if err != nil {
		return nil, err
	}
	var resp Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_content response", apiID)
	}
	return &resp, nil
}

// api request get_languages
