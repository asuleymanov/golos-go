package witness

import (
	"fmt"

	"github.com/asuleymanov/golos-go/transports"
)

const apiID = "witness_api"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = []string{}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetActiveWitnesses api request get_active_witnesses
func (api *API) GetActiveWitnesses() ([]*string, error) {
	var resp []*string
	err := api.call("get_active_witnesses", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetCurrentMedianHistoryPrice api request get_current_median_history_price
func (api *API) GetCurrentMedianHistoryPrice() (*CurrentMedianHistoryPrice, error) {
	var resp CurrentMedianHistoryPrice
	err := api.call("get_current_median_history_price", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetFeedHistory api request get_feed_history
func (api *API) GetFeedHistory() (*FeedHistory, error) {
	var resp FeedHistory
	err := api.call("get_feed_history", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetMinerQueue api request get_miner_queue
func (api *API) GetMinerQueue() ([]*string, error) {
	var resp []*string
	err := api.call("get_miner_queue", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetWitnessByAccount api request get_witness_by_account
func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	var resp Witness
	err := api.call("get_witness_by_account", []string{author}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetWitnessCount api request get_witness_count
func (api *API) GetWitnessCount() (*uint32, error) {
	var resp uint32
	err := api.call("get_witness_count", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetWitnessSchedule api request get_witness_schedule
func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	var resp WitnessSchedule
	err := api.call("get_witness_schedule", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetWitnesses api request get_witnesses
func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	var resp []*Witness
	err := api.call("get_witnesses", [][]uint32{id}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetWitnessByVote api request get_witnesses_by_vote
func (api *API) GetWitnessByVote(author string, limit uint) ([]*Witness, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_witnesses_by_vote -> limit must not exceed 1000", apiID)
	}
	var resp []*Witness
	err := api.call("get_witnesses_by_vote", []interface{}{author, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//LookupWitnessAccounts api request lookup_witness_accounts
func (api *API) LookupWitnessAccounts(author string, limit uint) ([]*string, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: lookup_witness_accounts -> limit must not exceed 1000", apiID)
	}
	var resp []*string
	err := api.call("lookup_witness_accounts", []interface{}{author, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
