package api

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/types"
)

//database_api

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
	err := api.call("database_api", "get_account_bandwidth", []interface{}{accountName, bandwidthType}, &resp)
	return &resp, err
}

//GetAccountCount api request get_account_count
func (api *API) GetAccountCount() (*uint32, error) {
	var resp uint32
	err := api.call("database_api", "get_account_count", transports.EmptyParams, &resp)
	return &resp, err
}

//GetAccounts api request get_accounts
func (api *API) GetAccounts(accountNames ...string) ([]*Account, error) {
	var resp []*Account
	err := api.call("database_api", "get_accounts", []interface{}{accountNames}, &resp)
	return resp, err
}

//GetBlock api request get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	err := api.call("database_api", "get_block", []uint32{blockNum}, &resp)
	resp.Number = blockNum
	return &resp, err
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	var resp BlockHeader
	err := api.call("database_api", "get_block_header", []uint32{blockNum}, &resp)
	resp.Number = blockNum
	return &resp, err
}

//GetChainProperties api request get_chain_properties
func (api *API) GetChainProperties() (*types.ChainProperties, error) {
	var resp types.ChainProperties
	err := api.call("database_api", "get_chain_properties", transports.EmptyParams, &resp)
	return &resp, err
}

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	var resp Config
	err := api.call("database_api", "get_config", transports.EmptyParams, &resp)
	return &resp, err
}

//GetConversionRequests api request get_conversion_requests
func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	var resp []*ConversionRequests
	err := api.call("database_api", "get_conversion_requests", []string{accountName}, &resp)
	return resp, err
}

//GetDatabaseInfo api request get_database_info
func (api *API) GetDatabaseInfo() (*DatabaseInfo, error) {
	var resp DatabaseInfo
	err := api.call("database_api", "get_database_info", transports.EmptyParams, &resp)
	return &resp, err
}

//GetDynamicGlobalProperties api request get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := api.call("database_api", "get_dynamic_global_properties", transports.EmptyParams, &resp)
	return &resp, err
}

//GetEscrow api request get_escrow
func (api *API) GetEscrow(from string, escrowID uint32) (*Escrow, error) {
	var resp Escrow
	err := api.call("database_api", "get_escrow", []interface{}{from, escrowID}, &resp)
	return &resp, err
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
	err := api.call("database_api", "get_expiring_vesting_delegations", params, &resp)
	return resp, err
}

//GetHardforkVersion api request get_hardfork_version
func (api *API) GetHardforkVersion() (*string, error) {
	var resp string
	err := api.call("database_api", "get_hardfork_version", transports.EmptyParams, &resp)
	return &resp, err
}

//GetNextScheduledHardfork api request get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	var resp NextScheduledHardfork
	err := api.call("database_api", "get_next_scheduled_hardfork", transports.EmptyParams, &resp)
	return &resp, err
}

//GetOwnerHistory api request get_owner_history
func (api *API) GetOwnerHistory(accountName string) ([]*OwnerHistory, error) {
	var resp []*OwnerHistory
	err := api.call("database_api", "get_owner_history", []interface{}{accountName}, &resp)
	return resp, err
}

//GetPotentialSignatures api request get_potential_signatures
func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]*string, error) {
	var resp []*string
	err := api.call("database_api", "get_potential_signatures", []interface{}{&trx}, &resp)
	return resp, err
}

//GetProposedTransaction api request get_proposed_transactions
func (api *API) GetProposedTransaction(account string, from, limit uint32) ([]*ProposalObject, error) {
	var resp []*ProposalObject
	err := api.call("database_api", "get_proposed_transactions", []interface{}{account, from, limit}, &resp)
	return resp, err
}

//GetRecoveryRequest api request get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*AccountRecoveryRequest, error) {
	var resp AccountRecoveryRequest
	err := api.call("database_api", "get_recovery_request", []interface{}{accountName}, &resp)
	return &resp, err
}

//GetRequiredSignatures api request get_required_signatures
func (api *API) GetRequiredSignatures(trx *types.Transaction, keys ...string) ([]*string, error) {
	var resp []*string
	err := api.call("database_api", "get_required_signatures", []interface{}{trx, keys}, &resp)
	return resp, err
}

//GetSavingsWithdrawFrom api request get_savings_withdraw_from
func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	err := api.call("database_api", "get_savings_withdraw_from", []interface{}{accountName}, &resp)
	return resp, err
}

//GetSavingsWithdrawTo api request get_savings_withdraw_to
func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	err := api.call("database_api", "get_savings_withdraw_to", []interface{}{accountName}, &resp)
	return resp, err
}

//GetTransactionHex api request get_transaction_hex
func (api *API) GetTransactionHex(trx *types.Transaction) (*string, error) {
	var resp string
	err := api.call("database_api", "get_transaction_hex", []interface{}{&trx}, &resp)
	return &resp, err
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
	err := api.call("database_api", "get_vesting_delegations", params, &resp)
	return resp, err
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
	err := api.call("database_api", "get_withdraw_routes", []interface{}{accountName, withdrawRouteType}, &resp)
	return resp, err
}

//LookupAccountNames api request lookup_account_names
func (api *API) LookupAccountNames(accountNames ...string) ([]*Account, error) {
	var resp []*Account
	err := api.call("database_api", "lookup_account_names", []interface{}{accountNames}, &resp)
	return resp, err
}

//LookupAccounts api request lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]*string, error) {
	var resp []*string
	err := api.call("database_api", "lookup_accounts", []interface{}{lowerBoundName, limit}, &resp)
	return resp, err
}

//GetVerifyAccountAuthority api request verify_account_authority
func (api *API) GetVerifyAccountAuthority(accountName string, keys ...string) (*bool, error) {
	var resp bool
	err := api.call("database_api", "verify_account_authority", []interface{}{accountName, keys}, &resp)
	return &resp, err
}

//GetVerifyAuthority api request verify_authority
func (api *API) GetVerifyAuthority(trx *types.Transaction) (*bool, error) {
	var resp bool
	err := api.call("database_api", "verify_authority", []interface{}{&trx}, &resp)
	return &resp, err
}

// Set callback to invoke as soon as a new block is applied
func (api *API) SetBlockAppliedCallback(notice func(header *BlockHeader, error error)) (err error) {
	err = api.setCallback("database_api", "set_block_applied_callback", func(raw json.RawMessage) {
		var header []BlockHeader
		if err := json.Unmarshal(raw, &header); err != nil {
			notice(nil, err)
		}
		for _, b := range header {
			notice(&b, nil)
		}
	})
	return
}
