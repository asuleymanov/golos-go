package client

import (
	"errors"
	"time"

	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/types"
)

//SetKeys you can specify keys for signing transactions.
func (client *Client) SetKeys(keys *Keys) {
	client.CurrentKeys = keys
}

//SetAsset returns data of type Asset
func SetAsset(amount float64, symbol string) *types.Asset {
	return &types.Asset{Amount: amount, Symbol: symbol}
}

//FollowersList returns the subscriber's list of subscribers
func (client *Client) FollowersList(username string) ([]string, error) {
	var followers []string
	fc, _ := client.Follow.GetFollowCount(username)
	fccount := fc.FollowerCount
	i := 0
	for i < fccount {
		req, err := client.Follow.GetFollowers(username, "", "blog", 1000)
		if err != nil {
			return followers, err
		}

		for _, v := range req {
			followers = append(followers, v.Follower)
		}
		i = i + 1000
	}

	return followers, nil
}

//FollowingList returns the list of user subscriptions
func (client *Client) FollowingList(username string) ([]string, error) {
	var following []string
	fc, _ := client.Follow.GetFollowCount(username)
	fccount := fc.FollowingCount
	i := 0
	for i < fccount {
		req, err := client.Follow.GetFollowing(username, "", "blog", 100)
		if err != nil {
			return following, err
		}

		for _, v := range req {
			following = append(following, v.Following)
		}
		i = i + 100
	}

	return following, nil
}

//GetVotingPower returns the POWER of the user based on the time of the last vote.
func (client *Client) GetVotingPower(username string) (int, error) {
	conf, errc := client.Database.GetConfig()
	if errc != nil {
		return 0, errc
	}

	acc, erra := client.Database.GetAccounts([]string{username})
	if erra != nil {
		return 0, erra
	}

	vp := acc[0].VotingPower
	lvt := acc[0].LastVoteTime
	dtn := time.Now()

	regen := conf.Steemit100Percent * int(dtn.Sub(*lvt.Time).Seconds()) / conf.SteemitVoteRegenerationSeconds
	power := (vp + regen) // 100
	if power > 10000 {
		power = 10000
	}
	return power, nil
}

//GetAuthorReward returns information about the reward received for publication.
func (client *Client) GetAuthorReward(username, permlink string, full bool) (*types.AuthorRewardOperation, error) {
	if !full {
		resp, err := client.AccountHistory.GetAccountHistory(username, -1, 1000)
		if err != nil {
			return nil, err
		}
		for k, v := range resp {
			if v.OperationType == "author_reward" {
				op := resp[k].Operation.Data()
				if op.(*types.AuthorRewardOperation).Permlink == permlink {
					return op.(*types.AuthorRewardOperation), nil
				}
			}
		}
		return nil, errors.New("In the last 1000 entries from the user's history, no data was found")
	}
	from := 1000
	limit := 1000
	var lastBlock uint32
	for {
		ans, err := client.AccountHistory.GetAccountHistory(username, int64(from), uint32(limit))
		if err != nil {
			return nil, err
		}

		if len(ans) == 0 {
			break
		}
		for k, v := range ans {
			if v.OperationType == "author_reward" {
				op := ans[k].Operation.Data()
				if op.(*types.AuthorRewardOperation).Permlink == permlink {
					return op.(*types.AuthorRewardOperation), nil
				}
			}
		}
		s := ans[len(ans)-1:]
		block := s[0].BlockNumber
		if block == lastBlock {
			break
		}

		lastBlock = block
		from = from + limit
	}
	return nil, errors.New("Data about the publication is not found in the entire history of the user's actions")
}

//GetCommentOptionsOperation generates CommentOptionsOperation depending on the incoming data
func GetCommentOptionsOperation(username, permlink string, options PCOptions) *types.CommentOptionsOperation {
	var ext []interface{}
	var av, acr bool
	var vMap *types.Asset
	symbol := "GBG"
	vMap = SetAsset(1000000.000, symbol)
	psd := options.Percent
	extens := []interface{}{}

	switch options.Percent {
	case 0:
		vMap = SetAsset(0.000, symbol)
		psd = 10000
	case 50:
		psd = 10000
	default:
		psd = 0
	}

	if options.AllowVotes == nil || *options.AllowVotes {
		av = true
	}

	if options.AllowCurationRewards == nil || *options.AllowCurationRewards {
		acr = true
	}

	if options.BeneficiaryList != nil && len(*options.BeneficiaryList) > 0 {
		var benList []types.Beneficiary
		var benef types.CommentPayoutBeneficiaries
		for _, val := range *options.BeneficiaryList {
			benList = append(benList, types.Beneficiary{val.Account, val.Weight})
		}
		benef.Beneficiaries = benList
		ext = append(ext, 0, benef)
	}

	if len(ext) > 0 {
		extens = []interface{}{ext}
	}

	return &types.CommentOptionsOperation{
		Author:               username,
		Permlink:             permlink,
		MaxAcceptedPayout:    vMap,
		PercentSteemDollars:  psd,
		AllowVotes:           av,
		AllowCurationRewards: acr,
		Extensions:           extens,
	}
}

//GetPostBandwidth returns the real (calculated) value of the post_bandwidth parameter.
func (client *Client) GetPostBandwidth(username string) (int64, error) {
	minutesPerDay := float64(1440)

	resp, err := client.Database.GetAccounts([]string{username})
	if err != nil {
		return 0, err
	}

	oldPostBandwidth := float64(resp[0].PostBandwidth)
	deltaTimeMinutes := float64(time.Now().Sub(*resp[0].LastRootPost.Time).Minutes())

	newPostBandwidth := ((minutesPerDay - deltaTimeMinutes) / minutesPerDay) * oldPostBandwidth

	return int64(newPostBandwidth), nil
}

func JSONTrxString(v *transactions.SignedTransaction) string {
	ans, _ := types.JSONMarshal(v)
	return string(ans)
}

func JSONOpString(v []types.Operation) string {
	var tx types.Operations

	tx = append(tx, v...)

	ans, _ := types.JSONMarshal(tx)
	return string(ans)
}
