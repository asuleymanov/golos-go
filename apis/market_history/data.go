package market_history

import (
	"github.com/asuleymanov/golos-go/types"
)

//MarketHistory structure for the GetMarketHistory function.
type MarketHistory struct {
	ID          uint32      `json:"id"`
	Open        *types.Time `json:"open"`
	Seconds     int         `json:"seconds"`
	HighSteem   int         `json:"high_steem"`
	HighSbd     int         `json:"high_sbd"`
	LowSteem    int         `json:"low_steem"`
	LowSbd      int         `json:"low_sbd"`
	OpenSteem   int         `json:"open_steem"`
	OpenSbd     int         `json:"open_sbd"`
	CloseSteem  int         `json:"close_steem"`
	CloseSbd    int         `json:"close_sbd"`
	SteemVolume int         `json:"steem_volume"`
	SbdVolume   int         `json:"sbd_volume"`
}

//OpenOrders structure for the GetOpenOrders function.
type OpenOrders struct {
	ID         *types.ID   `json:"id"`
	Created    types.Time  `json:"created"`
	Expiration types.Time  `json:"expiration"`
	Seller     string      `json:"seller"`
	Orderid    *types.Int  `json:"orderid"`
	ForSale    *types.Int  `json:"for_sale"`
	SellPrice  *OrderPrice `json:"sell_price"`
	RealPrice  string      `json:"real_price"`
	Rewarded   bool        `json:"rewarded"`
}

//OrderPrice additional structure for the function GetOpenOrders.
type OrderPrice struct {
	Base  *types.Asset `json:"base"`
	Quote *types.Asset `json:"quote"`
}

//OrderBook structure for the GetOrderBook function.
type OrderBook struct {
	Ask []*OrderBookAB `json:"asks"`
	Bid []*OrderBookAB `json:"bids"`
}

//OrderBookAB additional structure for the function GetOrderBook.
type OrderBookAB struct {
	OrderPrice *OrderPrice `json:"order_price"`
	RealPrice  string      `json:"real_price"`
	Steem      uint        `json:"steem"`
	Sbd        uint        `json:"sbd"`
	Created    *types.Time `json:"created"`
}

//Trades structure for the GetRecentTrades and GetTradeHistory functions.
type Trades struct {
	Date        *types.Time  `json:"date"`
	CurrentPays *types.Asset `json:"current_pays"`
	OpenPays    *types.Asset `json:"open_pays"`
}

//Ticker structure for the GetTicker function.
type Ticker struct {
	Latest        string       `json:"latest"`
	LowestAsk     string       `json:"lowest_ask"`
	HighestBid    string       `json:"highest_bid"`
	PercentChange string       `json:"percent_change"`
	SteemVolume   *types.Asset `json:"steem_volume"`
	SbdVolume     *types.Asset `json:"sbd_volume"`
}

//Volume structure for the GetVolume function.
type Volume struct {
	SteemVolume *types.Asset `json:"steem_volume"`
	SbdVolume   *types.Asset `json:"sbd_volume"`
}
