package types

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

// Add-on encode
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

func (exch *ExchRate) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(exch.Base)
	enc.Encode(exch.Quote)
	return enc.Err()
}

func (cp *ChainProperties) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	return enc.Err()
}

// encode VoteOperation{}
func (op *VoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Weight)
	return enc.Err()
}

// encode CommentOperation{}
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

// encode TransferOperation{}
func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// encode TransferToVestingOperation{}
func (op *TransferToVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToVesting.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	return enc.Err()
}

// encode WithdrawVestingOperation{}
func (op *WithdrawVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWithdrawVesting.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.VestingShares)
	return enc.Err()
}

// encode LimitOrderCreateOperation{}
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

// encode LimitOrderCancelOperation{}
func (op *LimitOrderCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCancel.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	return enc.Err()
}

// encode FeedPublishOperation{}
func (op *FeedPublishOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeFeedPublish.Code()))
	enc.Encode(op.Publisher)
	enc.Encode(op.ExchangeRate)
	return enc.Err()
}

// encode ConvertOperation{}
func (op *ConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeConvert.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.RequestID)
	enc.Encode(op.Amount)
	return enc.Err()
}

// encode AccountCreateOperation{}
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

// encode AccountUpdateOperation{}
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

// encode WitnessUpdateOperation{}
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

// encode AccountWitnessVoteOperation{}
func (op *AccountWitnessVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessVote.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Witness)
	enc.EncodeBool(op.Approve)
	return enc.Err()
}

// encode AccountWitnessProxyOperation{}
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
// encode DeleteCommentOperation{}
func (op *DeleteCommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeleteComment.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	return enc.Err()
}

// encode CustomJSONOperation{} in to file operation_custom_json.go
// encode CommentOptionsOperation{}
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

// encode SetWithdrawVestingRouteOperation{}
func (op *SetWithdrawVestingRouteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetWithdrawVestingRoute.Code()))
	enc.Encode(op.FromAccount)
	enc.Encode(op.ToAccount)
	enc.Encode(op.Percent)
	enc.EncodeBool(op.AutoVest)
	return enc.Err()
}

// encode LimitOrderCreate2Operation{}
func (op *LimitOrderCreate2Operation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate2.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.Orderid)
	enc.Encode(op.AmountToSell)
	enc.Encode(op.ExchangeRate)
	enc.EncodeBool(op.FillOrKill)
	enc.Encode(op.Expiration)
	return enc.Err()
}

// encode ChallengeAuthorityOperation{}
func (op *ChallengeAuthorityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChallengeAuthority.Code()))
	enc.Encode(op.Challenger)
	enc.Encode(op.Challenged)
	enc.EncodeBool(op.RequireOwner)
	return enc.Err()
}

// encode ProveAuthorityOperation{}
func (op *ProveAuthorityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProveAuthority.Code()))
	enc.Encode(op.Challenged)
	enc.EncodeBool(op.RequireOwner)
	return enc.Err()
}

// encode RequestAccountRecoveryOperation{}
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

// encode RecoverAccountOperation{}
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

// encode ChangeRecoveryAccountOperation{}
func (op *ChangeRecoveryAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChangeRecoveryAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewRecoveryAccount)
	enc.Encode(byte(0))
	return enc.Err()
}

// encode EscrowTransferOperation{}
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

// encode EscrowDisputeOperation{}
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

// encode EscrowReleaseOperation{}
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
// encode EscrowApproveOperation{}
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

// encode TransferToSavingsOperation{}
func (op *TransferToSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// encode TransferFromSavingsOperation{}
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

// encode CancelTransferFromSavingsOperation{}
func (op *CancelTransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCancelTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestID)
	return enc.Err()
}

// encode CustomBinaryOperation{}
// encode DeclineVotingRightsOperation{}
func (op *DeclineVotingRightsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeclineVotingRights.Code()))
	enc.Encode(op.Account)
	enc.EncodeBool(op.Decline)
	return enc.Err()
}

// encode ResetAccountOperation{}
func (op *ResetAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeResetAccount.Code()))
	enc.Encode(op.ResetAccount)
	enc.Encode(op.AccountToReset)
	enc.Encode(op.NewOwnerAuthority)
	return enc.Err()
}

// encode SetResetAccountOperation{}
func (op *SetResetAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetResetAccount.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.CurrentResetAccount)
	enc.Encode(op.ResetAccount)
	return enc.Err()
}

//Virtual Operation
// encode ClaimRewardBalanceOperation{}
// encode DelegateVestingSharesOperation{}
// encode AccountCreateWithDelegationOperation{}
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
