package types

//DelegationRewardOperation represents shutdown_witness operation data.
type DelegationRewardOperation struct {
	Owner string `json:"owner"`
}

//Type function that defines the type of operation DelegationRewardOperation.
func (op *DelegationRewardOperation) Type() OpType {
	return TypeDelegationReward
}

//Data returns the operation data DelegationRewardOperation.
func (op *DelegationRewardOperation) Data() interface{} {
	return op
}
