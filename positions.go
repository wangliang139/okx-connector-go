package okx_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

type PositionsService struct {
	c *Client

	instType string
	instId   string
	posId    string
}

func (s *PositionsService) InstType(instType string) *PositionsService {
	s.instType = instType
	return s
}

func (s *PositionsService) InstId(instId string) *PositionsService {
	s.instId = instId
	return s
}

func (s *PositionsService) PosId(posId string) *PositionsService {
	s.posId = posId
	return s
}

type Position struct {
	InstType               string            `json:"instType"`               // 产品类型
	MgnMode                string            `json:"mgnMode"`                // 保证金模式 cross：全仓 isolated：逐仓
	PosId                  string            `json:"posId"`                  // 持仓ID
	PosSide                string            `json:"posSide"`                // 持仓方向 long：开平仓模式开多，pos为正 short：开平仓模式开空，pos为正 net：买卖模式（交割/永续/期权：pos为正代表开多，pos为负代表开空。币币杠杆时，pos均为正，posCcy为交易货币时，代表开多；posCcy为计价货币时，代表开空。）
	Pos                    string            `json:"pos"`                    // 持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为0的仓位
	HedgedPos              string            `json:"hedgedPos"`              // 对冲持仓数量 仅在delta 中性策略模式的账户返回stgyType:1，对普通策略模式的账户返回""
	PosCcy                 string            `json:"posCcy"`                 // 仓位资产币种，仅适用于币币杠杆仓位
	AvailPos               string            `json:"availPos"`               // 可平仓数量，适用于 币币杠杆，期权 对于杠杆仓位，平仓时，杠杆还清负债后，余下的部分会视为币币交易，如果想要减少币币交易的数量，可通过"获取最大可用数量"接口获取只减仓的可用数量。
	AvgPx                  string            `json:"avgPx"`                  // 开仓均价 会随结算周期变化，特别是在交割合约全仓模式下，结算时开仓均价会更新为结算价格，同时新增头寸也会改变开仓均价。
	NonSettleAvgPx         string            `json:"nonSettleAvgPx"`         // 未结算均价 不受结算影响的加权开仓价格，仅在新增头寸时更新，和开仓均价的主要区别在于是否受到结算影响。仅适用于全仓交割
	Upl                    string            `json:"upl"`                    // 未实现收益（以标记价格计算）
	UplRatio               string            `json:"uplRatio"`               // 未实现收益率（以标记价格计算
	UplLastPx              string            `json:"uplLastPx"`              // 以最新成交价格计算的未实现收益，主要做展示使用，实际值还是 upl
	UplRatioLastPx         string            `json:"uplRatioLastPx"`         // 以最新成交价格计算的未实现收益率
	InstId                 string            `json:"instId"`                 // 产品ID，如 BTC-USDT-SWAP
	Lever                  string            `json:"lever"`                  // 杠杆倍数，不适用于期权以及组合保证金模式下的全仓仓位
	LiqPx                  string            `json:"liqPx"`                  // 预估强平价 不适用于期权
	MarkPx                 string            `json:"markPx"`                 // 最新标记价格
	Imr                    string            `json:"imr"`                    // 初始保证金，仅适用于全仓
	Margin                 string            `json:"margin"`                 // 保证金余额，可增减，仅适用于逐仓
	MgnRatio               string            `json:"mgnRatio"`               // 维持保证金率
	Mmr                    string            `json:"mmr"`                    // 维持保证金
	Liab                   string            `json:"liab"`                   // 负债额，仅适用于币币杠杆
	LiabCcy                string            `json:"liabCcy"`                // 负债币种，仅适用于币币杠杆
	Interest               string            `json:"interest"`               // 利息，已经生成的未扣利息
	TradeId                string            `json:"tradeId"`                // 最新成交ID
	OptVal                 string            `json:"optVal"`                 // 期权市值，仅适用于期权
	PendingCloseOrdLiabVal string            `json:"pendingCloseOrdLiabVal"` // 逐仓杠杆负债对应平仓挂单的数量
	NotionalUsd            string            `json:"notionalUsd"`            // 以美金价值为单位的持仓数量
	Adl                    string            `json:"adl"`                    // 自动减仓信号区 分为6档，从0到5，数字越小代表adl强度越弱 仅适用于交割/永续/期权
	Ccy                    string            `json:"ccy"`                    // 占用保证金的币种
	Last                   string            `json:"last"`                   // 最新成交价
	IdxPx                  string            `json:"idxPx"`                  // 最新指数价格
	UsdPx                  string            `json:"usdPx"`                  // 保证金币种的市场最新美金价格 仅适用于交割/永续/期权
	BePx                   string            `json:"bePx"`                   // 盈亏平衡价
	DeltaBS                string            `json:"deltaBS"`                // 美金本位持仓仓位delta，仅适用于期权
	DeltaPA                string            `json:"deltaPA"`                // 币本位持仓仓位delta，仅适用于期权
	GammaBS                string            `json:"gammaBS"`                // 美金本位持仓仓位gamma，仅适用于期权
	GammaPA                string            `json:"gammaPA"`                // 币本位持仓仓位gamma，仅适用于期权
	ThetaBS                string            `json:"thetaBS"`                // 美金本位持仓仓位theta，仅适用于期权
	ThetaPA                string            `json:"thetaPA"`                // 币本位持仓仓位theta，仅适用于期权
	VegaBS                 string            `json:"vegaBS"`                 // 美金本位持仓仓位vega，仅适用于期权
	VegaPA                 string            `json:"vegaPA"`                 // 币本位持仓仓位vega，仅适用于期权
	SpotInUseAmt           string            `json:"spotInUseAmt"`           // 现货对冲占用数量 适用于组合保证金模式
	SpotInUseCcy           string            `json:"spotInUseCcy"`           // 现货对冲占用币种，如 BTC 适用于组合保证金模式
	ClSpotInUseAmt         string            `json:"clSpotInUseAmt"`         // 用户自定义现货占用数量 适用于组合保证金模式
	MaxSpotInUseAmt        string            `json:"maxSpotInUseAmt"`        // 系统计算得到的最大可能现货占用数量 适用于组合保证金模式
	RealizedPnl            string            `json:"realizedPnl"`            // 已实现收益 仅适用于交割/永续/期权 realizedPnl=pnl+fee+fundingFee+liqPenalty+settledPnl
	SettledPnl             string            `json:"settledPnl"`             // 已结算收益 仅适用于全仓交割
	Pnl                    string            `json:"pnl"`                    // 平仓订单累计收益额(不包括手续费)
	Fee                    string            `json:"fee"`                    // 累计手续费金额，正数代表平台返佣 ，负数代表平台扣除
	FundingFee             string            `json:"fundingFee"`             // 累计资金费用
	LiqPenalty             string            `json:"liqPenalty"`             // 累计爆仓罚金，有值时为负数。
	CloseOrderAlgo         []*CloseAlgoOrder `json:"closeOrderAlgo"`         // 平仓策略委托订单。调用策略委托下单，且closeFraction=1 时，该数组才会有值。
	CTime                  string            `json:"cTime"`                  // 持仓创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime                  string            `json:"uTime"`                  // 最近一次持仓更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	BizRefId               string            `json:"bizRefId"`               // 外部业务id，如 体验券id
	BizRefType             string            `json:"bizRefType"`             // 外部业务类型
}

type CloseAlgoOrder struct {
	AlgoId          string `json:"algoId"`          // 策略委托单ID
	SlTriggerPx     string `json:"slTriggerPx"`     // 止损触发价
	SlTriggerPxType string `json:"slTriggerPxType"` // 止损触发价类型 	last：最新价格 	index：指数价格 	mark：标记价格
	TpTriggerPx     string `json:"tpTriggerPx"`     // 止盈触发价
	TpTriggerPxType string `json:"tpTriggerPxType"` // 止盈触发价类型 	last：最新价格 	index：指数价格 	mark：标记价格
	CloseFraction   string `json:"closeFraction"`   // 策略委托触发时，平仓的百分比。1 代表100%
}

func (s *PositionsService) Do(ctx context.Context, opts ...RequestOption) ([]*Position, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/positions",
		secType:  secTypeSigned,
	}
	if s.instType != "" {
		r.setParam("instType", s.instType)
	}
	if s.instId != "" {
		r.setParam("instId", s.instId)
	}
	if s.posId != "" {
		r.setParam("posId", s.posId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*Position)
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}
