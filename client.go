package golos

import (
	"errors"
	"net/url"

	"github.com/asuleymanov/golos-go/api"
	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/transports/http"
	"github.com/asuleymanov/golos-go/transports/websocket"
)

var (
	ErrInitializeTransport = errors.New("Failed to initialize transport.")
)

// Client can be used to access GOLOS remote APIs.
// There is a function for every available GOLOS API,
// for example, Client.API.GetDatabaseInfo() corresponds to database_api -> get_database_info.
type Client struct {
	cc transports.CallCloser

	asyncProtocol bool

	API *api.API

	Config api.Config

	// Current keys for operations
	CurrentKeys *Keys
}

// NewClient creates a new RPC client that use the given CallCloser internally.
// Initialize only server present API. Absent API initialized as nil value.
func NewClient(apiURL string) (*Client, error) {
	// Parse URL
	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	// Initializing Transport
	var call transports.CallCloser
	switch u.Scheme {
	case "wss", "ws":
		call, err = websocket.NewTransport(apiURL)
		if err != nil {
			return nil, err
		}
	case "https", "http":
		call, err = http.NewTransport(apiURL)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrInitializeTransport
	}
	client := &Client{cc: call}

	client.asyncProtocol = false

	client.API = api.NewAPI(client.cc)

	conf, err := client.API.GetConfig()
	if err != nil {
		return nil, err
	}
	client.Config = *conf

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}
