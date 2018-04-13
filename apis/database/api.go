package database

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
	"github.com/pkg/errors"
)

const apiID = "database_api"

type API struct {
	caller transports.Caller
}

func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = struct{}{}

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//GetActiveWitnesses api request get_active_witnesses
func (api *API) GetActiveWitnesses() ([]string, error) {
	raw, err := api.raw("get_active_witnesses", emptyParams)

	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_active_witnesses response", apiID)
	}
	return resp, nil
}

//GetMinerQueue api request get_miner_queue
func (api *API) GetMinerQueue() ([]string, error) {
	raw, err := api.raw("get_miner_queue", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_miner_queue response", apiID)
	}
	return resp, nil
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	raw, err := api.raw("get_block_header", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp BlockHeader
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_block_header response", apiID)
	}
	resp.Number = blockNum
	return &resp, nil
}

//GetBlock api request get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	raw, err := api.raw("get_block", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp Block
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_block response", apiID)
	}
	resp.Number = blockNum
	return &resp, nil
}

//GetOpsInBlock api request get_ops_in_block
func (api *API) GetOpsInBlock(blockNum uint32, onlyVirtual bool) ([]*types.OperationObject, error) {
	raw, err := api.raw("get_ops_in_block", []interface{}{blockNum, onlyVirtual})
	if err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_ops_in_block response", apiID)
	}
	return resp, nil
}

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	raw, err := api.raw("get_config", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp Config
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_config response", apiID)
	}
	return &resp, nil
}

//GetDynamicGlobalProperties api request get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	raw, err := api.raw("get_dynamic_global_properties", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp DynamicGlobalProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_dynamic_global_properties response", apiID)
	}
	return &resp, nil
}

//GetChainProperties api request get_chain_properties
func (api *API) GetChainProperties() (*ChainProperties, error) {
	raw, err := api.raw("get_chain_properties", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp ChainProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_chain_properties response", apiID)
	}
	return &resp, nil
}

//GetCurrentMedianHistoryPrice api request get_current_median_history_price
func (api *API) GetCurrentMedianHistoryPrice() (*CurrentMedianHistoryPrice, error) {
	raw, err := api.raw("get_current_median_history_price", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp CurrentMedianHistoryPrice
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_current_median_history_price response", apiID)
	}
	return &resp, nil
}

//GetFeedHistory api request get_feed_history
func (api *API) GetFeedHistory() (*FeedHistory, error) {
	raw, err := api.raw("get_feed_history", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp FeedHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_feed_history response", apiID)
	}
	return &resp, nil
}

//GetWitnessSchedule api request get_witness_schedule
func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	raw, err := api.raw("get_witness_schedule", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp WitnessSchedule
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witness_schedule response", apiID)
	}
	return &resp, nil
}

//GetHardforkVersion api request get_hardfork_version
func (api *API) GetHardforkVersion() (string, error) {
	raw, err := api.raw("get_hardfork_version", emptyParams)
	if err != nil {
		return "", err
	}
	var resp string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return "", errors.Wrapf(err, "golos: %v: failed to unmarshal get_hardfork_version response", apiID)
	}
	return resp, nil
}

//GetNextScheduledHardfork api request get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	raw, err := api.raw("get_next_scheduled_hardfork", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp NextScheduledHardfork
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_next_scheduled_hardfork response", apiID)
	}
	return &resp, nil
}

//GetAccounts api request get_accounts
func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	raw, err := api.raw("get_accounts", [][]string{accountNames})
	if err != nil {
		return nil, err
	}
	var resp []*Account
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_accounts response", apiID)
	}
	return resp, nil
}

//LookupAccountNames api request lookup_account_names
func (api *API) LookupAccountNames(accountNames []string) ([]*LookupAccountNames, error) {
	raw, err := api.raw("lookup_account_names", [][]string{accountNames})
	if err != nil {
		return nil, err
	}
	var resp []*LookupAccountNames
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal lookup_account_names response", apiID)
	}
	return resp, nil
}

//LookupAccounts api request lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]string, error) {
	raw, err := api.raw("lookup_accounts", []interface{}{lowerBoundName, limit})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal lookup_accounts response", apiID)
	}
	return resp, nil
}

//GetAccountCount api request get_account_count
func (api *API) GetAccountCount() (uint32, error) {
	raw, err := api.raw("get_account_count", emptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos: %v: failed to unmarshal get_account_count response", apiID)
	}
	return resp, nil
}

//GetOwnerHistory api request get_owner_history
func (api *API) GetOwnerHistory(accountName string) ([]*OwnerHistory, error) {
	raw, err := api.raw("get_owner_history", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*OwnerHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_owner_history response", apiID)
	}
	return resp, nil
}

//GetRecoveryRequest api request get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*json.RawMessage, error) {
	return api.raw("get_recovery_request", []interface{}{accountName})
}

//GetEscrow api request get_escrow
func (api *API) GetEscrow(from string, escrowID uint32) (*json.RawMessage, error) {
	return api.raw("get_escrow", []interface{}{from, escrowID})
}

//GetWithdrawRoutes api request get_withdraw_routes
func (api *API) GetWithdrawRoutes(accountName string, withdrawRouteType string) (*json.RawMessage, error) {
	return api.raw("get_withdraw_routes", []interface{}{accountName, withdrawRouteType})
}

//GetAccountBandwidth api request get_account_bandwidth
/*
bandwidthType:
post = 0
forum = 1
market = 2
old_forum = 3
old_market = 4
*/
func (api *API) GetAccountBandwidth(accountName string, bandwidthType uint32) (*Bandwidth, error) {
	raw, err := api.raw("get_account_bandwidth", []interface{}{accountName, bandwidthType})
	if err != nil {
		return nil, err
	}
	var resp *Bandwidth
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_account_bandwidth response", apiID)
	}
	return resp, nil
}

//GetSavingsWithdrawFrom api request get_savings_withdraw_from
func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.raw("get_savings_withdraw_from", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_savings_withdraw_from response", apiID)
	}
	return resp, nil
}

//GetSavingsWithdrawTo api request get_savings_withdraw_to
func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.raw("get_savings_withdraw_to", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_savings_withdraw_to response", apiID)
	}
	return resp, nil
}

//GetWitnesses api request get_witnesses
func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	raw, err := api.raw("get_witnesses", [][]uint32{id})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witnesses response", apiID)
	}
	return resp, nil
}

//GetConversionRequests api request get_conversion_requests
func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	raw, err := api.raw("get_conversion_requests", []string{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*ConversionRequests
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_conversion_requests response", apiID)
	}
	return resp, nil
}

//GetWitnessByAccount api request get_witness_by_account
func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	raw, err := api.raw("get_witness_by_account", []string{author})
	if err != nil {
		return nil, err
	}
	var resp Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witness_by_account response", apiID)
	}
	return &resp, nil
}

//GetWitnessByVote api request get_witnesses_by_vote
func (api *API) GetWitnessByVote(author string, limit uint) ([]*Witness, error) {
	if limit > 1000 {
		return nil, errors.New("GetWitnessByVote: limit must not exceed 1000")
	}
	raw, err := api.raw("get_witnesses_by_vote", []interface{}{author, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witnesses_by_vote response", apiID)
	}
	return resp, nil
}

//LookupWitnessAccounts api request lookup_witness_accounts
func (api *API) LookupWitnessAccounts(author string, limit uint) ([]string, error) {
	if limit > 1000 {
		return nil, errors.New("LookupWitnessAccounts: limit must not exceed 1000")
	}
	raw, err := api.raw("lookup_witness_accounts", []interface{}{author, limit})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal lookup_witness_accounts response", apiID)
	}
	return resp, nil
}

//GetWitnessCount api request get_witness_count
func (api *API) GetWitnessCount() (uint32, error) {
	raw, err := api.raw("get_witness_count", emptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witness_count response", apiID)
	}
	return resp, nil
}

//GetTransactionHex api request get_transaction_hex
func (api *API) GetTransactionHex(trx *types.Transaction) (*json.RawMessage, error) {
	return api.raw("get_transaction_hex", []interface{}{&trx})
}

//GetTransaction api request get_transaction
func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	raw, err := api.raw("get_transaction", []string{id})
	if err != nil {
		return nil, err
	}
	var resp types.Transaction
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_transaction response", apiID)
	}
	return &resp, nil
}

//get_required_signatures

//GetPotentialSignatures api request get_potential_signatures
func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]string, error) {
	raw, err := api.raw("get_potential_signatures", []interface{}{&trx})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_potential_signatures response", apiID)
	}
	return resp, nil
}

//GetVerifyAuthority api request verify_authority
func (api *API) GetVerifyAuthority(trx *types.Transaction) (bool, error) {
	raw, err := api.raw("verify_authority", []interface{}{&trx})
	if err != nil {
		return false, err
	}
	var resp bool
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return false, errors.Wrapf(err, "golos: %v: failed to unmarshal verify_authority response", apiID)
	}
	return resp, nil
}

//verify_account_authority

//GetAccountHistory api request get_account_history
func (api *API) GetAccountHistory(account string, from int64, limit uint32) ([]*types.OperationObject, error) {
	raw, err := api.raw("get_account_history", []interface{}{account, from, limit})
	if err != nil {
		return nil, err
	}
	var tmp1 [][]interface{}
	if err := json.Unmarshal([]byte(*raw), &tmp1); err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	for _, v := range tmp1 {
		byteData, errm := json.Marshal(&v[1])
		if errm != nil {
			return nil, errm
		}
		var tmp *types.OperationObject
		if err := json.Unmarshal(byteData, &tmp); err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}
