package witness

import (
	"github.com/asuleymanov/golos-go/types"
)

//ChainProperties
type ChainProperties struct {
	AccountCreationFee          *types.Asset `json:"account_creation_fee"`
	MaximumBlockSize            uint32       `json:"maximum_block_size"`
	SBDInterestRate             uint16       `json:"sbd_interest_rate"`
	CreateAccountMinGolosFee    *types.Asset `json:"create_account_min_golos_fee"`
	CreateAccountMinDelegation  *types.Asset `json:"create_account_min_delegation"`
	CreateAccountDelegationTime uint32       `json:"create_account_delegation_time"`
	MinDelegation               *types.Asset `json:"min_delegation"`
}

//CurrentMedianHistoryPrice
type CurrentMedianHistoryPrice struct {
	Base  *types.Asset `json:"base"`
	Quote *types.Asset `json:"quote"`
}

//FeedHistory
type FeedHistory struct {
	ID                   *types.Int                   `json:"id"`
	CurrentMedianHistory *CurrentMedianHistoryPrice   `json:"current_median_history"`
	PriceHistory         []*CurrentMedianHistoryPrice `json:"price_history"`
}

//WitnessSchedule
type WitnessSchedule struct {
	ID                            *types.Int       `json:"id"`
	CurrentVirtualTime            string           `json:"current_virtual_time"`
	NextShuffleBlockNum           *types.Int       `json:"next_shuffle_block_num"`
	CurrentShuffledWitnesses      string           `json:"current_shuffled_witnesses"`
	NumScheduledWitnesses         *types.Int       `json:"num_scheduled_witnesses"`
	Top19Weight                   *types.Int       `json:"top19_weight"`
	TimeshareWeight               *types.Int       `json:"timeshare_weight"`
	MinerWeight                   *types.Int       `json:"miner_weight"`
	WitnessPayNormalizationFactor *types.Int       `json:"witness_pay_normalization_factor"`
	MedianProps                   *ChainProperties `json:"median_props"`
	MajorityVersion               string           `json:"majority_version"`
}

//Witness
type Witness struct {
	ID                    *types.Int                 `json:"id"`
	Owner                 string                     `json:"owner"`
	Created               *types.Time                `json:"created"`
	URL                   string                     `json:"url"`
	Votes                 string                     `json:"votes"`
	VirtualLastUpdate     string                     `json:"virtual_last_update"`
	VirtualPosition       string                     `json:"virtual_position"`
	VirtualScheduledTime  string                     `json:"virtual_scheduled_time"`
	TotalMissed           *types.Int                 `json:"total_missed"`
	LastAslot             *types.Int                 `json:"last_aslot"`
	LastConfirmedBlockNum *types.Int                 `json:"last_confirmed_block_num"`
	PowWorker             *types.Int                 `json:"pow_worker"`
	SigningKey            string                     `json:"signing_key"`
	Props                 *ChainProperties           `json:"props"`
	SbdExchangeRate       *CurrentMedianHistoryPrice `json:"sbd_exchange_rate"`
	LastSbdExchangeUpdate *types.Time                `json:"last_sbd_exchange_update"`
	LastWork              string                     `json:"last_work"`
	RunningVersion        string                     `json:"running_version"`
	HardforkVersionVote   string                     `json:"hardfork_version_vote"`
	HardforkTimeVote      *types.Time                `json:"hardfork_time_vote"`
}
