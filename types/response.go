package types

import (
	"encoding/json"
	"fmt"
)

//OperationResponse of response when sending a transaction.
type OperationResponse struct {
	ID       string
	BlockNum int32
	TrxNum   int32
	Expired  bool
	JSONTrx  string
}

//RPCResponse response to a server request
type RPCResponse struct {
	Result *json.RawMessage `json:"result,omitempty"`
	Error  *RPCError        `json:"error,omitempty"`
	JSON   string           `json:"jsonrpc,omitempty"`
	ID     uint64           `json:"id"`
}

//RPCError description of error in response
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Code    int    `json:"code"`
		Name    string `json:"name"`
		Message string `json:"message"`
		Stack   []struct {
			Context struct {
				Level      string `json:"level"`
				File       string `json:"file"`
				Line       int    `json:"line"`
				Method     string `json:"method"`
				Hostname   string `json:"hostname"`
				ThreadName string `json:"thread_name"`
				Timestamp  string `json:"timestamp"`
			} `json:"context"`
			Format string      `json:"format"`
			Data   interface{} `json:"data"`
		} `json:"stack"`
	} `json:"data"`
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%d: %s\n %#v", e.Code, e.Message, e.Data)
}
