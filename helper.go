package client

import (
	"time"
)

func (api *Client) FollowersList(username string) ([]string, error) {
	var followers []string
	fc, _ := api.Follow.GetFollowCount(username)
	fccount := fc.FollowerCount
	i := 0
	for i < fccount {
		req, err := api.Follow.GetFollowers(username, "", "blog", 1000)
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

func (api *Client) FollowingList(username string) ([]string, error) {
	var following []string
	fc, _ := api.Follow.GetFollowCount(username)
	fccount := fc.FollowingCount
	i := 0
	for i < fccount {
		req, err := api.Follow.GetFollowing(username, "", "blog", 100)
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

func (api *Client) GetVotingPower(username string) (int, error) {
	conf, errc := api.Database.GetConfig()
	if errc != nil {
		return 0, errc
	}

	acc, erra := api.Database.GetAccounts([]string{username})
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
