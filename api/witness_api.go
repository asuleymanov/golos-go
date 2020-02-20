package api

import (
	"fmt"
)

//witness_api

//GetActiveWitnesses api request get_active_witnesses
func (api *API) GetActiveWitnesses() ([]*string, error) {
	var resp []*string
	err := api.call("witness_api", "get_active_witnesses", EmptyParams, &resp)
	return resp, err
}

//GetCurrentMedianHistoryPrice api request get_current_median_history_price
func (api *API) GetCurrentMedianHistoryPrice() (*Prices, error) {
	var resp Prices
	err := api.call("witness_api", "get_current_median_history_price", EmptyParams, &resp)
	return &resp, err
}

//GetFeedHistory api request get_feed_history
func (api *API) GetFeedHistory() (*FeedHistory, error) {
	var resp FeedHistory
	err := api.call("witness_api", "get_feed_history", EmptyParams, &resp)
	return &resp, err
}

//GetMinerQueue api request get_miner_queue
func (api *API) GetMinerQueue() ([]*string, error) {
	var resp []*string
	err := api.call("witness_api", "get_miner_queue", EmptyParams, &resp)
	return resp, err
}

//GetWitnessByAccount api request get_witness_by_account
func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	var resp Witness
	err := api.call("witness_api", "get_witness_by_account", []string{author}, &resp)
	return &resp, err
}

//GetWitnessCount api request get_witness_count
func (api *API) GetWitnessCount() (*uint64, error) {
	var resp uint64
	err := api.call("witness_api", "get_witness_count", EmptyParams, &resp)
	return &resp, err
}

//GetWitnessSchedule api request get_witness_schedule
func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	var resp WitnessSchedule
	err := api.call("witness_api", "get_witness_schedule", EmptyParams, &resp)
	return &resp, err
}

//GetWitnesses api request get_witnesses
func (api *API) GetWitnesses(id ...uint32) ([]*Witness, error) {
	var resp []*Witness
	err := api.call("witness_api", "get_witnesses", []interface{}{id}, &resp)
	return resp, err
}

//GetWitnessByVote api request get_witnesses_by_vote
func (api *API) GetWitnessByVote(author string, limit uint32) ([]*Witness, error) {
	if limit > 100 {
		return nil, fmt.Errorf("witness_api: get_witnesses_by_vote -> limit must not exceed 100")
	}
	var resp []*Witness
	err := api.call("witness_api", "get_witnesses_by_vote", []interface{}{author, limit}, &resp)
	return resp, err
}

//LookupWitnessAccounts api request lookup_witness_accounts
func (api *API) LookupWitnessAccounts(author string, limit uint32) ([]*string, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("witness_api: lookup_witness_accounts -> limit must not exceed 1000")
	}
	var resp []*string
	err := api.call("witness_api", "lookup_witness_accounts", []interface{}{author, limit}, &resp)
	return resp, err
}
