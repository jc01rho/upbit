package upbit

import "testing"

func TestTicker(t *testing.T) {
	u := NewUpbit("", "")

	tickers, remaining, e := u.GetTickers([]string{marketID})
	if e != nil {
		t.Fatalf("%s's GetTickers error : %s", marketID, e.Error())
	} else {
		t.Logf("GetTickers[remaining:%+v]", *remaining)
		for _, ticker := range tickers {
			t.Logf("%+v", *ticker)
		}
	}
}

func TestTickerWS(t *testing.T) {
	u := NewUpbit("", "")

	//go func() {
	//	u.GetTickersWSBlockingStringStreamForlogging("ticker1", []string{marketID})
	//}()
	//go func() {
	//	u.GetTickersWSBlockingStringStreamForlogging("ticker2", []string{marketID})
	//}()

	u.GetTickersWSBlockingStringStreamForlogging("ticker3", []string{marketID})
	//if e != nil {
	//	t.Fatalf("%s's GetTickersWSBlockingStringStreamForlogging error : %s", marketID, e.Error())
	//} else {
	//	for _, ticker := range tickers {
	//		t.Logf("%+v", *ticker)
	//	}
	//}
}
