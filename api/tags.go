package api

import (
	"github.com/asuleymanov/golos-go/types"
)

//tags

//GetDiscussionsByActive api request get_discussions_by_active
func (api *API) GetDiscussionsByActive(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_active", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByAuthorBeforeDate api request get_discussions_by_author_before_date
func (api *API) GetDiscussionsByAuthorBeforeDate(author, permlink string, date types.Time, limit uint32, opts ...uint32) ([]*Discussion, error) {
	var params []interface{}
	if len(opts) > 0 {
		params = []interface{}{author, permlink, date, limit, opts}
	} else {
		params = []interface{}{author, permlink, date, limit}
	}
	var resp []*Discussion
	err := api.call("tags", "get_discussions_by_author_before_date", params, &resp)
	return resp, err
}

//GetDiscussionsByBlog api request get_discussions_by_blog
func (api *API) GetDiscussionsByBlog(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_blog", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByCashout api request get_discussions_by_cashout
func (api *API) GetDiscussionsByCashout(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_cashout", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByChildren api request get_discussions_by_children
func (api *API) GetDiscussionsByChildren(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_children", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByComments api request get_discussions_by_comments
func (api *API) GetDiscussionsByComments(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_comments", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByCreated api request get_discussions_by_created
func (api *API) GetDiscussionsByCreated(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_created", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByFeed api request get_discussions_by_feed
func (api *API) GetDiscussionsByFeed(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_feed", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByHot api request get_discussions_by_hot
func (api *API) GetDiscussionsByHot(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_hot", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByPayout api request get_discussions_by_payout
func (api *API) GetDiscussionsByPayout(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_payout", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByPromoted api request get_discussions_by_promoted
func (api *API) GetDiscussionsByPromoted(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_promoted", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByTrending api request get_discussions_by_trending
func (api *API) GetDiscussionsByTrending(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_trending", []interface{}{q}, &resp)
	return resp, err
}

//GetDiscussionsByVotes api request get_discussions_by_votes
func (api *API) GetDiscussionsByVotes(query DiscussionQuery) ([]*Discussion, error) {
	var resp []*Discussion
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}
	err = api.call("tags", "get_discussions_by_votes", []interface{}{q}, &resp)
	return resp, err
}

//GetLanguages api request get_languages
func (api *API) GetLanguages() ([]*string, error) {
	var resp []*string
	err := api.call("tags", "get_languages", EmptyParams, &resp)
	return resp, err
}

//GetTagsUsedByAuthor api request get_tags_used_by_author
func (api *API) GetTagsUsedByAuthor(accountName string) ([]*uint32, error) {
	var resp []*uint32
	err := api.call("tags", "get_tags_used_by_author", []interface{}{accountName}, &resp)
	return resp, err
}

//GetTrendingTags api request get_trending_tags
func (api *API) GetTrendingTags(afterTag string, limit uint32) ([]*TrendingTags, error) {
	var resp []*TrendingTags
	err := api.call("tags", "get_trending_tags", []interface{}{afterTag, limit}, &resp)
	return resp, err
}
