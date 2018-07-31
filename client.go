package client

import (
	"errors"
	// RPC
	"github.com/asuleymanov/golos-go/apis/account_by_key"
	"github.com/asuleymanov/golos-go/apis/account_history"
	"github.com/asuleymanov/golos-go/apis/database"
	"github.com/asuleymanov/golos-go/apis/follow"
	"github.com/asuleymanov/golos-go/apis/market_history"
	"github.com/asuleymanov/golos-go/apis/network_broadcast"
	"github.com/asuleymanov/golos-go/apis/operation_history"
	"github.com/asuleymanov/golos-go/apis/private_message"
	"github.com/asuleymanov/golos-go/apis/social_network"
	"github.com/asuleymanov/golos-go/apis/tags"
	"github.com/asuleymanov/golos-go/apis/witness"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/transports/websocket"
	"github.com/asuleymanov/golos-go/types"
)

// Client can be used to access Golos remote APIs.
// There is a public field for every Golos API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc transports.CallCloser

	AsyncProtocol bool

	// Fixed JSONMetadata added to posting all comments
	DefaultContentMetadata types.ContentMetadata

	// Database represents database_api.
	Database *database.API

	// Follow represents follow_api.
	Follow *follow.API

	// Follow represents market_history_api.
	MarketHistory *market_history.API

	// NetworkBroadcast represents network_broadcast_api.
	NetworkBroadcast *network_broadcast.API

	// AccountByKey represents account_by_key.
	AccountByKey *account_by_key.API

	// SocialNetwork represents social_network.
	SocialNetwork *social_network.API

	// PrivateMessage represents social_network.
	PrivateMessage *private_message.API

	Witness *witness.API

	AccountHistory *account_history.API

	OperationHistory *operation_history.API

	Tags *tags.API

	//Chain Id
	Chain *transactions.Chain

	// Current keys for operations
	CurrentKeys *Keys
}

// NewClient creates a new RPC client that use the given CallCloser internally.
// Initialize only server present API. Absent API initialized as nil value.
func NewClient(url []string, chain string, options ...websocket.Option) (*Client, error) {
	call, err := initClient(url, options...)
	if err != nil {
		return nil, err
	}
	client := &Client{cc: call}

	client.AsyncProtocol = false

	client.Database = database.NewAPI(client.cc)

	client.Follow = follow.NewAPI(client.cc)

	client.MarketHistory = market_history.NewAPI(client.cc)

	client.NetworkBroadcast = network_broadcast.NewAPI(client.cc)

	client.AccountByKey = account_by_key.NewAPI(client.cc)

	client.SocialNetwork = social_network.NewAPI(client.cc)

	client.PrivateMessage = private_message.NewAPI(client.cc)

	client.Witness = witness.NewAPI(client.cc)

	client.AccountHistory = account_history.NewAPI(client.cc)

	client.OperationHistory = operation_history.NewAPI(client.cc)

	client.Tags = tags.NewAPI(client.cc)

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

func initClient(url []string, options ...websocket.Option) (*websocket.Transport, error) {
	// Инициализация Websocket
	t, err := websocket.NewTransport(url, options...)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func initChainID(str string) (*transactions.Chain, error) {
	var chainID transactions.Chain
	// Определяем ChainId
	switch str {
	case "golos":
		chainID = *transactions.GolosChain
	case "test":
		chainID = *transactions.TestChain
	default:
		return nil, errors.New("Chain not found")
	}
	return &chainID, nil
}

func (client *Client) GenCommentMetadata(meta *types.ContentMetadata) *types.ContentMetadata {
	if client.DefaultContentMetadata != nil {
		for k := range client.DefaultContentMetadata {
			_, ok := (*meta)[k]
			if !ok {
				// Set fixed value only if value not exists
				(*meta)[k] = client.DefaultContentMetadata[k]
			}
		}
	}
	return meta
}
