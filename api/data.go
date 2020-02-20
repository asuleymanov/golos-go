package api

import (
	"github.com/asuleymanov/golos-go/types"
)

//Bandwidth structure for the GetAccountBandwidth function
type Bandwidth struct {
	ID                  types.Int  `json:"id"`
	Account             string     `json:"account"`
	Type                string     `json:"type"`
	AverageBandwidth    int64      `json:"average_bandwidth"`
	LifetimeBandwidth   int64      `json:"lifetime_bandwidth"`
	LastBandwidthUpdate types.Time `json:"last_bandwidth_update"`
}

//Account structure for the GetAccounts and LookupAccountNames function
type Account struct {
	ID                            types.Int       `json:"id"`
	Name                          string          `json:"name"`
	Owner                         types.Authority `json:"owner"`
	Active                        types.Authority `json:"active"`
	Posting                       types.Authority `json:"posting"`
	MemoKey                       string          `json:"memo_key"`
	JSONMetadata                  string          `json:"json_metadata"`
	Proxy                         string          `json:"proxy"`
	LastOwnerUpdate               types.Time      `json:"last_owner_update"`
	LastAccountUpdate             types.Time      `json:"last_account_update"`
	Created                       types.Time      `json:"created"`
	Mined                         bool            `json:"mined"`
	OwnerChallenged               bool            `json:"owner_challenged"`
	ActiveChallenged              bool            `json:"active_challenged"`
	LastOwnerProved               types.Time      `json:"last_owner_proved"`
	LastActiveProved              types.Time      `json:"last_active_proved"`
	RecoveryAccount               string          `json:"recovery_account"`
	LastAccountRecovery           types.Time      `json:"last_account_recovery"`
	ResetAccount                  string          `json:"reset_account"`
	CommentCount                  uint32          `json:"comment_count"`
	LifetimeVoteCount             uint32          `json:"lifetime_vote_count"`
	PostCount                     uint32          `json:"post_count"`
	CanVote                       bool            `json:"can_vote"`
	VotingPower                   uint16          `json:"voting_power"`
	LastVoteTime                  types.Time      `json:"last_vote_time"`
	Balance                       types.Asset     `json:"balance"`
	SavingsBalance                types.Asset     `json:"savings_balance"`
	SbdBalance                    types.Asset     `json:"sbd_balance"`
	SbdSeconds                    types.UInt64    `json:"sbd_seconds"`
	SbdSecondsLastUpdate          types.Time      `json:"sbd_seconds_last_update"`
	SbdLastInterestPayment        types.Time      `json:"sbd_last_interest_payment"`
	SavingsSbdBalance             types.Asset     `json:"savings_sbd_balance"`
	SavingsSbdSeconds             types.UInt64    `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   types.Time      `json:"savings_sbd_seconds_last_update"`
	SavingsSbdLastInterestPayment types.Time      `json:"savings_sbd_last_interest_payment"`
	SavingsWithdrawRequests       uint8           `json:"savings_withdraw_requests"`
	VestingShares                 types.Asset     `json:"vesting_shares"`
	DelegatedVestingShares        types.Asset     `json:"delegated_vesting_shares"`
	ReceivedVestingShares         types.Asset     `json:"received_vesting_shares"`
	VestingWithdrawRate           types.Asset     `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         types.Time      `json:"next_vesting_withdrawal"`
	Withdrawn                     int64           `json:"withdrawn"`
	ToWithdraw                    int64           `json:"to_withdraw"`
	WithdrawRoutes                uint16          `json:"withdraw_routes"`
	BenefactionRrewards           int64           `json:"benefaction_rewards"`
	CurationRewards               int64           `json:"curation_rewards"`
	DelegationRewards             int64           `json:"delegation_rewards"`
	PostingRewards                int64           `json:"posting_rewards"`
	ProxiedVsfVotes               []int64         `json:"proxied_vsf_votes"`
	WitnessesVotedFor             uint16          `json:"witnesses_voted_for"`
	AverageBandwidth              int64           `json:"average_bandwidth"`
	AverageMarketBandwidth        int64           `json:"average_market_bandwidth"`
	LifetimeBandwidth             int64           `json:"lifetime_bandwidth"`
	LifetimeMarketBandwidth       int64           `json:"lifetime_market_bandwidth"`
	LastBandwidthUpdate           types.Time      `json:"last_bandwidth_update"`
	LastMarketBandwidthUpdate     types.Time      `json:"last_market_bandwidth_update"`
	LastComment                   types.Time      `json:"last_comment"`
	LastPost                      types.Time      `json:"last_post"`
	PostBandwidth                 int64           `json:"post_bandwidth"`
	WitnessVotes                  []string        `json:"witness_votes"`
	Reputation                    int64           `json:"reputation"`
	PostsCapacity                 uint32          `json:"posts_capacity"`
	CommentsCapacity              uint32          `json:"comments_capacity"`
	VotingCapacity                uint32          `json:"voting_capacity"`
	ReferrerAccount               string          `json:"referrer_account"`
	ReferrerInterestRate          uint16          `json:"referrer_interest_rate"`
	ReferralEndDate               types.Time      `json:"referral_end_date"`
	ReferralBreakFee              types.Asset     `json:"referral_break_fee"`
	LastActiveOperation           types.Time      `json:"last_active_operation"`
}

//Block structure for the GetBlock function
type Block struct {
	Number                uint32              `json:"-"`
	Timestamp             types.Time          `json:"timestamp"`
	Witness               string              `json:"witness"`
	WitnessSignature      string              `json:"witness_signature"`
	TransactionMerkleRoot string              `json:"transaction_merkle_root"`
	Previous              string              `json:"previous"`
	Extensions            []interface{}       `json:"extensions"`
	Transactions          []types.Transaction `json:"transactions"`
}

//BlockHeader structure for the GetBlockHeader function
type BlockHeader struct {
	Number                uint32        `json:"-"`
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	Witness               string        `json:"witness"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Extensions            []interface{} `json:"extensions"`
}

//Config structure for the GetConfig function.
type Config struct {
	BuildTestnet                   bool        `json:"STEEMIT_BUILD_TESTNET"`
	GrapheneCurrentDBVersion       string      `json:"GRAPHENE_CURRENT_DB_VERSION"`
	SbdSymbol                      types.Int   `json:"SBD_SYMBOL"`
	Percent100                     int         `json:"STEEMIT_100_PERCENT"`
	Percent1                       types.Int   `json:"STEEMIT_1_PERCENT"`
	AddressPrefix                  string      `json:"STEEMIT_ADDRESS_PREFIX"`
	AprPercentMultiplyPerBlock     string      `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	AprPercentMultiplyPerHour      string      `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR"`
	AprPercentMultiplyPerRound     string      `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND"`
	AprPercentShiftPerBlock        types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK"`
	AprPercentShiftPerHour         types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_HOUR"`
	AprPercentShiftPerRound        types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_ROUND"`
	BandwidthAverageWindowSeconds  types.Int   `json:"STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	BandwidthPrecision             types.Int   `json:"STEEMIT_BANDWIDTH_PRECISION"`
	BlockchainPrecision            types.Int   `json:"STEEMIT_BLOCKCHAIN_PRECISION"`
	BlockchainPrecisionDigits      types.Int   `json:"STEEMIT_BLOCKCHAIN_PRECISION_DIGITS"`
	BlockchainHardforkVersion      string      `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	BlockchainVersion              string      `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	BlockInterval                  uint        `json:"STEEMIT_BLOCK_INTERVAL"`
	BlocksPerDay                   types.Int   `json:"STEEMIT_BLOCKS_PER_DAY"`
	BlocksPerHour                  types.Int   `json:"STEEMIT_BLOCKS_PER_HOUR"`
	BlocksPerYear                  types.Int   `json:"STEEMIT_BLOCKS_PER_YEAR"`
	CashoutWindowSeconds           types.Int   `json:"STEEMIT_CASHOUT_WINDOW_SECONDS"`
	ChainID                        string      `json:"STEEMIT_CHAIN_ID"`
	ContentAprPercent              types.Int   `json:"STEEMIT_CONTENT_APR_PERCENT"`
	ConversionDelay                string      `json:"STEEMIT_CONVERSION_DELAY"`
	CurateAprPercent               types.Int   `json:"STEEMIT_CURATE_APR_PERCENT"`
	DefaultSbdInterestRate         types.Int   `json:"STEEMIT_DEFAULT_SBD_INTEREST_RATE"`
	FeedHistoryWindow              types.Int   `json:"STEEMIT_FEED_HISTORY_WINDOW"`
	FeedIntervalBlocks             types.Int   `json:"STEEMIT_FEED_INTERVAL_BLOCKS"`
	FreeTransactionsWithNewAccount types.Int   `json:"STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT"`
	GenesisTime                    string      `json:"STEEMIT_GENESIS_TIME"`
	HardforkRequiredWitnesses      types.Int   `json:"STEEMIT_HARDFORK_REQUIRED_WITNESSES"`
	InitMinerName                  string      `json:"STEEMIT_INIT_MINER_NAME"`
	InitPublicKeyStr               string      `json:"STEEMIT_INIT_PUBLIC_KEY_STR"`
	InitSupply                     types.Int   `json:"STEEMIT_INIT_SUPPLY"`
	InitTime                       string      `json:"STEEMIT_INIT_TIME"`
	IrreversibleThreshold          types.Int   `json:"STEEMIT_IRREVERSIBLE_THRESHOLD"`
	LiquidityAprPercent            types.Int   `json:"STEEMIT_LIQUIDITY_APR_PERCENT"`
	LiquidityRewardBlocks          types.Int   `json:"STEEMIT_LIQUIDITY_REWARD_BLOCKS"`
	LiquidityRewardPeriodSec       types.Int   `json:"STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC"`
	LiquidityTimeoutSec            string      `json:"STEEMIT_LIQUIDITY_TIMEOUT_SEC"`
	MaxAccountNameLength           types.Int   `json:"STEEMIT_MAX_ACCOUNT_NAME_LENGTH"`
	MaxAccountWitnessVotes         types.Int   `json:"STEEMIT_MAX_ACCOUNT_WITNESS_VOTES"`
	MaxAssetWhitelistAuthorities   types.Int   `json:"STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	MaxAuthorityMembership         types.Int   `json:"STEEMIT_MAX_AUTHORITY_MEMBERSHIP"`
	MaxBlockSize                   types.Int   `json:"STEEMIT_MAX_BLOCK_SIZE"`
	MaxCashoutWindowSeconds        types.Int   `json:"STEEMIT_MAX_CASHOUT_WINDOW_SECONDS"`
	MaxCommentDepth                types.Int   `json:"STEEMIT_MAX_COMMENT_DEPTH"`
	MaxFeedAge                     types.Int   `json:"STEEMIT_MAX_FEED_AGE"`
	MaxInstanceID                  string      `json:"STEEMIT_MAX_INSTANCE_ID"`
	MaxMemoSize                    types.Int   `json:"STEEMIT_MAX_MEMO_SIZE"`
	MaxWitnesses                   types.Int   `json:"STEEMIT_MAX_WITNESSES"`
	MaxMinerWitnesses              types.Int   `json:"STEEMIT_MAX_MINER_WITNESSES"`
	MaxProxyRecursionDepth         types.Int   `json:"STEEMIT_MAX_PROXY_RECURSION_DEPTH"`
	MaxRationDecayRate             types.Int   `json:"STEEMIT_MAX_RATION_DECAY_RATE"`
	MaxReserveRatio                types.Int   `json:"STEEMIT_MAX_RESERVE_RATIO"`
	MaxRunnerWitnesses             types.Int   `json:"STEEMIT_MAX_RUNNER_WITNESSES"`
	MaxShareSupply                 string      `json:"STEEMIT_MAX_SHARE_SUPPLY"`
	MaxSigCheckDepth               types.Int   `json:"STEEMIT_MAX_SIG_CHECK_DEPTH"`
	MaxTimeUntilExpiration         types.Int   `json:"STEEMIT_MAX_TIME_UNTIL_EXPIRATION"`
	MaxTransactionSize             types.Int   `json:"STEEMIT_MAX_TRANSACTION_SIZE"`
	MaxUndoHistory                 types.Int   `json:"STEEMIT_MAX_UNDO_HISTORY"`
	MaxURLLength                   types.Int   `json:"STEEMIT_MAX_URL_LENGTH"`
	MaxVoteChanges                 types.Int   `json:"STEEMIT_MAX_VOTE_CHANGES"`
	MaxVotedWitnesses              types.Int   `json:"STEEMIT_MAX_VOTED_WITNESSES"`
	MaxWithdrawRoutes              types.Int   `json:"STEEMIT_MAX_WITHDRAW_ROUTES"`
	MaxWitnessURLLength            types.Int   `json:"STEEMIT_MAX_WITNESS_URL_LENGTH"`
	MinAccountCreationFee          types.Int   `json:"STEEMIT_MIN_ACCOUNT_CREATION_FEE"`
	MinAccountNameLength           types.Int   `json:"STEEMIT_MIN_ACCOUNT_NAME_LENGTH"`
	MinBlockSizeLimit              types.Int   `json:"STEEMIT_MIN_BLOCK_SIZE_LIMIT"`
	MinContentReward               types.Asset `json:"STEEMIT_MIN_CONTENT_REWARD"`
	MinCurateReward                types.Asset `json:"STEEMIT_MIN_CURATE_REWARD"`
	MinerAccount                   string      `json:"STEEMIT_MINER_ACCOUNT"`
	MinerPayPercent                types.Int   `json:"STEEMIT_MINER_PAY_PERCENT"`
	MinFeeds                       types.Int   `json:"STEEMIT_MIN_FEEDS"`
	MiningReward                   types.Asset `json:"STEEMIT_MINING_REWARD"`
	MiningTime                     string      `json:"STEEMIT_MINING_TIME"`
	MinLiquidityReward             types.Asset `json:"STEEMIT_MIN_LIQUIDITY_REWARD"`
	MinLiquidityRewardPeriodSec    types.Int   `json:"STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC"`
	MinPayoutSbd                   types.Asset `json:"STEEMIT_MIN_PAYOUT_SBD"`
	MinPowReward                   types.Asset `json:"STEEMIT_MIN_POW_REWARD"`
	MinProducerReward              types.Asset `json:"STEEMIT_MIN_PRODUCER_REWARD"`
	MinRation                      types.Int   `json:"STEEMIT_MIN_RATION"`
	MinTransactionExpirationLimit  types.Int   `json:"STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	MinTransactionSizeLimit        types.Int   `json:"STEEMIT_MIN_TRANSACTION_SIZE_LIMIT"`
	MinUndoHistory                 types.Int   `json:"STEEMIT_MIN_UNDO_HISTORY"`
	NullAccount                    string      `json:"STEEMIT_NULL_ACCOUNT"`
	NumInitMiners                  types.Int   `json:"STEEMIT_NUM_INIT_MINERS"`
	PowAprPercent                  types.Int   `json:"STEEMIT_POW_APR_PERCENT"`
	ProducerAprPercent             types.Int   `json:"STEEMIT_PRODUCER_APR_PERCENT"`
	ProxyToSelfAccount             string      `json:"STEEMIT_PROXY_TO_SELF_ACCOUNT"`
	SbdInterestCompoundIntervalSec types.Int   `json:"STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC"`
	SecondsPerYear                 types.Int   `json:"STEEMIT_SECONDS_PER_YEAR"`
	ReverseAuctionWindowSeconds    types.Int   `json:"STEEMIT_REVERSE_AUCTION_WINDOW_SECONDS"`
	StartMinerVotingBlock          types.Int   `json:"STEEMIT_START_MINER_VOTING_BLOCK"`
	StartVestingBlock              types.Int   `json:"STEEMIT_START_VESTING_BLOCK"`
	Symbol                         string      `json:"STEEMIT_SYMBOL"`
	TempAccount                    string      `json:"STEEMIT_TEMP_ACCOUNT"`
	UpvoteLockout                  types.Int   `json:"STEEMIT_UPVOTE_LOCKOUT"`
	VestingWithdrawIntervals       types.Int   `json:"STEEMIT_VESTING_WITHDRAW_INTERVALS"`
	VestingWithdrawIntervalSeconds types.Int   `json:"STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	VoteChangeLockoutPeriod        types.Int   `json:"STEEMIT_VOTE_CHANGE_LOCKOUT_PERIOD"`
	VoteRegenerationSeconds        int         `json:"STEEMIT_VOTE_REGENERATION_SECONDS"`
	SteemSymbol                    string      `json:"STEEM_SYMBOL"`
	VestsSymbol                    string      `json:"VESTS_SYMBOL"`
	BlockchainName                 string      `json:"BLOCKCHAIN_NAME"`
}

//ConversionRequests structure for the GetConversionRequests function.
type ConversionRequests struct {
	ID             types.Int   `json:"id"`
	Owner          string      `json:"owner"`
	Requestid      uint32      `json:"requestid"`
	Amount         types.Asset `json:"amount"`
	ConversionDate types.Time  `json:"conversion_date"`
}

//DatabaseInfo structure for the GetDatabaseInfo function.
type DatabaseInfo struct {
	TotalSize    uint64              `json:"total_size"`
	FreeSize     uint64              `json:"free_size"`
	ReservedSize uint64              `json:"reserved_size"`
	UsedSize     uint64              `json:"used_size"`
	IndexList    []DatabaseInfoIndex `json:"index_list"`
}

//DatabaseInfoIndex additional structure for the function GetDatabaseInfo.
type DatabaseInfoIndex struct {
	Name        string `json:"name"`
	RecordCount uint32 `json:"record_count"`
}

//DynamicGlobalProperties structure for the GetDynamicGlobalProperties function.
type DynamicGlobalProperties struct {
	ID                           types.Int        `json:"id"`
	HeadBlockNumber              uint32           `json:"head_block_number"`
	HeadBlockID                  string           `json:"head_block_id"`
	Time                         types.Time       `json:"time"`
	CurrentWitness               string           `json:"current_witness"`
	TotalPow                     uint64           `json:"total_pow"`
	NumPowWitnesses              uint32           `json:"num_pow_witnesses"`
	VirtualSupply                types.Asset      `json:"virtual_supply"`
	CurrentSupply                types.Asset      `json:"current_supply"`
	ConfidentialSupply           types.Asset      `json:"confidential_supply"`
	CurrentSBDSupply             types.Asset      `json:"current_sbd_supply"`
	ConfidentialSBDSupply        types.Asset      `json:"confidential_sbd_supply"`
	TotalVersingFund             types.Asset      `json:"total_vesting_fund_steem"`
	TotalVestingShares           types.Asset      `json:"total_vesting_shares"`
	TotalRewardFund              types.Asset      `json:"total_reward_fund_steem"`
	TotalRewardShares2           types.UInt64     `json:"total_reward_shares2"`
	SBDInterestRate              uint16           `json:"sbd_interest_rate"`
	SBDPrintRate                 uint16           `json:"sbd_print_rate"`
	SBDDebtPercent               uint16           `json:"sbd_debt_percent"`
	IsForcedMinPrice             bool             `json:"is_forced_min_price"`
	AverageBlockSize             uint32           `json:"average_block_size"`
	MaximumBlockSize             uint32           `json:"maximum_block_size"`
	CurrentAslot                 uint64           `json:"current_aslot"`
	RecentSlotsFilled            types.UInt64     `json:"recent_slots_filled"`
	ParticipationCount           uint8            `json:"participation_count"`
	LastIrreversibleBlockNum     uint32           `json:"last_irreversible_block_num"`
	MaxVirtualBandwidth          uint64           `json:"max_virtual_bandwidth"`
	CurrentReserveRatio          uint64           `json:"current_reserve_ratio"`
	CustomOpsBandwidthMultiplier uint16           `json:"custom_ops_bandwidth_multiplier"`
	TransitBlockNum              uint32           `json:"transit_block_num"`
	TransitWitnesses             string           `json:"transit_witnesses"`
	WorkerRequests               []WorkerRequests `json:"worker_requests"`
}

//VestingDelegationExpiration structure for the GetExpiringVestingDelegations function.
type VestingDelegationExpiration struct {
	ID            types.Int   `json:"id"`
	Delegator     string      `json:"delegator"`
	VestingShares types.Asset `json:"vesting_shares"`
	Expiration    types.Time  `json:"expiration"`
}

//NextScheduledHardfork structure for the GetNextScheduledHardfork function.
type NextScheduledHardfork struct {
	HfVersion string     `json:"hf_version"`
	LiveTime  types.Time `json:"live_time"`
}

//OwnerHistory structure for the GetOwnerHistory function.
type OwnerHistory struct {
	ID                     types.Int       `json:"id"`
	Account                string          `json:"account"`
	PreviousOwnerAuthority types.Authority `json:"previous_owner_authority"`
	LastValidTime          string          `json:"last_valid_time"`
}

//ProposalObject structure for the GetProposedTransaction function.
type ProposalObject struct {
	Author                    string           `json:"author"`
	Title                     string           `json:"title"`
	Memo                      string           `json:"memo"`
	ExpirationTime            types.Time       `json:"expiration_time"`
	ReviewPeriodTime          types.Time       `json:"review_period_time,omitempty"`
	ProposedOperations        types.Operations `json:"proposed_operations"`
	RequiredActiveApprovals   []string         `json:"required_active_approvals"`
	AvailableActiveApprovals  []string         `json:"available_active_approvals"`
	RequiredOwnerApprovals    []string         `json:"required_owner_approvals"`
	AvailableOwnerApprovals   []string         `json:"available_owner_approvals"`
	RequiredPostingApprovals  []string         `json:"required_posting_approvals"`
	AvailablePostingApprovals []string         `json:"available_posting_approvals"`
	AvailableKeyApprovals     []string         `json:"available_key_approvals"`
}

//SavingsWithdraw structure for the GetSavingsWithdrawFrom and GetSavingsWithdrawTo functions.
type SavingsWithdraw struct {
	ID        types.Int   `json:"id"`
	From      string      `json:"from"`
	To        string      `json:"to"`
	Memo      string      `json:"memo"`
	RequestID uint32      `json:"request_id"`
	Amount    types.Asset `json:"amount"`
	Complete  types.Time  `json:"complete"`
}

//VestingDelegation structure for the GetVestingDelegations function.
type VestingDelegation struct {
	ID                types.Int   `json:"id"`
	Delegator         string      `json:"delegator"`
	Delegatee         string      `json:"delegatee"`
	VestingShares     types.Asset `json:"vesting_shares"`
	InterestRate      uint16      `json:"interest_rate"`
	MinDelegationTime types.Time  `json:"min_delegation_time"`
}

//WithdrawVestingRoutes structure for the GetWithdrawRoutes function.
type WithdrawVestingRoutes struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

//Escrow structure for the GetEscrow function.
type Escrow struct {
	ID                   types.Int   `json:"id"`
	EscrowID             uint32      `json:"escrow_id"`
	From                 string      `json:"from"`
	To                   string      `json:"to"`
	Agent                string      `json:"agent"`
	RatificationDeadline types.Time  `json:"ratification_deadline"`
	EscrowExpiration     types.Time  `json:"escrow_expiration"`
	SbdBalance           types.Asset `json:"sbd_balance"`
	SteemBalance         types.Asset `json:"steem_balance"`
	Fee                  types.Asset `json:"pending_fee"`
	ToApproved           bool        `json:"to_approved"`
	AgentApproved        bool        `json:"agent_approved"`
	Disputed             bool        `json:"disputed"`
	IsApproved           bool        `json:"is_approved"`
}

//AccountRecoveryRequest structure for the GetRecoveryRequest function.
type AccountRecoveryRequest struct {
	ID                types.Int       `json:"id"`
	AccountToRecover  string          `json:"account_to_recover"`
	NewOwnerAuthority types.Authority `json:"new_owner_authority"`
	Expires           types.Time      `json:"expires"`
}

//Blogs structure for the GetBlog function
type Blogs struct {
	Comment            CommentData `json:"comment"`
	Blog               string      `json:"blog"`
	ReblogOn           types.Time  `json:"reblog_on"`
	EntryID            uint32      `json:"entry_id"`
	ReblogTitle        string      `json:"reblog_title"`
	ReblogBody         string      `json:"reblog_body"`
	ReblogJSONMetadata string      `json:"reblog_json_metadata"`
}

//CommentData additional structure for the function GetBlog.
type CommentData struct {
	ID                             types.Int           `json:"id"`
	Title                          string              `json:"title"`
	Body                           string              `json:"body"`
	JSONMetadata                   string              `json:"json_metadata"`
	ParentAuthor                   string              `json:"parent_author"`
	ParentPermlink                 string              `json:"parent_permlink"`
	Author                         string              `json:"author"`
	Permlink                       string              `json:"permlink"`
	Category                       string              `json:"category"`
	LastUpdate                     types.Time          `json:"last_update"`
	Created                        types.Time          `json:"created"`
	Active                         types.Time          `json:"active"`
	LastPayout                     types.Time          `json:"last_payout"`
	Depth                          uint8               `json:"depth"`
	Children                       uint32              `json:"children"`
	ChildrenRshares2               types.UInt64        `json:"children_rshares2"`
	NetRshares                     int64               `json:"net_rshares"`
	AbsRshares                     int64               `json:"abs_rshares"`
	VoteRshares                    int64               `json:"vote_rshares"`
	ChildrenAbsRshares             int64               `json:"children_abs_rshares"`
	CashoutTime                    types.Time          `json:"cashout_time"`
	MaxCashoutTime                 types.Time          `json:"max_cashout_time"`
	TotalVoteWeight                uint64              `json:"total_vote_weight"`
	RewardWeight                   uint16              `json:"reward_weight"`
	TotalPayoutValue               types.Asset         `json:"total_payout_value"`
	BeneficiaryPayoutValue         types.Asset         `json:"beneficiary_payout_value"`
	BeneficiaryGestsPayoutValue    types.Asset         `json:"beneficiary_gests_payout_value"`
	CuratorPayoutValue             types.Asset         `json:"curator_payout_value"`
	CuratorGestsPayoutValue        types.Asset         `json:"curator_gests_payout_value"`
	AuthorRewards                  int64               `json:"author_rewards"`
	AuthorGBGPayoutValue           types.Asset         `json:"author_gbg_payout_value"`
	AuthorGolosPayoutValue         types.Asset         `json:"author_golos_payout_value"`
	AuthorGestsPayoutValue         types.Asset         `json:"author_gests_payout_value"`
	NetVotes                       int32               `json:"net_votes"`
	Mode                           string              `json:"mode"`
	RootComment                    types.Int           `json:"root_comment"`
	CurationRewardCurve            uint8               `json:"curation_reward_curve"`
	AuctionWindowRewardDestination int8                `json:"auction_window_reward_destination"`
	AuctionWindowSize              uint16              `json:"auction_window_size"`
	AuctionWindowWeight            uint64              `json:"auction_window_weight"`
	VotesInAuctionWindowWeight     uint64              `json:"votes_in_auction_window_weight"`
	RootTitle                      string              `json:"root_title"`
	MaxAcceptedPayout              types.Asset         `json:"max_accepted_payout"`
	PercentSteemDollars            uint16              `json:"percent_steem_dollars"`
	AllowReplies                   bool                `json:"allow_replies"`
	AllowVotes                     bool                `json:"allow_votes"`
	AllowCurationRewards           bool                `json:"allow_curation_rewards"`
	CurationRewardsPercent         uint16              `json:"curation_rewards_percent"`
	Beneficiaries                  []types.Beneficiary `json:"beneficiaries"`
}

//BlogAuthor structure for the GetBlogAuthors function
type BlogAuthor struct {
	Author string `json:"author"`
	Count  uint32 `json:"count"`
}

//BlogEntries structure for the GetBlogEntries function
type BlogEntries struct {
	Author             string     `json:"author"`
	Permlink           string     `json:"permlink"`
	Blog               string     `json:"blog"`
	ReblogOn           types.Time `json:"reblog_on"`
	EntryID            uint32     `json:"entry_id"`
	ReblogTitle        string     `json:"reblog_title"`
	ReblogBody         string     `json:"reblog_body"`
	ReblogJSONMetadata string     `json:"reblog_json_metadata"`
}

//Feeds structure for the GetFeed function
type Feeds struct {
	Comment       CommentData   `json:"comment"`
	ReblogBy      []string      `json:"reblog_by"`
	ReblogEntries []ReblogEntry `json:"reblog_entries"`
	ReblogOn      types.Time    `json:"reblog_on"`
	EntryID       uint32        `json:"entry_id"`
}

//ReblogEntry additional structure for the function Feeds, FeedEntry.
type ReblogEntry struct {
	Author       string `json:"author"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	JSONMetadata string `json:"json_metadata"`
}

//FeedEntry structure for the GetFeedEntries function
type FeedEntry struct {
	Author        string        `json:"author"`
	Permlink      string        `json:"permlink"`
	ReblogBy      []string      `json:"reblog_by"`
	ReblogEntries []ReblogEntry `json:"reblog_entries"`
	ReblogOn      types.Time    `json:"reblog_on"`
	EntryID       uint32        `json:"entry_id"`
}

//FollowCount structure for the GetFollowCount function
type FollowCount struct {
	Account        string `json:"account"`
	FollowerCount  uint32 `json:"follower_count"`
	FollowingCount uint32 `json:"following_count"`
	Limit          uint32 `json:"limit"`
}

//FollowObject structure for the GetFollowers and GetFollowing functions
type FollowObject struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []uint32 `json:"what"`
}

//BroadcastResponse structure for the BroadcastTransactionSynchronous function
type BroadcastResponse struct {
	ID       string `json:"id"`
	BlockNum int32  `json:"block_num"`
	TrxNum   int32  `json:"trx_num"`
	Expired  bool   `json:"expired"`
}

//MarketHistory structure for the GetMarketHistory function.
type MarketHistory struct {
	ID          types.Int  `json:"id"`
	Open        types.Time `json:"open"`
	Seconds     uint32     `json:"seconds"`
	HighSteem   int64      `json:"high_steem"`
	HighSbd     int64      `json:"high_sbd"`
	LowSteem    int64      `json:"low_steem"`
	LowSbd      int64      `json:"low_sbd"`
	OpenSteem   int64      `json:"open_steem"`
	OpenSbd     int64      `json:"open_sbd"`
	CloseSteem  int64      `json:"close_steem"`
	CloseSbd    int64      `json:"close_sbd"`
	SteemVolume int64      `json:"steem_volume"`
	SbdVolume   int64      `json:"sbd_volume"`
}

//OpenOrders structure for the GetOpenOrders function.
type OpenOrders struct {
	ID         types.Int  `json:"id"`
	Created    types.Time `json:"created"`
	Expiration types.Time `json:"expiration"`
	Seller     string     `json:"seller"`
	Orderid    uint32     `json:"orderid"`
	ForSale    int64      `json:"for_sale"`
	SellPrice  Prices     `json:"sell_price"`
	RealPrice  float64    `json:"real_price"`
	Rewarded   bool       `json:"rewarded"`
}

//Price additional structure for the function GetOpenOrders.
type Prices struct {
	Base  types.Asset `json:"base"`
	Quote types.Asset `json:"quote"`
}

//OrderBook structure for the GetOrderBook function.
type OrderBook struct {
	Ask []Order `json:"asks"`
	Bid []Order `json:"bids"`
}

//Order additional structure for the function OrderBook.
type Order struct {
	Price float64 `json:"price"`
	Steem int64   `json:"steem"`
	Sbd   int64   `json:"sbd"`
}

//OrderBookExtended structure for the GetOrderBookExtended function.
type OrderBookExtended struct {
	Ask []OrderExtended `json:"asks"`
	Bid []OrderExtended `json:"bids"`
}

//OrderExtended additional structure for the function OrderBookExtended.
type OrderExtended struct {
	Price     Prices     `json:"price"`
	RealPrice float64    `json:"real_price"`
	Steem     int64      `json:"steem"`
	Sbd       int64      `json:"sbd"`
	Created   types.Time `json:"created"`
}

//Trades structure for the GetRecentTrades and GetTradeHistory functions.
type Trades struct {
	Date        types.Time  `json:"date"`
	CurrentPays types.Asset `json:"current_pays"`
	OpenPays    types.Asset `json:"open_pays"`
}

//Ticker structure for the GetTicker function.
type Ticker struct {
	Latest        float64     `json:"latest"`
	LowestAsk     float64     `json:"lowest_ask"`
	HighestBid    float64     `json:"highest_bid"`
	PercentChange float64     `json:"percent_change"`
	SteemVolume   types.Asset `json:"steem_volume"`
	SbdVolume     types.Asset `json:"sbd_volume"`
}

//Volume structure for the GetVolume function.
type Volume struct {
	SteemVolume types.Asset `json:"steem_volume"`
	SbdVolume   types.Asset `json:"sbd_volume"`
}

//Message structure for the GetInbox, GetOutbox and GetThread functions.
type Message struct {
	ID               types.Int  `json:"id"`
	From             string     `json:"from"`
	To               string     `json:"to"`
	Nonce            uint64     `json:"nonce"`
	FromMemoKey      string     `json:"from_memo_key"`
	ToMemoKey        string     `json:"to_memo_key"`
	Checksum         uint32     `json:"checksum"`
	EncryptedMessage string     `json:"encrypted_message"`
	CreateDate       types.Time `json:"create_date"`
	ReceiveDate      types.Time `json:"receive_date"`
	ReadDate         types.Time `json:"read_date"`
	RemoveDate       types.Time `json:"remove_date"`
}

//MessageSettings structure for the GetSettings function.
type MessageSettings struct {
	IgnoreMessagesFromUnknownContact bool `json:"ignore_messages_from_unknown_contact"`
}

//MessageContact structure for the GetContactInfo and GetContacts functions.
type MessageContact struct {
	Owner        string `json:"owner"`
	Contact      string `json:"contact"`
	JsonMetadata string `json:"json_metadata"`
	LocalType    string `json:"local_type"`
	RemoteType   string `json:"remote_type"`
	Size         uint32 `json:"size"`
}

//Votes structure for the GetAccountVotes function.
type Votes struct {
	Authorperm string     `json:"authorperm"`
	Weight     uint64     `json:"weight"`
	Rshares    int64      `json:"rshares"`
	Percent    int16      `json:"percent"`
	Time       types.Time `json:"time"`
}

//DiscussionQuery structure used in queries.
type DiscussionQuery struct {
	Limit            uint32   `json:"limit,omitempty"`
	SelectTags       []string `json:"select_tags,omitempty"`
	FilterTags       []string `json:"filter_tags,omitempty"`
	SelectCategories []string `json:"select_categories,omitempty"`
	SelectLanguages  []string `json:"select_languages,omitempty"`
	FilterLanguages  []string `json:"filter_languages,omitempty"`
	TruncateBody     uint32   `json:"truncate_body,omitempty"`
	VoteLimit        uint32   `json:"vote_limit,omitempty"`
	VoteOffset       uint32   `json:"vote_offset,omitempty"`
	SelectAuthors    []string `json:"select_authors,omitempty"`
	StartAuthor      string   `json:"start_author,omitempty"`
	StartPermlink    string   `json:"start_permlink,omitempty"`
	ParentAuthor     string   `json:"parent_author,omitempty"`
	ParentPermlink   string   `json:"parent_permlink,omitempty"`
}

//Discussion structure for the GetAllContentReplies, GetContent, GetContentReplies and GetRepliesByLastUpdate function.
type Discussion struct {
	CommentData
	url                               string        `json:"url"`
	PendingAuthorPayoutValue          types.Asset   `json:"pending_author_payout_value"`
	PendingAuthorPayoutGBGValue       types.Asset   `json:"pending_author_payout_gbg_value"`
	PendingAuthorPayoutGestsValue     types.Asset   `json:"pending_author_payout_gests_value"`
	PendingAuthorPayoutGolosValue     types.Asset   `json:"pending_author_payout_golos_value"`
	PendingBenefactorPayoutValue      types.Asset   `json:"pending_benefactor_payout_value"`
	PendingBenefactorPayoutGestsValue types.Asset   `json:"pending_benefactor_payout_gests_value"`
	PendingCuratorPayoutValue         types.Asset   `json:"pending_curator_payout_value"`
	PendingCuratorPayoutGestsValue    types.Asset   `json:"pending_curator_payout_gests_value"`
	PendingPayoutValue                types.Asset   `json:"pending_payout_value"`
	TotalPendingPayoutValue           types.Asset   `json:"total_pending_payout_value"`
	ActiveVotes                       []VoteState   `json:"active_votes"`
	ActiveVotesCount                  uint32        `json:"active_votes_count"`
	Replies                           []string      `json:"replies"`
	AuthorReputation                  int64         `json:"author_reputation"`
	Promoted                          types.Asset   `json:"promoted"`
	Hot                               float64       `json:"hot"`
	Trending                          float64       `json:"trending"`
	BodyLength                        uint32        `json:"body_length"`
	RebloggedBy                       []string      `json:"reblogged_by"`
	FirstRebloggedBy                  string        `json:"first_reblogged_by"`
	FirstRebloggedOn                  types.Time    `json:"first_reblogged_on"`
	ReblogAuthor                      string        `json:"reblog_author"`
	ReblogTitle                       string        `json:"reblog_title"`
	ReblogBody                        string        `json:"reblog_body"`
	ReblogJSONMetadata                string        `json:"reblog_json_metadata"`
	ReblogEntries                     []ReblogEntry `json:"reblog_entries"`
}

//VoteState additional structure for the functions GetDiscussionsByActive, GetDiscussionsByAuthorBeforeDate, GetDiscussionsByBlog, GetDiscussionsByCashout, GetDiscussionsByChildren, GetDiscussionsByComments, GetDiscussionsByCreated, GetDiscussionsByFeed, GetDiscussionsByHot, GetDiscussionsByPayout, GetDiscussionsByPromoted, GetDiscussionsByTrending and GetDiscussionsByVotes.
type VoteState struct {
	Voter      string     `json:"voter"`
	Weight     uint64     `json:"weight"`
	Rshares    int64      `json:"rshares"`
	Percent    int16      `json:"percent"`
	Reputation int64      `json:"reputation"`
	Time       types.Time `json:"time"`
}

//TrendingTags structure for the GetTrendingTags function.
type TrendingTags struct {
	Name                  string      `json:"name"`
	TotalChildrenRshares2 float64     `json:"total_children_rshares2"`
	TotalPayouts          types.Asset `json:"total_payouts"`
	NetVotes              int32       `json:"net_votes"`
	TopPosts              uint32      `json:"top_posts"`
	Comments              uint32      `json:"comments"`
}

//FeedHistory structure for the GetFeedHistory function.
type FeedHistory struct {
	ID                   types.Int `json:"id"`
	CurrentMedianHistory Prices    `json:"current_median_history"`
	PriceHistory         []Prices  `json:"price_history"`
}

//WitnessSchedule structure for the GetWitnessSchedule function.
type WitnessSchedule struct {
	ID                            types.Int             `json:"id"`
	CurrentVirtualTime            types.UInt64          `json:"current_virtual_time"`
	NextShuffleBlockNum           uint32                `json:"next_shuffle_block_num"`
	CurrentShuffledWitnesses      string                `json:"current_shuffled_witnesses"`
	NumScheduledWitnesses         uint8                 `json:"num_scheduled_witnesses"`
	Top19Weight                   uint8                 `json:"top19_weight"`
	TimeshareWeight               uint8                 `json:"timeshare_weight"`
	MinerWeight                   uint8                 `json:"miner_weight"`
	WitnessPayNormalizationFactor uint32                `json:"witness_pay_normalization_factor"`
	MedianProps                   types.ChainProperties `json:"median_props"`
	MajorityVersion               string                `json:"majority_version"`
}

//Witness structure for the GetWitnessByAccount, GetWitnesses and GetWitnessByVote function.
type Witness struct {
	ID                    types.Int             `json:"id"`
	Owner                 string                `json:"owner"`
	Created               types.Time            `json:"created"`
	URL                   string                `json:"url"`
	Votes                 int64                 `json:"votes"`
	VirtualLastUpdate     types.UInt64          `json:"virtual_last_update"`
	VirtualPosition       types.UInt64          `json:"virtual_position"`
	VirtualScheduledTime  types.UInt64          `json:"virtual_scheduled_time"`
	TotalMissed           uint32                `json:"total_missed"`
	LastAslot             uint64                `json:"last_aslot"`
	LastConfirmedBlockNum uint64                `json:"last_confirmed_block_num"`
	PowWorker             uint64                `json:"pow_worker"`
	SigningKey            string                `json:"signing_key"`
	Props                 types.ChainProperties `json:"props"`
	SbdExchangeRate       Prices                `json:"sbd_exchange_rate"`
	LastSbdExchangeUpdate types.Time            `json:"last_sbd_exchange_update"`
	LastWork              string                `json:"last_work"`
	RunningVersion        string                `json:"running_version"`
	HardforkVersionVote   string                `json:"hardfork_version_vote"`
	HardforkTimeVote      types.Time            `json:"hardfork_time_vote"`
	TransitToCyberwayVote types.Time            `json:"transit_to_cyberway_vote"`
}

//AccountReputation structure for the GetAccountReputations function.
type AccountReputation struct {
	Account    string `json:"account"`
	Reputation int64  `json:"reputation"`
}

//WorkerRequests structure for the GetWorkerRequests function.
type WorkerRequests struct {
	Post              string      `json:"post"`
	Worker            string      `json:"worker"`
	State             int8        `json:"state"`
	Created           types.Time  `json:"created"`
	Modified          types.Time  `json:"modified"`
	NetRshares        int64       `json:"net_rshares"`
	RequiredAmountMin types.Asset `json:"required_amount_min"`
	RequiredAmountMax types.Asset `json:"required_amount_max"`
	VestReward        bool        `json:"vest_reward"`
	Duration          uint32      `json:"duration"`
	VoteEndTime       types.Time  `json:"vote_end_time"`
	CreationFee       types.Asset `json:"creation_fee"`
	Upvotes           uint16      `json:"upvotes"`
	Downvotes         uint16      `json:"downvotes"`
	StakeRshares      int64       `json:"stake_rshares"`
	StakeTotal        int64       `json:"stake_total"`
	RemainingPayment  types.Asset `json:"remaining_payment"`
}

//WorkerVotes structure for the GetWorkerRequestVotes function.
type WorkerVotes struct {
	ID          types.Int `json:"id"`
	Voter       string    `json:"voter"`
	Post        string    `json:"post"`
	VotePercent int16     `json:"vote_percent"`
	Rshares     int64     `json:"rshares"`
	Stake       int64     `json:"stake"`
}

type WorkerRequestQuery struct {
	Limit         uint32   `json:"limit"`
	StartAuthor   []string `json:"start_author,omitempty"`
	StartPermlink []string `json:"start_permlink,omitempty"`
	SelectAuthors []string `json:"select_authors"`
	SelectStates  []int8   `json:"select_states"`
}
