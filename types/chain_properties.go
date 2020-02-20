package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//ChainProperties is an additional structure used by other structures.
type ChainProperties struct {
	AccountCreationFee              Asset  `json:"account_creation_fee"`
	MaximumBlockSize                uint32 `json:"maximum_block_size"`
	SBDInterestRate                 uint16 `json:"sbd_interest_rate"`
	CreateAccountMinGolosFee        Asset  `json:"create_account_min_golos_fee,omitempty"`
	CreateAccountMinDelegation      Asset  `json:"create_account_min_delegation,omitempty"`
	CreateAccountDelegationTime     uint32 `json:"create_account_delegation_time,omitempty"`
	MinDelegation                   Asset  `json:"min_delegation,omitempty"`
	MaxReferralInterestRate         uint16 `json:"max_referral_interest_rate,omitempty"`
	MaxReferralTermSec              uint32 `json:"max_referral_term_sec,omitempty"`
	MinReferralBreakFee             Asset  `json:"min_referral_break_fee,omitempty"`
	MaxReferralBreakFee             Asset  `json:"max_referral_break_fee,omitempty"`
	PostsWindow                     uint16 `json:"posts_window,omitempty"`
	PostsPerWindow                  uint16 `json:"posts_per_window,omitempty"`
	CommentsWindow                  uint16 `json:"comments_window,omitempty"`
	CommentsPerWindow               uint16 `json:"comments_per_window,omitempty"`
	VotesWindow                     uint16 `json:"votes_window,omitempty"`
	VotesPerWindow                  uint16 `json:"votes_per_window,omitempty"`
	AuctionWindowSize               uint16 `json:"auction_window_size,omitempty"`
	MaxDelegatedVestingInterestRate uint16 `json:"max_delegated_vesting_interest_rate,omitempty"`
	CustomOpsBandwidthMultiplier    uint16 `json:"custom_ops_bandwidth_multiplier,omitempty"`
	MinCurationPercent              uint16 `json:"min_curation_percent,omitempty"`
	MaxCurationPercent              uint16 `json:"max_curation_percent,omitempty"`
	CurationRewardCurve             uint64 `json:"curation_reward_curve,omitempty"`
	AllowDistributeAuctionReward    bool   `json:"allow_distribute_auction_reward,omitempty"`
	AllowReturnAuctionRewardToFund  bool   `json:"allow_return_auction_reward_to_fund,omitempty"`
	WorkerRewardPercent             uint16 `json:"worker_reward_percent,omitempty"`
	WitnessRewardPercent            uint16 `json:"witness_reward_percent,omitempty"`
	VestingRewardPercent            uint16 `json:"vesting_reward_percent,omitempty"`
	WorkerRequestCreationFee        Asset  `json:"worker_request_creation_fee,omitempty"`
	WorkerRequestApproveMinPercent  uint16 `json:"worker_request_approve_min_percent,omitempty"`
	SbdDebtConvertRate              uint16 `json:"sbd_debt_convert_rate,omitempty"`
	VoteRegenerationPerDay          uint32 `json:"vote_regeneration_per_day,omitempty"`
	WitnessSkippingResetTime        uint32 `json:"witness_skipping_reset_time,omitempty"`
	WitnessIdlenessTime             uint32 `json:"witness_idleness_time,omitempty"`
	AccountIdlenessTime             uint32 `json:"account_idleness_time,omitempty"`
}

//MarshalTransaction is a function of converting type ChainProperties to bytes.
func (cp *ChainProperties) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	enc.Encode(cp.CreateAccountMinGolosFee)
	enc.Encode(cp.CreateAccountMinDelegation)
	enc.Encode(cp.CreateAccountDelegationTime)
	enc.Encode(cp.MinDelegation)
	enc.Encode(cp.MaxReferralInterestRate)
	enc.Encode(cp.MaxReferralTermSec)
	enc.Encode(cp.MinReferralBreakFee)
	enc.Encode(cp.MaxReferralBreakFee)
	enc.Encode(cp.PostsWindow)
	enc.Encode(cp.PostsPerWindow)
	enc.Encode(cp.CommentsWindow)
	enc.Encode(cp.CommentsPerWindow)
	enc.Encode(cp.VotesWindow)
	enc.Encode(cp.VotesPerWindow)
	enc.Encode(cp.AuctionWindowSize)
	enc.Encode(cp.MaxDelegatedVestingInterestRate)
	enc.Encode(cp.CustomOpsBandwidthMultiplier)
	enc.Encode(cp.MinCurationPercent)
	enc.Encode(cp.MaxCurationPercent)
	enc.Encode(cp.CurationRewardCurve)
	enc.Encode(cp.AllowDistributeAuctionReward)
	enc.Encode(cp.AllowReturnAuctionRewardToFund)
	return enc.Err()
}

// Очень старое возможно даже можно будет удалить.

//ChainProperties is an additional structure used by other structures.
type ChainPropertiesOLD struct {
	AccountCreationFee *Asset `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

//MarshalTransaction is a function of converting type ChainPropertiesOLD to bytes.
func (cp *ChainPropertiesOLD) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	return enc.Err()
}
