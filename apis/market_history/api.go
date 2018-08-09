package market_history

import (
	"encoding/json"

	"github.com/asuleymanov/golos-go/transports"
	"github.com/pkg/errors"
)

const apiID = "market_history"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = []string{}

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//GetMarketHistory api request get_market_history
func (api *API) GetMarketHistory(bSec uint32, start, end string) ([]*MarketHistory, error) {
	raw, err := api.raw("get_market_history", []interface{}{bSec, start, end})
	if err != nil {
		return nil, err
	}
	var resp []*MarketHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_market_history response", apiID)
	}
	return resp, nil
}

//GetMarketHistoryBuckets api request get_market_history_buckets
func (api *API) GetMarketHistoryBuckets() ([]uint32, error) {
	raw, err := api.raw("get_market_history_buckets", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp []uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_market_history_buckets response", apiID)
	}
	return resp, nil
}

//GetOpenOrders api request get_open_orders
func (api *API) GetOpenOrders(accountName string) ([]*OpenOrders, error) {
	raw, err := api.raw("get_open_orders", []string{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*OpenOrders
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_open_orders response", apiID)
	}
	return resp, nil
}

//GetOrderBook api request get_order_book
func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, errors.Errorf("%v: get_order_book -> limit must not exceed 1000", apiID)
	}
	raw, err := api.raw("get_order_book", []interface{}{limit})
	if err != nil {
		return nil, err
	}
	var resp *OrderBook
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_order_book response", apiID)
	}
	return resp, nil
}

//GetOrderBookExtended api request get_order_book_extended

//GetRecentTrades api request get_recent_trades
func (api *API) GetRecentTrades(limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, errors.Errorf("%v: get_order_book -> limit must not exceed 1000", apiID)
	}
	raw, err := api.raw("get_recent_trades", []interface{}{limit})
	if err != nil {
		return nil, err
	}
	var resp []*Trades
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_recent_trades response", apiID)
	}
	return resp, nil
}

//GetTicker api request get_ticker
func (api *API) GetTicker() (*Ticker, error) {
	raw, err := api.raw("get_ticker", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp *Ticker
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_ticker response", apiID)
	}
	return resp, nil
}

//GetTradeHistory api request get_trade_history
func (api *API) GetTradeHistory(start, end string, limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, errors.Errorf("%v: get_order_book -> limit must not exceed 1000", apiID)
	}
	raw, err := api.raw("get_trade_history", []interface{}{start, end, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Trades
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_trade_history response", apiID)
	}
	return resp, nil
}

//GetVolume api request get_volume
func (api *API) GetVolume() (*Volume, error) {
	raw, err := api.raw("get_volume", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp *Volume
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "%v: failed to unmarshal get_volume response", apiID)
	}
	return resp, nil
}
