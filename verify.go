package client

import (
	// Stdlib
	"fmt"
	"log"
	"strings"
	"time"

	// Vendor
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/types"
	"github.com/pkg/errors"
)

//We check whether there is a voter on the list of those who have already voted for the weight of the vote.
func (api *Client) VerifyVoterWeight(author, permlink, voter string, weight int) bool {
	ans, err := api.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Voter: "))
		return false
	} else {
		for _, v := range ans {
			if v.Voter == voter && v.Percent == weight {
				return true
			}
		}
		return false
	}
}

//We check whether there is a voter on the list of those who have already voted without taking into account the weight of the vote.
func (api *Client) VerifyVoter(author, permlink, voter string) bool {
	ans, err := api.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Voter: "))
		return false
	} else {
		for _, v := range ans {
			if v.Voter == voter {
				return true
			}
		}
		return false
	}
}

//We check whether there are voted
func (api *Client) VerifyVotes(author, permlink string) bool {
	ans, err := api.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Votes: "))
		return false
	} else {
		if len(ans) > 0 {
			return true
		} else {
			return false
		}
	}
}

//We check whether the entry in GOLOS is a comment.
func (api *Client) VerifyComments(author, permlink string) bool {
	ans, err := api.Database.GetContentReplies(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Comments: "))
		return false
	} else {
		if len(ans) > 0 {
			return true
		} else {
			return false
		}
	}
}

//Check if the user made a repost entry in GOLOS
func (api *Client) VerifyReblogs(author, permlink, rebloger string) bool {
	ans, err := api.Follow.GetRebloggedBy(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Reblogs: "))
		return false
	} else {
		for _, v := range ans {
			if v == rebloger {
				return true
			}
		}
		return false
	}
}

//Check if one user is signed for the second in GOLOS
func (api *Client) VerifyFollow(follower, following string) bool {
	ans, err := api.Follow.GetFollowing(follower, following, "blog", 1)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Follow: "))
		return false
	} else {
		for _, v := range ans {
			if (v.Follower == follower) && (v.Following == following) {
				return true
			} else {
				return false
			}
		}
		return false
	}
}

//Check if there is an entry in GOLOS
func (api *Client) VerifyPost(author, permlink string) bool {
	ans, err := api.Database.GetContent(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Post: "))
		return false
	} else {
		if (ans.Author == author) && (ans.Permlink == permlink) {
			return true
		} else {
			return false
		}
		return false
	}
}

//We check whether the user has delegated the opportunity to use his post by using operations from a given user.
func (api *Client) VerifyDelegatePostingKeySign(from_user, to_user string) bool {
	acc, err := api.Database.GetAccounts([]string{from_user})
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Delegate Vote Sign: "))
		return false
	} else if len(acc) == 1 {
		for _, v := range acc[0].Posting.AccountAuths {
			tu := strings.Split(strings.Replace(strings.Replace(fmt.Sprintf("%v", v), "[", "", -1), "]", "", -1), " ")
			if tu[0] == to_user {
				if tu[1] == fmt.Sprintf("%v", acc[0].Posting.WeightThreshold) {
					return true
				}
			}
		}
		return false
	} else {
		return false
	}
}

//Check whether the post of the user is his first post in GOLOS
func (api *Client) VerifyFirstPost(username string) bool {
	d := time.Now()
	cont, err := api.Database.GetDiscussionsByAuthorBeforeDate(username, "", d.Format("2006-01-02T00:00:00"), 100)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify First Post: "))
		return false
	} else {
		if len(cont) > 1 {
			return false
		} else {
			return true
		}
		return false
	}
}

//Check if the user exists in GOLOS
func (api *Client) VerifyUser(username string) bool {
	acc, err := api.Database.GetAccounts([]string{username})
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify User: "))
		return false
	} else if len(acc) == 1 {
		return true
	} else {
		return false
	}
}

//We check the possibility of execution of the signed transaction for its execution in GOLOS.
//The check is performed using the standard GetVerifyAuthority API.
func (api *Client) VerifyTrx(username string, strx types.Operation) (bool, error) {
	// Получение необходимых параметров
	props, err := api.Database.GetDynamicGlobalProperties()
	if err != nil {
		return false, err
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return false, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	privKeys := api.SigningKeys(username, strx)

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, api.Chain); err != nil {
		return false, err
	}

	// Отправка транзакции
	resp, err := api.Database.GetVerifyAuthority(tx.Transaction)

	if err != nil {
		return false, err
	} else {
		return resp, nil
	}
}
