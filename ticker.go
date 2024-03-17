package upbit

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
	"upbit/model"
	"upbit/model/quotation"
	"upbit/model/upbitWebsocket"
)

// GetTickers 현재가 정보. 최대 100개의 정보를 반환
//
// [QUERY PARAMS]
//
// markets : REQUIRED. 마켓 코드 목록 (ex. KRW-BTC,BTC-BCC)
func (u *Upbit) _getWS(ticketName string, markets []string) *websocket.Conn {
	c, _ := u.createWebSocket()

	if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
		log.Fatalf("Ping attempt but WS closed: %v", err)

	}
	//defer tictiemTickerker.Stop()

	//_ = c.WriteJSON(_createWSRequestBody(ticketName, upbitWebsocket.TICKER, markets))

	return c

}

// call with goroutine
func (u *Upbit) _processWS(websocketConn *websocket.Conn, ticketName string, markets []string) {
	defer websocketConn.Close()

	for {
		_, message, err := websocketConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
	}

}

func (u *Upbit) GetTickersWSChannel(markets []string) ([]*quotation.Ticker, *model.Remaining, error) {

}

func (u *Upbit) GetTickersWSProcess() {

}

func (u *Upbit) GetTickersWSBlockingStringStreamForlogging(ticketName string, markets []string) ([]*quotation.TickerWebSocket, error) {
	if len(markets) == 0 || markets == nil {
		return nil, fmt.Errorf("invalid markets")
	}

	c, _ := u.createWebSocket()

	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	_ = c.WriteJSON(_createWSRequestBody(ticketName, upbitWebsocket.TICKER, markets))

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {

			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("recv: %s, type: %d", message, mt)
		}
	}()

	for {
		select {
		case <-done:
			return nil, nil
		case <-ticker.C:
			if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Fatalf("Ping attempt but WS closed: %v", err)

			}

		}

	}

	return nil, nil
}

func (u *Upbit) GetTickers(markets []string) ([]*quotation.Ticker, *model.Remaining, error) {
	if len(markets) == 0 || markets == nil {
		return nil, nil, fmt.Errorf("invalid markets")
	}

	api, e := GetApiInfo(FuncGetTickers)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"markets": markets,
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, values, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	tickers, e := quotation.TickersFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return tickers, model.RemainingFromHeader(resp.Header), nil
}
