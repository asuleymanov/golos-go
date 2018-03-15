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

func (api *Client) Vote(username, authorname, permlink string, weight int) (*OperResp, error) {
	if weight > 10000 {
		weight = 10000
	}
	if api.VerifyVoterWeight(authorname, permlink, username, weight) {
		return nil, errors.New("The voter is on the list")
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
	return &OperResp{NameOper: "Vote", Bresp: resp}, err
}

func (api *Client) MultiVote(username, author, permlink string, arrvote []ArrVote) (*OperResp, error) {
	var trx []types.Operation
	var arrvotes []ArrVote

	for _, v := range arrvote {
		if api.VerifyDelegatePostingKeySign(v.User, username) && !api.VerifyVoter(author, permlink, v.User) {
			arrvotes = append(arrvotes, v)
		}
	}

	if len(arrvotes) == 0 {
		return nil, errors.New("Error Multi_Vote : All users from the list have already voted.")
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
	return &OperResp{NameOper: "MultiVote", Bresp: resp}, err
}

func (api *Client) Comment(username, authorname, ppermlink, body string, o *PCOptions) (*OperResp, error) {
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

	resp, err := api.SendTrx(username, trx)
	return &OperResp{NameOper: "Comment", PermLink: permlink, Bresp: resp}, err
}

func (api *Client) Post(authorname, title, body, permlink, ptag, post_image string, tags []string, o *PCOptions) (*OperResp, error) {
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

	resp, err := api.SendTrx(authorname, trx)
	return &OperResp{NameOper: "Post", PermLink: permlink, Bresp: resp}, err
}

func (api *Client) DeleteComment(authorname, permlink string) (*OperResp, error) {
	if api.VerifyVotes(authorname, permlink) {
		return nil, errors.New("You can not delete already there are voted")
	}
	if api.VerifyComments(authorname, permlink) {
		return nil, errors.New("You can not delete already have comments")
	}
	var trx []types.Operation

	tx := &types.DeleteCommentOperation{
		Author:   authorname,
		Permlink: permlink,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(authorname, trx)
	return &OperResp{NameOper: "Delete Comment/Post", Bresp: resp}, err
}

//Subscribe to the user
func (api *Client) Follows(follower, following string) (*OperResp, error) {
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
	return &OperResp{NameOper: "Follows", Bresp: resp}, err
}

//Unsubscribe from user
func (api *Client) Unfollow(follower, following string) (*OperResp, error) {
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
	return &OperResp{NameOper: "Unfollow", Bresp: resp}, err
}

//Ignore user
func (api *Client) Ignore(follower, following string) (*OperResp, error) {
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
	return &OperResp{NameOper: "Unfollow", Bresp: resp}, err
}

//Undo ignore the user
func (api *Client) Notice(follower, following string) (*OperResp, error) {
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
	return &OperResp{NameOper: "Notice", Bresp: resp}, err
}

//Repost records
func (api *Client) Reblog(username, authorname, permlink string) (*OperResp, error) {
	if api.VerifyReblogs(authorname, permlink, username) {
		return nil, errors.New("The user already did repost")
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
	return &OperResp{NameOper: "Reblog", Bresp: resp}, err
}

//Operation of voting for the delegate.
func (api *Client) AccountWitnessVote(username, witness_name string, approv bool) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.AccountWitnessVoteOperation{
		Account: username,
		Witness: witness_name,
		Approve: approv,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(username, trx)
	return &OperResp{NameOper: "AccountWitnessVote", Bresp: resp}, err
}

//Transfer of the right to vote for delegates to another user.
func (api *Client) AccountWitnessProxy(username, proxy string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.AccountWitnessProxyOperation{
		Account: username,
		Proxy:   proxy,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(username, trx)
	return &OperResp{NameOper: "AccountWitnessProxy", Bresp: resp}, err
}

//Transfer of funds to any user.
func (api *Client) Transfer(from_name, to_name, memo, ammount string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferOperation{
		From:   from_name,
		To:     to_name,
		Amount: ammount,
		Memo:   memo,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from_name, trx)
	return &OperResp{NameOper: "Transfer", Bresp: resp}, err
}

//Multiple funds transfer in one transaction.
func (api *Client) MultiTransfer(username string, arrtrans []ArrTransfer) (*OperResp, error) {
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
	return &OperResp{NameOper: "MultiTransfer", Bresp: resp}, err
}

//Checking the user's key for the possibility of operations in GOLOS.
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

func (api *Client) LimitOrderCancel(owner string, orderid uint32) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.LimitOrderCancelOperation{
		Owner:   owner,
		OrderID: orderid,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(owner, trx)
	return &OperResp{NameOper: "LimitOrderCancel", Bresp: resp}, err
}

func (api *Client) LimitOrderCreate(owner, sell, buy string, orderid uint32) (*OperResp, error) {
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
	return &OperResp{NameOper: "LimitOrderCreate", Bresp: resp}, err
}

func (api *Client) Convert(owner, amount string, requestid uint32) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.ConvertOperation{
		Owner:     owner,
		RequestID: requestid,
		Amount:    amount,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(owner, trx)
	return &OperResp{NameOper: "Convert", Bresp: resp}, err
}

func (api *Client) TransferToVesting(from, to, amount string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferToVestingOperation{
		From:   from,
		To:     to,
		Amount: amount,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from, trx)
	return &OperResp{NameOper: "TransferToVesting", Bresp: resp}, err
}

func (api *Client) WithdrawVesting(account, vshares string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.WithdrawVestingOperation{
		Account:       account,
		VestingShares: vshares,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(account, trx)
	return &OperResp{NameOper: "WithdrawVesting", Bresp: resp}, err
}

func (api *Client) ChangeRecoveryAccount(accounttorecover, newrecoveryaccount string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.ChangeRecoveryAccountOperation{
		AccountToRecover:   accounttorecover,
		NewRecoveryAccount: newrecoveryaccount,
		Extensions:         []interface{}{},
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(accounttorecover, trx)
	return &OperResp{NameOper: "ChangeRecoveryAccount", Bresp: resp}, err
}

func (api *Client) TransferToSavings(from, to, amount, memo string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferToSavingsOperation{
		From:   from,
		To:     to,
		Amount: amount,
		Memo:   memo,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from, trx)
	return &OperResp{NameOper: "TransferToSavings", Bresp: resp}, err
}

func (api *Client) TransferFromSavings(from, to, amount, memo string, requestid uint32) (*OperResp, error) {
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
	return &OperResp{NameOper: "TransferFromSavings", Bresp: resp}, err
}

func (api *Client) CancelTransferFromSavings(from string, requestid uint32) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.CancelTransferFromSavingsOperation{
		From:      from,
		RequestId: requestid,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(from, trx)
	return &OperResp{NameOper: "CancelTransferFromSavings", Bresp: resp}, err
}

func (api *Client) DeclineVotingRights(account string, decline bool) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.DeclineVotingRightsOperation{
		Account: account,
		Decline: decline,
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(account, trx)
	return &OperResp{NameOper: "DeclineVotingRights", Bresp: resp}, err
}

func (api *Client) FeedPublish(publisher, base, quote string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.FeedPublishOperation{
		Publisher: publisher,
		ExchangeRate: &types.ExchRate{
			Base:  base,
			Quote: quote,
		},
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(publisher, trx)
	return &OperResp{NameOper: "FeedPublish", Bresp: resp}, err
}

func (api *Client) WitnessUpdate(owner, url, blocksigningkey, accountcreationfee string, maxblocksize uint32, sbdinterestrate uint16) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.WitnessUpdateOperation{
		Owner:           owner,
		Url:             url,
		BlockSigningKey: blocksigningkey,
		Props: &types.ChainProperties{
			AccountCreationFee: accountcreationfee,
			MaximumBlockSize:   maxblocksize,
			SBDInterestRate:    sbdinterestrate,
		},
		Fee: "0.000 GOLOS",
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(owner, trx)
	return &OperResp{NameOper: "WitnessUpdate", Bresp: resp}, err
}

func (api *Client) AccountCreate(creator, newAccountName, password string, fee string) (*OperResp, error) {
	type Keys struct {
		Private string
		Public  string
	}

	var trx []types.Operation
	var listKeys = make(map[string]Keys)

	empty := map[string]int64{}
	roles := [4]string{"owner", "active", "posting", "memo"}

	for _, val := range roles {
		priv := GetPrivateKey(newAccountName, val, password)
		pub := GetPublicKey("GLS", priv)
		listKeys[val] = Keys{Private: priv, Public: pub}
	}

	owner := types.Authority{
		WeightThreshold: 1,
		AccountAuths:    empty,
		KeyAuths:        map[string]int64{listKeys["owner"].Public: 1},
	}

	active := types.Authority{
		WeightThreshold: 1,
		AccountAuths:    empty,
		KeyAuths:        map[string]int64{listKeys["active"].Public: 1},
	}

	posting := types.Authority{
		WeightThreshold: 1,
		AccountAuths:    empty,
		KeyAuths:        map[string]int64{listKeys["posting"].Public: 1},
	}

	tx := &types.AccountCreateOperation{
		Fee:            fee,
		Creator:        creator,
		NewAccountName: newAccountName,
		Owner:          &owner,
		Active:         &active,
		Posting:        &posting,
		MemoKey:        listKeys["memo"].Public,
		JsonMetadata:   "",
	}

	trx = append(trx, tx)
	resp, err := api.SendTrx(creator, trx)
	return &OperResp{NameOper: "AccountCreate", Bresp: resp}, err
}
