package okx_connector

import (
	"context"
	"encoding/json"
	"fmt"
)

type WsEventArg struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

// WsKlineEvent define websocket kline event
type WsKlineEvent struct {
	Arg  WsEventArg `json:"arg"`
	Data [][]string `json:"data"`
}

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func (client *WebsocketStreamClient) WsKlineServe(ctx context.Context, symbols []string, channel KlineChannel, handler WsKlineHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	if !channel.Valid() {
		return nil, nil, fmt.Errorf("invalid channel: %s", channel)
	}

	conn, err := client.dial(ctx, "ws/v5/business")
	if err != nil {
		return
	}

	var args []SubOpArg
	for _, symbol := range symbols {
		args = append(args, SubOpArg{Channel: ToPtr(string(channel)), InstId: &symbol})
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
func (client *WebsocketStreamClient) WsDepthServe(ctx context.Context, symbols []string, channel DepthChannel, handler WsDepthHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	if !channel.Valid() {
		return nil, nil, fmt.Errorf("invalid channel: %s", channel)
	}
	conn, err := client.dial(ctx, "ws/v5/public")
	if err != nil {
		return
	}

	var args []SubOpArg
	for _, symbol := range symbols {
		args = append(args, SubOpArg{Channel: ToPtr(string(channel)), InstId: &symbol})
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
	Source  string `json:"source"`
	SeqId   int    `json:"seqId"`
}

type WsTradeEvent struct {
	Arg  WsEventArg `json:"arg"`
	Data []AggTrade `json:"data"`
}

// WsTradeServe serve websocket trade handler with a symbol and interval like 15m, 30s
func (client *WebsocketStreamClient) WsTradeServe(ctx context.Context, symbols []string, handler WsTradeHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	channel := "trades"
	conn, err := client.dial(ctx, "ws/v5/public")
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

type WsTickerHandler func(event *WsTickerEvent)

type WsTickerEvent struct {
	Arg  WsEventArg `json:"arg"`
	Data []Ticker   `json:"data"`
}

type Ticker struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	Last      string `json:"last"`      // 最新成交价
	LastSz    string `json:"lastSz"`    // 最新成交的数量，0 代表没有成交量
	AskPx     string `json:"askPx"`     // 卖一价
	AskSz     string `json:"askSz"`     // 卖一价对应的量
	BidPx     string `json:"bidPx"`     // 买一价
	BidSz     string `json:"bidSz"`     // 买一价对应的数量
	Open24h   string `json:"open24h"`   // 24小时开盘价
	High24h   string `json:"high24h"`   // 24小时最高价
	Low24h    string `json:"low24h"`    // 24小时最低价
	VolCcy24h string `json:"volCcy24h"` // 24小时成交量，以币为单位
	Vol24h    string `json:"vol24h"`    // 24小时成交量，以张为单位
	SodUtc0   string `json:"sodUtc0"`   // UTC+0 时开盘价
	SodUtc8   string `json:"sodUtc8"`   // UTC+8 时开盘价
	Ts        string `json:"ts"`        // 数据产生时间，Unix时间戳的毫秒数格式，如 1597026383085
}

func (client *WebsocketStreamClient) WsTickerServe(ctx context.Context, symbols []string, handler WsTickerHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	channel := "tickers"
	conn, err := client.dial(ctx, "ws/v5/public")
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
		event := new(WsTickerEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}
