package api

import (
	"errors"
)

//follow

//GetAccountReputations api request get_account_reputations
func (api *API) GetAccountReputations(accounts ...string) ([]*AccountReputation, error) {
	var resp []*AccountReputation
	err := api.call("follow", "get_account_reputations", []interface{}{accounts}, &resp)
	return resp, err
}

//GetBlog api request get_blog
func (api *API) GetBlog(accountName string, entryID, limit uint32) ([]*Blogs, error) {
	if limit > 500 {
		return nil, errors.New("follow: get_blog -> limit must not exceed 500")
	}
	var resp []*Blogs
	err := api.call("follow", "get_blog", []interface{}{accountName, entryID, limit}, &resp)
	return resp, err
}

//GetBlogAuthors api request get_blog_authors
func (api *API) GetBlogAuthors(author string) ([]*BlogAuthor, error) {
	var resp []*BlogAuthor
	err := api.call("follow", "get_blog_authors", []interface{}{author}, &resp)
	return resp, err
}

//GetBlogEntries api request get_blog_entries
func (api *API) GetBlogEntries(accountName string, entryID, limit uint32) ([]*BlogEntries, error) {
	if limit > 500 {
		return nil, errors.New("follow: get_blog_entries -> limit must not exceed 500")
	}
	var resp []*BlogEntries
	err := api.call("follow", "get_blog_entries", []interface{}{accountName, entryID, limit}, &resp)
	return resp, err
}

//GetFeed api request get_feed
func (api *API) GetFeed(accountName string, entryID, limit uint32) ([]*Feeds, error) {
	if limit > 500 {
		return nil, errors.New("follow: get_feed -> limit must not exceed 500")
	}
	var resp []*Feeds
	err := api.call("follow", "get_feed", []interface{}{accountName, entryID, limit}, &resp)
	return resp, err
}

//GetFeedEntries api request get_feed_entries
func (api *API) GetFeedEntries(accountName string, entryID, limit uint32) ([]*FeedEntry, error) {
	if limit > 500 {
		return nil, errors.New("follow: get_feed_entries -> limit must not exceed 500")
	}
	var resp []*FeedEntry
	err := api.call("follow", "get_feed_entries", []interface{}{accountName, entryID, limit}, &resp)
	return resp, err
}

//GetFollowCount api request get_follow_count
func (api *API) GetFollowCount(accountName string) (*FollowCount, error) {
	var resp FollowCount
	err := api.call("follow", "get_follow_count", []interface{}{accountName}, &resp)
	return &resp, err
}

//GetFollowers api request get_followers
/*
kind:
0 = undefined
1 = blog
2 = ignore
*/
func (api *API) GetFollowers(accountName, start string, kind, limit uint32) ([]*FollowObject, error) {
	if limit > 1000 {
		return nil, errors.New("follow: get_followers -> limit must not exceed 1000")
	}

	var resp []*FollowObject
	err := api.call("follow", "get_followers", []interface{}{accountName, start, kind, limit}, &resp)
	return resp, err
}

//GetFollowing api request get_following
/*
kind:
0 = undefined
1 = blog
2 = ignore
*/
func (api *API) GetFollowing(accountName, start string, kind, limit uint32) ([]*FollowObject, error) {
	if limit > 1000 {
		return nil, errors.New("follow: get_following -> limit must not exceed 1000")
	}

	var resp []*FollowObject
	err := api.call("follow", "get_following", []interface{}{accountName, start, kind, limit}, &resp)
	return resp, err
}

//GetRebloggedBy api request get_reblogged_by
func (api *API) GetRebloggedBy(author, permlink string) ([]*string, error) {
	var resp []*string
	err := api.call("follow", "get_reblogged_by", []interface{}{author, permlink}, &resp)
	return resp, err
}
