package api

import (
	"fmt"

	"github.com/asuleymanov/golos-go/transports"
)

//market_history

//GetMarketHistory api request get_market_history
func (api *API) GetMarketHistory(bSec uint32, start, end string) ([]*MarketHistory, error) {
	var resp []*MarketHistory
	err := api.call("market_history", "get_market_history", []interface{}{bSec, start, end}, &resp)
	return resp, err
}

//GetMarketHistoryBuckets api request get_market_history_buckets
func (api *API) GetMarketHistoryBuckets() ([]*uint32, error) {
	var resp []*uint32
	err := api.call("market_history", "get_market_history_buckets", transports.EmptyParams, &resp)
	return resp, err
}

//GetOpenOrders api request get_open_orders
func (api *API) GetOpenOrders(accountName string) ([]*OpenOrders, error) {
	var resp []*OpenOrders
	err := api.call("market_history", "get_open_orders", []string{accountName}, &resp)
	return resp, err
}

//GetOrderBook api request get_order_book
func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("market_history: get_order_book -> limit must not exceed 1000")
	}
	var resp OrderBook
	err := api.call("market_history", "get_order_book", []interface{}{limit}, &resp)
	return &resp, err
}

//GetOrderBookExtended api request get_order_book_extended

//GetRecentTrades api request get_recent_trades
func (api *API) GetRecentTrades(limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("market_history: get_order_book -> limit must not exceed 1000")
	}
	var resp []*Trades
	err := api.call("market_history", "get_recent_trades", []interface{}{limit}, &resp)
	return resp, err
}

//GetTicker api request get_ticker
func (api *API) GetTicker() (*Ticker, error) {
	var resp Ticker
	err := api.call("market_history", "get_ticker", transports.EmptyParams, &resp)
	return &resp, err
}

//GetTradeHistory api request get_trade_history
func (api *API) GetTradeHistory(start, end string, limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("market_history: get_order_book -> limit must not exceed 1000")
	}
	var resp []*Trades
	err := api.call("market_history", "get_trade_history", []interface{}{start, end, limit}, &resp)
	return resp, err
}

//GetVolume api request get_volume
func (api *API) GetVolume() (*Volume, error) {
	var resp Volume
	err := api.call("market_history", "get_volume", transports.EmptyParams, &resp)
	return &resp, err
}
