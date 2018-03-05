package client

import (
	// Stdlib
	"log"
	"strconv"
	"strings"
	"time"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/wif"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/translit"
	"github.com/asuleymanov/golos-go/types"
)

const fdt = `"20060102t150405"`

func (api *Client) Vote(username, authorname, permlink string, weight int) error {
	if weight > 10000 {
		weight = 10000
	}
	if api.VerifyVoterWeight(authorname, permlink, username, weight) {
		return errors.New("The voter is on the list")
	}

	var trx []types.Operation

	tx := &types.VoteOperation{
		Voter:    username,
		Author:   authorname,
		Permlink: permlink,
		Weight:   types.Int16(weight),
	}
	trx = append(trx, tx)

	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Vote: ")
	} else {
		log.Println("[Vote] Block -> ", resp.BlockNum, " User -> ", username)
		return nil
	}
}

func (api *Client) MultiVote(username, author, permlink string, arrvote []ArrVote) error {
	var trx []types.Operation
	var arrvotes []ArrVote

	for _, v := range arrvote {
		if api.VerifyDelegatePostingKeySign(v.User, username) && !api.VerifyVoter(author, permlink, v.User) {
			arrvotes = append(arrvotes, v)
		}
	}

	if len(arrvotes) == 0 {
		return errors.New("Error Multi_Vote : All users from the list have already voted.")
	}

	for _, val := range arrvotes {
		txt := &types.VoteOperation{
			Voter:    val.User,
			Author:   author,
			Permlink: permlink,
			Weight:   types.Int16(val.Weight),
		}
		trx = append(trx, txt)
	}

	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Multi_Vote: ")
	} else {
		log.Println("[Multi_Vote] Block -> ", resp.BlockNum, " User/Permlink -> ", author, "/", permlink)
		return nil
	}
}

func (api *Client) Comment(username, authorname, ppermlink, body string, v *PCVote, o *PCOptions) error {
	var trx []types.Operation

	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + authorname + "-" + ppermlink + "-" + times
	permlink = strings.Replace(permlink, ".", "-", -1)

	tx := &types.CommentOperation{
		ParentAuthor:   authorname,
		ParentPermlink: ppermlink,
		Author:         username,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JsonMetadata:   "{\"app\":\"golos\"}",
	}
	trx = append(trx, tx)

	if o != nil {
		symbol := "GBG"
		MAP := "1000000.000 " + symbol
		PSD := o.Percent
		if o.Percent == 0 {
			MAP = "0.000 " + symbol
			PSD = 10000
		} else if o.Percent == 50 {
			PSD = 10000
		} else {
			PSD = 0
		}

		txo := &types.CommentOptionsOperation{
			Author:               username,
			Permlink:             permlink,
			MaxAcceptedPayout:    MAP,
			PercentSteemDollars:  PSD,
			AllowVotes:           true,
			AllowCurationRewards: true,
			Extensions:           []interface{}{},
		}
		trx = append(trx, txo)
	}

	if v != nil && v.Weight != 0 {
		txv := &types.VoteOperation{
			Voter:    username,
			Author:   username,
			Permlink: permlink,
			Weight:   types.Int16(v.Weight),
		}
		trx = append(trx, txv)
	}

	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Comment : ")
	} else {
		log.Println("[Comment] Block -> ", resp.BlockNum, " User -> ", username)
		return nil
	}
}

func (api *Client) CommentLink(username, authorname, ppermlink, body string, v *PCVote, o *PCOptions) (string, error) {
	var trx []types.Operation

	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + authorname + "-" + ppermlink + "-" + times
	permlink = strings.Replace(permlink, ".", "-", -1)

	tx := &types.CommentOperation{
		ParentAuthor:   authorname,
		ParentPermlink: ppermlink,
		Author:         username,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JsonMetadata:   "{\"app\":\"golos\"}",
	}
	trx = append(trx, tx)

	if o != nil {
		symbol := "GBG"
		MAP := "1000000.000 " + symbol
		PSD := o.Percent
		if o.Percent == 0 {
			MAP = "0.000 " + symbol
			PSD = 10000
		} else if o.Percent == 50 {
			PSD = 10000
		} else {
			PSD = 0
		}

		txo := &types.CommentOptionsOperation{
			Author:               username,
			Permlink:             permlink,
			MaxAcceptedPayout:    MAP,
			PercentSteemDollars:  PSD,
			AllowVotes:           true,
			AllowCurationRewards: true,
			Extensions:           []interface{}{},
		}
		trx = append(trx, txo)
	}

	if v != nil && v.Weight != 0 {
		txv := &types.VoteOperation{
			Voter:    username,
			Author:   username,
			Permlink: permlink,
			Weight:   types.Int16(v.Weight),
		}
		trx = append(trx, txv)
	}

	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return "", errors.Wrapf(err, "Error Comment : ")
	} else {
		log.Println("[Comment] Block -> ", resp.BlockNum, " User -> ", username)
		return permlink, nil
	}
}

func (api *Client) DeleteComment(authorname, permlink string) error {
	if api.VerifyVotes(authorname, permlink) {
		return errors.New("You can not delete already there are voted")
	}
	if api.VerifyComments(authorname, permlink) {
		return errors.New("You can not delete already have comments")
	}
	var trx []types.Operation

	tx := &types.DeleteCommentOperation{
		Author:   authorname,
		Permlink: permlink,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(authorname, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Delete Comment: ")
	} else {
		log.Println("[Delete Comment] Block -> ", resp.BlockNum, " User -> ", authorname)
		return nil
	}
}

func (api *Client) Post(authorname, title, body, permlink, ptag, post_image string, tags []string, v *PCVote, o *PCOptions) error {
	if permlink == "" {
		permlink = translit.EncodeTitle(title)
	} else {
		permlink = translit.EncodeTitle(permlink)
	}
	tag := translit.EncodeTags(tags)
	if ptag == "" {
		ptag = translit.EncodeTag(tags[0])
	} else {
		ptag = translit.EncodeTag(ptag)
	}

	json_meta := "{\"tags\":["
	for k, v := range tag {
		if k != len(tags)-1 {
			json_meta = json_meta + "\"" + v + "\","
		} else {
			json_meta = json_meta + "\"" + v + "\"]"
		}
	}
	if post_image != "" {
		json_meta = json_meta + ",\"image\":[\"" + post_image + "\"]"
	}
	json_meta = json_meta + ",\"app\":\"golos\"}"

	var trx []types.Operation
	txp := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         authorname,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JsonMetadata:   json_meta,
	}
	trx = append(trx, txp)

	if o != nil {
		symbol := "GBG"
		MAP := "1000000.000 " + symbol
		PSD := o.Percent
		if o.Percent == 0 {
			MAP = "0.000 " + symbol
			PSD = 10000
		} else if o.Percent == 50 {
			PSD = 10000
		} else {
			PSD = 0
		}

		txo := &types.CommentOptionsOperation{
			Author:               authorname,
			Permlink:             permlink,
			MaxAcceptedPayout:    MAP,
			PercentSteemDollars:  PSD,
			AllowVotes:           true,
			AllowCurationRewards: true,
			Extensions:           []interface{}{},
		}
		trx = append(trx, txo)
	}

	if v != nil && v.Weight != 0 {
		txv := &types.VoteOperation{
			Voter:    authorname,
			Author:   authorname,
			Permlink: permlink,
			Weight:   types.Int16(v.Weight),
		}
		trx = append(trx, txv)
	}

	resp, err := api.SendTrx(authorname, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Post : ")
	} else {
		log.Println("[Post] Block -> ", resp.BlockNum, " User -> ", authorname)
		return nil
	}
}

func (api *Client) Follows(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[\"blog\"]}]"
	var trx []types.Operation

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(follower, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Follow: ")
	} else {
		log.Println("[Follow] Block -> ", resp.BlockNum, " Follower user -> ", follower, " Following user -> ", following)
		return nil
	}
}

func (api *Client) Unfollow(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[]}]"
	var trx []types.Operation

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(follower, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Unfollow: ")
	} else {
		log.Println("[Unfollow] Block -> ", resp.BlockNum, " Unfollower user -> ", follower, " Unfollowing user -> ", following)
		return nil
	}
}

func (api *Client) Ignore(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[\"ignore\"]}]"
	var trx []types.Operation

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(follower, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Ignore: ")
	} else {
		log.Println("[Ignore] Block -> ", resp.BlockNum, " Ignore user -> ", follower, " Ignoring user -> ", following)
		return nil
	}
}

func (api *Client) Notice(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[]}]"
	var trx []types.Operation

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(follower, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Notice: ")
	} else {
		log.Println("[Notice] Block -> ", resp.BlockNum, " Notice user -> ", follower, " Noticing user -> ", following)
		return nil
	}
}

func (api *Client) Reblog(username, authorname, permlink string) error {
	if api.VerifyReblogs(authorname, permlink, username) {
		return errors.New("The user already did repost")
	}
	json_string := "[\"reblog\",{\"account\":\"" + username + "\",\"author\":\"" + authorname + "\",\"permlink\":\"" + permlink + "\"}]"
	var trx []types.Operation

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{username},
		ID:                   "follow",
		JSON:                 json_string,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("[Reblog] Block -> ", resp.BlockNum, " Reblog user -> ", username, " Rebloging -> ", authorname, "/", permlink)
		return nil
	}
}

func (api *Client) AccountWitnessVote(username, witness_name string, approv bool) error {
	var trx []types.Operation

	tx := &types.AccountWitnessVoteOperation{
		Account: username,
		Witness: witness_name,
		Approve: approv,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return errors.Wrapf(err, "Error AccountWitnessVote: ")
	} else {
		log.Println("[AccountWitnessVote] Block -> ", resp.BlockNum, " User -> ", username)
		return nil
	}
}

func (api *Client) AccountWitnessProxy(username, proxy string) error {
	var trx []types.Operation

	tx := &types.AccountWitnessProxyOperation{
		Account: username,
		Proxy:   proxy,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return errors.Wrapf(err, "Error AccountWitnessProxy: ")
	} else {
		log.Println("[AccountWitnessProxy] Block -> ", resp.BlockNum, " User -> ", username)
		return nil
	}
}

func (api *Client) Transfer(from_name, to_name, memo, ammount string) error {
	var trx []types.Operation

	tx := &types.TransferOperation{
		From:   from_name,
		To:     to_name,
		Amount: ammount,
		Memo:   memo,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from_name, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Transfer: ")
	} else {
		log.Println("[Transfer] Block -> ", resp.BlockNum, " From user -> ", from_name, " To user -> ", to_name)
		return nil
	}
}

func (api *Client) MultiTransfer(username string, arrtrans []ArrTransfer) error {
	var trx []types.Operation

	for _, val := range arrtrans {
		txt := &types.TransferOperation{
			From:   username,
			To:     val.To,
			Amount: val.Ammount,
			Memo:   val.Memo,
		}
		trx = append(trx, txt)
	}

	resp, err := api.SendTrx(username, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Multi_Transfer: ")
	} else {
		log.Println("[Multi_Transfer] Block -> ", resp.BlockNum, " From user -> ", username)
		return nil
	}
}

func (api *Client) Login(username, key string) bool {
	json_string := "[\"login\",{\"account\":\"" + username + "\",\"app\":\"golos-go\"}]"

	strx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{username},
		ID:                   "login",
		JSON:                 json_string,
	}

	props, err := api.Database.GetDynamicGlobalProperties()
	if err != nil {
		return false
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return false
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	var keys [][]byte
	privKey, _ := wif.Decode(string([]byte(key)))
	keys = append(keys, privKey)

	// Подписываем транзакцию
	if err := tx.Sign(keys, api.Chain); err != nil {
		return false
	}

	// Отправка транзакции
	resp, err := api.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return false
	} else {
		log.Println("[Login] Block -> ", resp.BlockNum, " User -> ", username)
		return true
	}
}

func (api *Client) LimitOrderCancel(owner string, orderid uint32) error {
	var trx []types.Operation

	tx := &types.LimitOrderCancelOperation{
		Owner:   owner,
		OrderID: orderid,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(owner, trx)
	if err != nil {
		return errors.Wrapf(err, "Error LimitOrderCancel: ")
	} else {
		log.Println("[LimitOrderCancel] Block -> ", resp.BlockNum, " LimitOrderCancel user -> ", owner)
		return nil
	}
}

func (api *Client) LimitOrderCreate(owner, sell, buy string, orderid uint32) error {
	var trx []types.Operation

	expiration := time.Now().Add(3600000 * time.Second).UTC()
	fok := false

	tx := &types.LimitOrderCreateOperation{
		Owner:        owner,
		OrderID:      orderid,
		AmountToSell: sell,
		MinToReceive: buy,
		FillOrKill:   fok,
		Expiration:   &types.Time{&expiration},
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(owner, trx)
	if err != nil {
		return errors.Wrapf(err, "Error LimitOrderCreate: ")
	} else {
		log.Println("[LimitOrderCreate] Block -> ", resp.BlockNum, " LimitOrderCreate user -> ", owner)
		return nil
	}
}

func (api *Client) Convert(owner, amount string, requestid uint32) error {
	var trx []types.Operation

	tx := &types.ConvertOperation{
		Owner:     owner,
		RequestID: requestid,
		Amount:    amount,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(owner, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Convert: ")
	} else {
		log.Println("[Convert] Block -> ", resp.BlockNum, " Convert user -> ", owner)
		return nil
	}
}

func (api *Client) TransferToVesting(from, to, amount string) error {
	var trx []types.Operation

	tx := &types.TransferToVestingOperation{
		From:   from,
		To:     to,
		Amount: amount,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from, trx)
	if err != nil {
		return errors.Wrapf(err, "Error TransferToVesting: ")
	} else {
		log.Println("[TransferToVesting] Block -> ", resp.BlockNum, " TransferToVesting user -> ", from)
		return nil
	}
}

func (api *Client) WithdrawVesting(account, vshares string) error {
	var trx []types.Operation

	tx := &types.WithdrawVestingOperation{
		Account:       account,
		VestingShares: vshares,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(account, trx)
	if err != nil {
		return errors.Wrapf(err, "Error WithdrawVesting: ")
	} else {
		log.Println("[WithdrawVesting] Block -> ", resp.BlockNum, " WithdrawVesting user -> ", account)
		return nil
	}
}

func (api *Client) ChangeRecoveryAccount(accounttorecover, newrecoveryaccount string) error {
	var trx []types.Operation

	tx := &types.ChangeRecoveryAccountOperation{
		AccountToRecover:   accounttorecover,
		NewRecoveryAccount: newrecoveryaccount,
		Extensions:         []interface{}{},
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(accounttorecover, trx)
	if err != nil {
		return errors.Wrapf(err, "Error ChangeRecoveryAccount: ")
	} else {
		log.Println("[ChangeRecoveryAccount] Block -> ", resp.BlockNum, " ChangeRecoveryAccount user -> ", accounttorecover)
		return nil
	}
}

func (api *Client) TransferToSavings(from, to, amount, memo string) error {
	var trx []types.Operation

	tx := &types.TransferToSavingsOperation{
		From:   from,
		To:     to,
		Amount: amount,
		Memo:   memo,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from, trx)
	if err != nil {
		return errors.Wrapf(err, "Error TransferToSavings: ")
	} else {
		log.Println("[TransferToSavings] Block -> ", resp.BlockNum, " TransferToSavings user -> ", from)
		return nil
	}
}

func (api *Client) TransferFromSavings(from, to, amount, memo string, requestid uint32) error {
	var trx []types.Operation

	tx := &types.TransferFromSavingsOperation{
		From:      from,
		RequestId: requestid,
		To:        to,
		Amount:    amount,
		Memo:      memo,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from, trx)
	if err != nil {
		return errors.Wrapf(err, "Error TransferFromSavings: ")
	} else {
		log.Println("[TransferFromSavings] Block -> ", resp.BlockNum, " TransferFromSavings user -> ", from)
		return nil
	}
}

func (api *Client) CancelTransferFromSavings(from string, requestid uint32) error {
	var trx []types.Operation

	tx := &types.CancelTransferFromSavingsOperation{
		From:      from,
		RequestId: requestid,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from, trx)
	if err != nil {
		return errors.Wrapf(err, "Error CancelTransferFromSavings: ")
	} else {
		log.Println("[CancelTransferFromSavings] Block -> ", resp.BlockNum, " CancelTransferFromSavings user -> ", from)
		return nil
	}
}

func (api *Client) DeclineVotingRights(account string, decline bool) error {
	var trx []types.Operation

	tx := &types.DeclineVotingRightsOperation{
		Account: account,
		Decline: decline,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(account, trx)
	if err != nil {
		return errors.Wrapf(err, "Error DeclineVotingRights: ")
	} else {
		log.Println("[DeclineVotingRights] Block -> ", resp.BlockNum, " DeclineVotingRights user -> ", account)
		return nil
	}
}

func (api *Client) FeedPublish(publisher, base, quote string) error {
	var trx []types.Operation

	tx := &types.FeedPublishOperation{
		Publisher: publisher,
		ExchangeRate: types.ExchRate{
			Base:  base,
			Quote: quote,
		},
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(publisher, trx)
	if err != nil {
		return errors.Wrapf(err, "Error FeedPublish: ")
	} else {
		log.Println("[FeedPublish] Block -> ", resp.BlockNum, " FeedPublish user -> ", publisher)
		return nil
	}
}

func (api *Client) WitnessUpdate(owner, url, blocksigningkey, accountcreationfee string, maxblocksize uint32, sbdinterestrate uint16) error {
	var trx []types.Operation

	tx := &types.WitnessUpdateOperation{
		Owner:           owner,
		Url:             url,
		BlockSigningKey: blocksigningkey,
		Props: types.ChainProperties{
			AccountCreationFee: accountcreationfee,
			MaximumBlockSize:   maxblocksize,
			SBDInterestRate:    sbdinterestrate,
		},
		Fee: "0.000 GOLOS",
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(owner, trx)
	if err != nil {
		return errors.Wrapf(err, "Error WitnessUpdate: ")
	} else {
		log.Println("[WitnessUpdate] Block -> ", resp.BlockNum, " WitnessUpdate user -> ", owner)
		return nil
	}
}
