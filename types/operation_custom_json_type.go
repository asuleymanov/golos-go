package types

// struct PrivateMessageOperation{}
type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

// struct PrivateMessageOperation{}
type ReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

// struct PrivateMessageOperation{}
type PrivateMessageOperation struct {
	From             string `json:"from"`
	To               string `json:"to"`
	FromMemoKey      string `json:"from_memo_key"`
	ToMemoKey        string `json:"to_memo_key"`
	SentTime         uint64 `json:"sent_time"`
	Checksum         uint32 `json:"checksum"`
	EncryptedMessage string `json:"encrypted_message"`
}
