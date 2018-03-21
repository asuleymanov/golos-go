package client

import (
	"errors"
	"time"

	"github.com/asuleymanov/golos-go/types"
)

//Returns the subscriber's list of subscribers
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

//Returns the list of user subscriptions
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

//Returns the POWER of the user based on the time of the last vote.
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

func (client *Client) GetAuthorReward(username, permlink string) (*types.AuthorRewardOperation, error) {
	resp, err := client.Database.GetAccountHistory(username, -1, 1000)
	if err != nil {
		return nil, err
	} else {
		for k, v := range resp {
			if v.OperationType == "author_reward" {
				op := resp[k].Operation.Data()
				if op.(*types.AuthorRewardOperation).Permlink == permlink {
					return op.(*types.AuthorRewardOperation), nil
				}
			}
		}
		return nil, errors.New("In the last 1000 entries from the user's history, no data was found.")
	}
}
