package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//ChainPropertiesOLD is an additional structure used by other structures.
type ChainPropertiesOLD struct {
	AccountCreationFee *Asset `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

//ChainProperties is an additional structure used by other structures.
type ChainProperties struct {
	AccountCreationFee          *Asset `json:"account_creation_fee"`
	MaximumBlockSize            uint32 `json:"maximum_block_size"`
	SBDInterestRate             uint16 `json:"sbd_interest_rate"`
	CreateAccountMinGolosFee    *Asset `json:"create_account_min_golos_fee"`
	CreateAccountMinDelegation  *Asset `json:"create_account_min_delegation"`
	CreateAccountDelegationTime uint32 `json:"create_account_delegation_time"`
	MinDelegation               *Asset `json:"min_delegation"`
}

//MarshalTransaction is a function of converting type ChainPropertiesOLD to bytes.
func (cp *ChainPropertiesOLD) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	return enc.Err()
}

//MarshalTransaction is a function of converting type ChainProperties to bytes.
func (cp *ChainProperties) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	enc.Encode(cp.CreateAccountMinGolosFee)
	enc.Encode(cp.CreateAccountMinDelegation)
	enc.Encode(cp.CreateAccountDelegationTime)
	enc.Encode(cp.MinDelegation)
	return enc.Err()
}