package types

//AuctionWindowRewardOperation represents auction_window_reward operation data.
type AuctionWindowRewardOperation struct {
	Reward          *Asset `json:"reward"`
	CommentAuthor   string `json:"comment_author"`
	CommentPermlink string `json:"comment_permlink"`
}

//Type function that defines the type of operation AuctionWindowRewardOperation.
func (op *AuctionWindowRewardOperation) Type() OpType {
	return TypeAuctionWindowReward
}

//Data returns the operation data AuctionWindowRewardOperation.
func (op *AuctionWindowRewardOperation) Data() interface{} {
	return op
}
