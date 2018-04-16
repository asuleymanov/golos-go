package database

import (
	"github.com/asuleymanov/golos-go/types"
)

//BlockHeader
type BlockHeader struct {
	Number                uint32        `json:"-"`
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	Witness               string        `json:"witness"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Extensions            []interface{} `json:"extensions"`
}

//Block
type Block struct {
	Number                uint32               `json:"-"`
	Timestamp             *types.Time          `json:"timestamp"`
	Witness               string               `json:"witness"`
	WitnessSignature      string               `json:"witness_signature"`
	TransactionMerkleRoot string               `json:"transaction_merkle_root"`
	Previous              string               `json:"previous"`
	Extensions            [][]interface{}      `json:"extensions"`
	Transactions          []*types.Transaction `json:"transactions"`
}

//Config
type Config struct {
	SteemitBuildTestnet                   bool         `json:"STEEMIT_BUILD_TESTNET"`
	GrapheneCurrentDBVersion              string       `json:"GRAPHENE_CURRENT_DB_VERSION"`
	SbdSymbol                             *types.Int   `json:"SBD_SYMBOL"`
	Steemit100Percent                     int          `json:"STEEMIT_100_PERCENT"`
	Steemit1Percent                       *types.Int   `json:"STEEMIT_1_PERCENT"`
	SteemitAddressPrefix                  string       `json:"STEEMIT_ADDRESS_PREFIX"`
	SteemitAprPercentMultiplyPerBlock     string       `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	SteemitAprPercentMultiplyPerHour      string       `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR"`
	SteemitAprPercentMultiplyPerRound     string       `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND"`
	SteemitAprPercentShiftPerBlock        *types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK"`
	SteemitAprPercentShiftPerHour         *types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_HOUR"`
	SteemitAprPercentShiftPerRound        *types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_ROUND"`
	SteemitBandwidthAverageWindowSeconds  *types.Int   `json:"STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	SteemitBandwidthPrecision             *types.Int   `json:"STEEMIT_BANDWIDTH_PRECISION"`
	SteemitBlockchainPrecision            *types.Int   `json:"STEEMIT_BLOCKCHAIN_PRECISION"`
	SteemitBlockchainPrecisionDigits      *types.Int   `json:"STEEMIT_BLOCKCHAIN_PRECISION_DIGITS"`
	SteemitBlockchainHardforkVersion      string       `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	SteemitBlockchainVersion              string       `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	SteemitBlockInterval                  uint         `json:"STEEMIT_BLOCK_INTERVAL"`
	SteemitBlocksPerDay                   *types.Int   `json:"STEEMIT_BLOCKS_PER_DAY"`
	SteemitBlocksPerHour                  *types.Int   `json:"STEEMIT_BLOCKS_PER_HOUR"`
	SteemitBlocksPerYear                  *types.Int   `json:"STEEMIT_BLOCKS_PER_YEAR"`
	SteemitCashoutWindowSeconds           *types.Int   `json:"STEEMIT_CASHOUT_WINDOW_SECONDS"`
	SteemitChainID                        string       `json:"STEEMIT_CHAIN_ID"`
	SteemitContentAprPercent              *types.Int   `json:"STEEMIT_CONTENT_APR_PERCENT"`
	SteemitConversionDelay                string       `json:"STEEMIT_CONVERSION_DELAY"`
	SteemitCurateAprPercent               *types.Int   `json:"STEEMIT_CURATE_APR_PERCENT"`
	SteemitDefaultSbdInterestRate         *types.Int   `json:"STEEMIT_DEFAULT_SBD_INTEREST_RATE"`
	SteemitFeedHistoryWindow              *types.Int   `json:"STEEMIT_FEED_HISTORY_WINDOW"`
	SteemitFeedIntervalBlocks             *types.Int   `json:"STEEMIT_FEED_INTERVAL_BLOCKS"`
	SteemitFreeTransactionsWithNewAccount *types.Int   `json:"STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT"`
	SteemitGenesisTime                    string       `json:"STEEMIT_GENESIS_TIME"`
	SteemitHardforkRequiredWitnesses      *types.Int   `json:"STEEMIT_HARDFORK_REQUIRED_WITNESSES"`
	SteemitInitMinerName                  string       `json:"STEEMIT_INIT_MINER_NAME"`
	SteemitInitPublicKeyStr               string       `json:"STEEMIT_INIT_PUBLIC_KEY_STR"`
	SteemitInitSupply                     *types.Int   `json:"STEEMIT_INIT_SUPPLY"`
	SteemitInitTime                       string       `json:"STEEMIT_INIT_TIME"`
	SteemitIrreversibleThreshold          *types.Int   `json:"STEEMIT_IRREVERSIBLE_THRESHOLD"`
	SteemitLiquidityAprPercent            *types.Int   `json:"STEEMIT_LIQUIDITY_APR_PERCENT"`
	SteemitLiquidityRewardBlocks          *types.Int   `json:"STEEMIT_LIQUIDITY_REWARD_BLOCKS"`
	SteemitLiquidityRewardPeriodSec       *types.Int   `json:"STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemitLiquidityTimeoutSec            string       `json:"STEEMIT_LIQUIDITY_TIMEOUT_SEC"`
	SteemitMaxAccountNameLength           *types.Int   `json:"STEEMIT_MAX_ACCOUNT_NAME_LENGTH"`
	SteemitMaxAccountWitnessVotes         *types.Int   `json:"STEEMIT_MAX_ACCOUNT_WITNESS_VOTES"`
	SteemitMaxAssetWhitelistAuthorities   *types.Int   `json:"STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	SteemitMaxAuthorityMembership         *types.Int   `json:"STEEMIT_MAX_AUTHORITY_MEMBERSHIP"`
	SteemitMaxBlockSize                   *types.Int   `json:"STEEMIT_MAX_BLOCK_SIZE"`
	SteemitMaxCashoutWindowSeconds        *types.Int   `json:"STEEMIT_MAX_CASHOUT_WINDOW_SECONDS"`
	SteemitMaxCommentDepth                *types.Int   `json:"STEEMIT_MAX_COMMENT_DEPTH"`
	SteemitMaxFeedAge                     *types.Int   `json:"STEEMIT_MAX_FEED_AGE"`
	SteemitMaxInstanceID                  string       `json:"STEEMIT_MAX_INSTANCE_ID"`
	SteemitMaxMemoSize                    *types.Int   `json:"STEEMIT_MAX_MEMO_SIZE"`
	SteemitMaxWitnesses                   *types.Int   `json:"STEEMIT_MAX_WITNESSES"`
	SteemitMaxMinerWitnesses              *types.Int   `json:"STEEMIT_MAX_MINER_WITNESSES"`
	SteemitMaxProxyRecursionDepth         *types.Int   `json:"STEEMIT_MAX_PROXY_RECURSION_DEPTH"`
	SteemitMaxRationDecayRate             *types.Int   `json:"STEEMIT_MAX_RATION_DECAY_RATE"`
	SteemitMaxReserveRatio                *types.Int   `json:"STEEMIT_MAX_RESERVE_RATIO"`
	SteemitMaxRunnerWitnesses             *types.Int   `json:"STEEMIT_MAX_RUNNER_WITNESSES"`
	SteemitMaxShareSupply                 string       `json:"STEEMIT_MAX_SHARE_SUPPLY"`
	SteemitMaxSigCheckDepth               *types.Int   `json:"STEEMIT_MAX_SIG_CHECK_DEPTH"`
	SteemitMaxTimeUntilExpiration         *types.Int   `json:"STEEMIT_MAX_TIME_UNTIL_EXPIRATION"`
	SteemitMaxTransactionSize             *types.Int   `json:"STEEMIT_MAX_TRANSACTION_SIZE"`
	SteemitMaxUndoHistory                 *types.Int   `json:"STEEMIT_MAX_UNDO_HISTORY"`
	SteemitMaxURLLength                   *types.Int   `json:"STEEMIT_MAX_URL_LENGTH"`
	SteemitMaxVoteChanges                 *types.Int   `json:"STEEMIT_MAX_VOTE_CHANGES"`
	SteemitMaxVotedWitnesses              *types.Int   `json:"STEEMIT_MAX_VOTED_WITNESSES"`
	SteemitMaxWithdrawRoutes              *types.Int   `json:"STEEMIT_MAX_WITHDRAW_ROUTES"`
	SteemitMaxWitnessURLLength            *types.Int   `json:"STEEMIT_MAX_WITNESS_URL_LENGTH"`
	SteemitMinAccountCreationFee          *types.Int   `json:"STEEMIT_MIN_ACCOUNT_CREATION_FEE"`
	SteemitMinAccountNameLength           *types.Int   `json:"STEEMIT_MIN_ACCOUNT_NAME_LENGTH"`
	SteemitMinBlockSizeLimit              *types.Int   `json:"STEEMIT_MIN_BLOCK_SIZE_LIMIT"`
	SteemitMinContentReward               *types.Asset `json:"STEEMIT_MIN_CONTENT_REWARD"`
	SteemitMinCurateReward                *types.Asset `json:"STEEMIT_MIN_CURATE_REWARD"`
	SteemitMinerAccount                   string       `json:"STEEMIT_MINER_ACCOUNT"`
	SteemitMinerPayPercent                *types.Int   `json:"STEEMIT_MINER_PAY_PERCENT"`
	SteemitMinFeeds                       *types.Int   `json:"STEEMIT_MIN_FEEDS"`
	SteemitMiningReward                   *types.Asset `json:"STEEMIT_MINING_REWARD"`
	SteemitMiningTime                     string       `json:"STEEMIT_MINING_TIME"`
	SteemitMinLiquidityReward             *types.Asset `json:"STEEMIT_MIN_LIQUIDITY_REWARD"`
	SteemitMinLiquidityRewardPeriodSec    *types.Int   `json:"STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemitMinPayoutSbd                   *types.Asset `json:"STEEMIT_MIN_PAYOUT_SBD"`
	SteemitMinPowReward                   *types.Asset `json:"STEEMIT_MIN_POW_REWARD"`
	SteemitMinProducerReward              *types.Asset `json:"STEEMIT_MIN_PRODUCER_REWARD"`
	SteemitMinRation                      *types.Int   `json:"STEEMIT_MIN_RATION"`
	SteemitMinTransactionExpirationLimit  *types.Int   `json:"STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	SteemitMinTransactionSizeLimit        *types.Int   `json:"STEEMIT_MIN_TRANSACTION_SIZE_LIMIT"`
	SteemitMinUndoHistory                 *types.Int   `json:"STEEMIT_MIN_UNDO_HISTORY"`
	SteemitNullAccount                    string       `json:"STEEMIT_NULL_ACCOUNT"`
	SteemitNumInitMiners                  *types.Int   `json:"STEEMIT_NUM_INIT_MINERS"`
	SteemitPowAprPercent                  *types.Int   `json:"STEEMIT_POW_APR_PERCENT"`
	SteemitProducerAprPercent             *types.Int   `json:"STEEMIT_PRODUCER_APR_PERCENT"`
	SteemitProxyToSelfAccount             string       `json:"STEEMIT_PROXY_TO_SELF_ACCOUNT"`
	SteemitSbdInterestCompoundIntervalSec *types.Int   `json:"STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC"`
	SteemitSecondsPerYear                 *types.Int   `json:"STEEMIT_SECONDS_PER_YEAR"`
	SteemitReverseAuctionWindowSeconds    *types.Int   `json:"STEEMIT_REVERSE_AUCTION_WINDOW_SECONDS"`
	SteemitStartMinerVotingBlock          *types.Int   `json:"STEEMIT_START_MINER_VOTING_BLOCK"`
	SteemitStartVestingBlock              *types.Int   `json:"STEEMIT_START_VESTING_BLOCK"`
	SteemitSymbol                         string       `json:"STEEMIT_SYMBOL"`
	SteemitTempAccount                    string       `json:"STEEMIT_TEMP_ACCOUNT"`
	SteemitUpvoteLockout                  *types.Int   `json:"STEEMIT_UPVOTE_LOCKOUT"`
	SteemitVestingWithdrawIntervals       *types.Int   `json:"STEEMIT_VESTING_WITHDRAW_INTERVALS"`
	SteemitVestingWithdrawIntervalSeconds *types.Int   `json:"STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	SteemitVoteChangeLockoutPeriod        *types.Int   `json:"STEEMIT_VOTE_CHANGE_LOCKOUT_PERIOD"`
	SteemitVoteRegenerationSeconds        int          `json:"STEEMIT_VOTE_REGENERATION_SECONDS"`
	SteemSymbol                           string       `json:"STEEM_SYMBOL"`
	VestsSymbol                           string       `json:"VESTS_SYMBOL"`
	BlockchainName                        string       `json:"BLOCKCHAIN_NAME"`
}

//DynamicGlobalProperties
type DynamicGlobalProperties struct {
	Time                     *types.Time  `json:"time"`
	TotalPow                 *types.Int   `json:"total_pow"`
	NumPowWitnesses          *types.Int   `json:"num_pow_witnesses"`
	CurrentReserveRatio      *types.Int   `json:"current_reserve_ratio"`
	ID                       *types.ID    `json:"id"`
	CurrentSupply            *types.Asset `json:"current_supply"`
	CurrentSBDSupply         *types.Asset `json:"current_sbd_supply"`
	MaximumBlockSize         *types.Int   `json:"maximum_block_size"`
	RecentSlotsFilled        *types.Int   `json:"recent_slots_filled"`
	CurrentWitness           string       `json:"current_witness"`
	TotalRewardShares2       *types.Int   `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int   `json:"average_block_size"`
	CurrentAslot             *types.Int   `json:"current_aslot"`
	LastIrreversibleBlockNum uint32       `json:"last_irreversible_block_num"`
	TotalVestingShares       *types.Asset `json:"total_vesting_shares"`
	TotalVersingFundSteem    *types.Asset `json:"total_vesting_fund_steem"`
	HeadBlockID              string       `json:"head_block_id"`
	HeadBlockNumber          uint32       `json:"head_block_number"`
	VirtualSupply            *types.Asset `json:"virtual_supply"`
	ConfidentialSupply       *types.Asset `json:"confidential_supply"`
	ConfidentialSBDSupply    *types.Asset `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     *types.Asset `json:"total_reward_fund_steem"`
	TotalActivityFundSteem   string       `json:"total_activity_fund_steem"`
	TotalActivityFundShares  *types.Int   `json:"total_activity_fund_shares"`
	SBDInterestRate          *types.Int   `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      *types.Int   `json:"max_virtual_bandwidth"`
}

//ChainProperties
type ChainProperties struct {
	AccountCreationFee *types.Asset `json:"account_creation_fee"`
	MaximumBlockSize   *types.Int   `json:"maximum_block_size"`
	SbdInterestRate    *types.Int   `json:"sbd_interest_rate"`
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

//NextScheduledHardfork
type NextScheduledHardfork struct {
	HfVersion string      `json:"hf_version"`
	LiveTime  *types.Time `json:"live_time"`
}

//Account
type Account struct {
	ID                            *types.Int             `json:"id"`
	Name                          string                 `json:"name"`
	Owner                         *types.Authority       `json:"owner"`
	Active                        *types.Authority       `json:"active"`
	Posting                       *types.Authority       `json:"posting"`
	MemoKey                       string                 `json:"memo_key"`
	JSONMetadata                  *types.AccountMetadata `json:"json_metadata"`
	Proxy                         string                 `json:"proxy"`
	LastOwnerUpdate               *types.Time            `json:"last_owner_update"`
	LastAccountUpdate             *types.Time            `json:"last_account_update"`
	Created                       *types.Time            `json:"created"`
	Mined                         bool                   `json:"mined"`
	OwnerChallenged               bool                   `json:"owner_challenged"`
	ActiveChallenged              bool                   `json:"active_challenged"`
	LastOwnerProved               *types.Time            `json:"last_owner_proved"`
	LastActiveProved              *types.Time            `json:"last_active_proved"`
	RecoveryAccount               string                 `json:"recovery_account"`
	LastAccountRecovery           *types.Time            `json:"last_account_recovery"`
	ResetAccount                  string                 `json:"reset_account"`
	CommentCount                  *types.Int             `json:"comment_count"`
	LifetimeVoteCount             *types.Int             `json:"lifetime_vote_count"`
	PostCount                     *types.Int             `json:"post_count"`
	CanVote                       bool                   `json:"can_vote"`
	VotingPower                   int                    `json:"voting_power"`
	LastVoteTime                  *types.Time            `json:"last_vote_time"`
	Balance                       *types.Asset           `json:"balance"`
	SavingsBalance                *types.Asset           `json:"savings_balance"`
	SbdBalance                    *types.Asset           `json:"sbd_balance"`
	SbdSeconds                    string                 `json:"sbd_seconds"`
	SbdSecondsLastUpdate          *types.Time            `json:"sbd_seconds_last_update"`
	SbdLastInterestPayment        *types.Time            `json:"sbd_last_interest_payment"`
	SavingsSbdBalance             *types.Asset           `json:"savings_sbd_balance"`
	SavingsSbdSeconds             string                 `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   *types.Time            `json:"savings_sbd_seconds_last_update"`
	SavingsSbdLastInterestPayment *types.Time            `json:"savings_sbd_last_interest_payment"`
	SavingsWithdrawRequests       *types.Int             `json:"savings_withdraw_requests"`
	VestingShares                 *types.Asset           `json:"vesting_shares"`
	VestingWithdrawRate           *types.Asset           `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         *types.Time            `json:"next_vesting_withdrawal"`
	Withdrawn                     *types.Int             `json:"withdrawn"`
	ToWithdraw                    *types.Int             `json:"to_withdraw"`
	WithdrawRoutes                *types.Int             `json:"withdraw_routes"`
	CurationRewards               *types.Int             `json:"curation_rewards"`
	PostingRewards                *types.Int             `json:"posting_rewards"`
	ProxiedVsfVotes               []*types.Int           `json:"proxied_vsf_votes"`
	WitnessesVotedFor             *types.Int             `json:"witnesses_voted_for"`
	AverageBandwidth              *types.Int             `json:"average_bandwidth"`
	LifetimeBandwidth             *types.Int64           `json:"lifetime_bandwidth"`
	LastBandwidthUpdate           *types.Time            `json:"last_bandwidth_update"`
	AverageMarketBandwidth        *types.Int             `json:"average_market_bandwidth"`
	LastMarketBandwidthUpdate     *types.Time            `json:"last_market_bandwidth_update"`
	LastPost                      *types.Time            `json:"last_post"`
	LastRootPost                  *types.Time            `json:"last_root_post"`
	PostBandwidth                 *types.Int             `json:"post_bandwidth"`
	NewAverageBandwidth           *types.Int64           `json:"new_average_bandwidth"`
	NewAverageMarketBandwidth     *types.Int64           `json:"new_average_market_bandwidth"`
	VestingBalance                *types.Asset           `json:"vesting_balance"`
	Reputation                    *types.Int64           `json:"reputation"`
	TransferHistory               []interface{}          `json:"transfer_history"`
	MarketHistory                 []interface{}          `json:"market_history"`
	PostHistory                   []interface{}          `json:"post_history"`
	VoteHistory                   []interface{}          `json:"vote_history"`
	OtherHistory                  []interface{}          `json:"other_history"`
	WitnessVotes                  []string               `json:"witness_votes"`
	TagsUsage                     []interface{}          `json:"tags_usage"`
	GuestBloggers                 []interface{}          `json:"guest_bloggers"`
	BlogCategory                  interface{}            `json:"blog_category"`
}

//SavingsWithdraw
type SavingsWithdraw struct {
	ID        *types.ID    `json:"id"`
	From      string       `json:"from"`
	To        string       `json:"to"`
	Memo      string       `json:"memo"`
	RequestID *types.Int   `json:"request_id"`
	Amount    *types.Asset `json:"amount"`
	Complete  *types.Time  `json:"complete"`
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

//ConversionRequests
type ConversionRequests struct {
	ID             *types.Int   `json:"id"`
	Owner          string       `json:"owner"`
	Requestid      *types.Int   `json:"requestid"`
	Amount         *types.Asset `json:"amount"`
	ConversionDate *types.Time  `json:"conversion_date"`
}

type Bandwidth struct {
	ID                  *types.Int   `json:"id"`
	Account             string       `json:"account"`
	Type                string       `json:"type"`
	AverageBandwidth    *types.Int   `json:"average_bandwidth"`
	LifetimeBandwidth   *types.Int64 `json:"lifetime_bandwidth"`
	LastBandwidthUpdate *types.Time  `json:"last_bandwidth_update"`
}

type LookupAccountNames struct {
	ID                            int                    `json:"id"`
	Name                          string                 `json:"name"`
	Owner                         *types.Authority       `json:"owner"`
	Active                        *types.Authority       `json:"active"`
	Posting                       *types.Authority       `json:"posting"`
	MemoKey                       string                 `json:"memo_key"`
	JSONMetadata                  *types.AccountMetadata `json:"json_metadata"`
	Proxy                         string                 `json:"proxy"`
	LastOwnerUpdate               string                 `json:"last_owner_update"`
	LastAccountUpdate             string                 `json:"last_account_update"`
	Created                       string                 `json:"created"`
	Mined                         bool                   `json:"mined"`
	OwnerChallenged               bool                   `json:"owner_challenged"`
	ActiveChallenged              bool                   `json:"active_challenged"`
	LastOwnerProved               string                 `json:"last_owner_proved"`
	LastActiveProved              string                 `json:"last_active_proved"`
	RecoveryAccount               string                 `json:"recovery_account"`
	LastAccountRecovery           string                 `json:"last_account_recovery"`
	ResetAccount                  string                 `json:"reset_account"`
	CommentCount                  *types.Int             `json:"comment_count"`
	LifetimeVoteCount             *types.Int             `json:"lifetime_vote_count"`
	PostCount                     *types.Int             `json:"post_count"`
	CanVote                       bool                   `json:"can_vote"`
	VotingPower                   *types.Int             `json:"voting_power"`
	LastVoteTime                  string                 `json:"last_vote_time"`
	Balance                       *types.Asset           `json:"balance"`
	SavingsBalance                *types.Asset           `json:"savings_balance"`
	SbdBalance                    *types.Asset           `json:"sbd_balance"`
	SbdSeconds                    string                 `json:"sbd_seconds"`
	SbdSecondsLastUpdate          string                 `json:"sbd_seconds_last_update"`
	SbdLastInterestPayment        string                 `json:"sbd_last_interest_payment"`
	SavingsSbdBalance             *types.Asset           `json:"savings_sbd_balance"`
	SavingsSbdSeconds             string                 `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   string                 `json:"savings_sbd_seconds_last_update"`
	SavingsSbdLastInterestPayment string                 `json:"savings_sbd_last_interest_payment"`
	SavingsWithdrawRequests       *types.Int             `json:"savings_withdraw_requests"`
	VestingShares                 *types.Asset           `json:"vesting_shares"`
	VestingWithdrawRate           *types.Asset           `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         string                 `json:"next_vesting_withdrawal"`
	Withdrawn                     *types.Int             `json:"withdrawn"`
	ToWithdraw                    *types.Int             `json:"to_withdraw"`
	WithdrawRoutes                *types.Int             `json:"withdraw_routes"`
	CurationRewards               *types.Int             `json:"curation_rewards"`
	PostingRewards                *types.Int             `json:"posting_rewards"`
	ProxiedVsfVotes               []interface{}          `json:"proxied_vsf_votes"`
	WitnessesVotedFor             *types.Int             `json:"witnesses_voted_for"`
	AverageBandwidth              *types.Int             `json:"average_bandwidth"`
	LifetimeBandwidth             string                 `json:"lifetime_bandwidth"`
	LastBandwidthUpdate           string                 `json:"last_bandwidth_update"`
	AverageMarketBandwidth        *types.Int             `json:"average_market_bandwidth"`
	LastMarketBandwidthUpdate     string                 `json:"last_market_bandwidth_update"`
	LastPost                      string                 `json:"last_post"`
	LastRootPost                  string                 `json:"last_root_post"`
	PostBandwidth                 *types.Int             `json:"post_bandwidth"`
	NewAverageBandwidth           string                 `json:"new_average_bandwidth"`
	NewAverageMarketBandwidth     string                 `json:"new_average_market_bandwidth"`
}

type OwnerHistory struct {
	ID                     int              `json:"id"`
	Account                string           `json:"account"`
	PreviousOwnerAuthority *types.Authority `json:"previous_owner_authority"`
	LastValidTime          string           `json:"last_valid_time"`
}
