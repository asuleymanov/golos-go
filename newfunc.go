package client

import (
	// Stdlib
	"strconv"
	"strings"
	"time"

	// RPC
	"github.com/asuleymanov/golos-go/translit"
	"github.com/asuleymanov/golos-go/types"
)

//TYPE

//PCOptionsNew for the Comment and Post functions.
//Sets the water to receive payment for a comment or post.
type PCOptionsNew struct {
	Permlink      string
	Ptag          string
	PostImage     string
	Tags          []string
	CommentOption *CommentOptions
}

type CommentOptions struct {
	Percent              uint16
	AllowVotes           *bool
	AllowCurationRewards *bool
	BeneficiaryList      []types.Beneficiary
}

//FUNCTION

//CommentNew for publication
func (client *Client) CommentNew(username, authorname, ppermlink, body string, opt *PCOptionsNew) (*OperResp, error) {
	var trx []types.Operation

	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + authorname + "-" + ppermlink + "-" + times
	permlink = strings.Replace(permlink, ".", "-", -1)
	if opt != nil && opt.Permlink != "" {
		permlink = opt.Permlink
	}

	jsonMeta := &types.ContentMetadata{"lib": "golos-go"}

	tx := &types.CommentOperation{
		ParentAuthor:   authorname,
		ParentPermlink: ppermlink,
		Author:         username,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JSONMetadata:   jsonMeta,
	}
	trx = append(trx, tx)

	if opt.CommentOption != nil {
		trx = append(trx, GetCommentOptionsOperationNew(username, permlink, *opt.CommentOption))
	}

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "Comment", PermLink: permlink, Bresp: resp}, err
}

//Post creating a publication
func (client *Client) PostNew(authorname, title, body string, opt *PCOptionsNew) (*OperResp, error) {
	permlink := ""
	ptag := ""
	tag := []string{}

	if opt != nil && opt.Permlink != "" {
		permlink = translit.EncodeTitle(opt.Permlink)
	} else {
		permlink = translit.EncodeTitle(title)
	}

	if opt != nil && opt.Tags != nil {
		tag = translit.EncodeTags(opt.Tags)
	}

	if opt != nil && opt.Ptag != "" {
		ptag = translit.EncodeTag(opt.Ptag)
	} else {
		ptag = translit.EncodeTag(opt.Tags[0])
	}

	jsonMeta := &types.ContentMetadata{
		"tags":  tag,
		"image": []string{opt.PostImage},
		"lib":   "golos-go",
	}

	var trx []types.Operation
	txp := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         authorname,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JSONMetadata:   jsonMeta,
	}
	trx = append(trx, txp)

	if opt.CommentOption != nil {
		trx = append(trx, GetCommentOptionsOperationNew(authorname, permlink, *opt.CommentOption))
	}

	resp, err := client.SendTrx(authorname, trx)
	return &OperResp{NameOper: "Post", PermLink: permlink, Bresp: resp}, err
}

//HELPER

//GetCommentOptionsOperation generates CommentOptionsOperation depending on the incoming data
func GetCommentOptionsOperationNew(username, permlink string, options CommentOptions) *types.CommentOptionsOperation {
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

	if options.BeneficiaryList != nil && len(options.BeneficiaryList) > 0 {
		var benList []types.Beneficiary
		var benef types.CommentPayoutBeneficiaries
		for _, val := range options.BeneficiaryList {
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
