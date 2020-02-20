package operations

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"
)

//FeedPublishOperation represents feed_publish operation data.
type FeedPublishOperation struct {
	Publisher    string          `json:"publisher"`
	ExchangeRate *types.ExchRate `json:"exchange_rate"`
}

//Type function that defines the type of operation FeedPublishOperation.
func (op *FeedPublishOperation) Type() OpType {
	return TypeFeedPublish
}

//Data returns the operation data FeedPublishOperation.
func (op *FeedPublishOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type FeedPublishOperation to bytes.
func (op *FeedPublishOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeFeedPublish.Code()))
	enc.Encode(op.Publisher)
	enc.Encode(op.ExchangeRate)
	return enc.Err()
}
