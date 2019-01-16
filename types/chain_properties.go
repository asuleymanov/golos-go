package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//ChainProperties is an additional structure used by other structures.
type ChainPropertiesOLD struct {
	AccountCreationFee *Asset `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

//ChainProperties is an additional structure used by other structures.
type ChainProperties struct {
	AccountCreationFee              *Asset `json:"account_creation_fee"`
	MaximumBlockSize                uint32 `json:"maximum_block_size"`
	SBDInterestRate                 uint16 `json:"sbd_interest_rate"`
	CreateAccountMinGolosFee        *Asset `json:"create_account_min_golos_fee"`
	CreateAccountMinDelegation      *Asset `json:"create_account_min_delegation"`
	CreateAccountDelegationTime     uint32 `json:"create_account_delegation_time"`
	MinDelegation                   *Asset `json:"min_delegation"`
	MaxReferralInterestRate         uint16 `json:"max_referral_interest_rate"`
	MaxReferralTermSec              uint32 `json:"max_referral_term_sec"`
	MinReferralBreakFee             *Asset `json:"min_referral_break_fee"`
	MaxReferralBreakFee             *Asset `json:"max_referral_break_fee"`
	PostsWindow                     uint16 `json:"posts_window"`
	PostsPerWindow                  uint16 `json:"posts_per_window"`
	CommentsWindow                  uint16 `json:"comments_window"`
	CommentsPerWindow               uint16 `json:"comments_per_window"`
	VotesWindow                     uint16 `json:"votes_window"`
	VotesPerWindow                  uint16 `json:"votes_per_window"`
	AuctionWindowSize               uint16 `json:"auction_window_size"`
	MaxDelegatedVestingInterestRate uint16 `json:"max_delegated_vesting_interest_rate"`
	CustomOpsBandwidthMultiplier    uint16 `json:"custom_ops_bandwidth_multiplier"`
	MinCurationPercent              uint16 `json:"min_curation_percent"`
	MaxCurationPercent              uint16 `json:"max_curation_percent"`
	CurationRewardCurve             uint64 `json:"curation_reward_curve"`
	AllowDistributeAuctionReward    bool   `json:"allow_distribute_auction_reward"`
	AllowReturnAuctionRewardToFund  bool   `json:"allow_return_auction_reward_to_fund"`
}

//MarshalTransaction is a function of converting type ChainPropertiesOLD to bytes.
func (cp *ChainPropertiesOLD) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	return enc.Err()
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
