package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/asuleymanov/golos-go/encoding/transaction"
)

//CustomJSONOperation represents custom_json operation data.
type CustomJSONOperation struct {
	RequiredAuths        []string `json:"required_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	ID                   string   `json:"id"`
	JSON                 string   `json:"json"`
}

//Type function that defines the type of operation.
func (op *CustomJSONOperation) Type() OpType {
	return TypeCustomJSON
}

//Data returns the operation data.
func (op *CustomJSONOperation) Data() interface{} {
	return op
}

//UnmarshalData unpacking the JSON parameter in the CustomJSONOperation type.
func (op *CustomJSONOperation) UnmarshalData() (interface{}, error) {
	// Get the corresponding data object template.
	template, ok := customJSONDataObjects[op.ID]
	if !ok {
		// In case there is no corresponding template, return nil.
		return nil, nil
	}

	// Clone the template.
	opData := reflect.New(reflect.Indirect(reflect.ValueOf(template)).Type()).Interface()

	// Prepare the whole operation tuple.
	var bodyReader io.Reader
	if op.JSON[0] == '[' {
		rawTuple := make([]json.RawMessage, 2)
		if err := json.NewDecoder(strings.NewReader(op.JSON)).Decode(&rawTuple); err != nil {
			return nil, fmt.Errorf("failed to unmarshal CustomJSONOperation.JSON: \n%v \n Error : %w", op.JSON, err)
		}
		if len(rawTuple) < 2 || rawTuple[1] == nil {
			return nil, fmt.Errorf("invalid CustomJSONOperation.JSON: \n%v", op.JSON)
		}
		bodyReader = bytes.NewReader([]byte(rawTuple[1]))
	} else {
		bodyReader = strings.NewReader(op.JSON)
	}

	// Unmarshal into the new object instance.
	if err := json.NewDecoder(bodyReader).Decode(opData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CustomJSONOperation.JSON: \n%v \n Error : %w", op.JSON, err)
	}

	return opData, nil
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
	case "DeleteReblogOperation":
		tmp = append(tmp, TypeDeleteReblog)
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

//MarshalTransaction is a function of converting type CustomJSONOperation to bytes.
func (op *CustomJSONOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCustomJSON.Code()))
	enc.EncodeArrString(op.RequiredAuths)
	enc.EncodeArrString(op.RequiredPostingAuths)
	enc.Encode(op.ID)
	enc.Encode(op.JSON)
	return enc.Err()
}
