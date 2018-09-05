package social_network

import (
	"github.com/asuleymanov/golos-go/types"
)

//Votes structure for the GetAccountVotes function.
type Votes struct {
	Authorperm string      `json:"authorperm"`
	Weight     *types.Int  `json:"weight"`
	Rshares    *types.Int  `json:"rshares"`
	Percent    uint        `json:"percent"`
	Time       *types.Time `json:"time"`
}

//Beneficiarie aditional structure for Content struct.
type Beneficiarie struct {
	Account string `json:"account"`
	Weight  int    `json:"weight"`
}

//VoteState structure for the GetActiveVotes function.
type VoteState struct {
	Voter   string      `json:"voter"`
	Weight  *types.Int  `json:"weight"`
	Rshares *types.Int  `json:"rshares"`
	Percent int         `json:"percent"`
	Time    *types.Time `json:"time"`
}

//Content structure for the GetContent, GetContentReplies and GetRepliesByLastUpdate function.
type Content struct {
		ID                          *types.ID    `json:"id"`
	Author                      string `json:"author"`
	Permlink                    string `json:"permlink"`
	ParentAuthor                string `json:"parent_author"`
	ParentPermlink              string `json:"parent_permlink"`
	Category                    string `json:"category"`
	Title                       string `json:"title"`
	Body                        string `json:"body"`
	JSONMetadata                *types.ContentMetadata `json:"json_metadata"`
	LastUpdate                  *types.Time `json:"last_update"`
	Created                     *types.Time `json:"created"`
	Active                      *types.Time `json:"active"`
	LastPayout                  *types.Time `json:"last_payout"`
	Depth                       *types.Int    `json:"depth"`
	Children                    *types.Int    `json:"children"`
	ChildrenRshares2            *types.Int `json:"children_rshares2"`
	NetRshares                  *types.Int `json:"net_rshares"`
	AbsRshares                  *types.Int `json:"abs_rshares"`
	VoteRshares                 *types.Int `json:"vote_rshares"`
	ChildrenAbsRshares          *types.Int `json:"children_abs_rshares"`
	CashoutTime                 *types.Time `json:"cashout_time"`
	MaxCashoutTime              *types.Time `json:"max_cashout_time"`
	TotalVoteWeight             *types.Int `json:"total_vote_weight"`
	RewardWeight                *types.Int    `json:"reward_weight"`
	TotalPayoutValue            *types.Asset `json:"total_payout_value"`
	BeneficiaryPayoutValue      *types.Asset `json:"beneficiary_payout_value"`
	BeneficiaryGestsPayoutValue *types.Asset `json:"beneficiary_gests_payout_value"`
	CuratorPayoutValue          *types.Asset `json:"curator_payout_value"`
	CuratorGestsPayoutValue     *types.Asset `json:"curator_gests_payout_value"`
	AuthorRewards               *types.Int    `json:"author_rewards"`
	AuthorGbgPayoutValue        *types.Asset `json:"author_gbg_payout_value"`
	AuthorGolosPayoutValue      *types.Asset `json:"author_golos_payout_value"`
	AuthorGestsPayoutValue      *types.Asset `json:"author_gests_payout_value"`
	NetVotes                    *types.Int    `json:"net_votes"`
	Mode                        string `json:"mode"`
	RootComment                 *types.Int    `json:"root_comment"`
	RootTitle                   string `json:"root_title"`
	MaxAcceptedPayout           *types.Asset `json:"max_accepted_payout"`
	PercentSteemDollars         *types.Int    `json:"percent_steem_dollars"`
	AllowReplies                bool   `json:"allow_replies"`
	AllowVotes                  bool   `json:"allow_votes"`
	AllowCurationRewards        bool   `json:"allow_curation_rewards"`
	Beneficiaries               []*Beneficiarie `json:"beneficiaries"`
	URL                               string        `json:"url"`
	PendingAuthorPayoutValue          *types.Asset        `json:"pending_author_payout_value"`
	PendingAuthorPayoutGbgValue       *types.Asset        `json:"pending_author_payout_gbg_value"`
	PendingAuthorPayoutGestsValue     *types.Asset        `json:"pending_author_payout_gests_value"`
	PendingAuthorPayoutGolosValue     *types.Asset        `json:"pending_author_payout_golos_value"`
	PendingBenefactorPayoutValue      *types.Asset        `json:"pending_benefactor_payout_value"`
	PendingBenefactorPayoutGestsValue *types.Asset        `json:"pending_benefactor_payout_gests_value"`
	PendingCuratorPayoutValue         *types.Asset        `json:"pending_curator_payout_value"`
	PendingCuratorPayoutGestsValue    *types.Asset        `json:"pending_curator_payout_gests_value"`
	PendingPayoutValue                *types.Asset        `json:"pending_payout_value"`
	TotalPendingPayoutValue           *types.Asset        `json:"total_pending_payout_value"`
	ActiveVotes                       []*VoteState `json:"active_votes"`
	ActiveVotesCount                  *types.Int           `json:"active_votes_count"`
	Replies                           []*Content `json:"replies"`
	AuthorReputation                  *types.Int        `json:"author_reputation"`
	Promoted                          *types.Asset        `json:"promoted"`
	BodyLength                        *types.Int           `json:"body_length"`
	RebloggedBy                       []interface{} `json:"reblogged_by"`
	ReblogEntries                     []interface{} `json:"reblog_entries"`
}

//IsStory operation that determines that Content is a publication.
func (content *Content) IsStory() bool {
	return content.ParentAuthor == ""
}
