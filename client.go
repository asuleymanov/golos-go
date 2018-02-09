package client

import (
	"errors"
	// RPC
	"github.com/asuleymanov/golos/apis/database"
	"github.com/asuleymanov/golos/apis/follow"
	"github.com/asuleymanov/golos/apis/market"
	"github.com/asuleymanov/golos/apis/networkbroadcast"
	"github.com/asuleymanov/golos/transactions"
	"github.com/asuleymanov/golos/transports"
	"github.com/asuleymanov/golos/transports/websocket"
)

var Key_List = make(map[string]Keys)

// Client can be used to access Steem remote APIs.
//
// There is a public field for every Steem API available,
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
}

// NewClient creates a new RPC client that use the given CallCloser internally.
func NewClient(url []string, chain string) (*Client, error) {
	call, err := initclient(url)
	if err != nil {
		return nil, err
	}
	client := &Client{cc: call}

	client.Database = database.NewAPI(client.cc)

	followAPI, err := follow.NewAPI(client.cc)
	if err != nil {
		return nil, err
	}
	client.Follow = followAPI

	marketAPI, err := market.NewAPI(client.cc)
	if err != nil {
		return nil, err
	}
	client.Market = marketAPI

	networkBroadcastAPI, err := networkbroadcast.NewAPI(client.cc)
	if err != nil {
		return nil, err
	}
	client.NetworkBroadcast = networkBroadcastAPI

	chainid, err := initChainId(chain)
	if err != nil {
		return nil, err
	}
	client.Chain = chainid

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}

func initclient(url []string) (*websocket.Transport, error) {
	// Инициализация Websocket
	t, err := websocket.NewTransport(url)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func initChainId(str string) (*transactions.Chain, error) {
	var ChainId transactions.Chain
	// Определяем ChainId
	switch str {
	case "golos":
		ChainId = *transactions.GolosChain
	case "test":
		ChainId = *transactions.TestChain
	default:
		return nil, errors.New("test")
	}
	return &ChainId, nil
}
