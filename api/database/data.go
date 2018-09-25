package database

import (
	"github.com/asuleymanov/golos-go/types"
)

//Bandwidth structure for the GetAccountBandwidth function
type Bandwidth struct {
	ID                  *types.Int   `json:"id"`
	Account             string       `json:"account"`
	Type                string       `json:"type"`
	AverageBandwidth    *types.Int   `json:"average_bandwidth"`
	LifetimeBandwidth   *types.Int64 `json:"lifetime_bandwidth"`
	LastBandwidthUpdate *types.Time  `json:"last_bandwidth_update"`
}

//Account structure for the GetAccounts and LookupAccountNames function
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
	DelegatedVestingShares        *types.Asset           `json:"delegated_vesting_shares"`
	ReceivedVestingShares         *types.Asset           `json:"received_vesting_shares"`
	VestingWithdrawRate           *types.Asset           `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         *types.Time            `json:"next_vesting_withdrawal"`
	Withdrawn                     *types.UInt64          `json:"withdrawn"`
	ToWithdraw                    *types.UInt64          `json:"to_withdraw"`
	WithdrawRoutes                *types.UInt32          `json:"withdraw_routes"`
	CurationRewards               *types.UInt32          `json:"curation_rewards"`
	PostingRewards                *types.UInt32          `json:"posting_rewards"`
	ProxiedVsfVotes               []types.UInt64         `json:"proxied_vsf_votes"`
	WitnessesVotedFor             *types.UInt32          `json:"witnesses_voted_for"`
	AverageBandwidth              *types.UInt64          `json:"average_bandwidth"`
	AverageMarketBandwidth        *types.UInt32          `json:"average_market_bandwidth"`
	LifetimeBandwidth             *types.UInt32          `json:"lifetime_bandwidth"`
	LifetimeMarketBandwidth       *types.UInt32          `json:"lifetime_market_bandwidth"`
	LastBandwidthUpdate           *types.Time            `json:"last_bandwidth_update"`
	LastMarketBandwidthUpdate     *types.Time            `json:"last_market_bandwidth_update"`
	LastPost                      *types.Time            `json:"last_post"`
	LastRootPost                  *types.Time            `json:"last_root_post"`
	PostBandwidth                 int                    `json:"post_bandwidth"`
	WitnessVotes                  []string               `json:"witness_votes"`
	Reputation                    *types.UInt64          `json:"reputation"`
}

//Block structure for the GetBlock function
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
	BuildTestnet                   bool         `json:"STEEMIT_BUILD_TESTNET"`
	GrapheneCurrentDBVersion       string       `json:"GRAPHENE_CURRENT_DB_VERSION"`
	SbdSymbol                      *types.Int   `json:"SBD_SYMBOL"`
	Percent100                     int          `json:"STEEMIT_100_PERCENT"`
	Percent1                       *types.Int   `json:"STEEMIT_1_PERCENT"`
	AddressPrefix                  string       `json:"STEEMIT_ADDRESS_PREFIX"`
	AprPercentMultiplyPerBlock     string       `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	AprPercentMultiplyPerHour      string       `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR"`
	AprPercentMultiplyPerRound     string       `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND"`
	AprPercentShiftPerBlock        *types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK"`
	AprPercentShiftPerHour         *types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_HOUR"`
	AprPercentShiftPerRound        *types.Int   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_ROUND"`
	BandwidthAverageWindowSeconds  *types.Int   `json:"STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	BandwidthPrecision             *types.Int   `json:"STEEMIT_BANDWIDTH_PRECISION"`
	BlockchainPrecision            *types.Int   `json:"STEEMIT_BLOCKCHAIN_PRECISION"`
	BlockchainPrecisionDigits      *types.Int   `json:"STEEMIT_BLOCKCHAIN_PRECISION_DIGITS"`
	BlockchainHardforkVersion      string       `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	BlockchainVersion              string       `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	BlockInterval                  uint         `json:"STEEMIT_BLOCK_INTERVAL"`
	BlocksPerDay                   *types.Int   `json:"STEEMIT_BLOCKS_PER_DAY"`
	BlocksPerHour                  *types.Int   `json:"STEEMIT_BLOCKS_PER_HOUR"`
	BlocksPerYear                  *types.Int   `json:"STEEMIT_BLOCKS_PER_YEAR"`
	CashoutWindowSeconds           *types.Int   `json:"STEEMIT_CASHOUT_WINDOW_SECONDS"`
	ChainID                        string       `json:"STEEMIT_CHAIN_ID"`
	ContentAprPercent              *types.Int   `json:"STEEMIT_CONTENT_APR_PERCENT"`
	ConversionDelay                string       `json:"STEEMIT_CONVERSION_DELAY"`
	CurateAprPercent               *types.Int   `json:"STEEMIT_CURATE_APR_PERCENT"`
	DefaultSbdInterestRate         *types.Int   `json:"STEEMIT_DEFAULT_SBD_INTEREST_RATE"`
	FeedHistoryWindow              *types.Int   `json:"STEEMIT_FEED_HISTORY_WINDOW"`
	FeedIntervalBlocks             *types.Int   `json:"STEEMIT_FEED_INTERVAL_BLOCKS"`
	FreeTransactionsWithNewAccount *types.Int   `json:"STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT"`
	GenesisTime                    string       `json:"STEEMIT_GENESIS_TIME"`
	HardforkRequiredWitnesses      *types.Int   `json:"STEEMIT_HARDFORK_REQUIRED_WITNESSES"`
	InitMinerName                  string       `json:"STEEMIT_INIT_MINER_NAME"`
	InitPublicKeyStr               string       `json:"STEEMIT_INIT_PUBLIC_KEY_STR"`
	InitSupply                     *types.Int   `json:"STEEMIT_INIT_SUPPLY"`
	InitTime                       string       `json:"STEEMIT_INIT_TIME"`
	IrreversibleThreshold          *types.Int   `json:"STEEMIT_IRREVERSIBLE_THRESHOLD"`
	LiquidityAprPercent            *types.Int   `json:"STEEMIT_LIQUIDITY_APR_PERCENT"`
	LiquidityRewardBlocks          *types.Int   `json:"STEEMIT_LIQUIDITY_REWARD_BLOCKS"`
	LiquidityRewardPeriodSec       *types.Int   `json:"STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC"`
	LiquidityTimeoutSec            string       `json:"STEEMIT_LIQUIDITY_TIMEOUT_SEC"`
	MaxAccountNameLength           *types.Int   `json:"STEEMIT_MAX_ACCOUNT_NAME_LENGTH"`
	MaxAccountWitnessVotes         *types.Int   `json:"STEEMIT_MAX_ACCOUNT_WITNESS_VOTES"`
	MaxAssetWhitelistAuthorities   *types.Int   `json:"STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	MaxAuthorityMembership         *types.Int   `json:"STEEMIT_MAX_AUTHORITY_MEMBERSHIP"`
	MaxBlockSize                   *types.Int   `json:"STEEMIT_MAX_BLOCK_SIZE"`
	MaxCashoutWindowSeconds        *types.Int   `json:"STEEMIT_MAX_CASHOUT_WINDOW_SECONDS"`
	MaxCommentDepth                *types.Int   `json:"STEEMIT_MAX_COMMENT_DEPTH"`
	MaxFeedAge                     *types.Int   `json:"STEEMIT_MAX_FEED_AGE"`
	MaxInstanceID                  string       `json:"STEEMIT_MAX_INSTANCE_ID"`
	MaxMemoSize                    *types.Int   `json:"STEEMIT_MAX_MEMO_SIZE"`
	MaxWitnesses                   *types.Int   `json:"STEEMIT_MAX_WITNESSES"`
	MaxMinerWitnesses              *types.Int   `json:"STEEMIT_MAX_MINER_WITNESSES"`
	MaxProxyRecursionDepth         *types.Int   `json:"STEEMIT_MAX_PROXY_RECURSION_DEPTH"`
	MaxRationDecayRate             *types.Int   `json:"STEEMIT_MAX_RATION_DECAY_RATE"`
	MaxReserveRatio                *types.Int   `json:"STEEMIT_MAX_RESERVE_RATIO"`
	MaxRunnerWitnesses             *types.Int   `json:"STEEMIT_MAX_RUNNER_WITNESSES"`
	MaxShareSupply                 string       `json:"STEEMIT_MAX_SHARE_SUPPLY"`
	MaxSigCheckDepth               *types.Int   `json:"STEEMIT_MAX_SIG_CHECK_DEPTH"`
	MaxTimeUntilExpiration         *types.Int   `json:"STEEMIT_MAX_TIME_UNTIL_EXPIRATION"`
	MaxTransactionSize             *types.Int   `json:"STEEMIT_MAX_TRANSACTION_SIZE"`
	MaxUndoHistory                 *types.Int   `json:"STEEMIT_MAX_UNDO_HISTORY"`
	MaxURLLength                   *types.Int   `json:"STEEMIT_MAX_URL_LENGTH"`
	MaxVoteChanges                 *types.Int   `json:"STEEMIT_MAX_VOTE_CHANGES"`
	MaxVotedWitnesses              *types.Int   `json:"STEEMIT_MAX_VOTED_WITNESSES"`
	MaxWithdrawRoutes              *types.Int   `json:"STEEMIT_MAX_WITHDRAW_ROUTES"`
	MaxWitnessURLLength            *types.Int   `json:"STEEMIT_MAX_WITNESS_URL_LENGTH"`
	MinAccountCreationFee          *types.Int   `json:"STEEMIT_MIN_ACCOUNT_CREATION_FEE"`
	MinAccountNameLength           *types.Int   `json:"STEEMIT_MIN_ACCOUNT_NAME_LENGTH"`
	MinBlockSizeLimit              *types.Int   `json:"STEEMIT_MIN_BLOCK_SIZE_LIMIT"`
	MinContentReward               *types.Asset `json:"STEEMIT_MIN_CONTENT_REWARD"`
	MinCurateReward                *types.Asset `json:"STEEMIT_MIN_CURATE_REWARD"`
	MinerAccount                   string       `json:"STEEMIT_MINER_ACCOUNT"`
	MinerPayPercent                *types.Int   `json:"STEEMIT_MINER_PAY_PERCENT"`
	MinFeeds                       *types.Int   `json:"STEEMIT_MIN_FEEDS"`
	MiningReward                   *types.Asset `json:"STEEMIT_MINING_REWARD"`
	MiningTime                     string       `json:"STEEMIT_MINING_TIME"`
	MinLiquidityReward             *types.Asset `json:"STEEMIT_MIN_LIQUIDITY_REWARD"`
	MinLiquidityRewardPeriodSec    *types.Int   `json:"STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC"`
	MinPayoutSbd                   *types.Asset `json:"STEEMIT_MIN_PAYOUT_SBD"`
	MinPowReward                   *types.Asset `json:"STEEMIT_MIN_POW_REWARD"`
	MinProducerReward              *types.Asset `json:"STEEMIT_MIN_PRODUCER_REWARD"`
	MinRation                      *types.Int   `json:"STEEMIT_MIN_RATION"`
	MinTransactionExpirationLimit  *types.Int   `json:"STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	MinTransactionSizeLimit        *types.Int   `json:"STEEMIT_MIN_TRANSACTION_SIZE_LIMIT"`
	MinUndoHistory                 *types.Int   `json:"STEEMIT_MIN_UNDO_HISTORY"`
	NullAccount                    string       `json:"STEEMIT_NULL_ACCOUNT"`
	NumInitMiners                  *types.Int   `json:"STEEMIT_NUM_INIT_MINERS"`
	PowAprPercent                  *types.Int   `json:"STEEMIT_POW_APR_PERCENT"`
	ProducerAprPercent             *types.Int   `json:"STEEMIT_PRODUCER_APR_PERCENT"`
	ProxyToSelfAccount             string       `json:"STEEMIT_PROXY_TO_SELF_ACCOUNT"`
	SbdInterestCompoundIntervalSec *types.Int   `json:"STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC"`
	SecondsPerYear                 *types.Int   `json:"STEEMIT_SECONDS_PER_YEAR"`
	ReverseAuctionWindowSeconds    *types.Int   `json:"STEEMIT_REVERSE_AUCTION_WINDOW_SECONDS"`
	StartMinerVotingBlock          *types.Int   `json:"STEEMIT_START_MINER_VOTING_BLOCK"`
	StartVestingBlock              *types.Int   `json:"STEEMIT_START_VESTING_BLOCK"`
	Symbol                         string       `json:"STEEMIT_SYMBOL"`
	TempAccount                    string       `json:"STEEMIT_TEMP_ACCOUNT"`
	UpvoteLockout                  *types.Int   `json:"STEEMIT_UPVOTE_LOCKOUT"`
	VestingWithdrawIntervals       *types.Int   `json:"STEEMIT_VESTING_WITHDRAW_INTERVALS"`
	VestingWithdrawIntervalSeconds *types.Int   `json:"STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	VoteChangeLockoutPeriod        *types.Int   `json:"STEEMIT_VOTE_CHANGE_LOCKOUT_PERIOD"`
	VoteRegenerationSeconds        int          `json:"STEEMIT_VOTE_REGENERATION_SECONDS"`
	SteemSymbol                    string       `json:"STEEM_SYMBOL"`
	VestsSymbol                    string       `json:"VESTS_SYMBOL"`
	BlockchainName                 string       `json:"BLOCKCHAIN_NAME"`
}

//ConversionRequests structure for the GetConversionRequests function.
type ConversionRequests struct {
	ID             *types.Int   `json:"id"`
	Owner          string       `json:"owner"`
	Requestid      *types.Int   `json:"requestid"`
	Amount         *types.Asset `json:"amount"`
	ConversionDate *types.Time  `json:"conversion_date"`
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
	SBDInterestRate          *types.Int   `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      *types.Int   `json:"max_virtual_bandwidth"`
	SBDPrintRate             uint16       `json:"sbd_print_rate"`
	ParticipationCount       uint16       `json:"participation_count"`
	VoteRegenerationPerDay   uint16       `json:"vote_regeneration_per_day"`
}

//VestingDelegationExpiration structure for the GetExpiringVestingDelegations function.
type VestingDelegationExpiration struct {
	ID            uint64       `json:"id"`
	Delegator     string       `json:"delegator"`
	VestingShares *types.Asset `json:"vesting_shares"`
	Expiration    *types.Time  `json:"expiration"`
}

//NextScheduledHardfork structure for the GetNextScheduledHardfork function.
type NextScheduledHardfork struct {
	HfVersion string      `json:"hf_version"`
	LiveTime  *types.Time `json:"live_time"`
}

//OwnerHistory structure for the GetOwnerHistory function.
type OwnerHistory struct {
	ID                     int              `json:"id"`
	Account                string           `json:"account"`
	PreviousOwnerAuthority *types.Authority `json:"previous_owner_authority"`
	LastValidTime          string           `json:"last_valid_time"`
}

//ProposalObject structure for the GetProposedTransaction function.
type ProposalObject struct {
	Author                    string           `json:"author"`
	Title                     string           `json:"title"`
	Memo                      string           `json:"memo"`
	ExpirationTime            *types.Time      `json:"expiration_time"`
	ReviewPeriodTime          *types.Time      `json:"review_period_time"`
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
	ID        *types.ID    `json:"id"`
	From      string       `json:"from"`
	To        string       `json:"to"`
	Memo      string       `json:"memo"`
	RequestID *types.Int   `json:"request_id"`
	Amount    *types.Asset `json:"amount"`
	Complete  *types.Time  `json:"complete"`
}

//VestingDelegation structure for the GetVestingDelegations function.
type VestingDelegation struct {
	ID                uint64       `json:"id"`
	Delegator         string       `json:"delegator"`
	Delegatee         string       `json:"delegatee"`
	VestingShares     *types.Asset `json:"vesting_shares"`
	MinDelegationTime *types.Time  `json:"min_delegation_time"`
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
	EscrowID             uint32       `json:"escrow_id"`
	From                 string       `json:"from"`
	To                   string       `json:"to"`
	Agent                string       `json:"agent"`
	RatificationDeadline *types.Time  `json:"ratification_deadline"`
	EscrowExpiration     *types.Time  `json:"escrow_expiration"`
	SbdAmount            *types.Asset `json:"sbd_amount"`
	SteemAmount          *types.Asset `json:"steem_amount"`
	Fee                  *types.Asset `json:"pending_fee"`
}

//AccountRecoveryRequest structure for the GetRecoveryRequest function.
type AccountRecoveryRequest struct {
	ID                uint             `json:"id"`
	AccountToRecover  string           `json:"account_to_recover"`
	NewOwnerAuthority *types.Authority `json:"new_owner_authority"`
	Extensions        []interface{}    `json:"extensions"`
}
