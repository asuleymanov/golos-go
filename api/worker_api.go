package api

import (
	"errors"
)

//worker_api

//GetWorkerRequests api request get_worker_requests
//sort:
//0 = by_created,
//1 = by_net_rshares,
//2 = by_upvotes,
//3 = by_downvotes
func (api *API) GetWorkerRequests(query WorkerRequestQuery, sort int8, fill_posts bool) ([]*WorkerRequests, error) {
	var resp []*WorkerRequests
	q, err := queryJson([]interface{}{query})
	if err != nil {
		return resp, err
	}

	err = api.call("worker_api", "get_worker_requests", []interface{}{q, sort, fill_posts}, &resp)
	return resp, err
}

//GetWorkerRequestVotes api request get_worker_request_votes
func (api *API) GetWorkerRequestVotes(author, permlink, start_voter string, limit uint32) ([]*WorkerVotes, error) {
	if limit > 50 {
		return nil, errors.New("worker_api: get_worker_request_votes -> limit must not exceed 50")
	}
	var resp []*WorkerVotes
	err := api.call("worker_api", "get_worker_request_votes", []interface{}{author, permlink, start_voter, limit}, &resp)
	return resp, err
}
