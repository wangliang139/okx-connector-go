package okx_connector

import (
	"context"
	"encoding/json"
	"fmt"
)

type WsEventArg struct {
	Channel    string `json:"channel,omitempty"`
	Uid        string `json:"uid,omitempty"`
	InstId     string `json:"instId,omitempty"`
	InstType   string `json:"instType,omitempty"`
	InstFamily string `json:"instFamily,omitempty"`
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
		args = append(args, SubOpArg{Channel: string(channel), InstId: &symbol})
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
		args = append(args, SubOpArg{Channel: string(channel), InstId: &symbol})
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
		args = append(args, SubOpArg{Channel: channel, InstId: &symbol})
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
		args = append(args, SubOpArg{Channel: channel, InstId: &symbol})
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

type WsAccountEvent struct {
	Arg       WsEventArg       `json:"arg"`
	EventType string           `json:"eventType"`
	CurPage   int              `json:"curPage"`
	LastPage  bool             `json:"lastPage"`
	Data      []AccountBalance `json:"data"`
}

type WsAccountHandler func(event *WsAccountEvent)

func (client *WebsocketStreamClient) WsAccountServe(ctx context.Context, handler WsAccountHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	channel := "account"
	conn, err := client.dial(ctx, "ws/v5/private")
	if err != nil {
		return
	}

	if err = conn.login(); err != nil {
		return
	}

	args := []SubOpArg{{Channel: channel}}

	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsAccountEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}

type WsPositionEvent struct {
	Arg       WsEventArg `json:"arg"`
	EventType string     `json:"eventType"`
	CurPage   int        `json:"curPage"`
	LastPage  bool       `json:"lastPage"`
	Data      []Position `json:"data"`
}

type WsPositionHandler func(event *WsPositionEvent)

func (client *WebsocketStreamClient) WsPositionServe(ctx context.Context, handler WsPositionHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	channel := "positions"
	conn, err := client.dial(ctx, "ws/v5/private")
	if err != nil {
		return
	}

	if err = conn.login(); err != nil {
		return
	}

	instType := "ANY"
	extraParams := "{\"updateInterval\": \"0\"}"
	args := []SubOpArg{{Channel: channel, InstType: &instType, ExtraParams: &extraParams}}

	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsPositionEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}

type WsOrderEvent struct {
	Arg  WsEventArg `json:"arg"`
	Data []Order    `json:"data"`
}

type WsOrderHandler func(event *WsOrderEvent)

func (client *WebsocketStreamClient) WsOrderServe(ctx context.Context, handler WsOrderHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	channel := "orders"
	conn, err := client.dial(ctx, "ws/v5/private")
	if err != nil {
		return
	}

	if err = conn.login(); err != nil {
		return
	}

	instType := "ANY"
	args := []SubOpArg{{Channel: channel, InstType: &instType}}

	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsOrderEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}

type WsFillsEvent struct {
	Arg  WsEventArg `json:"arg"`
	Data []Fill     `json:"data"`
}

type Fill struct {
	InstId   string `json:"instId"`
	FillSz   string `json:"fillSz"`
	FillPx   string `json:"fillPx"`
	Side     string `json:"side"`
	Ts       string `json:"ts"`
	OrdId    string `json:"ordId"`
	ClOrdId  string `json:"clOrdId"`
	TradeId  string `json:"tradeId"`
	ExecType string `json:"execType"`
	Count    string `json:"count"`
}

type WsFillsHandler func(event *WsFillsEvent)

func (client *WebsocketStreamClient) WsFillsServe(ctx context.Context, handler WsFillsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	channel := "fills"
	conn, err := client.dial(ctx, "ws/v5/private")
	if err != nil {
		return
	}

	if err = conn.login(); err != nil {
		return
	}

	args := []SubOpArg{{Channel: channel}}

	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsFillsEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}

type WsBalanceAndPositionEvent struct {
	Arg  WsEventArg      `json:"arg"`
	Data []BalAndPosData `json:"data"`
}

type BalData struct {
	Ccy     string `json:"ccy"`     // 币种
	CashBal string `json:"cashBal"` // 币种余额
	UTime   string `json:"uTime"`   // 币种余额信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
}

type PosData struct {
	PosId          string `json:"posId"`          // 持仓ID
	TradeId        string `json:"tradeId"`        // 最新成交ID
	InstId         string `json:"instId"`         // 交易产品ID，如 BTC-USD-180213
	InstType       string `json:"instType"`       // 交易产品类型 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	MgnMode        string `json:"mgnMode"`        // 保证金模式: isolated, cross
	AvgPx          string `json:"avgPx"`          // 开仓平均价
	Ccy            string `json:"ccy"`            // 占用保证金的币种
	PosSide        string `json:"posSide"`        // 持仓方向: long，short，net
	Pos            string `json:"pos"`            // 持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为0的仓位
	BaseBal        string `json:"baseBal"`        // 交易币余额,适用于 币币杠杆（逐仓一键借币模式）（已弃用）
	QuoteBal       string `json:"quoteBal"`       // 计价币余额,适用于 币币杠杆（逐仓一键借币模式）（已弃用）
	PosCcy         string `json:"posCcy"`         // 持仓数量币种,只适用于币币杠杆仓位。当是交割、永续、期权持仓时，该字段返回“”
	NonSettleAvgPx string `json:"nonSettleAvgPx"` // 未结算均价,不受结算影响的加权开仓价格，仅在新增头寸时更新，和开仓均价的主要区别在于是否受到结算影响。适用于全仓交割
	SettledPnl     string `json:"settledPnl"`     // 累计已结算收益（以结算价格计算）,适用于全仓交割
	UTime          string `json:"uTime"`          // 仓位信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
}

type TradeData struct {
	InstId  string `json:"instId"`  // 交易产品ID，如 BTC-USD-180213
	TradeId string `json:"tradeId"` // 最新成交ID
}

type BalAndPosData struct {
	PTime     string      `json:"pTime"`     // 推送时间，Unix时间戳的毫秒数格式，如 1597026383085
	EventType string      `json:"eventType"` // 事件类型
	BalData   []BalData   `json:"balData"`   // 余额数据
	PosData   []PosData   `json:"posData"`   // 持仓数据
	Trades    []TradeData `json:"trades"`    // 成交数据
}

type WsBalanceAndPositionHandler func(event *WsBalanceAndPositionEvent)

func (client *WebsocketStreamClient) WsBalanceAndPositionServe(ctx context.Context, handler WsBalanceAndPositionHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	channel := "balance_and_position"
	conn, err := client.dial(ctx, "ws/v5/private")
	if err != nil {
		return
	}

	if err = conn.login(); err != nil {
		return
	}

	args := []SubOpArg{{Channel: channel}}

	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		event := new(WsBalanceAndPositionEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return conn.serve(wsHandler, errHandler)
}

type WsUserDataHandler interface {
	HandleAccountEvent(event *WsAccountEvent)
	HandlePositionEvent(event *WsPositionEvent)
	HandleOrderEvent(event *WsOrderEvent)
	HandleBalanceAndPositionEvent(event *WsBalanceAndPositionEvent)
}

func (client *WebsocketStreamClient) WsUserDataServe(ctx context.Context, handler WsUserDataHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	conn, err := client.dial(ctx, "ws/v5/private")
	if err != nil {
		return
	}

	if err = conn.login(); err != nil {
		return
	}

	args := []SubOpArg{
		{Channel: "account"},
		{
			Channel:     "positions",
			InstType:    ToPtr("ANY"),
			ExtraParams: ToPtr("{\"updateInterval\": \"0\"}"),
		},
		{Channel: "orders", InstType: ToPtr("ANY")},
		{Channel: "balance_and_position"},
	}

	err = conn.subscribe(args)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		e := new(WsSimpleEvent)
		err := json.Unmarshal(message, e)
		if err != nil {
			errHandler(err)
			return
		}
		switch e.Arg.Channel {
		case "account":
			event := new(WsAccountEvent)
			err := json.Unmarshal(message, event)
			if err != nil {
				errHandler(err)
				return
			}
			handler.HandleAccountEvent(event)
		case "positions":
			event := new(WsPositionEvent)
			err := json.Unmarshal(message, event)
			if err != nil {
				errHandler(err)
				return
			}
			handler.HandlePositionEvent(event)
		case "orders":
			event := new(WsOrderEvent)
			err := json.Unmarshal(message, event)
			if err != nil {
				errHandler(err)
				return
			}
			handler.HandleOrderEvent(event)
		case "balance_and_position":
			event := new(WsBalanceAndPositionEvent)
			err := json.Unmarshal(message, event)
			if err != nil {
				errHandler(err)
				return
			}
			handler.HandleBalanceAndPositionEvent(event)
		}
	}
	return conn.serve(wsHandler, errHandler)
}

type WsSimpleEvent struct {
	Arg WsEventArg `json:"arg"`
}

type WsCommonHandler func(channel string, arg WsEventArg, rawMessage []byte)

func (client *WebsocketStreamClient) WsCommonServe(ctx context.Context, path string, isPrivate bool, channels []SubOpArg, handler WsCommonHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	if len(channels) == 0 {
		return nil, nil, fmt.Errorf("channels is empty")
	}

	conn, err := client.dial(ctx, path)
	if err != nil {
		return
	}

	if isPrivate {
		if err = conn.login(); err != nil {
			return
		}
	}

	err = conn.subscribe(channels)
	if err != nil {
		return
	}

	wsHandler := func(message []byte) {
		client.debug("Receive event: %s", message)
		e := new(WsSimpleEvent)
		err := json.Unmarshal(message, e)
		if err != nil {
			errHandler(err)
			return
		}
		handler(e.Arg.Channel, e.Arg, message)
	}
	return conn.serve(wsHandler, errHandler)
}
