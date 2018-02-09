package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

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
	enc.Encode(op.JsonMetadata)
	return enc.Err()
}

// encode TransferOperation{}
func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// encode TransferToVestingOperation{}
func (op *TransferToVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToVesting.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	return enc.Err()
}

// encode WithdrawVestingOperation{}
func (op *WithdrawVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWithdrawVesting.Code()))
	enc.Encode(op.Account)
	enc.EncodeMoney(op.VestingShares)
	return enc.Err()
}

// encode LimitOrderCreateOperation{}
func (op *LimitOrderCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	enc.EncodeMoney(op.AmountToSell)
	enc.EncodeMoney(op.MinToReceive)
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
	enc.EncodeMoney(op.ExchangeRate.Base)
	enc.EncodeMoney(op.ExchangeRate.Quote)
	return enc.Err()
}

// encode ConvertOperation{}
func (op *ConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeConvert.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.RequestID)
	enc.EncodeMoney(op.Amount)
	return enc.Err()
}

// encode AccountCreateOperation{}
// encode AccountUpdateOperation{}
// encode WitnessUpdateOperation{}
func (op *WitnessUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWitnessUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.Url)
	enc.EncodePubKey(op.BlockSigningKey)
	enc.EncodeMoney(op.Props.AccountCreationFee)
	enc.Encode(op.Props.MaximumBlockSize)
	enc.Encode(op.Props.SBDInterestRate)
	enc.EncodeMoney(op.Fee)
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
	enc.EncodeMoney(op.MaxAcceptedPayout)
	enc.Encode(op.PercentSteemDollars)
	enc.EncodeBool(op.AllowVotes)
	enc.EncodeBool(op.AllowCurationRewards)
	enc.Encode(byte(0))
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
// encode ChallengeAuthorityOperation{}
// encode ProveAuthorityOperation{}
// encode RequestAccountRecoveryOperation{}
// encode RecoverAccountOperation{}
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
// encode EscrowDisputeOperation{}
// encode EscrowReleaseOperation{}
// encode POW2Operation{}
// encode EscrowApproveOperation{}
// encode TransferToSavingsOperation{}
func (op *TransferToSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// encode TransferFromSavingsOperation{}
func (op *TransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestId)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// encode CancelTransferFromSavingsOperation{}
func (op *CancelTransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCancelTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestId)
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
// encode SetResetAccountOperation{}
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
