package upbitWebsocket

type RequestType string

const (
	TICKER    = RequestType("ticker")
	ORDERBOOK = RequestType("orderbook")
	TRADE     = RequestType("trade")
	MYTRADE   = RequestType("myTrade")
)
