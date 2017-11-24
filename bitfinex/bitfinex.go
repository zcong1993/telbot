package bitfinex

import (
	"context"
	"errors"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"log"
	"time"
)

// Bfx is bitfinex wrapper client
type Bfx struct {
	// Symbols are the ticker pairs you want to subscribe
	Symbols []string
	tickers map[string]float64
	ok      bool
}

// NewBfx create a Bfx instance
func NewBfx(Symbols []string) *Bfx {
	b := &Bfx{Symbols: Symbols, tickers: map[string]float64{}, ok: false}
	go b.run()
	return b
}

func (bfx *Bfx) run() {
	c := bitfinex.NewClient()
	err := c.Websocket.Connect()
	if err != nil {
		log.Fatal("Error connecting to web socket : ", err)
	}
	c.Websocket.SetReadTimeout(time.Second * 10)
	c.Websocket.AttachEventHandler(func(ev interface{}) {
		log.Printf("EVENT: %#v", ev)
	})
	c.Websocket.AttachEventHandler(func(ev interface{}) {
		log.Printf("EVENT: %#v", ev)
	})
	for _, symbol := range bfx.Symbols {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		msg := &bitfinex.PublicSubscriptionRequest{
			Event:   "subscribe",
			Channel: bitfinex.ChanTicker,
			Symbol:  bitfinex.TradingPrefix + symbol,
		}
		h := bfx.createTickerHandler(symbol)
		err = c.Websocket.Subscribe(ctx, msg, h)
		if err != nil {
			log.Fatal(err)
		}
	}
	for {
		select {
		case <-c.Websocket.Done():
			log.Printf("channel closed: %s", c.Websocket.Err())
			bfx.ok = false
			bfx.run()
			return
		}
	}
}

func (bfx *Bfx) createTickerHandler(symbol string) func(ev interface{}) {
	return func(ev interface{}) {
		t, ok := ev.([][]float64)
		if ok {
			last := t[0][6]
			bfx.tickers[symbol] = last
			bfx.ok = true
			log.Printf("PUBLIC MSG %s: %#v", symbol, last)
		} else {
			//log.Printf("PUBLIC MSG HEARTBEAT %s: %#v", symbol, ev)
		}
	}
}

// GetTicker can get all the last price now
func (bfx *Bfx) GetTicker() (map[string]float64, error) {
	if !bfx.ok {
		return map[string]float64{}, errors.New("not ok")
	}
	return bfx.tickers, nil
}
