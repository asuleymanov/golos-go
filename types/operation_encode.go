package types

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

// Add-on encode

//MarshalTransaction is a function of converting type Authority to bytes.
func (auth *Authority) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeNumber(uint32(auth.WeightThreshold))
	// encode AccountAuths as map[string]uint16
	enc.EncodeUVarint(uint64(len(auth.AccountAuths)))
	for k, v := range auth.AccountAuths {
		enc.EncodeString(k)
		enc.EncodeNumber(uint16(v))
	}
	// encode KeyAuths as map[PubKey]uint16
	enc.EncodeUVarint(uint64(len(auth.KeyAuths)))
	for k, v := range auth.KeyAuths {
		enc.EncodePubKey(k)
		enc.EncodeNumber(uint16(v))
	}
	return enc.Err()
}

//MarshalTransaction is a function of converting type ExchRate to bytes.
func (exch *ExchRate) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(exch.Base)
	enc.Encode(exch.Quote)
	return enc.Err()
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

//MarshalTransaction is a function of converting type VoteOperation to bytes.
func (op *VoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Weight)
	return enc.Err()
}

//MarshalTransaction is a function of converting type CommentOperation to bytes.
func (op *CommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeComment.Code()))
	if !op.IsStoryOperation() {
		enc.Encode(op.ParentAuthor)
	} else {
		enc.Encode(byte(0))
	}
	enc.Encode(op.ParentPermlink)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Title)
	enc.Encode(op.Body)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}

//MarshalTransaction is a function of converting type TransferOperation to bytes.
func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

//MarshalTransaction is a function of converting type TransferToVestingOperation to bytes.
func (op *TransferToVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToVesting.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	return enc.Err()
}

//MarshalTransaction is a function of converting type WithdrawVestingOperation to bytes.
func (op *WithdrawVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWithdrawVesting.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.VestingShares)
	return enc.Err()
}

//MarshalTransaction is a function of converting type LimitOrderCreateOperation to bytes.
func (op *LimitOrderCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	enc.Encode(op.AmountToSell)
	enc.Encode(op.MinToReceive)
	enc.EncodeBool(op.FillOrKill)
	enc.Encode(op.Expiration)
	return enc.Err()
}

//MarshalTransaction is a function of converting type LimitOrderCancelOperation to bytes.
func (op *LimitOrderCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCancel.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	return enc.Err()
}

//MarshalTransaction is a function of converting type FeedPublishOperation to bytes.
func (op *FeedPublishOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeFeedPublish.Code()))
	enc.Encode(op.Publisher)
	enc.Encode(op.ExchangeRate)
	return enc.Err()
}

//MarshalTransaction is a function of converting type ConvertOperation to bytes.
func (op *ConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeConvert.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.RequestID)
	enc.Encode(op.Amount)
	return enc.Err()
}

//MarshalTransaction is a function of converting type AccountCreateOperation to bytes.
func (op *AccountCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountCreate.Code()))
	enc.Encode(op.Fee)
	enc.EncodeString(op.Creator)
	enc.EncodeString(op.NewAccountName)
	enc.Encode(op.Owner)
	enc.Encode(op.Active)
	enc.Encode(op.Posting)
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}

//MarshalTransaction is a function of converting type AccountUpdateOperation to bytes.
func (op *AccountUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountUpdate.Code()))
	enc.EncodeString(op.Account)
	if op.Owner != nil {
		enc.Encode(op.Owner)
	}
	if op.Active != nil {
		enc.Encode(op.Active)
	}
	if op.Posting != nil {
		enc.Encode(op.Posting)
	}
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}

//MarshalTransaction is a function of converting type WitnessUpdateOperation to bytes.
func (op *WitnessUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWitnessUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.URL)
	enc.EncodePubKey(op.BlockSigningKey)
	enc.Encode(op.Props)
	enc.Encode(op.Fee)
	return enc.Err()
}

//MarshalTransaction is a function of converting type AccountWitnessVoteOperation to bytes.
func (op *AccountWitnessVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessVote.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Witness)
	enc.EncodeBool(op.Approve)
	return enc.Err()
}

//MarshalTransaction is a function of converting type AccountWitnessProxyOperation to bytes.
func (op *AccountWitnessProxyOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessProxy.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Proxy)
	return enc.Err()
}

// encode POWOperation{}
// encode CustomOperation{}
// encode ReportOverProductionOperation{}

//MarshalTransaction is a function of converting type DeleteCommentOperation to bytes.
func (op *DeleteCommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeleteComment.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	return enc.Err()
}

// encode CustomJSONOperation{} in to file operation_custom_json.go

//MarshalTransaction is a function of converting type CommentOptionsOperation to bytes.
func (op *CommentOptionsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCommentOptions.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.MaxAcceptedPayout)
	enc.Encode(op.PercentSteemDollars)
	enc.EncodeBool(op.AllowVotes)
	enc.EncodeBool(op.AllowCurationRewards)
	if len(op.Extensions) > 0 {
		//Parse Beneficiaries
		z, _ := json.Marshal(op.Extensions[0])
		var l []interface{}
		_ = json.Unmarshal(z, &l)
		z1, _ := json.Marshal(l[1])
		var d CommentPayoutBeneficiaries
		_ = json.Unmarshal(z1, &d)

		enc.Encode(byte(1))
		enc.Encode(byte(0))
		enc.EncodeUVarint(uint64(len(d.Beneficiaries)))
		for _, val := range d.Beneficiaries {
			enc.Encode(val.Account)
			enc.Encode(val.Weight)
		}
	} else {
		enc.Encode(byte(0))
	}
	return enc.Err()
}

//MarshalTransaction is a function of converting type SetWithdrawVestingRouteOperation to bytes.
func (op *SetWithdrawVestingRouteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetWithdrawVestingRoute.Code()))
	enc.Encode(op.FromAccount)
	enc.Encode(op.ToAccount)
	enc.Encode(op.Percent)
	enc.EncodeBool(op.AutoVest)
	return enc.Err()
}

//MarshalTransaction is a function of converting type LimitOrderCreate2Operation to bytes.
func (op *LimitOrderCreate2Operation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate2.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	enc.Encode(op.AmountToSell)
	enc.Encode(op.ExchangeRate)
	enc.EncodeBool(op.FillOrKill)
	enc.Encode(op.Expiration)
	return enc.Err()
}

//MarshalTransaction is a function of converting type ChallengeAuthorityOperation to bytes.
func (op *ChallengeAuthorityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChallengeAuthority.Code()))
	enc.Encode(op.Challenger)
	enc.Encode(op.Challenged)
	enc.EncodeBool(op.RequireOwner)
	return enc.Err()
}

//MarshalTransaction is a function of converting type ProveAuthorityOperation to bytes.
func (op *ProveAuthorityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProveAuthority.Code()))
	enc.Encode(op.Challenged)
	enc.EncodeBool(op.RequireOwner)
	return enc.Err()
}

//MarshalTransaction is a function of converting type RequestAccountRecoveryOperation to bytes.
func (op *RequestAccountRecoveryOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeRequestAccountRecovery.Code()))
	enc.Encode(op.RecoveryAccount)
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewOwnerAuthority)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}

//MarshalTransaction is a function of converting type RecoverAccountOperation to bytes.
func (op *RecoverAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeRecoverAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewOwnerAuthority)
	enc.Encode(op.RecentOwnerAuthority)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}

//MarshalTransaction is a function of converting type ChangeRecoveryAccountOperation to bytes.
func (op *ChangeRecoveryAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChangeRecoveryAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewRecoveryAccount)
	enc.Encode(byte(0))
	return enc.Err()
}

//MarshalTransaction is a function of converting type EscrowTransferOperation to bytes.
func (op *EscrowTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.SbdAmount)
	enc.Encode(op.SteemAmount)
	enc.Encode(op.EscrowID)
	enc.Encode(op.Agent)
	enc.Encode(op.Fee)
	enc.Encode(op.JSONMeta)
	enc.Encode(op.RatificationDeadline)
	enc.Encode(op.EscrowExpiration)
	return enc.Err()
}

//MarshalTransaction is a function of converting type EscrowDisputeOperation to bytes.
func (op *EscrowDisputeOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowDispute.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Agent)
	enc.Encode(op.Who)
	enc.Encode(op.EscrowID)
	return enc.Err()
}

//MarshalTransaction is a function of converting type EscrowReleaseOperation to bytes.
func (op *EscrowReleaseOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowRelease.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Agent)
	enc.Encode(op.Who)
	enc.Encode(op.Receiver)
	enc.Encode(op.EscrowID)
	enc.Encode(op.SbdAmount)
	enc.Encode(op.SteemAmount)
	return enc.Err()
}

// encode POW2Operation{}

//MarshalTransaction is a function of converting type EscrowApproveOperation to bytes.
func (op *EscrowApproveOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowApprove.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Agent)
	enc.Encode(op.Who)
	enc.Encode(op.EscrowID)
	enc.EncodeBool(op.Approve)
	return enc.Err()
}

//MarshalTransaction is a function of converting type TransferToSavingsOperation to bytes.
func (op *TransferToSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

//MarshalTransaction is a function of converting type TransferFromSavingsOperation to bytes.
func (op *TransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestID)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

//MarshalTransaction is a function of converting type CancelTransferFromSavingsOperation to bytes.
func (op *CancelTransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCancelTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestID)
	return enc.Err()
}

// encode CustomBinaryOperation{}

//MarshalTransaction is a function of converting type DeclineVotingRightsOperation to bytes.
func (op *DeclineVotingRightsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeclineVotingRights.Code()))
	enc.Encode(op.Account)
	enc.EncodeBool(op.Decline)
	return enc.Err()
}

//MarshalTransaction is a function of converting type ResetAccountOperation to bytes.
func (op *ResetAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeResetAccount.Code()))
	enc.Encode(op.ResetAccount)
	enc.Encode(op.AccountToReset)
	enc.Encode(op.NewOwnerAuthority)
	return enc.Err()
}

//MarshalTransaction is a function of converting type SetResetAccountOperation to bytes.
func (op *SetResetAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetResetAccount.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.CurrentResetAccount)
	enc.Encode(op.ResetAccount)
	return enc.Err()
}

//MarshalTransaction is a function of converting type DelegateVestingSharesOperation to bytes.
func (op *DelegateVestingSharesOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDelegateVestingShares.Code()))
	enc.Encode(op.Delegator)
	enc.Encode(op.Delegatee)
	enc.Encode(op.VestingShares)
	return enc.Err()
}

//MarshalTransaction is a function of converting type AccountCreateWithDelegationOperation to bytes.
func (op *AccountCreateWithDelegationOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountCreateWithDelegation.Code()))
	enc.Encode(op.Fee)
	enc.Encode(op.Delegation)
	enc.Encode(op.Creator)
	enc.Encode(op.NewAccountName)
	enc.Encode(op.Owner)
	enc.Encode(op.Active)
	enc.Encode(op.Posting)
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}

//MarshalTransaction is a function of converting type AccountMetadataOperation to bytes.
func (op *AccountMetadataOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountMetadata.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}

//MarshalTransaction is a function of converting type ProposalCreateOperation to bytes.
func (op *ProposalCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProposalCreate.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Title)
	enc.Encode(op.Memo)
	enc.Encode(op.ExpirationTime)
	enc.Encode(op.ProposedOperations)
	enc.Encode(byte(1))
	enc.Encode(op.ReviewPeriodTime)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}

//MarshalTransaction is a function of converting type ProposalUpdateOperation to bytes.
func (op *ProposalUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProposalUpdate.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Title)
	enc.EncodeArrString(op.ActiveApprovalsToAdd)
	enc.EncodeArrString(op.ActiveApprovalsToRemove)
	enc.EncodeArrString(op.OwnerApprovalsToAdd)
	enc.EncodeArrString(op.OwnerApprovalsToRemove)
	enc.EncodeArrString(op.PostingApprovalsToAdd)
	enc.EncodeArrString(op.PostingApprovalsToRemove)
	enc.EncodeArrString(op.KeyApprovalsToAdd)
	enc.EncodeArrString(op.KeyApprovalsToRemove)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}

//MarshalTransaction is a function of converting type ProposalDeleteOperation to bytes.
func (op *ProposalDeleteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProposalDelete.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Title)
	enc.Encode(op.Requester)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}

//MarshalTransaction is a function of converting type ChainPropertiesUpdateOperation to bytes.
func (op *ChainPropertiesUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChainPropertiesUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(byte(1))
	z, _ := json.Marshal(op.Props[1])
	var d ChainProperties
	_ = json.Unmarshal(z, &d)
	enc.Encode(&d)
	return enc.Err()
}

//Virtual Operation
// encode FillConvertRequestOperation{}
// encode AuthorRewardOperation{}
// encode CurationRewardOperation{}
// encode CommentRewardOperation{}
// encode LiquidityRewardOperation{}
// encode InterestOperation{}
// encode FillVestingWithdrawOperation{}
// encode FillOrderOperation{}
// encode ShutdownWitnessOperation{}
// encode FillTransferFromSavingsOperation{}
// encode HardforkOperation{}
// encode CommentPayoutUpdateOperation{}
// encode ReturnVestingDelegationOperation{}
// encode CommentBenefactorRewardOperation{}
