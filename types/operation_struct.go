package types

import (
	"encoding/json"
)

// Add-on struct

//ExchRate is an additional structure used by other structures.
type ExchRate struct {
	Base  *Asset `json:"base"`
	Quote *Asset `json:"quote"`
}

//POW is an additional structure used by other structures.
type POW struct {
	Worker    string `json:"worker"`
	Input     string `json:"input"`
	Signature string `json:"signature"`
	Work      string `json:"work"`
}

//ChainPropertiesOLD is an additional structure used by other structures.
type ChainPropertiesOLD struct {
	AccountCreationFee *Asset `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

//ChainProperties is an additional structure used by other structures.
type ChainProperties struct {
	AccountCreationFee          *Asset `json:"account_creation_fee"`
	MaximumBlockSize            uint32 `json:"maximum_block_size"`
	SBDInterestRate             uint16 `json:"sbd_interest_rate"`
	CreateAccountMinGolosFee    *Asset `json:"create_account_min_golos_fee"`
	CreateAccountMinDelegation  *Asset `json:"create_account_min_delegation"`
	CreateAccountDelegationTime uint32 `json:"create_account_delegation_time"`
	MinDelegation               *Asset `json:"min_delegation"`
}

//Authority is an additional structure used by other structures.
type Authority struct {
	AccountAuths    StringInt64Map `json:"account_auths"`
	KeyAuths        StringInt64Map `json:"key_auths"`
	WeightThreshold uint32         `json:"weight_threshold"`
}

//POW2Input is an additional structure used by other structures.
type POW2Input struct {
	WorkerAccount string `json:"worker_account"`
	PrevBlock     []byte `json:"prev_block"`
	Nonce         uint64 `json:"nonce"`
}

//Beneficiary is an additional structure used by other structures.
type Beneficiary struct {
	Account string `json:"account"`
	Weight  uint16 `json:"weight"`
}

//CommentPayoutBeneficiaries is an additional structure used by other structures.
type CommentPayoutBeneficiaries struct {
	Beneficiaries []Beneficiary `json:"beneficiaries"`
}

//VoteOperation represents vote operation data.
type VoteOperation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   Int16  `json:"weight"`
}

//Type function that defines the type of operation VoteOperation.
func (op *VoteOperation) Type() OpType {
	return TypeVote
}

//Data returns the operation data VoteOperation.
func (op *VoteOperation) Data() interface{} {
	return op
}

//CommentOperation represents comment operation data.
type CommentOperation struct {
	ParentAuthor   string           `json:"parent_author"`
	ParentPermlink string           `json:"parent_permlink"`
	Author         string           `json:"author"`
	Permlink       string           `json:"permlink"`
	Title          string           `json:"title"`
	Body           string           `json:"body"`
	JSONMetadata   *ContentMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation CommentOperation.
func (op *CommentOperation) Type() OpType {
	return TypeComment
}

//Data returns the operation data CommentOperation.
func (op *CommentOperation) Data() interface{} {
	return op
}

//IsStoryOperation function specifies the type of publication.
func (op *CommentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
}

//TransferOperation represents transfer operation data.
type TransferOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
	Memo   string `json:"memo"`
}

//Type function that defines the type of operation TransferOperation.
func (op *TransferOperation) Type() OpType {
	return TypeTransfer
}

//Data returns the operation data TransferOperation.
func (op *TransferOperation) Data() interface{} {
	return op
}

//TransferToVestingOperation represents transfer_to_vesting operation data.
type TransferToVestingOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
}

//Type function that defines the type of operation TransferToVestingOperation.
func (op *TransferToVestingOperation) Type() OpType {
	return TypeTransferToVesting
}

//Data returns the operation data TransferToVestingOperation.
func (op *TransferToVestingOperation) Data() interface{} {
	return op
}

//WithdrawVestingOperation represents withdraw_vesting operation data.
type WithdrawVestingOperation struct {
	Account       string `json:"account"`
	VestingShares *Asset `json:"vesting_shares"`
}

//Type function that defines the type of operation WithdrawVestingOperation.
func (op *WithdrawVestingOperation) Type() OpType {
	return TypeWithdrawVesting
}

//Data returns the operation data WithdrawVestingOperation.
func (op *WithdrawVestingOperation) Data() interface{} {
	return op
}

//LimitOrderCreateOperation represents limit_order_create operation data.
type LimitOrderCreateOperation struct {
	Owner        string `json:"owner"`
	OrderID      uint32 `json:"orderid"`
	AmountToSell *Asset `json:"amount_to_sell"`
	MinToReceive *Asset `json:"min_to_receive"`
	FillOrKill   bool   `json:"fill_or_kill"`
	Expiration   *Time  `json:"expiration"`
}

//Type function that defines the type of operation LimitOrderCreateOperation.
func (op *LimitOrderCreateOperation) Type() OpType {
	return TypeLimitOrderCreate
}

//Data returns the operation data LimitOrderCreateOperation.
func (op *LimitOrderCreateOperation) Data() interface{} {
	return op
}

//LimitOrderCancelOperation represents limit_order_cancel operation data.
type LimitOrderCancelOperation struct {
	Owner   string `json:"owner"`
	OrderID uint32 `json:"orderid"`
}

//Type function that defines the type of operation LimitOrderCancelOperation.
func (op *LimitOrderCancelOperation) Type() OpType {
	return TypeLimitOrderCancel
}

//Data returns the operation data LimitOrderCancelOperation.
func (op *LimitOrderCancelOperation) Data() interface{} {
	return op
}

//FeedPublishOperation represents feed_publish operation data.
type FeedPublishOperation struct {
	Publisher    string    `json:"publisher"`
	ExchangeRate *ExchRate `json:"exchange_rate"`
}

//Type function that defines the type of operation FeedPublishOperation.
func (op *FeedPublishOperation) Type() OpType {
	return TypeFeedPublish
}

//Data returns the operation data FeedPublishOperation.
func (op *FeedPublishOperation) Data() interface{} {
	return op
}

//ConvertOperation represents convert operation data.
type ConvertOperation struct {
	Owner     string `json:"owner"`
	RequestID uint32 `json:"requestid"`
	Amount    *Asset `json:"amount"`
}

//Type function that defines the type of operation ConvertOperation.
func (op *ConvertOperation) Type() OpType {
	return TypeConvert
}

//Data returns the operation data ConvertOperation.
func (op *ConvertOperation) Data() interface{} {
	return op
}

//AccountCreateOperation represents account_create operation data.
type AccountCreateOperation struct {
	Fee            *Asset           `json:"fee"`
	Creator        string           `json:"creator"`
	NewAccountName string           `json:"new_account_name"`
	Owner          *Authority       `json:"owner"`
	Active         *Authority       `json:"active"`
	Posting        *Authority       `json:"posting"`
	MemoKey        string           `json:"memo_key"`
	JSONMetadata   *AccountMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation AccountCreateOperation.
func (op *AccountCreateOperation) Type() OpType {
	return TypeAccountCreate
}

//Data returns the operation data AccountCreateOperation.
func (op *AccountCreateOperation) Data() interface{} {
	return op
}

//AccountUpdateOperation represents account_update operation data.
type AccountUpdateOperation struct {
	Account      string           `json:"account"`
	Owner        *Authority       `json:"owner"`
	Active       *Authority       `json:"active"`
	Posting      *Authority       `json:"posting"`
	MemoKey      string           `json:"memo_key"`
	JSONMetadata *AccountMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation AccountUpdateOperation.
func (op *AccountUpdateOperation) Type() OpType {
	return TypeAccountUpdate
}

//Data returns the operation data AccountUpdateOperation.
func (op *AccountUpdateOperation) Data() interface{} {
	return op
}

//WitnessUpdateOperation represents witness_update operation data.
type WitnessUpdateOperation struct {
	Owner           string              `json:"owner"`
	URL             string              `json:"url"`
	BlockSigningKey string              `json:"block_signing_key"`
	Props           *ChainPropertiesOLD `json:"props"`
	Fee             *Asset              `json:"fee"`
}

//Type function that defines the type of operation WitnessUpdateOperation.
func (op *WitnessUpdateOperation) Type() OpType {
	return TypeWitnessUpdate
}

//Data returns the operation data WitnessUpdateOperation.
func (op *WitnessUpdateOperation) Data() interface{} {
	return op
}

//AccountWitnessVoteOperation represents account_witness_vote operation data.
type AccountWitnessVoteOperation struct {
	Account string `json:"account"`
	Witness string `json:"witness"`
	Approve bool   `json:"approve"`
}

//Type function that defines the type of operation AccountWitnessVoteOperation.
func (op *AccountWitnessVoteOperation) Type() OpType {
	return TypeAccountWitnessVote
}

//Data returns the operation data AccountWitnessVoteOperation.
func (op *AccountWitnessVoteOperation) Data() interface{} {
	return op
}

//AccountWitnessProxyOperation represents account_witness_proxy operation data.
type AccountWitnessProxyOperation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
}

//Type function that defines the type of operation AccountWitnessProxyOperation.
func (op *AccountWitnessProxyOperation) Type() OpType {
	return TypeAccountWitnessProxy
}

//Data returns the operation data AccountWitnessProxyOperation.
func (op *AccountWitnessProxyOperation) Data() interface{} {
	return op
}

//POWOperation represents pow operation data.
type POWOperation struct {
	WorkerAccount string           `json:"worker_account"`
	BlockID       string           `json:"block_id"`
	Nonce         *Int             `json:"nonce"`
	Work          *POW             `json:"work"`
	Props         *ChainProperties `json:"props"`
}

//Type function that defines the type of operation POWOperation.
func (op *POWOperation) Type() OpType {
	return TypePOW
}

//Data returns the operation data POWOperation.
func (op *POWOperation) Data() interface{} {
	return op
}

//CustomOperation represents custom operation data.
type CustomOperation struct {
	RequiredAuths []string `json:"required_auths"`
	ID            uint16   `json:"id"`
	Datas         string   `json:"data"`
}

//Type function that defines the type of operation CustomOperation.
func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

//Data returns the operation data CustomOperation.
func (op *CustomOperation) Data() interface{} {
	return op
}

//ReportOverProductionOperation represents report_over_production operation data.
type ReportOverProductionOperation struct {
	Reporter string `json:"reporter"`
}

//Type function that defines the type of operation ReportOverProductionOperation.
func (op *ReportOverProductionOperation) Type() OpType {
	return TypeReportOverProduction
}

//Data returns the operation data ReportOverProductionOperation.
func (op *ReportOverProductionOperation) Data() interface{} {
	return op
}

//DeleteCommentOperation represents delete_comment operation data.
type DeleteCommentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//Type function that defines the type of operation DeleteCommentOperation.
func (op *DeleteCommentOperation) Type() OpType {
	return TypeDeleteComment
}

//Data returns the operation data DeleteCommentOperation.
func (op *DeleteCommentOperation) Data() interface{} {
	return op
}

// struct CustomJSONOperation{} in to file operation_custom_json.go

//CommentOptionsOperation represents comment_options operation data.
type CommentOptionsOperation struct {
	Author               string        `json:"author"`
	Permlink             string        `json:"permlink"`
	MaxAcceptedPayout    *Asset        `json:"max_accepted_payout"`
	PercentSteemDollars  uint16        `json:"percent_steem_dollars"`
	AllowVotes           bool          `json:"allow_votes"`
	AllowCurationRewards bool          `json:"allow_curation_rewards"`
	Extensions           []interface{} `json:"extensions"`
}

//Type function that defines the type of operation CommentOptionsOperation.
func (op *CommentOptionsOperation) Type() OpType {
	return TypeCommentOptions
}

//Data returns the operation data CommentOptionsOperation.
func (op *CommentOptionsOperation) Data() interface{} {
	return op
}

//SetWithdrawVestingRouteOperation represents set_withdraw_vesting_route operation data.
type SetWithdrawVestingRouteOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

//Type function that defines the type of operation SetWithdrawVestingRouteOperation.
func (op *SetWithdrawVestingRouteOperation) Type() OpType {
	return TypeSetWithdrawVestingRoute
}

//Data returns the operation data SetWithdrawVestingRouteOperation.
func (op *SetWithdrawVestingRouteOperation) Data() interface{} {
	return op
}

//LimitOrderCreate2Operation represents limit_order_create2 operation data.
type LimitOrderCreate2Operation struct {
	Owner        string    `json:"owner"`
	OrderID      uint32    `json:"orderid"`
	AmountToSell *Asset    `json:"amount_to_sell"`
	ExchangeRate *ExchRate `json:"exchange_rate"`
	FillOrKill   bool      `json:"fill_or_kill"`
	Expiration   *Time     `json:"expiration"`
}

//Type function that defines the type of operation LimitOrderCreate2Operation.
func (op *LimitOrderCreate2Operation) Type() OpType {
	return TypeLimitOrderCreate2
}

//Data returns the operation data LimitOrderCreate2Operation.
func (op *LimitOrderCreate2Operation) Data() interface{} {
	return op
}

//ChallengeAuthorityOperation represents challenge_authority operation data.
type ChallengeAuthorityOperation struct {
	Challenger   string `json:"challenger"`
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

//Type function that defines the type of operation ChallengeAuthorityOperation.
func (op *ChallengeAuthorityOperation) Type() OpType {
	return TypeChallengeAuthority
}

//Data returns the operation data ChallengeAuthorityOperation.
func (op *ChallengeAuthorityOperation) Data() interface{} {
	return op
}

//ProveAuthorityOperation represents prove_authority operation data.
type ProveAuthorityOperation struct {
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

//Type function that defines the type of operation ProveAuthorityOperation.
func (op *ProveAuthorityOperation) Type() OpType {
	return TypeProveAuthority
}

//Data returns the operation data ProveAuthorityOperation.
func (op *ProveAuthorityOperation) Data() interface{} {
	return op
}

//RequestAccountRecoveryOperation represents request_account_recovery operation data.
type RequestAccountRecoveryOperation struct {
	RecoveryAccount   string        `json:"recovery_account"`
	AccountToRecover  string        `json:"account_to_recover"`
	NewOwnerAuthority *Authority    `json:"new_owner_authority"`
	Extensions        []interface{} `json:"extensions"`
}

//Type function that defines the type of operation RequestAccountRecoveryOperation.
func (op *RequestAccountRecoveryOperation) Type() OpType {
	return TypeRequestAccountRecovery
}

//Data returns the operation data RequestAccountRecoveryOperation.
func (op *RequestAccountRecoveryOperation) Data() interface{} {
	return op
}

//RecoverAccountOperation represents recover_account operation data.
type RecoverAccountOperation struct {
	AccountToRecover     string        `json:"account_to_recover"`
	NewOwnerAuthority    *Authority    `json:"new_owner_authority"`
	RecentOwnerAuthority *Authority    `json:"recent_owner_authority"`
	Extensions           []interface{} `json:"extensions"`
}

//Type function that defines the type of operation RecoverAccountOperation.
func (op *RecoverAccountOperation) Type() OpType {
	return TypeRecoverAccount
}

//Data returns the operation data RecoverAccountOperation.
func (op *RecoverAccountOperation) Data() interface{} {
	return op
}

//ChangeRecoveryAccountOperation represents change_recovery_account operation data.
type ChangeRecoveryAccountOperation struct {
	AccountToRecover   string        `json:"account_to_recover"`
	NewRecoveryAccount string        `json:"new_recovery_account"`
	Extensions         []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ChangeRecoveryAccountOperation.
func (op *ChangeRecoveryAccountOperation) Type() OpType {
	return TypeChangeRecoveryAccount
}

//Data returns the operation data ChangeRecoveryAccountOperation.
func (op *ChangeRecoveryAccountOperation) Data() interface{} {
	return op
}

//EscrowTransferOperation represents escrow_transfer operation data.
type EscrowTransferOperation struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	SbdAmount            *Asset `json:"sbd_amount"`
	SteemAmount          *Asset `json:"steem_amount"`
	EscrowID             uint32 `json:"escrow_id"`
	Agent                string `json:"agent"`
	Fee                  *Asset `json:"fee"`
	JSONMeta             string `json:"json_meta"`
	RatificationDeadline *Time  `json:"ratification_deadline"`
	EscrowExpiration     *Time  `json:"escrow_expiration"`
}

//Type function that defines the type of operation EscrowTransferOperation.
func (op *EscrowTransferOperation) Type() OpType {
	return TypeEscrowTransfer
}

//Data returns the operation data EscrowTransferOperation.
func (op *EscrowTransferOperation) Data() interface{} {
	return op
}

//EscrowDisputeOperation represents escrow_dispute operation data.
type EscrowDisputeOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowID uint32 `json:"escrow_id"`
}

//Type function that defines the type of operation EscrowDisputeOperation.
func (op *EscrowDisputeOperation) Type() OpType {
	return TypeEscrowDispute
}

//Data returns the operation data EscrowDisputeOperation.
func (op *EscrowDisputeOperation) Data() interface{} {
	return op
}

//EscrowReleaseOperation represents escrow_release operation data.
type EscrowReleaseOperation struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Agent       string `json:"agent"`
	Who         string `json:"who"`
	Receiver    string `json:"receiver"`
	EscrowID    uint32 `json:"escrow_id"`
	SbdAmount   *Asset `json:"sbd_amount"`
	SteemAmount *Asset `json:"steem_amount"`
}

//Type function that defines the type of operation EscrowReleaseOperation.
func (op *EscrowReleaseOperation) Type() OpType {
	return TypeEscrowRelease
}

//Data returns the operation data EscrowReleaseOperation.
func (op *EscrowReleaseOperation) Data() interface{} {
	return op
}

//POW2Operation represents pow2 operation data.
type POW2Operation struct {
	Input      *POW2Input `json:"input"`
	PowSummary uint32     `json:"pow_summary"`
}

//Type function that defines the type of operation POW2Operation.
func (op *POW2Operation) Type() OpType {
	return TypePOW2
}

//Data returns the operation data POW2Operation.
func (op *POW2Operation) Data() interface{} {
	return op
}

//EscrowApproveOperation represents escrow_approve operation data.
type EscrowApproveOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowID uint32 `json:"escrow_id"`
	Approve  bool   `json:"approve"`
}

//Type function that defines the type of operation EscrowApproveOperation.
func (op *EscrowApproveOperation) Type() OpType {
	return TypeEscrowApprove
}

//Data returns the operation data EscrowApproveOperation.
func (op *EscrowApproveOperation) Data() interface{} {
	return op
}

//TransferToSavingsOperation represents transfer_to_savings operation data.
type TransferToSavingsOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
	Memo   string `json:"memo"`
}

//Type function that defines the type of operation TransferToSavingsOperation.
func (op *TransferToSavingsOperation) Type() OpType {
	return TypeTransferToSavings
}

//Data returns the operation data TransferToSavingsOperation.
func (op *TransferToSavingsOperation) Data() interface{} {
	return op
}

//TransferFromSavingsOperation represents transfer_from_savings operation data.
type TransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestID uint32 `json:"request_id"`
	To        string `json:"to"`
	Amount    *Asset `json:"amount"`
	Memo      string `json:"memo"`
}

//Type function that defines the type of operation TransferFromSavingsOperation.
func (op *TransferFromSavingsOperation) Type() OpType {
	return TypeTransferFromSavings
}

//Data returns the operation data TransferFromSavingsOperation.
func (op *TransferFromSavingsOperation) Data() interface{} {
	return op
}

//CancelTransferFromSavingsOperation represents cancel_transfer_from_savings operation data.
type CancelTransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestID uint32 `json:"request_id"`
}

//Type function that defines the type of operation CancelTransferFromSavingsOperation.
func (op *CancelTransferFromSavingsOperation) Type() OpType {
	return TypeCancelTransferFromSavings
}

//Data returns the operation data CancelTransferFromSavingsOperation.
func (op *CancelTransferFromSavingsOperation) Data() interface{} {
	return op
}

//CustomBinaryOperation represents custom_binary operation data.
type CustomBinaryOperation struct {
	RequiredOwnerAuths   []string `json:"required_owner_auths"`
	RequiredActiveAuths  []string `json:"required_active_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	RequiredAuths        []string `json:"required_auths"`
	ID                   string   `json:"id"`
	Datas                []byte   `json:"data"`
}

//Type function that defines the type of operation CustomBinaryOperation.
func (op *CustomBinaryOperation) Type() OpType {
	return TypeCustomBinary
}

//Data returns the operation data CustomBinaryOperation.
func (op *CustomBinaryOperation) Data() interface{} {
	return op
}

//DeclineVotingRightsOperation represents decline_voting_rights operation data.
type DeclineVotingRightsOperation struct {
	Account string `json:"account"`
	Decline bool   `json:"decline"`
}

//Type function that defines the type of operation DeclineVotingRightsOperation.
func (op *DeclineVotingRightsOperation) Type() OpType {
	return TypeDeclineVotingRights
}

//Data returns the operation data DeclineVotingRightsOperation.
func (op *DeclineVotingRightsOperation) Data() interface{} {
	return op
}

//ResetAccountOperation represents reset_account operation data.
type ResetAccountOperation struct {
	ResetAccount      string     `json:"reset_account"`
	AccountToReset    string     `json:"Account_to_reset"`
	NewOwnerAuthority *Authority `json:"new_owner_authority"`
}

//Type function that defines the type of operation ResetAccountOperation.
func (op *ResetAccountOperation) Type() OpType {
	return TypeResetAccount
}

//Data returns the operation data ResetAccountOperation.
func (op *ResetAccountOperation) Data() interface{} {
	return op
}

//SetResetAccountOperation represents set_reset_account operation data.
type SetResetAccountOperation struct {
	Account             string `json:"account"`
	CurrentResetAccount string `json:"current_reset_account"`
	ResetAccount        string `json:"reset_account"`
}

//Type function that defines the type of operation SetResetAccountOperation.
func (op *SetResetAccountOperation) Type() OpType {
	return TypeSetResetAccount
}

//Data returns the operation data SetResetAccountOperation.
func (op *SetResetAccountOperation) Data() interface{} {
	return op
}

//DelegateVestingSharesOperation represents delegate_vesting_shares operation data.
type DelegateVestingSharesOperation struct {
	Delegator     string `json:"delegator"`
	Delegatee     string `json:"delegatee"`
	VestingShares *Asset `json:"vesting_shares"`
}

//Type function that defines the type of operation DelegateVestingSharesOperation.
func (op *DelegateVestingSharesOperation) Type() OpType {
	return TypeDelegateVestingShares
}

//Data returns the operation data DelegateVestingSharesOperation.
func (op *DelegateVestingSharesOperation) Data() interface{} {
	return op
}

//AccountCreateWithDelegationOperation represents account_create_with_delegation operation data.
type AccountCreateWithDelegationOperation struct {
	Fee            *Asset           `json:"fee"`
	Delegation     *Asset           `json:"delegation"`
	Creator        string           `json:"creator"`
	NewAccountName string           `json:"new_account_name"`
	Owner          *Authority       `json:"owner"`
	Active         *Authority       `json:"active"`
	Posting        *Authority       `json:"posting"`
	MemoKey        string           `json:"memo_key"`
	JSONMetadata   *AccountMetadata `json:"json_metadata"`
	Extensions     []interface{}    `json:"extensions"`
}

//Type function that defines the type of operation AccountCreateWithDelegationOperation.
func (op *AccountCreateWithDelegationOperation) Type() OpType {
	return TypeAccountCreateWithDelegation
}

//Data returns the operation data AccountCreateWithDelegationOperation.
func (op *AccountCreateWithDelegationOperation) Data() interface{} {
	return op
}

//AccountMetadataOperation represents account_metadata operation data.
type AccountMetadataOperation struct {
	Account      string           `json:"account"`
	JSONMetadata *AccountMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation AccountMetadataOperation.
func (op *AccountMetadataOperation) Type() OpType {
	return TypeAccountMetadata
}

//Data returns the operation data AccountMetadataOperation.
func (op *AccountMetadataOperation) Data() interface{} {
	return op
}

//ProposalCreateOperation represents proposal_create operation data.
type ProposalCreateOperation struct {
	Author             string             `json:"author"`
	Title              string             `json:"title"`
	Memo               string             `json:"memo"`
	ProposedOperations ProposalOperations `json:"proposed_operations"`
	ExpirationTime     *Time              `json:"expiration_time"`
	ReviewPeriodTime   *Time              `json:"review_period_time"`
	Extensions         []interface{}      `json:"extensions"`
}

//Type function that defines the type of operation ProposalCreateOperation.
func (op *ProposalCreateOperation) Type() OpType {
	return TypeProposalCreate
}

//Data returns the operation data ProposalCreateOperation.
func (op *ProposalCreateOperation) Data() interface{} {
	return op
}

//ProposalUpdateOperation represents proposal_update operation data.
type ProposalUpdateOperation struct {
	Author                   string        `json:"author"`
	Title                    string        `json:"title"`
	ActiveApprovalsToAdd     []string      `json:"active_approvals_to_add"`
	ActiveApprovalsToRemove  []string      `json:"active_approvals_to_remove"`
	OwnerApprovalsToAdd      []string      `json:"owner_approvals_to_add"`
	OwnerApprovalsToRemove   []string      `json:"owner_approvals_to_remove"`
	PostingApprovalsToAdd    []string      `json:"posting_approvals_to_add"`
	PostingApprovalsToRemove []string      `json:"posting_approvals_to_remove"`
	KeyApprovalsToAdd        []string      `json:"key_approvals_to_add"`
	KeyApprovalsToRemove     []string      `json:"key_approvals_to_remove"`
	Extensions               []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ProposalUpdateOperation.
func (op *ProposalUpdateOperation) Type() OpType {
	return TypeProposalUpdate
}

//Data returns the operation data ProposalUpdateOperation.
func (op *ProposalUpdateOperation) Data() interface{} {
	return op
}

//ProposalDeleteOperation represents proposal_delete operation data.
type ProposalDeleteOperation struct {
	Author     string        `json:"author"`
	Title      string        `json:"title"`
	Requester  string        `json:"requester"`
	Extensions []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ProposalDeleteOperation.
func (op *ProposalDeleteOperation) Type() OpType {
	return TypeProposalDelete
}

//Data returns the operation data ProposalDeleteOperation.
func (op *ProposalDeleteOperation) Data() interface{} {
	return op
}

//ChainPropertiesUpdateOperation represents chain_properties_update operation data.
type ChainPropertiesUpdateOperation struct {
	Owner string        `json:"owner"`
	Props []interface{} `json:"props"`
}

//Type function that defines the type of operation ChainPropertiesUpdateOperation.
func (op *ChainPropertiesUpdateOperation) Type() OpType {
	return TypeChainPropertiesUpdate
}

//Data returns the operation data ChainPropertiesUpdateOperation.
func (op *ChainPropertiesUpdateOperation) Data() interface{} {
	return op
}

//Virtual Operation

//FillConvertRequestOperation represents fill_convert_request operation data.
type FillConvertRequestOperation struct {
	Owner     string `json:"owner"`
	Requestid uint32 `json:"requestid"`
	AmountIn  *Asset `json:"amount_in"`
	AmountOut *Asset `json:"amount_out"`
}

//Type function that defines the type of operation FillConvertRequestOperation.
func (op *FillConvertRequestOperation) Type() OpType {
	return TypeFillConvertRequest
}

//Data returns the operation data FillConvertRequestOperation.
func (op *FillConvertRequestOperation) Data() interface{} {
	return op
}

//AuthorRewardOperation represents author_reward operation data.
type AuthorRewardOperation struct {
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	SbdPayout     *Asset `json:"sbd_payout"`
	SteemPayout   *Asset `json:"steem_payout"`
	VestingPayout *Asset `json:"vesting_payout"`
}

//Type function that defines the type of operation AuthorRewardOperation.
func (op *AuthorRewardOperation) Type() OpType {
	return TypeAuthorReward
}

//Data returns the operation data AuthorRewardOperation.
func (op *AuthorRewardOperation) Data() interface{} {
	return op
}

//CurationRewardOperation represents curation_reward operation data.
type CurationRewardOperation struct {
	Curator         string `json:"curator"`
	Reward          *Asset `json:"reward"`
	CommentAuthor   string `json:"comment_author"`
	CommentPermlink string `json:"comment_permlink"`
}

//Type function that defines the type of operation CurationRewardOperation.
func (op *CurationRewardOperation) Type() OpType {
	return TypeCurationReward
}

//Data returns the operation data CurationRewardOperation.
func (op *CurationRewardOperation) Data() interface{} {
	return op
}

//CommentRewardOperation represents comment_reward operation data.
type CommentRewardOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Payout   *Asset `json:"payout"`
}

//Type function that defines the type of operation CommentRewardOperation.
func (op *CommentRewardOperation) Type() OpType {
	return TypeCommentReward
}

//Data returns the operation data CommentRewardOperation.
func (op *CommentRewardOperation) Data() interface{} {
	return op
}

//LiquidityRewardOperation represents liquidity_reward operation data.
type LiquidityRewardOperation struct {
	Owner  string `json:"owner"`
	Payout *Asset `json:"payout"`
}

//Type function that defines the type of operation LiquidityRewardOperation.
func (op *LiquidityRewardOperation) Type() OpType {
	return TypeLiquidityReward
}

//Data returns the operation data LiquidityRewardOperation.
func (op *LiquidityRewardOperation) Data() interface{} {
	return op
}

//InterestOperation represents interest operation data.
type InterestOperation struct {
	Owner    string `json:"owner"`
	Interest *Asset `json:"interest"`
}

//Type function that defines the type of operation InterestOperation.
func (op *InterestOperation) Type() OpType {
	return TypeInterest
}

//Data returns the operation data InterestOperation.
func (op *InterestOperation) Data() interface{} {
	return op
}

//FillVestingWithdrawOperation represents fill_vesting_withdraw operation data.
type FillVestingWithdrawOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Withdrawn   *Asset `json:"withdrawn"`
	Deposited   *Asset `json:"deposited"`
}

//Type function that defines the type of operation FillVestingWithdrawOperation.
func (op *FillVestingWithdrawOperation) Type() OpType {
	return TypeFillVestingWithdraw
}

//Data returns the operation data FillVestingWithdrawOperation.
func (op *FillVestingWithdrawOperation) Data() interface{} {
	return op
}

//FillOrderOperation represents fill_order operation data.
type FillOrderOperation struct {
	CurrentOwner   string `json:"current_owner"`
	CurrentOrderid uint32 `json:"current_orderid"`
	CurrentPays    *Asset `json:"current_pays"`
	OpenOwner      string `json:"open_owner"`
	OpenOrderid    uint32 `json:"open_orderid"`
	OpenPays       *Asset `json:"open_pays"`
}

//Type function that defines the type of operation FillOrderOperation.
func (op *FillOrderOperation) Type() OpType {
	return TypeFillOrder
}

//Data returns the operation data FillOrderOperation.
func (op *FillOrderOperation) Data() interface{} {
	return op
}

//ShutdownWitnessOperation represents shutdown_witness operation data.
type ShutdownWitnessOperation struct {
	Owner string `json:"owner"`
}

//Type function that defines the type of operation ShutdownWitnessOperation.
func (op *ShutdownWitnessOperation) Type() OpType {
	return TypeShutdownWitness
}

//Data returns the operation data ShutdownWitnessOperation.
func (op *ShutdownWitnessOperation) Data() interface{} {
	return op
}

//FillTransferFromSavingsOperation represents fill_transfer_from_savings operation data.
type FillTransferFromSavingsOperation struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    *Asset `json:"amount"`
	RequestID uint32 `json:"request_id"`
	Memo      string `json:"memo"`
}

//Type function that defines the type of operation FillTransferFromSavingsOperation.
func (op *FillTransferFromSavingsOperation) Type() OpType {
	return TypeFillTransferFromSavings
}

//Data returns the operation data FillTransferFromSavingsOperation.
func (op *FillTransferFromSavingsOperation) Data() interface{} {
	return op
}

//HardforkOperation represents hardfork operation data.
type HardforkOperation struct {
	HardforkID uint32 `json:"hardfork_id"`
}

//Type function that defines the type of operation HardforkOperation.
func (op *HardforkOperation) Type() OpType {
	return TypeHardfork
}

//Data returns the operation data HardforkOperation.
func (op *HardforkOperation) Data() interface{} {
	return op
}

//CommentPayoutUpdateOperation represents comment_payout_update operation data.
type CommentPayoutUpdateOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//Type function that defines the type of operation CommentPayoutUpdateOperation.
func (op *CommentPayoutUpdateOperation) Type() OpType {
	return TypeCommentPayoutUpdate
}

//Data returns the operation data CommentPayoutUpdateOperation.
func (op *CommentPayoutUpdateOperation) Data() interface{} {
	return op
}

//CommentBenefactorRewardOperation represents comment_benefactor_reward operation data.
type CommentBenefactorRewardOperation struct {
	Benefactor string `json:"benefactor"`
	Author     string `json:"author"`
	Permlink   string `json:"permlink"`
	Reward     *Asset `json:"reward"`
}

//Type function that defines the type of operation CommentBenefactorRewardOperation.
func (op *CommentBenefactorRewardOperation) Type() OpType {
	return TypeCommentBenefactorReward
}

//Data returns the operation data CommentBenefactorRewardOperation.
func (op *CommentBenefactorRewardOperation) Data() interface{} {
	return op
}

//ReturnVestingDelegationOperation represents return_vesting_delegation operation data.
type ReturnVestingDelegationOperation struct {
	Account       string `json:"account"`
	VestingShares *Asset `json:"vesting_shares"`
}

//Type function that defines the type of operation ReturnVestingDelegationOperation.
func (op *ReturnVestingDelegationOperation) Type() OpType {
	return TypeReturnVestingDelegation
}

//Data returns the operation data ReturnVestingDelegationOperation.
func (op *ReturnVestingDelegationOperation) Data() interface{} {
	return op
}

//UnknownOperation represents Unknown operation data.
type UnknownOperation struct {
	kind OpType
	data *json.RawMessage
}

//Type function that defines the type of operation UnknownOperation.
func (op *UnknownOperation) Type() OpType {
	return op.kind
}

//Data returns the operation data UnknownOperation.
func (op *UnknownOperation) Data() interface{} {
	return op.data
}
