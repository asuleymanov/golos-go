package private_message

import (
	"github.com/asuleymanov/golos-go/types"
)

type Message struct {
	ID               *types.Int  `json:"id"`
	From             string      `json:"from"`
	To               string      `json:"to"`
	FromMemoKey      string      `json:"from_memo_key"`
	ToMemoKey        string      `json:"to_memo_key"`
	SentTime         uint64      `json:"sent_time"`
	ReceiveTime      *types.Time `json:"receive_time"`
	Checksum         uint32      `json:"checksum"`
	EncryptedMessage string      `json:"encrypted_message"`
}
