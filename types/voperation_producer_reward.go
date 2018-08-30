package types

//ProducerRewardOperationOperation represents producer_reward operation data.
type ProducerRewardOperationOperation struct {
	Producer       string `json:"producer"`
	VestingShares *Asset `json:"vesting_shares"`
}

//Type function that defines the type of operation ProducerRewardOperationOperation.
func (op *ProducerRewardOperationOperation) Type() OpType {
	return TypeProducerRewardOperation
}

//Data returns the operation data ProducerRewardOperationOperation.
func (op *ProducerRewardOperationOperation) Data() interface{} {
	return op
}
