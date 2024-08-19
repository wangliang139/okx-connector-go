package okx_connector

import (
	"encoding/json"
	"log"
)

type WsEventArg struct {
	Channel string
	InstId  string
}

// WsKlineEvent define websocket kline event
type WsKlineEvent struct {
	Event  string  `json:"e"`
	Time   int64   `json:"E"`
	Symbol string  `json:"s"`
	Kline  WsKline `json:"k"`
}

// WsKline define websocket kline
type WsKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Symbol               string `json:"s"`
	Interval             string `json:"i"`
	FirstTradeID         int64  `json:"f"`
	LastTradeID          int64  `json:"L"`
	Open                 string `json:"o"`
	Close                string `json:"c"`
	High                 string `json:"h"`
	Low                  string `json:"l"`
	Volume               string `json:"v"`
	TradeNum             int64  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func (client *WebsocketStreamClient) WsKlineServe(symbols []string, channel string, handler WsKlineHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	err = client.dail()
	if err != nil {
		return
	}

	var args []SubOpArg
	for _, symbol := range symbols {
		args = append(args, SubOpArg{Channel: &channel, InstId: &symbol})
	}
	err = client.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		log.Printf("Receive event: %s", message)
		event := new(WsKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return client.serve(wsHandler, errHandler)
}

type WsDepthHandler func(event *WsDepthEvent)

type SymbolDepth struct {
	Ts        string     `json:"ts"`
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Checksum  int        `json:"checksum"`
	PrevSeqId int        `json:"prevSeqId"`
	SeqId     int        `json:"seqId"`
}

type WsDepthEvent struct {
	Arg    WsEventArg    `json:"arg"`
	Action string        `json:"action"`
	Data   []SymbolDepth `json:"data"`
}

// WsDepthServe serve websocket depth handler with a symbol and interval like 15m, 30s
func (client *WebsocketStreamClient) WsDepthServe(symbols []string, channel string, handler WsDepthHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	err = client.dail()
	if err != nil {
		return
	}

	var args []SubOpArg
	for _, symbol := range symbols {
		args = append(args, SubOpArg{Channel: &channel, InstId: &symbol})
	}
	err = client.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		log.Printf("Receive event: %s", message)
		event := new(WsDepthEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return client.serve(wsHandler, errHandler)
}
