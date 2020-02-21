package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/asuleymanov/golos-go/types"
)

type Transport struct {
	Url    string
	client http.Client

	requestID uint64
	reqMutex  sync.Mutex
}

func NewTransport(url string) (*Transport, error) {
	timeout := time.Duration(5 * time.Second)

	return &Transport{
		client: http.Client{
			Timeout: timeout,
		},
		Url: url,
	}, check(url)
}

func (caller *Transport) Call(method string, args []interface{}, reply interface{}) error {
	caller.reqMutex.Lock()
	defer caller.reqMutex.Unlock()

	// increase request id
	if caller.requestID == math.MaxUint64 {
		caller.requestID = 0
	}
	caller.requestID++

	request := types.RPCRequest{
		Method: method,
		JSON:   "2.0",
		ID:     caller.requestID,
		Params: args,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp, err := caller.client.Post(caller.Url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body : %s", err)
	}

	var rpcResponse types.RPCResponse
	if err = json.Unmarshal(respBody, &rpcResponse); err != nil {
		return fmt.Errorf("failed to unmarshal response: %+v \n Error : %s", string(respBody), err)
	}

	if rpcResponse.Error != nil {
		return rpcResponse.Error
	}

	if rpcResponse.Result != nil {
		if err := json.Unmarshal(*rpcResponse.Result, reply); err != nil {
			return fmt.Errorf("failed to unmarshal rpc result: %+v \n Error : %s", string(*rpcResponse.Result), err)
		}
	}

	return nil
}

func (caller *Transport) SetCallback(api string, method string, notice func(args json.RawMessage)) error {
	panic("not supported")
}

func (caller *Transport) Close() error {
	return nil
}

func check(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error. URL: %s STATUS: %s\n", url, resp.StatusCode)
	}
	return nil
}
