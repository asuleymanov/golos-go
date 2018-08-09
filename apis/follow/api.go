package follow

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
)

const apiID = "follow"

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

//GetAccountReputations api request get_account_reputations
func (api *API) GetAccountReputations(accounts []string) ([]uint32, error) {
	raw, err := api.raw("get_account_reputations", []interface{}{accounts})
	if err != nil {
		return nil, err
	}
	var resp []uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_account_reputations response", apiID)
	}
	return resp, nil
}

//GetBlog api request get_blog
func (api *API) GetBlog(accountName string, entryID uint32, limit uint16) ([]*Blogs, error) {
	if limit > 500 {
		return nil, errors.Errorf("%v: get_blog -> limit must not exceed 500", apiID)
	}
	raw, err := api.raw("get_blog", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Blogs
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_feed response", apiID)
	}
	return resp, nil
}

//GetBlogAuthors api request get_blog_authors
func (api *API) GetBlogAuthors(author string) (*BlogAuthors, error) {
	raw, err := api.raw("get_blog_authors", []interface{}{author})
	if err != nil {
		return nil, err
	}
	var resp BlogAuthors
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "market_history_api: failed to unmarshal get_blog_authors response", apiID)
	}
	return &resp, nil
}

//GetBlogEntries api request get_blog_entries
func (api *API) GetBlogEntries(accountName string, entryID uint32, limit uint16) ([]*BlogEntries, error) {
	if limit > 500 {
		return nil, errors.Errorf("%v: get_blog_entries -> limit must not exceed 500", apiID)
	}
	raw, err := api.raw("get_blog_entries", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*BlogEntries
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_feed_entries response", apiID)
	}
	return resp, nil
}

//GetFeed api request get_feed
func (api *API) GetFeed(accountName string, entryID uint32, limit uint16) ([]*Feeds, error) {
	if limit > 500 {
		return nil, errors.Errorf("%v: get_feed -> limit must not exceed 500", apiID)
	}
	raw, err := api.raw("get_feed", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Feeds
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_feed response", apiID)
	}
	return resp, nil
}

//GetFeedEntries api request get_feed_entries
func (api *API) GetFeedEntries(accountName string, entryID uint32, limit uint16) ([]*FeedEntry, error) {
	if limit > 500 {
		return nil, errors.Errorf("%v: get_feed_entries -> limit must not exceed 500", apiID)
	}
	raw, err := api.raw("get_feed_entries", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*FeedEntry
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_feed_entries response", apiID)
	}
	return resp, nil
}

//GetFollowCount api request get_follow_count
func (api *API) GetFollowCount(accountName string) (*FollowCount, error) {
	raw, err := api.raw("get_follow_count", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp *FollowCount
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_follow_count response", apiID)
	}
	return resp, nil
}

//GetFollowers api request get_followers
/*
kind:
undefined
blog
ignore
*/
func (api *API) GetFollowers(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
if limit > 1000 {
		return nil, errors.Errorf("%v: get_followers -> limit must not exceed 1000", apiID)
	}
	if kind !="undefined" || kind !="blog" || kind !="ignore" {
		return nil, errors.Errorf("%v: get_followers -> kind can take values only \"undefined\", \"blog\" and \"ignore\".", apiID)
	}
	raw, err := api.raw("get_followers", []interface{}{accountName, start, kind, limit})
	if err != nil {
		return nil, err
	}
	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_followers response", apiID)
	}
	return resp, nil
}

//GetFollowing api request get_following
/*
kind:
undefined
blog
ignore
*/
func (api *API) GetFollowing(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
if limit > 1000 {
		return nil, errors.Errorf("%v: get_following -> limit must not exceed 1000", apiID)
	}
	if kind !="undefined" || kind !="blog" || kind !="ignore" {
		return nil, errors.Errorf("%v: get_following -> kind can take values only \"undefined\", \"blog\" and \"ignore\".", apiID)
	}
	raw, err := api.raw("get_following", []interface{}{accountName, start, kind, limit})
	if err != nil {
		return nil, err
	}
	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_following response", apiID)
	}
	return resp, nil
}

//GetRebloggedBy api request get_reblogged_by
func (api *API) GetRebloggedBy(author, permlink string) ([]string, error) {
	raw, err := api.raw("get_reblogged_by", []interface{}{author, permlink})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "market_history_api: failed to unmarshal get_reblogged_by response", apiID)
	}
	return resp, nil
}
