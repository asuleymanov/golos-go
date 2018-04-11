package types

import (
	"encoding/json"
	"strconv"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

type ContentMetadata struct {
	Tags   []string `json:"tags,omitempty"`
	Image  []string `json:"image,omitempty"`
	Lib    string   `json:"lib,omitempty"`
	App    string   `json:"app,omitempty"`
	Format string   `json:"format,omitempty"`
}

type rawContentMetadata struct {
	Tags   []string `json:"tags,omitempty"`
	Image  []string `json:"image,omitempty"`
	Lib    string   `json:"lib,omitempty"`
	App    string   `json:"app,omitempty"`
	Format string   `json:"format,omitempty"`
}

func (op *ContentMetadata) UnmarshalJSON(p []byte) error {
	var raw rawContentMetadata

	str, _ := strconv.Unquote(string(p))

	if err := json.Unmarshal([]byte(str), &raw); err != nil {
		return err
	}

	op.Tags = raw.Tags
	op.Image = raw.Image
	op.Lib = raw.Lib
	op.App = raw.App
	op.Format = raw.Format
	return nil
}

func (op *ContentMetadata) MarshalJSON() ([]byte, error) {
	ans, err := json.Marshal(&rawContentMetadata{
		Tags:   op.Tags,
		Image:  op.Image,
		Lib:    op.Lib,
		App:    op.App,
		Format: op.Format,
	})
	if err != nil {
		return []byte{}, err
	}
	return []byte(strconv.Quote(string(ans))), nil
}

func (op *ContentMetadata) MarshalTransaction(encoder *transaction.Encoder) error {
	ans, err := json.Marshal(op)
	if err != nil {
		return err
	}

	str, err := strconv.Unquote(string(ans))
	if err != nil {
		return err
	}

	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeString(str)
	return enc.Err()
}
