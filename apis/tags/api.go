package tags

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
)

const apiID = "tags"

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

//GetTrendingTags api request get_trending_tags
func (api *API) GetTrendingTags(afterTag string, limit uint32) ([]*TrendingTags, error) {
	raw, err := api.raw("get_trending_tags", []interface{}{afterTag, limit})
	if err != nil {
		return nil, err
	}
	var resp []*TrendingTags
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_trending_tags response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByTrending api request get_discussions_by_trending
func (api *API) GetDiscussionsByTrending(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_trending", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_trending response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByCreated api request get_discussions_by_created
func (api *API) GetDiscussionsByCreated(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_created", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_created response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByActive api request get_discussions_by_active
func (api *API) GetDiscussionsByActive(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_active", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_active response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByCashout api request get_discussions_by_cashout
func (api *API) GetDiscussionsByCashout(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_cashout", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_cashout response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByPayout api request get_discussions_by_payout
func (api *API) GetDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_payout", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_payout response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByVotes api request get_discussions_by_votes
func (api *API) GetDiscussionsByVotes(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_votes", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_votes response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByChildren api request get_discussions_by_children
func (api *API) GetDiscussionsByChildren(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_children", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_children response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByHot api request get_discussions_by_hot
func (api *API) GetDiscussionsByHot(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_hot", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_hot response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByFeed api request get_discussions_by_feed
func (api *API) GetDiscussionsByFeed(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_feed", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_feed response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByBlog api request get_discussions_by_blog
func (api *API) GetDiscussionsByBlog(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_blog", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_blog response", apiID)
	}
	return resp, nil
}

//GetDiscussionsByComments api request get_discussions_by_comments
func (api *API) GetDiscussionsByComments(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_comments", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_comments response", apiID)
	}
	return resp, nil
}

//GetTagsUsedByAuthor api request get_tags_used_by_author
func (api *API) GetTagsUsedByAuthor(accountName string) (*json.RawMessage, error) {
	return api.raw("get_tags_used_by_author", []interface{}{accountName})
}

//GetDiscussionsByPromoted api request get_discussions_by_promoted
func (api *API) GetDiscussionsByPromoted(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.raw("get_discussions_by_promoted", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_promoted response", apiID)
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
	raw, err := api.raw("get_discussions_by_author_before_date", params)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_discussions_by_author_before_date response", apiID)
	}
	return resp, nil
}

// api request get_languages
