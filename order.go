package okx_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

type OpenOrdersService struct {
	c *Client

	instType   string
	instId     string
	instFamily string
	ordType    string
	state      string
	after      string
	before     string
	limit      string
}

func (s *OpenOrdersService) InstId(instId string) *OpenOrdersService {
	s.instId = instId
	return s
}

func (s *OpenOrdersService) InstFamily(instFamily string) *OpenOrdersService {
	s.instFamily = instFamily
	return s
}

func (s *OpenOrdersService) OrdType(ordType string) *OpenOrdersService {
	s.ordType = ordType
	return s
}

func (s *OpenOrdersService) State(state string) *OpenOrdersService {
	s.state = state
	return s
}

func (s *OpenOrdersService) After(after string) *OpenOrdersService {
	s.after = after
	return s
}

func (s *OpenOrdersService) Before(before string) *OpenOrdersService {
	s.before = before
	return s
}

func (s *OpenOrdersService) Limit(limit string) *OpenOrdersService {
	s.limit = limit
	return s
}

func (s *OpenOrdersService) InstType(instType string) *OpenOrdersService {
	s.instType = instType
	return s
}

type Order struct {
	InstType          string             `json:"instType"`          // 产品类型
	InstId            string             `json:"instId"`            // 产品ID
	TgtCcy            string             `json:"tgtCcy"`            // 币币市价单委托数量sz的单位
	Ccy               string             `json:"ccy"`               // 保证金币种，适用于逐仓杠杆及合约模式下的全仓杠杆订单以及交割、永续和期权合约订单。
	OrdId             string             `json:"ordId"`             // 订单ID
	ClOrdId           string             `json:"clOrdId"`           // 客户自定义订单ID
	Tag               string             `json:"tag"`               // 订单标签
	Px                string             `json:"px"`                // 委托价格，对于期权，以币(如BTC, ETH)为单位
	PxUsd             string             `json:"pxUsd"`             // 期权价格，以USD为单位 仅适用于期权，其他业务线返回空字符串""
	PxVol             string             `json:"pxVol"`             // 期权订单的隐含波动率 仅适用于期权，其他业务线返回空字符串""
	PxType            string             `json:"pxType"`            // 期权的价格类型 px：代表按价格下单，单位为币 (请求参数 px 的数值单位是BTC或ETH) pxVol：代表按pxVol下单 pxUsd：代表按照pxUsd下单，单位为USD (请求参数px 的数值单位是USD)
	Sz                string             `json:"sz"`                // 委托数量
	Pnl               string             `json:"pnl"`               // 收益(不包括手续费) 适用于有成交的平仓订单，其他情况均为0
	OrdType           string             `json:"ordType"`           // 订单类型
	Side              string             `json:"side"`              // 订单方向
	PosSide           string             `json:"posSide"`           // 持仓方向
	TdMode            string             `json:"tdMode"`            // 交易模式
	AccFillSz         string             `json:"accFillSz"`         // 累计成交数量
	FillPx            string             `json:"fillPx"`            // 最新成交价格。如果还没成交，系统返回""。
	TradeId           string             `json:"tradeId"`           // 最新成交ID
	FillSz            string             `json:"fillSz"`            // 最新成交数量
	FillTime          string             `json:"fillTime"`          // 最新成交时间
	AvgPx             string             `json:"avgPx"`             // 成交均价。如果还没成交，系统返回0。
	State             string             `json:"state"`             // 订单状态
	Lever             string             `json:"lever"`             // 杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	AttachAlgoClOrdId string             `json:"attachAlgoClOrdId"` // 下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx       string             `json:"tpTriggerPx"`       // 止盈触发价
	TpTriggerPxType   string             `json:"tpTriggerPxType"`   // 止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格
	SlTriggerPx       string             `json:"slTriggerPx"`       // 止损触发价
	SlTriggerPxType   string             `json:"slTriggerPxType"`   // 止损触发价类型 last：最新价格 index：指数价格 mark：标记价格
	SlOrdPx           string             `json:"slOrdPx"`           // 止损委托价
	TpOrdPx           string             `json:"tpOrdPx"`           // 止盈委托价
	AttachAlgoOrds    []*AttachAlgoOrder `json:"attachAlgoOrds"`    // 下单附带止盈止损信息
	LinkedAlgoOrd     *struct {
		AlgoId string `json:"algoId"` // 策略订单唯一标识
	} `json:"linkedAlgoOrd"` // 止损订单信息，仅适用于包含限价止盈单的双向止盈止损订单，触发后生成的普通订单
	StpId              string `json:"stpId"`              // 自成交保护ID 如果自成交保护不适用则返回""（已弃用）
	StpMode            string `json:"stpMode"`            // 自成交保护模式
	FeeCcy             string `json:"feeCcy"`             // 手续费币种 对于币币和杠杆的挂单卖单，表示计价币种；其他情况下，表示收取手续费的币种。
	Fee                string `json:"fee"`                // 手续费金额 对于币币和杠杆（除挂单卖单外）：平台收取的累计手续费，始终为负数。 对于币币和杠杆的挂单卖单、交割、永续和期权：累计手续费和返佣（币币和杠杆挂单卖单始终以计价币种计算）。
	RebateCcy          string `json:"rebateCcy"`          // 返佣币种 对于币币和杠杆的挂单卖单，表示交易币种；其他情况下，表示支付返佣的币种。
	Rebate             string `json:"rebate"`             // 返佣金额，仅适用于币币和杠杆 对于挂单卖单：以交易币种为单位的累计手续费和返佣金额。 其他情况下，表示挂单返佣金额，始终为正数，如无返佣则返回""。
	Source             string `json:"source"`             // 订单来源 6：计划委托策略触发后的生成的普通单 7：止盈止损策略触发后的生成的普通单
	Category           string `json:"category"`           // 订单种类
	ReduceOnly         string `json:"reduceOnly"`         // 是否只减仓，true 或 false
	QuickMgnType       string `json:"quickMgnType"`       // 一键借币类型，仅适用于杠杆逐仓的一键借币模式 manual：手动，auto_borrow：自动借币，auto_repay：自动还币
	AlgoClOrdId        string `json:"algoClOrdId"`        // 客户自定义策略订单ID。策略订单触发，且策略单有algoClOrdId是有值，否则为""
	AlgoId             string `json:"algoId"`             // 策略委托单ID，策略订单触发时有值，否则为""
	IsTpLimit          string `json:"isTpLimit"`          // 是否为限价止盈，true 或 false.
	UTime              string `json:"uTime"`              // 订单状态更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	CTime              string `json:"cTime"`              // 订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	CancelSource       string `json:"cancelSource"`       // 订单取消来源的原因枚举值代码
	CancelSourceReason string `json:"cancelSourceReason"` // 订单取消来源的对应具体原因
	TradeQuoteCcy      string `json:"tradeQuoteCcy"`      // 用于交易的计价币种。
}

type AttachAlgoOrder struct {
	AttachAlgoId         string `json:"attachAlgoId"`         // 附带止盈止损的订单ID，改单时，可用来标识该笔附带止盈止损订单。下止盈止损委托单时，该值不会传给 algoId
	AttachAlgoClOrdId    string `json:"attachAlgoClOrdId"`    // 下单附带止盈止损时，客户自定义的策略订单ID
	TpOrdKind            string `json:"tpOrdKind"`            // 止盈订单类型 condition: 条件单 limit: 限价单
	TpTriggerPx          string `json:"tpTriggerPx"`          // 止盈触发价
	TpTriggerRatio       string `json:"tpTriggerRatio"`       // 止盈触发比例，0.3 代表 30% 仅适用于交割/永续合约
	TpTriggerPxType      string `json:"tpTriggerPxType"`      // 止盈触发价类型 last: 最新价格 index: 指数价格 mark: 标记价格
	TpOrdPx              string `json:"tpOrdPx"`              // 止盈委托价
	SlTriggerPx          string `json:"slTriggerPx"`          // 止损触发价
	SlTriggerRatio       string `json:"slTriggerRatio"`       // 止损触发比例，0.3 代表 30% 仅适用于交割/永续合约
	SlTriggerPxType      string `json:"slTriggerPxType"`      // 止损触发价类型 last: 最新价格 index: 指数价格 mark: 标记价格
	SlOrdPx              string `json:"slOrdPx"`              // 止损委托价
	Sz                   string `json:"sz"`                   // 张数。仅适用于“多笔止盈”的止盈订单
	FailCode             string `json:"failCode"`             // 委托失败的错误码，默认为"" 委托失败时有值，如 51020
	FailReason           string `json:"failReason"`           // 委托失败的原因，默认为"" 委托失败时有值
	AmendPxOnTriggerType string `json:"amendPxOnTriggerType"` // 是否启用开仓价止损，仅适用于分批止盈的止损订单 0：不开启，默认值 1：开启
}

func (s *OpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*Order, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/orders-pending",
		secType:  secTypeSigned,
	}
	if s.instType != "" {
		r.setParam("instType", s.instType)
	}
	if s.instId != "" {
		r.setParam("instId", s.instId)
	}
	if s.instFamily != "" {
		r.setParam("instFamily", s.instFamily)
	}
	if s.ordType != "" {
		r.setParam("ordType", s.ordType)
	}
	if s.state != "" {
		r.setParam("state", s.state)
	}
	if s.after != "" {
		r.setParam("after", s.after)
	}
	if s.before != "" {
		r.setParam("before", s.before)
	}
	if s.limit != "" {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*Order)
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}

type OrderService struct {
	c *Client

	instId  string
	ordId   string
	clOrdId string
}

func (s *OrderService) InstId(instId string) *OrderService {
	s.instId = instId
	return s
}

func (s *OrderService) OrdId(ordId string) *OrderService {
	s.ordId = ordId
	return s
}

func (s *OrderService) ClOrdId(clOrdId string) *OrderService {
	s.clOrdId = clOrdId
	return s
}

func (s *OrderService) Do(ctx context.Context, opts ...RequestOption) ([]*Order, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/order",
		secType:  secTypeSigned,
	}
	if s.instId != "" {
		r.setParam("instId", s.instId)
	}
	if s.ordId != "" {
		r.setParam("ordId", s.ordId)
	}
	if s.clOrdId != "" {
		r.setParam("clOrdId", s.clOrdId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*Order)
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}

type Orders7DHistoryService struct {
	c *Client

	instType   string
	instId     string
	instFamily string
	ordType    string
	state      string
	category   string
	after      string
	before     string
	begin      string
	end        string
	limit      string
}

func (s *Orders7DHistoryService) InstId(instId string) *Orders7DHistoryService {
	s.instId = instId
	return s
}

func (s *Orders7DHistoryService) InstFamily(instFamily string) *Orders7DHistoryService {
	s.instFamily = instFamily
	return s
}

func (s *Orders7DHistoryService) OrdType(ordType string) *Orders7DHistoryService {
	s.ordType = ordType
	return s
}

func (s *Orders7DHistoryService) State(state string) *Orders7DHistoryService {
	s.state = state
	return s
}

func (s *Orders7DHistoryService) Category(category string) *Orders7DHistoryService {
	s.category = category
	return s
}

func (s *Orders7DHistoryService) After(after string) *Orders7DHistoryService {
	s.after = after
	return s
}

func (s *Orders7DHistoryService) Before(before string) *Orders7DHistoryService {
	s.before = before
	return s
}

func (s *Orders7DHistoryService) Begin(begin string) *Orders7DHistoryService {
	s.begin = begin
	return s
}

func (s *Orders7DHistoryService) End(end string) *Orders7DHistoryService {
	s.end = end
	return s
}

func (s *Orders7DHistoryService) Limit(limit string) *Orders7DHistoryService {
	s.limit = limit
	return s
}

func (s *Orders7DHistoryService) InstType(instType string) *Orders7DHistoryService {
	s.instType = instType
	return s
}

func (s *Orders7DHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]*Order, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/orders-history",
		secType:  secTypeSigned,
	}
	if s.instType != "" {
		r.setParam("instType", s.instType)
	}
	if s.instId != "" {
		r.setParam("instId", s.instId)
	}
	if s.instFamily != "" {
		r.setParam("instFamily", s.instFamily)
	}
	if s.ordType != "" {
		r.setParam("ordType", s.ordType)
	}
	if s.state != "" {
		r.setParam("state", s.state)
	}
	if s.category != "" {
		r.setParam("category", s.category)
	}
	if s.after != "" {
		r.setParam("after", s.after)
	}
	if s.before != "" {
		r.setParam("before", s.before)
	}
	if s.begin != "" {
		r.setParam("begin", s.begin)
	}
	if s.end != "" {
		r.setParam("end", s.end)
	}
	if s.limit != "" {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*Order)
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}

type OrdersHistory3MService struct {
	c *Client

	instType   string
	instId     string
	instFamily string
	ordType    string
	state      string
	category   string
	after      string
	before     string
	begin      string
	end        string
	limit      string
}

func (s *OrdersHistory3MService) InstId(instId string) *OrdersHistory3MService {
	s.instId = instId
	return s
}

func (s *OrdersHistory3MService) InstFamily(instFamily string) *OrdersHistory3MService {
	s.instFamily = instFamily
	return s
}

func (s *OrdersHistory3MService) OrdType(ordType string) *OrdersHistory3MService {
	s.ordType = ordType
	return s
}

func (s *OrdersHistory3MService) State(state string) *OrdersHistory3MService {
	s.state = state
	return s
}

func (s *OrdersHistory3MService) Category(category string) *OrdersHistory3MService {
	s.category = category
	return s
}

func (s *OrdersHistory3MService) After(after string) *OrdersHistory3MService {
	s.after = after
	return s
}

func (s *OrdersHistory3MService) Before(before string) *OrdersHistory3MService {
	s.before = before
	return s
}

func (s *OrdersHistory3MService) Begin(begin string) *OrdersHistory3MService {
	s.begin = begin
	return s
}

func (s *OrdersHistory3MService) End(end string) *OrdersHistory3MService {
	s.end = end
	return s
}

func (s *OrdersHistory3MService) Limit(limit string) *OrdersHistory3MService {
	s.limit = limit
	return s
}

func (s *OrdersHistory3MService) InstType(instType string) *OrdersHistory3MService {
	s.instType = instType
	return s
}

func (s *OrdersHistory3MService) Do(ctx context.Context, opts ...RequestOption) ([]*Order, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/orders-history-archive",
		secType:  secTypeSigned,
	}
	if s.instType != "" {
		r.setParam("instType", s.instType)
	}
	if s.instId != "" {
		r.setParam("instId", s.instId)
	}
	if s.instFamily != "" {
		r.setParam("instFamily", s.instFamily)
	}
	if s.ordType != "" {
		r.setParam("ordType", s.ordType)
	}
	if s.state != "" {
		r.setParam("state", s.state)
	}
	if s.category != "" {
		r.setParam("category", s.category)
	}
	if s.after != "" {
		r.setParam("after", s.after)
	}
	if s.before != "" {
		r.setParam("before", s.before)
	}
	if s.begin != "" {
		r.setParam("begin", s.begin)
	}
	if s.end != "" {
		r.setParam("end", s.end)
	}
	if s.limit != "" {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*Order)
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}
