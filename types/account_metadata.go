package types

import (
	"encoding/json"
	"strconv"
)

type AccountMetadata struct {
	Profile ProfileJSON `json:"profile"`
}

type rawAccountMetadata struct {
	Profile ProfileJSON `json:"profile"`
}

type ProfileJSON struct {
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	CoverImage   string `json:"cover_image"`
	Gender       string `json:"gender"`
	About        string `json:"about"`
	Location     string `json:"location"`
	Website      string `json:"website"`
}

func (metadata *AccountMetadata) UnmarshalJSON(p []byte) error {
	var raw rawAccountMetadata

	str, _ := strconv.Unquote(string(p))

	if err := json.Unmarshal([]byte(str), &raw); err != nil {
		return err
	}

	metadata.Profile = ProfileJSON{
		Name:         raw.Profile.Name,
		ProfileImage: raw.Profile.ProfileImage,
		CoverImage:   raw.Profile.CoverImage,
		Gender:       raw.Profile.Gender,
		About:        raw.Profile.About,
		Location:     raw.Profile.Location,
		Website:      raw.Profile.Website,
	}
	return nil
}
