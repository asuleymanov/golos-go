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

//VerifyVoterWeight check whether there is a voter on the list of those who have already voted for the weight of the vote.
func (client *Client) VerifyVoterWeight(author, permlink, voter string, weight int) bool {
	ans, err := client.SocialNetwork.GetActiveVotes(author, permlink)

	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Voter: "))
		return false
	}
	for _, v := range ans {
		if v.Voter == voter && v.Percent == weight {
			return true
		}
	}
	return false
}

//VerifyVoter check whether there is a voter on the list of those who have already voted without taking into account the weight of the vote.
func (client *Client) VerifyVoter(author, permlink, voter string) bool {
	ans, err := client.SocialNetwork.GetActiveVotes(author, permlink)

	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Voter: "))
		return false
	}
	for _, v := range ans {
		if v.Voter == voter {
			return true
		}
	}
	return false
}

//VerifyVotes check whether there are voted
func (client *Client) VerifyVotes(author, permlink string) bool {
	ans, err := client.SocialNetwork.GetActiveVotes(author, permlink)

	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Votes: "))
		return false
	}
	if len(ans) > 0 {
		return true
	}
	return false
}

//VerifyComments check whether the entry in GOLOS is a comment.
func (client *Client) VerifyComments(author, permlink string) bool {
	ans, err := client.SocialNetwork.GetContentReplies(author, permlink)

	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Comments: "))
		return false
	}
	if len(ans) > 0 {
		return true
	}
	return false
}

//VerifyReblogs сheck if the user made a repost entry in GOLOS
func (client *Client) VerifyReblogs(author, permlink, rebloger string) bool {
	ans, err := client.Follow.GetRebloggedBy(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Reblogs: "))
		return false
	}
	for _, v := range ans {
		if v == rebloger {
			return true
		}
	}
	return false
}

//VerifyFollow сheck if one user is signed for the second in GOLOS
func (client *Client) VerifyFollow(follower, following string) bool {
	ans, err := client.Follow.GetFollowing(follower, following, "blog", 1)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Follow: "))
		return false
	}
	for _, v := range ans {
		if (v.Follower == follower) && (v.Following == following) {
			return true
		}
	}
	return false
}

//VerifyPost сheck if there is an entry in GOLOS
func (client *Client) VerifyPost(author, permlink string) bool {
	ans, err := client.SocialNetwork.GetContent(author, permlink)

	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Post: "))
		return false
	}
	if (ans.Author == author) && (ans.Permlink == permlink) {
		return true
	}
	return false
}

//VerifyDelegatePostingKeySign check whether the user has delegated the opportunity to use his post by using operations from a given user.
func (client *Client) VerifyDelegatePostingKeySign(fromUser, toUser string) bool {
	acc, err := client.Database.GetAccounts([]string{fromUser})
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Delegate Vote Sign: "))
		return false
	}

	if len(acc) == 1 {
		for _, v := range acc[0].Posting.AccountAuths {
			tu := strings.Split(strings.Replace(strings.Replace(fmt.Sprintf("%v", v), "[", "", -1), "]", "", -1), " ")
			if tu[0] == toUser {
				if tu[1] == fmt.Sprintf("%v", acc[0].Posting.WeightThreshold) {
					return true
				}
			}
		}
	}
	return false
}

//VerifyFirstPost сheck whether the post of the user is his first post in GOLOS
func (client *Client) VerifyFirstPost(username string) bool {
	d := time.Now()
	cont, err := client.Tags.GetDiscussionsByAuthorBeforeDate(username, "", d.Format("2006-01-02T00:00:00"), 100)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify First Post: "))
		return false
	}

	return len(cont) <= 1
}

//VerifyUser сheck if the user exists in GOLOS
func (client *Client) VerifyUser(username string) bool {
	acc, err := client.Database.GetAccounts([]string{username})
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify User: "))
		return false
	}

	return len(acc) == 1
}

//VerifyTrx check the possibility of execution of the signed transaction for its execution in GOLOS.
//The check is performed using the standard GetVerifyAuthority API.
func (client *Client) VerifyTrx(username string, strx types.Operation) (bool, error) {
	// Получение необходимых параметров
	props, err := client.Database.GetDynamicGlobalProperties()
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
	privKeys, err := client.SigningKeys(strx)
	if err != nil {
		return false, err
	}

	// Подписываем транзакцию
	if err := tx.Sign(privKeys, client.Chain); err != nil {
		return false, err
	}

	// Отправка транзакции
	resp, err := client.Database.GetVerifyAuthority(tx.Transaction)

	if err != nil {
		return false, err
	}
	return resp, nil
}
