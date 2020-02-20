package types

import (
	"encoding/json"
)

//RPCRequest server request
type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
	JSON   string      `json:"jsonrpc"`
	ID     uint64      `json:"id"`
}

//RPCIncoming incoming request from server
type RPCIncoming struct {
	ID     uint64          `json:"id"`
	JSON   string          `json:"jsonrpc"`
	Result json.RawMessage `json:"result"`
}
