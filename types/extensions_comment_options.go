package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//Beneficiary is an additional structure used by other structures.
type Beneficiary struct {
	Account string `json:"account"`
	Weight  uint16 `json:"weight"`
}

//CommentPayoutBeneficiaries is an additional structure used by other structures.
type CommentPayoutBeneficiaries struct {
	Beneficiaries []Beneficiary `json:"beneficiaries"`
}

//MarshalTransaction is a function of converting type CommentPayoutBeneficiaries to bytes.
func (cp *CommentPayoutBeneficiaries) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(byte(0))
	enc.EncodeUVarint(uint64(len(cp.Beneficiaries)))
	for _, val := range cp.Beneficiaries {
		enc.Encode(val.Account)
		enc.Encode(val.Weight)
	}
	return enc.Err()
}

//CommentAuctionWindowRewardDestination is an additional structure used by other structures.
/*
Destination:
to_reward_fund - 0
to_curators - 1
to_author - 2
*/
type CommentAuctionWindowRewardDestination struct {
	Destination int `json:"destination"`
}

//MarshalTransaction is a function of converting type CommentAuctionWindowRewardDestination to bytes.
func (cp *CommentAuctionWindowRewardDestination) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.Destination)
	return enc.Err()
}

//CommentCurationRewardsPercent is an additional structure used by other structures.
type CommentCurationRewardsPercent struct {
	Percent uint16 `json:"percent"`
}

//MarshalTransaction is a function of converting type CommentCurationRewardsPercent to bytes.
func (cp *CommentCurationRewardsPercent) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.Percent)
	return enc.Err()
}
