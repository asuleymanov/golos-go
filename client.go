package client

import (
	"errors"
	// RPC
	"github.com/asuleymanov/golos-go/apis/database"
	"github.com/asuleymanov/golos-go/apis/follow"
	"github.com/asuleymanov/golos-go/apis/network_broadcast"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/transports/websocket"
	"github.com/asuleymanov/golos-go/apis/market"
)

// Client can be used to access Golos remote APIs.
// There is a public field for every Golos API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc transports.CallCloser

	// Database represents database_api.
	Database *database.API

	// Follow represents follow_api.
	Follow *follow.API

	// Follow represents market_history_api.
	Market *market.API

	// NetworkBroadcast represents network_broadcast_api.
	NetworkBroadcast *networkbroadcast.API

	//Chain Id
	Chain *transactions.Chain

	// Current keys for operations
	CurrentKeys *Keys
}

// NewClient creates a new RPC client that use the given CallCloser internally.
// Initialize only server present API. Absent API initialized as nil value.
func NewClient(url []string, chain string) (*Client, error) {
	call, err := initclient(url)
	if err != nil {
		return nil, err
	}
	client := &Client{cc: call}

	client.Database = database.NewAPI(client.cc)

	client.Follow, err = follow.NewAPI(client.cc)
	if err != nil {
		client.Follow = nil
	}

	client.Market, err = market.NewAPI(client.cc)
	if err != nil {
		client.Market = nil
	}

	client.NetworkBroadcast, err = networkbroadcast.NewAPI(client.cc)
	if err != nil {
		client.NetworkBroadcast = nil
	}

	client.Chain, err = initChainID(chain)
	if err != nil {
		client.Chain = transactions.GolosChain
	}

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}

//SetKeys you can specify keys for signing transactions.
func (client *Client) SetKeys(keys *Keys) {
	client.CurrentKeys = keys
}

func initclient(url []string) (*websocket.Transport, error) {
	// Инициализация Websocket
	t, err := websocket.NewTransport(url)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func initChainID(str string) (*transactions.Chain, error) {
	var ChainID transactions.Chain
	// Определяем ChainId
	switch str {
	case "golos":
		ChainID = *transactions.GolosChain
	case "test":
		ChainID = *transactions.TestChain
	default:
		return nil, errors.New("Chain not found")
	}
	return &ChainID, nil
}
