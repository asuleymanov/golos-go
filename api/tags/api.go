package tags

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
)

const apiID = "tags"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = []struct{}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetDiscussionsByActive api request get_discussions_by_active
func (api *API) GetDiscussionsByActive(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_active", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByAuthorBeforeDate api request get_discussions_by_author_before_date
func (api *API) GetDiscussionsByAuthorBeforeDate(author, permlink, date string, limit uint32, opts ...interface{}) ([]*Content, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{author, permlink, date, limit, opts[0]}
	} else {
		params = []interface{}{author, permlink, date, limit}
	}
	var resp []*Content
	err := api.call("get_discussions_by_author_before_date", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByBlog api request get_discussions_by_blog
func (api *API) GetDiscussionsByBlog(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_blog", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByCashout api request get_discussions_by_cashout
func (api *API) GetDiscussionsByCashout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_cashout", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByChildren api request get_discussions_by_children
func (api *API) GetDiscussionsByChildren(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_children", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByComments api request get_discussions_by_comments
func (api *API) GetDiscussionsByComments(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_comments", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByCreated api request get_discussions_by_created
func (api *API) GetDiscussionsByCreated(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_created", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByFeed api request get_discussions_by_feed
func (api *API) GetDiscussionsByFeed(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_feed", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByHot api request get_discussions_by_hot
func (api *API) GetDiscussionsByHot(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_hot", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByPayout api request get_discussions_by_payout
func (api *API) GetDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_payout", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByPromoted api request get_discussions_by_promoted
func (api *API) GetDiscussionsByPromoted(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_promoted", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByTrending api request get_discussions_by_trending
func (api *API) GetDiscussionsByTrending(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_trending", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByVotes api request get_discussions_by_votes
func (api *API) GetDiscussionsByVotes(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_votes", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetLanguages api request get_languages
func (api *API) GetLanguages() ([]*string, error) {
	var resp []*string
	err := api.call("get_languages", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetTagsUsedByAuthor api request get_tags_used_by_author
func (api *API) GetTagsUsedByAuthor(accountName string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_tags_used_by_author", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetTrendingTags api request get_trending_tags
func (api *API) GetTrendingTags(afterTag string, limit uint32) ([]*TrendingTags, error) {
	var resp []*TrendingTags
	err := api.call("get_trending_tags", []interface{}{afterTag, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
