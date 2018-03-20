package database

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
	"github.com/pkg/errors"
)

const APIID = "database_api"

type API struct {
	caller transports.Caller
}

func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var EmptyParams = []string{}

func (api *API) Raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{APIID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to call %v\n", APIID, method)
	}
	return &resp, nil
}

//get_active_witnesses
func (api *API) GetActiveWitnesses() ([]string, error) {
	raw, err := api.Raw("get_active_witnesses", EmptyParams)

	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_active_witnesses response", APIID)
	}
	return resp, nil
}

//get_miner_queue
func (api *API) GetMinerQueue() ([]string, error) {
	raw, err := api.Raw("get_miner_queue", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_miner_queue response", APIID)
	}
	return resp, nil
}

//get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	raw, err := api.Raw("get_block_header", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp BlockHeader
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_block_header response", APIID)
	}
	resp.Number = blockNum
	return &resp, nil
}

//get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	raw, err := api.Raw("get_block", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp Block
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_block response", APIID)
	}
	resp.Number = blockNum
	return &resp, nil
}

//get_ops_in_block
func (api *API) GetOpsInBlock(blockNum uint32, only_virtual bool) ([]*types.OperationObject, error) {
	raw, err := api.Raw("get_ops_in_block", []interface{}{blockNum, only_virtual})
	if err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_ops_in_block response", APIID)
	}
	return resp, nil
}

//get_config
func (api *API) GetConfig() (*Config, error) {
	raw, err := api.Raw("get_config", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp Config
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_config response", APIID)
	}
	return &resp, nil
}

//get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	raw, err := api.Raw("get_dynamic_global_properties", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp DynamicGlobalProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_dynamic_global_properties response", APIID)
	}
	return &resp, nil
}

//get_chain_properties
func (api *API) GetChainProperties() (*ChainProperties, error) {
	raw, err := api.Raw("get_chain_properties", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp ChainProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_chain_properties response", APIID)
	}
	return &resp, nil
}

//get_current_median_history_price
func (api *API) GetCurrentMedianHistoryPrice() (*CurrentMedianHistoryPrice, error) {
	raw, err := api.Raw("get_current_median_history_price", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp CurrentMedianHistoryPrice
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_current_median_history_price response", APIID)
	}
	return &resp, nil
}

//get_feed_history
func (api *API) GetFeedHistory() (*FeedHistory, error) {
	raw, err := api.Raw("get_feed_history", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp FeedHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_feed_history response", APIID)
	}
	return &resp, nil
}

//get_witness_schedule
func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	raw, err := api.Raw("get_witness_schedule", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp WitnessSchedule
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witness_schedule response", APIID)
	}
	return &resp, nil
}

//get_hardfork_version
func (api *API) GetHardforkVersion() (string, error) {
	raw, err := api.Raw("get_hardfork_version", EmptyParams)
	if err != nil {
		return "", err
	}
	var resp string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return "", errors.Wrapf(err, "golos: %v: failed to unmarshal get_hardfork_version response", APIID)
	}
	return resp, nil
}

//get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	raw, err := api.Raw("get_next_scheduled_hardfork", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp NextScheduledHardfork
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_next_scheduled_hardfork response", APIID)
	}
	return &resp, nil
}

//get_accounts
func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	raw, err := api.Raw("get_accounts", [][]string{accountNames})
	if err != nil {
		return nil, err
	}
	var resp []*Account
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_accounts response", APIID)
	}
	return resp, nil
}

//lookup_account_names
func (api *API) LookupAccountNames(accountNames []string) ([]*LookupAccountNames, error) {
	raw, err := api.Raw("lookup_account_names", [][]string{accountNames})
	if err != nil {
		return nil, err
	}
	var resp []*LookupAccountNames
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal lookup_account_names response", APIID)
	}
	return resp, nil
}

//lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]string, error) {
	raw, err := api.Raw("lookup_accounts", []interface{}{lowerBoundName, limit})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal lookup_accounts response", APIID)
	}
	return resp, nil
}

//get_account_count
func (api *API) GetAccountCount() (uint32, error) {
	raw, err := api.Raw("get_account_count", EmptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos: %v: failed to unmarshal get_account_count response", APIID)
	}
	return resp, nil
}

//get_owner_history
func (api *API) GetOwnerHistory(accountName string) ([]*OwnerHistory, error) {
	raw, err := api.Raw("get_owner_history", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*OwnerHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_owner_history response", APIID)
	}
	return resp, nil
}

//get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_recovery_request", []interface{}{accountName})
}

//get_escrow
func (api *API) GetEscrow(from string, escrow_id uint32) (*json.RawMessage, error) {
	return api.Raw("get_escrow", []interface{}{from, escrow_id})
}

//get_withdraw_routes
func (api *API) GetWithdrawRoutes(accountName string, withdraw_route_type string) (*json.RawMessage, error) {
	return api.Raw("get_withdraw_routes", []interface{}{accountName, withdraw_route_type})
}

//get_account_bandwidth
/*
bandwidth_type:
post = 0
forum = 1
market = 2
old_forum = 3
old_market = 4
*/
func (api *API) GetAccountBandwidth(accountName string, bandwidth_type uint32) (*Bandwidth, error) {
	raw, err := api.Raw("get_account_bandwidth", []interface{}{accountName, bandwidth_type})
	if err != nil {
		return nil, err
	}
	var resp *Bandwidth
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_account_bandwidth response", APIID)
	}
	return resp, nil
}

//get_savings_withdraw_from
func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.Raw("get_savings_withdraw_from", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_savings_withdraw_from response", APIID)
	}
	return resp, nil
}

//get_savings_withdraw_to
func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.Raw("get_savings_withdraw_to", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_savings_withdraw_to response", APIID)
	}
	return resp, nil
}

//get_witnesses
func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	raw, err := api.Raw("get_witnesses", [][]uint32{id})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witnesses response", APIID)
	}
	return resp, nil
}

//get_conversion_requests
func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	raw, err := api.Raw("get_conversion_requests", []string{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*ConversionRequests
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_conversion_requests response", APIID)
	}
	return resp, nil
}

//get_witness_by_account
func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	raw, err := api.Raw("get_witness_by_account", []string{author})
	if err != nil {
		return nil, err
	}
	var resp Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witness_by_account response", APIID)
	}
	return &resp, nil
}

//get_witnesses_by_vote
func (api *API) GetWitnessByVote(author string, limit uint) ([]*Witness, error) {
	if limit > 1000 {
		return nil, errors.New("GetWitnessByVote: limit must not exceed 1000")
	}
	raw, err := api.Raw("get_witnesses_by_vote", []interface{}{author, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witnesses_by_vote response", APIID)
	}
	return resp, nil
}

//lookup_witness_accounts
func (api *API) LookupWitnessAccounts(author string, limit uint) ([]string, error) {
	if limit > 1000 {
		return nil, errors.New("LookupWitnessAccounts: limit must not exceed 1000")
	}
	raw, err := api.Raw("lookup_witness_accounts", []interface{}{author, limit})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal lookup_witness_accounts response", APIID)
	}
	return resp, nil
}

//get_witness_count
func (api *API) GetWitnessCount() (uint32, error) {
	raw, err := api.Raw("get_witness_count", EmptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos: %v: failed to unmarshal get_witness_count response", APIID)
	}
	return resp, nil
}

//get_transaction_hex
func (api *API) GetTransactionHex(trx *types.Transaction) (*json.RawMessage, error) {
	return api.Raw("get_transaction_hex", []interface{}{&trx})
}

//get_transaction
func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	raw, err := api.Raw("get_transaction", []string{id})
	if err != nil {
		return nil, err
	}
	var resp types.Transaction
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_transaction response", APIID)
	}
	return &resp, nil
}

//get_required_signatures
//get_potential_signatures
func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]string, error) {
	raw, err := api.Raw("get_potential_signatures", []interface{}{&trx})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos: %v: failed to unmarshal get_potential_signatures response", APIID)
	}
	return resp, nil
}

//verify_authority
func (api *API) GetVerifyAuthority(trx *types.Transaction) (bool, error) {
	raw, err := api.Raw("verify_authority", []interface{}{&trx})
	if err != nil {
		return false, err
	}
	var resp bool
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return false, errors.Wrapf(err, "golos: %v: failed to unmarshal verify_authority response", APIID)
	}
	return resp, nil
}

//verify_account_authority
//get_account_history
func (api *API) GetAccountHistory(account string, from uint64, limit uint32) ([]*types.OperationObject, error) {
	raw, err := api.Raw("get_account_history", []interface{}{account, from, limit})
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
