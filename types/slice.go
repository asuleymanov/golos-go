package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

//StringSlice type from parameter JSON
type StringSlice []string

//UnmarshalJSON unpacking the JSON parameter in the StringSlice type.
func (ss *StringSlice) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] == '[' {
		var v []string
		if err := json.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("failed to unmarshal string slice : %s", err)
		}
		*ss = v
	} else {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("failed to unmarshal string slice : %s", err)
		}
		*ss = strings.Split(v, " ")
	}
	return nil
}
