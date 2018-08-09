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

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
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
	var resp Bandwidth
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_account_bandwidth response", apiID)
	}
	return &resp, nil
}

//GetAccountCount api request get_account_count
func (api *API) GetAccountCount() (*uint32, error) {
	raw, err := api.raw("get_account_count", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_account_count response", apiID)
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
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_accounts response", apiID)
	}
	return resp, nil
}

//GetBlock api request get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	raw, err := api.raw("get_block", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp Block
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_block response", apiID)
	}
	resp.Number = blockNum
	return &resp, nil
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	raw, err := api.raw("get_block_header", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp BlockHeader
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_block_header response", apiID)
	}
	resp.Number = blockNum
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
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_chain_properties response", apiID)
	}
	return &resp, nil
}

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	raw, err := api.raw("get_config", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp Config
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_config response", apiID)
	}
	return &resp, nil
}

//GetConversionRequests api request get_conversion_requests
func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	raw, err := api.raw("get_conversion_requests", []string{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*ConversionRequests
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_conversion_requests response", apiID)
	}
	return resp, nil
}

//GetDatabaseInfo api request get_database_info
func (api *API) GetDatabaseInfo() (*DatabaseInfo, error) {
	raw, err := api.raw("get_database_info", []interface{}{})
	if err != nil {
		return nil, err
	}
	var resp DatabaseInfo
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_database_info response", apiID)
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
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_dynamic_global_properties response", apiID)
	}
	return &resp, nil
}

//GetEscrow api request get_escrow
func (api *API) GetEscrow(from string, escrowID uint32) (*Escrow, error) {
	raw, err := api.raw("get_escrow", []interface{}{from, escrowID})
	if err != nil {
		return nil, err
	}
	var resp Escrow
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_escrow response", apiID)
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
	raw, err := api.raw("get_expiring_vesting_delegations", params)
	if err != nil {
		return nil, err
	}
	var resp []*VestingDelegationExpiration
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_expiring_vesting_delegations response", apiID)
	}
	return resp, nil
}

//GetHardforkVersion api request get_hardfork_version
func (api *API) GetHardforkVersion() (*string, error) {
	raw, err := api.raw("get_hardfork_version", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_hardfork_version response", apiID)
	}
	return &resp, nil
}

//GetNextScheduledHardfork api request get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	raw, err := api.raw("get_next_scheduled_hardfork", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp NextScheduledHardfork
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_next_scheduled_hardfork response", apiID)
	}
	return &resp, nil
}

//GetOwnerHistory api request get_owner_history
func (api *API) GetOwnerHistory(accountName string) ([]*OwnerHistory, error) {
	raw, err := api.raw("get_owner_history", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*OwnerHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_owner_history response", apiID)
	}
	return resp, nil
}

//GetPotentialSignatures api request get_potential_signatures
func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]*string, error) {
	raw, err := api.raw("get_potential_signatures", []interface{}{&trx})
	if err != nil {
		return nil, err
	}
	var resp []*string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_potential_signatures response", apiID)
	}
	return resp, nil
}

//GetProposedTransaction api request get_proposed_transactions
func (api *API) GetProposedTransaction(account string, from, limit uint32) ([]*ProposalObject, error) {
	raw, err := api.raw("get_proposed_transactions", []interface{}{account, from, limit})
	if err != nil {
		return nil, err
	}
	var resp []*ProposalObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_proposed_transactions response", apiID)
	}
	return resp, nil
}

//GetRecoveryRequest api request get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*AccountRecoveryRequest, error) {
	raw, err := api.raw("get_recovery_request", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp AccountRecoveryRequest
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_recovery_request response", apiID)
	}
	return &resp, nil
}

//GetRequiredSignatures api request get_required_signatures
func (api *API) GetRequiredSignatures(trx *types.Transaction, keys []string) ([]*string, error) {
	raw, err := api.raw("get_required_signatures", []interface{}{trx, keys})
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}
	var resp []*string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_required_signatures response", apiID)
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
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_savings_withdraw_from response", apiID)
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
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_savings_withdraw_to response", apiID)
	}
	return resp, nil
}

//GetTransactionHex api request get_transaction_hex
func (api *API) GetTransactionHex(trx *types.Transaction) (*string, error) {
	raw, err := api.raw("get_transaction_hex", []interface{}{&trx})
	if err != nil {
		return nil, err
	}
	var resp string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_transaction_hex response", apiID)
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
	raw, err := api.raw("get_vesting_delegations", params)
	if err != nil {
		return nil, err
	}
	var resp []*VestingDelegation
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal verify_authority response", apiID)
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
	raw, err := api.raw("get_withdraw_routes", []interface{}{accountName, withdrawRouteType})
	if err != nil {
		return nil, err
	}
	var resp []*WithdrawVestingRoutes
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_withdraw_routes response", apiID)
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
		return nil, errors.Wrapf(err, "%v: failed to unmarshal lookup_account_names response", apiID)
	}
	return resp, nil
}

//LookupAccounts api request lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]*string, error) {
	raw, err := api.raw("lookup_accounts", []interface{}{lowerBoundName, limit})
	if err != nil {
		return nil, err
	}
	var resp []*string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal lookup_accounts response", apiID)
	}
	return resp, nil
}

//GetVerifyAccountAuthority api request verify_account_authority
func (api *API) GetVerifyAccountAuthority(accountName string, keys []string) (*json.RawMessage, error) {
	raw, err := api.raw("verify_account_authority", []interface{}{accountName, keys})
	if err != nil {
		return nil, err
	}
	return raw, nil
}

//GetVerifyAuthority api request verify_authority
func (api *API) GetVerifyAuthority(trx *types.Transaction) (*bool, error) {
	raw, err := api.raw("verify_authority", []interface{}{&trx})
	if err != nil {
		return nil, err
	}
	var resp bool
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal verify_authority response", apiID)
	}
	return &resp, nil
}
