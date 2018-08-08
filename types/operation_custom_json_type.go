package types

import (
	"encoding/json"
	"errors"
	"reflect"
)

//FollowOperation the structure for the operation CustomJSONOperation.
type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

//ReblogOperation the structure for the operation CustomJSONOperation.
type ReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//LoginOperation the structure for the operation CustomJSONOperation.
type LoginOperation struct {
	Account string `json:"account"`
}

//PrivateMessageOperation the structure for the operation CustomJSONOperation.
type PrivateMessageOperation struct {
	From             string `json:"from"`
	To               string `json:"to"`
	FromMemoKey      string `json:"from_memo_key"`
	ToMemoKey        string `json:"to_memo_key"`
	SentTime         uint64 `json:"sent_time"`
	Checksum         uint32 `json:"checksum"`
	EncryptedMessage string `json:"encrypted_message"`
}

//MarshalCustomJSON generate a row from the structure fields.
func MarshalCustomJSON(v interface{}) (string, error) {
	var tmp []interface{}

	typeInterface := reflect.TypeOf(v).Name()
	switch typeInterface {
	case "FollowOperation":
		tmp = append(tmp, TypeFollow)
	case "ReblogOperation":
		tmp = append(tmp, TypeReblog)
	case "LoginOperation":
		tmp = append(tmp, TypeLogin)
	case "PrivateMessageOperation":
		tmp = append(tmp, TypePrivateMessage)
	default:
		return "", errors.New("Unknown type")
	}

	tmp = append(tmp, v)

	b, err := json.Marshal(tmp)
	if err != nil {
		return "", err
	}

	return string(b), nil //strings.Replace(string(b), "\"", "\\\"", -1), nil
}
