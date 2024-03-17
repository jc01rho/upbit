package upbit

import (
	"fmt"
	"log"
	"net/url"
	"upbit/model"
	"upbit/model/quotation"
)

// GetTickers 현재가 정보. 최대 100개의 정보를 반환
//
// [QUERY PARAMS]
//
// markets : REQUIRED. 마켓 코드 목록 (ex. KRW-BTC,BTC-BCC)
func (u *Upbit) GetTickersWS(markets []string) ([]*quotation.TickerWebSocket, error) {
	if len(markets) == 0 || markets == nil {
		return nil, fmt.Errorf("invalid markets")
	}

	//api, e := GetApiInfo(FuncGetTickersWS)
	//if e != nil {
	//	return nil, e
	//}
	//
	//var values = url.Values{
	//	"markets": markets,
	//}
	//
	//req, e := u.createRequest(api.Method, BaseURI+api.Url, values, api.Section)
	//if e != nil {
	//	return nil, e
	//}
	//
	//resp, e := u.do(req, api.Group)
	//if e != nil {
	//	return nil, e
	//}
	//defer resp.Body.Close()
	c, _ := u.createWebSocket()

	wsReqList := make([]interface{}, 0)
	wsReqList = append(wsReqList, &model.WebSocketTicket{

		Ticket: "ticker",
	})
	wsReqList = append(wsReqList, &model.WebSocketType{

		Type:           "ticker",
		Codes:          markets,
		IsOnlySnapshot: nil,
		IsOnlyRealtime: nil,
	})

	//c.SetReadDeadline(time.Now().Add(1))
	//c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(1)); return nil })

	c.WriteMessage(1, []byte("PING"))
	c.WriteJSON(wsReqList)

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
		}

	}

	//tickers, e := quotation.TickerWSsFromJSON(resp.Body)
	//if e != nil {
	//	return nil, e
	//}

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
