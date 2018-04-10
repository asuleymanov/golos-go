package types

import (
	"encoding/json"
	"strconv"
)

type ContentMetadata struct {
	Tags   []string `json:"tags"`
	Image  []string `json:"image"`
	Lib    string   `json:"lib"`
	App    string   `json:"app"`
	Format string   `json:"format"`
}

type rawContentMetadata struct {
	Tags   []string `json:"tags"`
	Image  []string `json:"image"`
	Lib    string   `json:"lib"`
	App    string   `json:"app"`
	Format string   `json:"format"`
}

func (metadata *ContentMetadata) UnmarshalJSON(p []byte) error {
	var raw rawContentMetadata

	str, _ := strconv.Unquote(string(p))

	if err := json.Unmarshal([]byte(str), &raw); err != nil {
		return err
	}

	metadata.Tags = raw.Tags
	metadata.Image = raw.Image
	metadata.Lib = raw.Lib
	metadata.App = raw.App
	metadata.Format = raw.Format
	return nil
}
