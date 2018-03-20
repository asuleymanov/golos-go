package client

import (
	"errors"
	// RPC
	"github.com/asuleymanov/golos-go/apis/accountbykey"
	"github.com/asuleymanov/golos-go/apis/database"
	"github.com/asuleymanov/golos-go/apis/follow"
	"github.com/asuleymanov/golos-go/apis/markethistory"
	"github.com/asuleymanov/golos-go/apis/networkbroadcast"
	"github.com/asuleymanov/golos-go/apis/socialnetwork"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/transports/websocket"
)

var Key_List = make(map[string]Keys)

// Client can be used to access Golos remote APIs.
// There is a public field for every Golos API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc transports.CallCloser

	// Database represents database_api.
	Database *database.API

	// Follow represents follow.
	Follow *follow.API

	// MarketHistory represents market_history.
	MarketHistory *market_history.API

	// NetworkBroadcast represents network_broadcast_api.
	NetworkBroadcast *network_broadcast.API

	// AccountByKey represents account_by_key.
	AccountByKey *account_by_key.API

	// SocialNetwork represents social_network.
	SocialNetwork *social_network.API

	//Chain Id
	Chain *transactions.Chain
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

	client.Follow = follow.NewAPI(client.cc)

	client.MarketHistory = market_history.NewAPI(client.cc)

	client.NetworkBroadcast = network_broadcast.NewAPI(client.cc)

	client.AccountByKey = account_by_key.NewAPI(client.cc)

	client.SocialNetwork = social_network.NewAPI(client.cc)

	client.Chain, err = initChainId(chain)
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
		return nil, errors.New("Chain not found")
	}
	return &ChainId, nil
}
