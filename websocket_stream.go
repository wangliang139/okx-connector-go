package okx_connector

import (
	"context"
	"encoding/json"
)

type WsEventArg struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
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
func (client *WebsocketStreamClient) WsKlineServe(ctx context.Context, symbols []string, channel string, handler WsKlineHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	conn, err := client.dail(ctx)
	if err != nil {
		return
	}

	var args []SubOpArg
	for _, symbol := range symbols {
		args = append(args, SubOpArg{Channel: &channel, InstId: &symbol})
	}
	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
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
func (client *WebsocketStreamClient) WsDepthServe(ctx context.Context, symbols []string, channel string, handler WsDepthHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	conn, err := client.dail(ctx)
	if err != nil {
		return
	}

	var args []SubOpArg
	for _, symbol := range symbols {
		args = append(args, SubOpArg{Channel: &channel, InstId: &symbol})
	}
	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsDepthEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}

type WsTradeHandler func(event *WsTradeEvent)

type AggTrade struct {
	InstId  string `json:"instId"`
	TradeId string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	Ts      string `json:"ts"`
	Count   string `json:"count"`
}

type WsTradeEvent struct {
	Arg  WsEventArg `json:"arg"`
	Data []AggTrade `json:"data"`
}

// WsTradeServe serve websocket trade handler with a symbol and interval like 15m, 30s
func (client *WebsocketStreamClient) WsTradeServe(ctx context.Context, symbols []string, channel string, handler WsTradeHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	conn, err := client.dail(ctx)
	if err != nil {
		return
	}

	var args []SubOpArg
	for _, symbol := range symbols {
		args = append(args, SubOpArg{Channel: &channel, InstId: &symbol})
	}
	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}
