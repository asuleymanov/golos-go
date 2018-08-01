package types

import (
	"encoding/json"
)

// Add-on struct
type ExchRate struct {
	Base  *Asset `json:"base"`
	Quote *Asset `json:"quote"`
}

type POW struct {
	Worker    string `json:"worker"`
	Input     string `json:"input"`
	Signature string `json:"signature"`
	Work      string `json:"work"`
}

type ChainPropertiesOLD struct {
	AccountCreationFee *Asset `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

type ChainProperties struct {
	AccountCreationFee          *Asset `json:"account_creation_fee"`
	MaximumBlockSize            uint32 `json:"maximum_block_size"`
	SBDInterestRate             uint16 `json:"sbd_interest_rate"`
	CreateAccountMinGolosFee    *Asset `json:"create_account_min_golos_fee"`
	CreateAccountMinDelegation  *Asset `json:"create_account_min_delegation"`
	CreateAccountDelegationTime uint32 `json:"create_account_delegation_time"`
	MinDelegation               *Asset `json:"min_delegation"`
}

type Authority struct {
	AccountAuths    StringInt64Map `json:"account_auths"`
	KeyAuths        StringInt64Map `json:"key_auths"`
	WeightThreshold uint32         `json:"weight_threshold"`
}

type POW2Input struct {
	WorkerAccount string `json:"worker_account"`
	PrevBlock     []byte `json:"prev_block"`
	Nonce         uint64 `json:"nonce"`
}

type Beneficiary struct {
	Account string `json:"account"`
	Weight  uint16 `json:"weight"`
}

type CommentPayoutBeneficiaries struct {
	Beneficiaries []Beneficiary `json:"beneficiaries"`
}

// struct VoteOperation{}
type VoteOperation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   Int16  `json:"weight"`
}

func (op *VoteOperation) Type() OpType {
	return TypeVote
}

func (op *VoteOperation) Data() interface{} {
	return op
}

// struct CommentOperation{}
type CommentOperation struct {
	ParentAuthor   string           `json:"parent_author"`
	ParentPermlink string           `json:"parent_permlink"`
	Author         string           `json:"author"`
	Permlink       string           `json:"permlink"`
	Title          string           `json:"title"`
	Body           string           `json:"body"`
	JSONMetadata   *ContentMetadata `json:"json_metadata"`
}

func (op *CommentOperation) Type() OpType {
	return TypeComment
}

func (op *CommentOperation) Data() interface{} {
	return op
}

func (op *CommentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
}

// struct TransferOperation{}
type TransferOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *TransferOperation) Type() OpType {
	return TypeTransfer
}

func (op *TransferOperation) Data() interface{} {
	return op
}

// struct TransferToVestingOperation{}
type TransferToVestingOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
}

func (op *TransferToVestingOperation) Type() OpType {
	return TypeTransferToVesting
}

func (op *TransferToVestingOperation) Data() interface{} {
	return op
}

// struct WithdrawVestingOperation{}
type WithdrawVestingOperation struct {
	Account       string `json:"account"`
	VestingShares *Asset `json:"vesting_shares"`
}

func (op *WithdrawVestingOperation) Type() OpType {
	return TypeWithdrawVesting
}

func (op *WithdrawVestingOperation) Data() interface{} {
	return op
}

// struct LimitOrderCreateOperation{}
type LimitOrderCreateOperation struct {
	Owner        string `json:"owner"`
	OrderID      uint32 `json:"orderid"`
	AmountToSell *Asset `json:"amount_to_sell"`
	MinToReceive *Asset `json:"min_to_receive"`
	FillOrKill   bool   `json:"fill_or_kill"`
	Expiration   *Time  `json:"expiration"`
}

func (op *LimitOrderCreateOperation) Type() OpType {
	return TypeLimitOrderCreate
}

func (op *LimitOrderCreateOperation) Data() interface{} {
	return op
}

// struct LimitOrderCancelOperation{}
type LimitOrderCancelOperation struct {
	Owner   string `json:"owner"`
	OrderID uint32 `json:"orderid"`
}

func (op *LimitOrderCancelOperation) Type() OpType {
	return TypeLimitOrderCancel
}

func (op *LimitOrderCancelOperation) Data() interface{} {
	return op
}

// struct FeedPublishOperation{}
type FeedPublishOperation struct {
	Publisher    string    `json:"publisher"`
	ExchangeRate *ExchRate `json:"exchange_rate"`
}

func (op *FeedPublishOperation) Type() OpType {
	return TypeFeedPublish
}

func (op *FeedPublishOperation) Data() interface{} {
	return op
}

// struct ConvertOperation{}
type ConvertOperation struct {
	Owner     string `json:"owner"`
	RequestID uint32 `json:"requestid"`
	Amount    *Asset `json:"amount"`
}

func (op *ConvertOperation) Type() OpType {
	return TypeConvert
}

func (op *ConvertOperation) Data() interface{} {
	return op
}

// struct AccountCreateOperation{}
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

func (op *AccountCreateOperation) Type() OpType {
	return TypeAccountCreate
}

func (op *AccountCreateOperation) Data() interface{} {
	return op
}

// struct AccountUpdateOperation{}
type AccountUpdateOperation struct {
	Account      string           `json:"account"`
	Owner        *Authority       `json:"owner"`
	Active       *Authority       `json:"active"`
	Posting      *Authority       `json:"posting"`
	MemoKey      string           `json:"memo_key"`
	JSONMetadata *AccountMetadata `json:"json_metadata"`
}

func (op *AccountUpdateOperation) Type() OpType {
	return TypeAccountUpdate
}

func (op *AccountUpdateOperation) Data() interface{} {
	return op
}

// struct WitnessUpdateOperation{}
type WitnessUpdateOperation struct {
	Owner           string              `json:"owner"`
	URL             string              `json:"url"`
	BlockSigningKey string              `json:"block_signing_key"`
	Props           *ChainPropertiesOLD `json:"props"`
	Fee             *Asset              `json:"fee"`
}

func (op *WitnessUpdateOperation) Type() OpType {
	return TypeWitnessUpdate
}

func (op *WitnessUpdateOperation) Data() interface{} {
	return op
}

// struct AccountWitnessVoteOperation{}
type AccountWitnessVoteOperation struct {
	Account string `json:"account"`
	Witness string `json:"witness"`
	Approve bool   `json:"approve"`
}

func (op *AccountWitnessVoteOperation) Type() OpType {
	return TypeAccountWitnessVote
}

func (op *AccountWitnessVoteOperation) Data() interface{} {
	return op
}

// struct AccountWitnessProxyOperation{}
type AccountWitnessProxyOperation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
}

func (op *AccountWitnessProxyOperation) Type() OpType {
	return TypeAccountWitnessProxy
}

func (op *AccountWitnessProxyOperation) Data() interface{} {
	return op
}

// struct POWOperation{}
type POWOperation struct {
	WorkerAccount string           `json:"worker_account"`
	BlockID       string           `json:"block_id"`
	Nonce         *Int             `json:"nonce"`
	Work          *POW             `json:"work"`
	Props         *ChainProperties `json:"props"`
}

func (op *POWOperation) Type() OpType {
	return TypePOW
}

func (op *POWOperation) Data() interface{} {
	return op
}

// struct CustomOperation{}
type CustomOperation struct {
	RequiredAuths []string `json:"required_auths"`
	ID            uint16   `json:"id"`
	Datas         string   `json:"data"`
}

func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

func (op *CustomOperation) Data() interface{} {
	return op
}

// struct ReportOverProductionOperation{}
type ReportOverProductionOperation struct {
	Reporter string `json:"reporter"`
}

func (op *ReportOverProductionOperation) Type() OpType {
	return TypeReportOverProduction
}

func (op *ReportOverProductionOperation) Data() interface{} {
	return op
}

// struct DeleteCommentOperation{}
type DeleteCommentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *DeleteCommentOperation) Type() OpType {
	return TypeDeleteComment
}

func (op *DeleteCommentOperation) Data() interface{} {
	return op
}

// struct CustomJSONOperation{} in to file operation_custom_json.go
// struct CommentOptionsOperation{}
type CommentOptionsOperation struct {
	Author               string        `json:"author"`
	Permlink             string        `json:"permlink"`
	MaxAcceptedPayout    *Asset        `json:"max_accepted_payout"`
	PercentSteemDollars  uint16        `json:"percent_steem_dollars"`
	AllowVotes           bool          `json:"allow_votes"`
	AllowCurationRewards bool          `json:"allow_curation_rewards"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *CommentOptionsOperation) Type() OpType {
	return TypeCommentOptions
}

func (op *CommentOptionsOperation) Data() interface{} {
	return op
}

// struct SetWithdrawVestingRouteOperation{}
type SetWithdrawVestingRouteOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

func (op *SetWithdrawVestingRouteOperation) Type() OpType {
	return TypeSetWithdrawVestingRoute
}

func (op *SetWithdrawVestingRouteOperation) Data() interface{} {
	return op
}

// struct LimitOrderCreate2Operation{}
type LimitOrderCreate2Operation struct {
	Owner        string    `json:"owner"`
	OrderID      uint32    `json:"orderid"`
	AmountToSell *Asset    `json:"amount_to_sell"`
	ExchangeRate *ExchRate `json:"exchange_rate"`
	FillOrKill   bool      `json:"fill_or_kill"`
	Expiration   *Time     `json:"expiration"`
}

func (op *LimitOrderCreate2Operation) Type() OpType {
	return TypeLimitOrderCreate2
}

func (op *LimitOrderCreate2Operation) Data() interface{} {
	return op
}

// struct ChallengeAuthorityOperation{}
type ChallengeAuthorityOperation struct {
	Challenger   string `json:"challenger"`
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ChallengeAuthorityOperation) Type() OpType {
	return TypeChallengeAuthority
}

func (op *ChallengeAuthorityOperation) Data() interface{} {
	return op
}

// struct ProveAuthorityOperation{}
type ProveAuthorityOperation struct {
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ProveAuthorityOperation) Type() OpType {
	return TypeProveAuthority
}

func (op *ProveAuthorityOperation) Data() interface{} {
	return op
}

// struct RequestAccountRecoveryOperation{}
type RequestAccountRecoveryOperation struct {
	RecoveryAccount   string        `json:"recovery_account"`
	AccountToRecover  string        `json:"account_to_recover"`
	NewOwnerAuthority *Authority    `json:"new_owner_authority"`
	Extensions        []interface{} `json:"extensions"`
}

func (op *RequestAccountRecoveryOperation) Type() OpType {
	return TypeRequestAccountRecovery
}

func (op *RequestAccountRecoveryOperation) Data() interface{} {
	return op
}

// struct RecoverAccountOperation{}
type RecoverAccountOperation struct {
	AccountToRecover     string        `json:"account_to_recover"`
	NewOwnerAuthority    *Authority    `json:"new_owner_authority"`
	RecentOwnerAuthority *Authority    `json:"recent_owner_authority"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *RecoverAccountOperation) Type() OpType {
	return TypeRecoverAccount
}

func (op *RecoverAccountOperation) Data() interface{} {
	return op
}

// struct ChangeRecoveryAccountOperation{}
type ChangeRecoveryAccountOperation struct {
	AccountToRecover   string        `json:"account_to_recover"`
	NewRecoveryAccount string        `json:"new_recovery_account"`
	Extensions         []interface{} `json:"extensions"`
}

func (op *ChangeRecoveryAccountOperation) Type() OpType {
	return TypeChangeRecoveryAccount
}

func (op *ChangeRecoveryAccountOperation) Data() interface{} {
	return op
}

// struct EscrowTransferOperation{}
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

func (op *EscrowTransferOperation) Type() OpType {
	return TypeEscrowTransfer
}

func (op *EscrowTransferOperation) Data() interface{} {
	return op
}

// struct EscrowDisputeOperation{}
type EscrowDisputeOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowID uint32 `json:"escrow_id"`
}

func (op *EscrowDisputeOperation) Type() OpType {
	return TypeEscrowDispute
}

func (op *EscrowDisputeOperation) Data() interface{} {
	return op
}

// struct EscrowReleaseOperation{}
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

func (op *EscrowReleaseOperation) Type() OpType {
	return TypeEscrowRelease
}

func (op *EscrowReleaseOperation) Data() interface{} {
	return op
}

// struct POW2Operation{}
type POW2Operation struct {
	Input      *POW2Input `json:"input"`
	PowSummary uint32     `json:"pow_summary"`
}

func (op *POW2Operation) Type() OpType {
	return TypePOW2
}

func (op *POW2Operation) Data() interface{} {
	return op
}

// struct EscrowApproveOperation{}
type EscrowApproveOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowID uint32 `json:"escrow_id"`
	Approve  bool   `json:"approve"`
}

func (op *EscrowApproveOperation) Type() OpType {
	return TypeEscrowApprove
}

func (op *EscrowApproveOperation) Data() interface{} {
	return op
}

// struct TransferToSavingsOperation{}
type TransferToSavingsOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *TransferToSavingsOperation) Type() OpType {
	return TypeTransferToSavings
}

func (op *TransferToSavingsOperation) Data() interface{} {
	return op
}

// struct TransferFromSavingsOperation{}
type TransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestID uint32 `json:"request_id"`
	To        string `json:"to"`
	Amount    *Asset `json:"amount"`
	Memo      string `json:"memo"`
}

func (op *TransferFromSavingsOperation) Type() OpType {
	return TypeTransferFromSavings
}

func (op *TransferFromSavingsOperation) Data() interface{} {
	return op
}

// struct CancelTransferFromSavingsOperation{}
type CancelTransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestID uint32 `json:"request_id"`
}

func (op *CancelTransferFromSavingsOperation) Type() OpType {
	return TypeCancelTransferFromSavings
}

func (op *CancelTransferFromSavingsOperation) Data() interface{} {
	return op
}

// struct CustomBinaryOperation{}
type CustomBinaryOperation struct {
	RequiredOwnerAuths   []string `json:"required_owner_auths"`
	RequiredActiveAuths  []string `json:"required_active_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	RequiredAuths        []string `json:"required_auths"`
	ID                   string   `json:"id"`
	Datas                []byte   `json:"data"`
}

func (op *CustomBinaryOperation) Type() OpType {
	return TypeCustomBinary
}

func (op *CustomBinaryOperation) Data() interface{} {
	return op
}

// struct DeclineVotingRightsOperation{}
type DeclineVotingRightsOperation struct {
	Account string `json:"account"`
	Decline bool   `json:"decline"`
}

func (op *DeclineVotingRightsOperation) Type() OpType {
	return TypeDeclineVotingRights
}

func (op *DeclineVotingRightsOperation) Data() interface{} {
	return op
}

// struct ResetAccountOperation{}
type ResetAccountOperation struct {
	ResetAccount      string     `json:"reset_account"`
	AccountToReset    string     `json:"Account_to_reset"`
	NewOwnerAuthority *Authority `json:"new_owner_authority"`
}

func (op *ResetAccountOperation) Type() OpType {
	return TypeResetAccount
}

func (op *ResetAccountOperation) Data() interface{} {
	return op
}

// struct SetResetAccountOperation{}
type SetResetAccountOperation struct {
	Account             string `json:"account"`
	CurrentResetAccount string `json:"current_reset_account"`
	ResetAccount        string `json:"reset_account"`
}

func (op *SetResetAccountOperation) Type() OpType {
	return TypeSetResetAccount
}

func (op *SetResetAccountOperation) Data() interface{} {
	return op
}

// struct DelegateVestingSharesOperation{}
type DelegateVestingSharesOperation struct {
	Delegator     string `json:"delegator"`
	Delegatee     string `json:"delegatee"`
	VestingShares *Asset `json:"vesting_shares"`
}

func (op *DelegateVestingSharesOperation) Type() OpType {
	return TypeDelegateVestingShares
}

func (op *DelegateVestingSharesOperation) Data() interface{} {
	return op
}

// struct AccountCreateWithDelegationOperation{}
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

func (op *AccountCreateWithDelegationOperation) Type() OpType {
	return TypeAccountCreateWithDelegation
}

func (op *AccountCreateWithDelegationOperation) Data() interface{} {
	return op
}

// struct AccountMetadataOperation{}
type AccountMetadataOperation struct {
	Account      string           `json:"account"`
	JSONMetadata *AccountMetadata `json:"json_metadata"`
}

func (op *AccountMetadataOperation) Type() OpType {
	return TypeAccountMetadata
}

func (op *AccountMetadataOperation) Data() interface{} {
	return op
}

// struct ProposalCreateOperation{}
type ProposalCreateOperation struct {
	Author             string              `json:"author"`
	Title              string              `json:"title"`
	Memo               string              `json:"memo"`
	ProposedOperations []ProposalOperation `json:"proposed_operations"`
	ExpirationTime     *Time               `json:"expiration_time"`
	ReviewPeriodTime   *Time               `json:"review_period_time"`
	Extensions         []interface{}       `json:"extensions"`
}

func (op *ProposalCreateOperation) Type() OpType {
	return TypeProposalCreate
}

func (op *ProposalCreateOperation) Data() interface{} {
	return op
}

// struct ProposalUpdateOperation{}
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

func (op *ProposalUpdateOperation) Type() OpType {
	return TypeProposalUpdate
}

func (op *ProposalUpdateOperation) Data() interface{} {
	return op
}

// struct ProposalDeleteOperation{}
type ProposalDeleteOperation struct {
	Author     string        `json:"author"`
	Title      string        `json:"title"`
	Requester  string        `json:"requester"`
	Extensions []interface{} `json:"extensions"`
}

func (op *ProposalDeleteOperation) Type() OpType {
	return TypeProposalDelete
}

func (op *ProposalDeleteOperation) Data() interface{} {
	return op
}

// struct ChainPropertiesUpdateOperation{}
type ChainPropertiesUpdateOperation struct {
	Owner string        `json:"owner"`
	Props []interface{} `json:"props"`
}

func (op *ChainPropertiesUpdateOperation) Type() OpType {
	return TypeChainPropertiesUpdate
}

func (op *ChainPropertiesUpdateOperation) Data() interface{} {
	return op
}

//Virtual Operation
// struct FillConvertRequestOperation{}
type FillConvertRequestOperation struct {
	Owner     string `json:"owner"`
	Requestid uint32 `json:"requestid"`
	AmountIn  *Asset `json:"amount_in"`
	AmountOut *Asset `json:"amount_out"`
}

func (op *FillConvertRequestOperation) Type() OpType {
	return TypeFillConvertRequest
}

func (op *FillConvertRequestOperation) Data() interface{} {
	return op
}

// struct AuthorRewardOperation{}
type AuthorRewardOperation struct {
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	SbdPayout     *Asset `json:"sbd_payout"`
	SteemPayout   *Asset `json:"steem_payout"`
	VestingPayout *Asset `json:"vesting_payout"`
}

func (op *AuthorRewardOperation) Type() OpType {
	return TypeAuthorReward
}

func (op *AuthorRewardOperation) Data() interface{} {
	return op
}

// struct CurationRewardOperation{}
type CurationRewardOperation struct {
	Curator         string `json:"curator"`
	Reward          *Asset `json:"reward"`
	CommentAuthor   string `json:"comment_author"`
	CommentPermlink string `json:"comment_permlink"`
}

func (op *CurationRewardOperation) Type() OpType {
	return TypeCurationReward
}

func (op *CurationRewardOperation) Data() interface{} {
	return op
}

// struct CommentRewardOperation{}
type CommentRewardOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Payout   *Asset `json:"payout"`
}

func (op *CommentRewardOperation) Type() OpType {
	return TypeCommentReward
}

func (op *CommentRewardOperation) Data() interface{} {
	return op
}

// struct LiquidityRewardOperation{}
type LiquidityRewardOperation struct {
	Owner  string `json:"owner"`
	Payout *Asset `json:"payout"`
}

func (op *LiquidityRewardOperation) Type() OpType {
	return TypeLiquidityReward
}

func (op *LiquidityRewardOperation) Data() interface{} {
	return op
}

// struct InterestOperation{}
type InterestOperation struct {
	Owner    string `json:"owner"`
	Interest *Asset `json:"interest"`
}

func (op *InterestOperation) Type() OpType {
	return TypeInterest
}

func (op *InterestOperation) Data() interface{} {
	return op
}

// struct FillVestingWithdrawOperation{}
type FillVestingWithdrawOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Withdrawn   *Asset `json:"withdrawn"`
	Deposited   *Asset `json:"deposited"`
}

func (op *FillVestingWithdrawOperation) Type() OpType {
	return TypeFillVestingWithdraw
}

func (op *FillVestingWithdrawOperation) Data() interface{} {
	return op
}

// struct FillOrderOperation{}
type FillOrderOperation struct {
	CurrentOwner   string `json:"current_owner"`
	CurrentOrderid uint32 `json:"current_orderid"`
	CurrentPays    *Asset `json:"current_pays"`
	OpenOwner      string `json:"open_owner"`
	OpenOrderid    uint32 `json:"open_orderid"`
	OpenPays       *Asset `json:"open_pays"`
}

func (op *FillOrderOperation) Type() OpType {
	return TypeFillOrder
}

func (op *FillOrderOperation) Data() interface{} {
	return op
}

// struct ShutdownWitnessOperation{}
type ShutdownWitnessOperation struct {
	Owner string `json:"owner"`
}

func (op *ShutdownWitnessOperation) Type() OpType {
	return TypeShutdownWitness
}

func (op *ShutdownWitnessOperation) Data() interface{} {
	return op
}

// struct FillTransferFromSavingsOperation{}
type FillTransferFromSavingsOperation struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    *Asset `json:"amount"`
	RequestID uint32 `json:"request_id"`
	Memo      string `json:"memo"`
}

func (op *FillTransferFromSavingsOperation) Type() OpType {
	return TypeFillTransferFromSavings
}

func (op *FillTransferFromSavingsOperation) Data() interface{} {
	return op
}

// struct HardforkOperation{}
type HardforkOperation struct {
	HardforkID uint32 `json:"hardfork_id"`
}

func (op *HardforkOperation) Type() OpType {
	return TypeHardfork
}

func (op *HardforkOperation) Data() interface{} {
	return op
}

// struct CommentPayoutUpdateOperation{}
type CommentPayoutUpdateOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *CommentPayoutUpdateOperation) Type() OpType {
	return TypeCommentPayoutUpdate
}

func (op *CommentPayoutUpdateOperation) Data() interface{} {
	return op
}

// struct CommentBenefactorRewardOperation{}
type CommentBenefactorRewardOperation struct {
	Benefactor string `json:"benefactor"`
	Author     string `json:"author"`
	Permlink   string `json:"permlink"`
	Reward     *Asset `json:"reward"`
}

func (op *CommentBenefactorRewardOperation) Type() OpType {
	return TypeCommentBenefactorReward
}

func (op *CommentBenefactorRewardOperation) Data() interface{} {
	return op
}

// struct ReturnVestingDelegationOperation{}
type ReturnVestingDelegationOperation struct {
	Account       string `json:"account"`
	VestingShares *Asset `json:"vesting_shares"`
}

func (op *ReturnVestingDelegationOperation) Type() OpType {
	return TypeReturnVestingDelegation
}

func (op *ReturnVestingDelegationOperation) Data() interface{} {
	return op
}

// struct UnknownOperation{}
type UnknownOperation struct {
	kind OpType
	data *json.RawMessage
}

func (op *UnknownOperation) Type() OpType {
	return op.kind
}

func (op *UnknownOperation) Data() interface{} {
	return op.data
}
