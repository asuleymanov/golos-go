package social_network

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/asuleymanov/golos-go/types"
)

//DiscussionQuery
type DiscussionQuery struct {
	Tag            string   `json:"tag"`
	Limit          uint32   `json:"limit"`
	FilterTags     []string `json:"filter_tags"`
	StartAuthor    string   `json:"start_author,omitempty"`
	StartPermlink  string   `json:"start_permlink,omitempty"`
	ParentAuthor   string   `json:"parent_author,omitempty"`
	ParentPermlink string   `json:"parent_permlink"`
}

//Content
type Content struct {
	ID                      *types.ID        `json:"id"`
	Author                  string           `json:"author"`
	Permlink                string           `json:"permlink"`
	Category                string           `json:"category"`
	ParentAuthor            string           `json:"parent_author"`
	ParentPermlink          string           `json:"parent_permlink"`
	Title                   string           `json:"title"`
	Body                    string           `json:"body"`
	JsonMetadata            *ContentMetadata `json:"json_metadata"`
	LastUpdate              *types.Time      `json:"last_update"`
	Created                 *types.Time      `json:"created"`
	Active                  *types.Time      `json:"active"`
	LastPayout              *types.Time      `json:"last_payout"`
	Depth                   *types.Int       `json:"depth"`
	Children                *types.Int       `json:"children"`
	ChildrenRshares2        *types.Int       `json:"children_rshares2"`
	NetRshares              *types.Int       `json:"net_rshares"`
	AbsRshares              *types.Int       `json:"abs_rshares"`
	VoteRshares             *types.Int       `json:"vote_rshares"`
	ChildrenAbsRshares      *types.Int       `json:"children_abs_rshares"`
	CashoutTime             *types.Time      `json:"cashout_time"`
	MaxCashoutTime          *types.Time      `json:"max_cashout_time"`
	TotalVoteWeight         *types.Int       `json:"total_vote_weight"`
	RewardWeight            *types.Int       `json:"reward_weight"`
	TotalPayoutValue        string           `json:"total_payout_value"`
	CuratorPayoutValue      string           `json:"curator_payout_value"`
	AuthorRewards           *types.Int       `json:"author_rewards"`
	NetVotes                *types.Int       `json:"net_votes"`
	RootComment             *types.Int       `json:"root_comment"`
	Mode                    string           `json:"mode"`
	MaxAcceptedPayout       string           `json:"max_accepted_payout"`
	PercentSteemDollars     *types.Int       `json:"percent_steem_dollars"`
	AllowReplies            bool             `json:"allow_replies"`
	AllowVotes              bool             `json:"allow_votes"`
	AllowCurationRewards    bool             `json:"allow_curation_rewards"`
	URL                     string           `json:"url"`
	RootTitle               string           `json:"root_title"`
	PendingPayoutValue      string           `json:"pending_payout_value"`
	TotalPendingPayoutValue string           `json:"total_pending_payout_value"`
	ActiveVotes             []*VoteState     `json:"active_votes"`
	Replies                 []*Content       `json:"replies"`
	AuthorReputation        *types.Int       `json:"author_reputation"`
	Promoted                string           `json:"promoted"`
	BodyLength              *types.Int       `json:"body_length"`
	RebloggedBy             []interface{}    `json:"reblogged_by"`
}

func (content *Content) IsStory() bool {
	return content.ParentAuthor == ""
}

//ContentMetadata
type ContentMetadata struct {
	Flag  bool
	Users []string
	Tags  []string
	Image []string
}

type ContentMetadataRaw struct {
	Users types.StringSlice `json:"users"`
	Tags  types.StringSlice `json:"tags"`
	Image types.StringSlice `json:"image"`
}

func (metadata *ContentMetadata) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	switch unquoted {
	case "true":
		metadata.Flag = true
		return nil
	case "false":
		metadata.Flag = false
		return nil
	}

	if len(unquoted) == 0 {
		var value ContentMetadata
		metadata = &value
		return nil
	}

	var raw ContentMetadataRaw
	if err := json.NewDecoder(strings.NewReader(unquoted)).Decode(&raw); err != nil {
		return err
	}

	metadata.Users = raw.Users
	metadata.Tags = raw.Tags
	metadata.Image = raw.Image

	return nil
}

//TrendingTags
type TrendingTags struct {
	Name                  string     `json:"name"`
	TotalChildrenRshares2 string     `json:"total_children_rshares2"`
	TotalPayouts          string     `json:"total_payouts"`
	NetVotes              *types.Int `json:"net_votes"`
	TopPosts              *types.Int `json:"top_posts"`
	Comments              *types.Int `json:"comments"`
}

//Categories
type Categories struct {
	ID           *types.Int `json:"id"`
	Name         string     `json:"name"`
	AbsRshares   string     `json:"abs_rshares"`
	TotalPayouts string     `json:"total_payouts"`
	Discussions  *types.Int `json:"discussions"`
	LastUpdate   string     `json:"last_update"`
}

//VoteState
type VoteState struct {
	Voter   string      `json:"voter"`
	Weight  *types.Int  `json:"weight"`
	Rshares *types.Int  `json:"rshares"`
	Percent int         `json:"percent"`
	Time    *types.Time `json:"time"`
}

//Votes
type Votes struct {
	Authorperm string      `json:"authorperm"`
	Weight     *types.Int  `json:"weight"`
	Rshares    *types.Int  `json:"rshares"`
	Percent    uint        `json:"percent"`
	Time       *types.Time `json:"time"`
}
