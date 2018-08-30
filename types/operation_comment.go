package types

import (
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//CommentOperation represents comment operation data.
type CommentOperation struct {
	ParentAuthor   string           `json:"parent_author"`
	ParentPermlink string           `json:"parent_permlink"`
	Author         string           `json:"author"`
	Permlink       string           `json:"permlink"`
	Title          string           `json:"title"`
	Body           string           `json:"body"`
	JSONMetadata   *ContentMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation CommentOperation.
func (op *CommentOperation) Type() OpType {
	return TypeComment
}

//Data returns the operation data CommentOperation.
func (op *CommentOperation) Data() interface{} {
	return op
}

//IsStoryOperation function specifies the type of publication.
func (op *CommentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
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
