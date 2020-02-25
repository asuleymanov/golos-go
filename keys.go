package golos

import (
	"errors"

	"github.com/asuleymanov/golos-go/encoding/wif"
	"github.com/asuleymanov/golos-go/operations"
)

var (
	//OpTypeKey include a description of the operation and the key needed to sign it
	OpTypeKey = make(map[operations.OpType][]string)
)

//Keys is used as a keystroke for a specific user.
//Only a few keys can be set.
type Keys struct {
	PKey []string
	AKey []string
	OKey []string
	MKey []string
}

func init() {
	OpTypeKey["vote"] = []string{"posting"}
	OpTypeKey["comment"] = []string{"posting"}
	OpTypeKey["transfer"] = []string{"active"}
	OpTypeKey["transfer_to_vesting"] = []string{"active"}
	OpTypeKey["withdraw_vesting"] = []string{"active"}
	OpTypeKey["limit_order_create"] = []string{"active"}
	OpTypeKey["limit_order_cancel"] = []string{"active"}
	OpTypeKey["feed_publish"] = []string{"active"}
	OpTypeKey["convert"] = []string{"active"}
	OpTypeKey["account_create"] = []string{"active"}
	OpTypeKey["account_update"] = []string{"active"}
	OpTypeKey["witness_update"] = []string{"active"}
	OpTypeKey["account_witness_vote"] = []string{"posting"}
	OpTypeKey["account_witness_proxy"] = []string{"posting"}
	OpTypeKey["delete_comment"] = []string{"posting"}
	OpTypeKey["custom_json"] = []string{"posting"}
	OpTypeKey["comment_options"] = []string{"posting"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"active"}
	OpTypeKey["limit_order_create2"] = []string{"active"}
	OpTypeKey["challenge_authority"] = []string{"owner"}
	OpTypeKey["prove_authority"] = []string{"active"}
	OpTypeKey["request_account_recovery"] = []string{"active"}
	OpTypeKey["recover_account"] = []string{"active"}
	OpTypeKey["change_recovery_account"] = []string{"posting"}
	OpTypeKey["escrow_transfer"] = []string{"active"}
	OpTypeKey["escrow_dispute"] = []string{"active"}
	OpTypeKey["escrow_release"] = []string{"active"}
	OpTypeKey["escrow_approve"] = []string{"active"}
	OpTypeKey["transfer_to_savings"] = []string{"active"}
	OpTypeKey["transfer_from_savings"] = []string{"active"}
	OpTypeKey["cancel_transfer_from_savings"] = []string{"active"}
	OpTypeKey["decline_voting_rights"] = []string{"owner"}
	OpTypeKey["reset_account"] = []string{"active"}
	OpTypeKey["set_reset_account"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares"] = []string{"active"}
	OpTypeKey["account_create_with_delegation"] = []string{"active"}
	OpTypeKey["account_metadata"] = []string{"posting"}
	OpTypeKey["proposal_create"] = []string{"active"}
	OpTypeKey["proposal_update"] = []string{"active"}
	OpTypeKey["proposal_delete"] = []string{"active"}
	OpTypeKey["chain_properties_update"] = []string{"active"}
	OpTypeKey["break_free_referral"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares_with_interest"] = []string{"active"}
	OpTypeKey["reject_vesting_shares_delegation"] = []string{"active"}
	OpTypeKey["transit_to_cyberway"] = []string{"active"}
	OpTypeKey["worker_request"] = []string{"active"}
	OpTypeKey["worker_request_delete"] = []string{"active"}
	OpTypeKey["worker_request_vote"] = []string{"active"}
}

//SigningKeys returns the key from the CurrentKeys
func (client *Client) SigningKeys(trx operations.Operation) ([][]byte, error) {
	var keys [][]byte

	if client.CurrentKeys == nil {
		return nil, errors.New("Client Keys not initialized. Use SetKeys method")
	}

	opKeys := OpTypeKey[trx.Type()]
	for _, val := range opKeys {
		switch val {
		case "posting":
			for _, keyStr := range client.CurrentKeys.PKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Posting Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "active":
			for _, keyStr := range client.CurrentKeys.AKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Active Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "owner":
			for _, keyStr := range client.CurrentKeys.OKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Owner Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "memo":
			for _, keyStr := range client.CurrentKeys.MKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Memo Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		}
	}

	return keys, nil
}
