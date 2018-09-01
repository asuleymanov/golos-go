package database

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
	"github.com/pkg/errors"
)

const apiID = "database_api"

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
	var resp Bandwidth
	err := api.call("get_account_bandwidth", []interface{}{accountName, bandwidthType}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetAccountCount api request get_account_count
func (api *API) GetAccountCount() (*uint32, error) {
	var resp uint32
	err := api.call("get_account_count", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetAccounts api request get_accounts
func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	var resp []*Account
	err := api.call("get_accounts", [][]string{accountNames}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetBlock api request get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	err := api.call("get_block", []uint32{blockNum}, &resp)
	if err != nil {
		return nil, err
	}
	resp.Number = blockNum
	return &resp, nil
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	var resp BlockHeader
	err := api.call("get_block_header", []uint32{blockNum}, &resp)
	if err != nil {
		return nil, err
	}
	resp.Number = blockNum
	return &resp, nil
}

//GetChainProperties api request get_chain_properties
func (api *API) GetChainProperties() (*ChainProperties, error) {
	var resp ChainProperties
	err := api.call("get_chain_properties", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	var resp Config
	err := api.call("get_config", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetConversionRequests api request get_conversion_requests
func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	var resp []*ConversionRequests
	err := api.call("get_conversion_requests", []string{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDatabaseInfo api request get_database_info
func (api *API) GetDatabaseInfo() (*DatabaseInfo, error) {
	var resp DatabaseInfo
	err := api.call("get_database_info", []interface{}{}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetDynamicGlobalProperties api request get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := api.call("get_dynamic_global_properties", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetEscrow api request get_escrow
func (api *API) GetEscrow(from string, escrowID uint32) (*Escrow, error) {
	var resp Escrow
	err := api.call("get_escrow", []interface{}{from, escrowID}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetExpiringVestingDelegations api request get_expiring_vesting_delegations
func (api *API) GetExpiringVestingDelegations(account string, from types.Time, opts ...interface{}) ([]*VestingDelegationExpiration, error) {
	params := []interface{}{account, from}
	switch len(opts) {
	case 0:
		params = append(params, 100)
	default:
		params = append(params, opts...)
	}
	var resp []*VestingDelegationExpiration
	err := api.call("get_expiring_vesting_delegations", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetHardforkVersion api request get_hardfork_version
func (api *API) GetHardforkVersion() (*string, error) {
	var resp string
	err := api.call("get_hardfork_version", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetNextScheduledHardfork api request get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	var resp NextScheduledHardfork
	err := api.call("get_next_scheduled_hardfork", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetOwnerHistory api request get_owner_history
func (api *API) GetOwnerHistory(accountName string) ([]*OwnerHistory, error) {
	var resp []*OwnerHistory
	err := api.call("get_owner_history", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetPotentialSignatures api request get_potential_signatures
func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]*string, error) {
	var resp []*string
	err := api.call("get_potential_signatures", []interface{}{&trx}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetProposedTransaction api request get_proposed_transactions
func (api *API) GetProposedTransaction(account string, from, limit uint32) ([]*ProposalObject, error) {
	var resp []*ProposalObject
	err := api.call("get_proposed_transactions", []interface{}{account, from, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetRecoveryRequest api request get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*AccountRecoveryRequest, error) {
	var resp AccountRecoveryRequest
	err := api.call("get_recovery_request", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetRequiredSignatures api request get_required_signatures
func (api *API) GetRequiredSignatures(trx *types.Transaction, keys []string) ([]*string, error) {
	var resp []*string
	err := api.call("get_required_signatures", []interface{}{trx, keys}, &resp)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}
	return resp, nil
}

//GetSavingsWithdrawFrom api request get_savings_withdraw_from
func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	err := api.call("get_savings_withdraw_from", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetSavingsWithdrawTo api request get_savings_withdraw_to
func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	err := api.call("get_savings_withdraw_to", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetTransactionHex api request get_transaction_hex
func (api *API) GetTransactionHex(trx *types.Transaction) (*string, error) {
	var resp string
	err := api.call("get_transaction_hex", []interface{}{&trx}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetVestingDelegations api request get_vesting_delegations
func (api *API) GetVestingDelegations(account, from string, opts ...interface{}) ([]*VestingDelegation, error) {
	params := []interface{}{account, from}
	switch len(opts) {
	case 0:
		params = append(params, 100, "delegated")
	case 1:
		params = append(params, opts[0], "delegated")
	default:
		params = append(params, opts...)
	}
	var resp []*VestingDelegation
	err := api.call("get_vesting_delegations", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetWithdrawRoutes api request get_withdraw_routes
/*
withdrawRouteType:
incoming
outgoing
all
*/
func (api *API) GetWithdrawRoutes(accountName, withdrawRouteType string) ([]*WithdrawVestingRoutes, error) {
	var resp []*WithdrawVestingRoutes
	err := api.call("get_withdraw_routes", []interface{}{accountName, withdrawRouteType}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//LookupAccountNames api request lookup_account_names
func (api *API) LookupAccountNames(accountNames []string) ([]*Account, error) {
	var resp []*Account
	err := api.call("lookup_account_names", [][]string{accountNames}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//LookupAccounts api request lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]*string, error) {
	var resp []*string
	err := api.call("lookup_accounts", []interface{}{lowerBoundName, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetVerifyAccountAuthority api request verify_account_authority
func (api *API) GetVerifyAccountAuthority(accountName string, keys []string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("verify_account_authority", []interface{}{accountName, keys}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetVerifyAuthority api request verify_authority
func (api *API) GetVerifyAuthority(trx *types.Transaction) (*bool, error) {
	var resp bool
	err := api.call("verify_authority", []interface{}{&trx}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
