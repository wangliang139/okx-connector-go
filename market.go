package okx_connector

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/bytedance/sonic"
)

// Test Connectivity endpoint (GET /api/v5/public/time)
type PingService struct {
	c *Client
}

// Send the request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/time",
		secType:  secTypeNone,
	}
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Test Connectivity endpoint (GET /api/v5/public/time)
type SystemTimeService struct {
	c *Client
}

// Send the request
func (s *SystemTimeService) Do(ctx context.Context, opts ...RequestOption) (time *int64, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/time",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	t, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

type SystemStatusService struct {
	c *Client
}

type SystemStatusResponse struct{}

// Send the request
func (s *SystemStatusService) Do(ctx context.Context, opts ...RequestOption) (res *SystemStatusResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/system/status",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SystemStatusResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SymbolInfoService struct {
	c          *Client
	instType   string
	uly        *string
	instFamily *string
	instId     *string
}

func (s *SymbolInfoService) InstType(instType string) *SymbolInfoService {
	s.instType = instType
	return s
}

func (s *SymbolInfoService) Uly(uly string) *SymbolInfoService {
	s.uly = &uly
	return s
}

func (s *SymbolInfoService) InstFamily(instFamily string) *SymbolInfoService {
	s.instFamily = &instFamily
	return s
}

func (s *SymbolInfoService) InstId(instId string) *SymbolInfoService {
	s.instId = &instId
	return s
}

type SymbolInfo struct {
	InstType     string `json:"instType"`
	InstId       string `json:"instId"`
	Uly          string `json:"uly"`
	InstFamily   string `json:"instFamily"`
	BaseCcy      string `json:"baseCcy"`
	QuoteCcy     string `json:"quoteCcy"`
	SettleCcy    string `json:"settleCcy"`
	CtVal        string `json:"ctVal"`
	CtMult       string `json:"ctMult"`
	CtValCcy     string `json:"ctValCcy"`
	OptType      string `json:"optType"`
	Stk          string `json:"stk"`
	ListTime     string `json:"listTime"`
	ExpTime      string `json:"expTime"`
	Lever        string `json:"lever"`
	TickSz       string `json:"tickSz"`
	LotSz        string `json:"lotSz"`
	MinSz        string `json:"minSz"`
	CtType       string `json:"ctType"`
	Linear       string `json:"linear"`
	Inverse      string `json:"inverse"`
	State        string `json:"state"`
	Live         string `json:"live"`
	Suspend      string `json:"suspend"`
	Preopen      string `json:"preopen"`
	Test         string `json:"test"`
	RuleType     string `json:"ruleType"`
	Normal       string `json:"normal"`
	PreMarket    string `json:"pre_market"`
	PosLmtAmt    string `json:"posLmtAmt"`
	PosLmtPct    string `json:"posLmtPct"`
	MaxLmtSz     string `json:"maxLmtSz"`
	MaxMktSz     string `json:"maxMktSz"`
	MaxLmtAmt    string `json:"maxLmtAmt"`
	MaxMktAmt    string `json:"maxMktAmt"`
	MaxTwapSz    string `json:"maxTwapSz"`
	MaxIcebergSz string `json:"maxIcebergSz"`
	MaxTriggerSz string `json:"maxTriggerSz"`
	MaxStopSz    string `json:"maxStopSz"`
}

// Send the request
func (s *SymbolInfoService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/instruments",
		secType:  secTypeSigned,
	}
	r.setParam("instType", s.instType)
	if s.uly != nil {
		r.setParam("uly", *s.uly)
	}
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	symbols := new([]*SymbolInfo)
	err = sonic.Unmarshal(data, symbols)
	if err != nil {
		return nil, err
	}
	return *symbols, nil
}

type MarketKlinesService struct {
	c      *Client
	instId string
	bar    *string
	after  *int64
	before *int64
	limit  *int
}

func (s *MarketKlinesService) InstId(instId string) *MarketKlinesService {
	s.instId = instId
	return s
}

func (s *MarketKlinesService) Bar(bar string) *MarketKlinesService {
	s.bar = &bar
	return s
}

func (s *MarketKlinesService) After(after int64) *MarketKlinesService {
	s.after = &after
	return s
}

func (s *MarketKlinesService) Before(before int64) *MarketKlinesService {
	s.before = &before
	return s
}

func (s *MarketKlinesService) Limit(limit int) *MarketKlinesService {
	s.limit = &limit
	return s
}

type Kline struct {
	Ts          string
	Open        string
	High        string
	Low         string
	Close       string
	Vol         string
	VolCcy      string
	VolCcyQuote string
	Confirm     bool
}

// Send the request
func (s *MarketKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/candles",
		secType:  secTypeNone,
	}
	r.setParam("instId", s.instId)
	if s.bar != nil {
		r.setParam("bar", *s.bar)
	}
	if s.after != nil {
		r.setParam("after", *s.after)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var array [][]string
	err = sonic.Unmarshal(data, &array)
	if err != nil {
		return nil, err
	}
	candles := make([]*Kline, 0)
	for _, item := range array {
		candle := &Kline{
			Ts:          item[0],
			Open:        item[1],
			High:        item[2],
			Low:         item[3],
			Close:       item[4],
			Vol:         item[5],
			VolCcy:      item[6],
			VolCcyQuote: item[7],
			Confirm:     item[8] == "1",
		}
		candles = append(candles, candle)
	}
	return candles, nil
}

type MarketKlinesHisService struct {
	c      *Client
	instId string
	bar    *string
	after  *int64
	before *int64
	limit  *int
}

func (s *MarketKlinesHisService) InstId(instId string) *MarketKlinesHisService {
	s.instId = instId
	return s
}

func (s *MarketKlinesHisService) Bar(bar string) *MarketKlinesHisService {
	s.bar = &bar
	return s
}

func (s *MarketKlinesHisService) After(after int64) *MarketKlinesHisService {
	s.after = &after
	return s
}

func (s *MarketKlinesHisService) Before(before int64) *MarketKlinesHisService {
	s.before = &before
	return s
}

func (s *MarketKlinesHisService) Limit(limit int) *MarketKlinesHisService {
	s.limit = &limit
	return s
}

// Send the request
func (s *MarketKlinesHisService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/history-candles",
		secType:  secTypeNone,
	}
	r.setParam("instId", s.instId)
	if s.bar != nil {
		r.setParam("bar", *s.bar)
	}
	if s.after != nil {
		r.setParam("after", *s.after)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var array [][]string
	err = sonic.Unmarshal(data, &array)
	if err != nil {
		return nil, err
	}
	candles := make([]*Kline, 0)
	for _, item := range array {
		candle := &Kline{
			Ts:          item[0],
			Open:        item[1],
			High:        item[2],
			Low:         item[3],
			Close:       item[4],
			Vol:         item[5],
			VolCcy:      item[6],
			VolCcyQuote: item[7],
			Confirm:     item[8] == "1",
		}
		candles = append(candles, candle)
	}
	return candles, nil
}

type SymbolQuotationService struct {
	c      *Client
	instId *string
}

func (s *SymbolQuotationService) InstId(instId string) *SymbolQuotationService {
	s.instId = &instId
	return s
}

type Quotation struct {
	Ts        string `json:"ts"`
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
}

// Send the request
func (s *SymbolQuotationService) Do(ctx context.Context, opts ...RequestOption) (res *Quotation, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/ticker",
		secType:  secTypeNone,
	}
	if s.instId == nil {
		return nil, errors.New("instId is required")
	}
	r.setParam("instId", *s.instId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var quotations []*Quotation
	if err := sonic.Unmarshal(data, &quotations); err != nil {
		return nil, err
	}
	if len(quotations) == 0 {
		return nil, nil
	}
	return quotations[0], nil
}

type MarketDepthService struct {
	c      *Client
	instId *string
	size   *int
}

func (s *MarketDepthService) InstId(instId string) *MarketDepthService {
	s.instId = &instId
	return s
}

func (s *MarketDepthService) Size(size int) *MarketDepthService {
	s.size = &size
	return s
}

type Depth struct {
	Ts   string     `json:"ts"`
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

// Send the request
func (s *MarketDepthService) Do(ctx context.Context, opts ...RequestOption) (*Depth, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/books",
		secType:  secTypeNone,
	}
	if s.instId == nil {
		return nil, errors.New("instId is required")
	}
	r.setParam("instId", *s.instId)
	if s.size != nil {
		if *s.size < 1 || *s.size > 400 {
			return nil, errors.New("size must be between 1 and 400")
		}
		r.setParam("sz", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	depth := new([]*Depth)
	if err := sonic.Unmarshal(data, depth); err != nil {
		return nil, err
	}
	if len(*depth) > 0 {
		return (*depth)[0], nil
	}
	return nil, nil
}

type MarketDepthFullService struct {
	c      *Client
	instId *string
	size   *int
}

func (s *MarketDepthFullService) InstId(instId string) *MarketDepthFullService {
	s.instId = &instId
	return s
}

func (s *MarketDepthFullService) Size(size int) *MarketDepthFullService {
	s.size = &size
	return s
}

// Send the request
func (s *MarketDepthFullService) Do(ctx context.Context, opts ...RequestOption) (res *Depth, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/books-full",
		secType:  secTypeNone,
	}
	if s.instId == nil {
		return nil, errors.New("instId is required")
	}
	r.setParam("instId", *s.instId)
	if s.size != nil {
		if *s.size < 1 || *s.size > 5000 {
			return nil, errors.New("size must be between 1 and 5000")
		}
		r.setParam("sz", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	depth := new([]*Depth)
	if err := sonic.Unmarshal(data, depth); err != nil {
		return nil, err
	}
	if len(*depth) > 0 {
		return (*depth)[0], nil
	}
	return nil, nil
}

type AnnouncementTypeService struct {
	c *Client
}

type AnnouncementType struct {
	AnnType     string `json:"annType"`
	AnnTypeDesc string `json:"annTypeDesc"`
}

// Send the request
func (s *AnnouncementTypeService) Do(ctx context.Context, opts ...RequestOption) (res []*AnnouncementType, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/support/announcement-types",
		secType:  secTypeNone,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AnnouncementType)
	err = sonic.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

type AnnouncementService struct {
	c       *Client
	annType *string
	page    *int
}

func (s *AnnouncementService) AnnType(annType string) *AnnouncementService {
	s.annType = &annType
	return s
}

func (s *AnnouncementService) Page(page int) *AnnouncementService {
	s.page = &page
	return s
}

type AnnouncementResponse struct {
	TotalPage string          `json:"totalPage"`
	Details   []*Announcement `json:"details"`
}

type Announcement struct {
	Title   string `json:"title"`
	AnnType string `json:"annType"`
	PTime   string `json:"pTime"`
	Url     string `json:"url"`
}

func (s *AnnouncementService) Do(ctx context.Context, opts ...RequestOption) (res []*AnnouncementResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/support/announcements",
		secType:  secTypeSigned,
	}
	if s.annType != nil {
		r.setParam("annType", *s.annType)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AnnouncementResponse)
	err = sonic.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

// /api/v5/market/trades
type MarketTradesService struct {
	c *Client

	instId *string
	limit  *int
}

func (s *MarketTradesService) InstId(instId string) *MarketTradesService {
	s.instId = &instId
	return s
}

func (s *MarketTradesService) Limit(limit int) *MarketTradesService {
	s.limit = &limit
	return s
}

type Trade struct {
	InstId  string `json:"instId"`
	TradeId string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	Source  string `json:"source"`
	Ts      string `json:"ts"`
}

// Send the request
func (s *MarketTradesService) Do(ctx context.Context, opts ...RequestOption) ([]*Trade, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/trades",
		secType:  secTypeNone,
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*Trade)
	err = sonic.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

// /api/v5/market/tickers
type MarketTickersService struct {
	c *Client

	instType   string
	instFamily *string
}

func (s *MarketTickersService) InstType(instType string) *MarketTickersService {
	s.instType = instType
	return s
}

func (s *MarketTickersService) InstFamily(instFamily string) *MarketTickersService {
	s.instFamily = &instFamily
	return s
}

type MarketTicker struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	Last      string `json:"last"`      // 最新成交价
	LastSz    string `json:"lastSz"`    // 最新成交的数量，0 代表没有成交量
	AskPx     string `json:"askPx"`     // 卖一价
	AskSz     string `json:"askSz"`     // 卖一价的挂单数数量
	BidPx     string `json:"bidPx"`     // 买一价
	BidSz     string `json:"bidSz"`     // 买一价的挂单数量
	Open24h   string `json:"open24h"`   // 24小时开盘价
	High24h   string `json:"high24h"`   // 24小时最高价
	Low24h    string `json:"low24h"`    // 24小时最低价
	VolCcy24h string `json:"volCcy24h"` // 24小时成交量，以币为单位
	Vol24h    string `json:"vol24h"`    // 24小时成交量，以张为单位
	SodUtc0   string `json:"sodUtc0"`   // UTC 0 时开盘价
	SodUtc8   string `json:"sodUtc8"`   // UTC+8 时开盘价
	Ts        string `json:"ts"`        // ticker数据产生时间，Unix时间戳的毫秒数格式，如 1597026383085
}

// Send the request
func (s *MarketTickersService) Do(ctx context.Context, opts ...RequestOption) ([]*MarketTicker, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/tickers",
		secType:  secTypeNone,
	}
	if s.instType == "" {
		return nil, errors.New("instType is required")
	}
	r.setParam("instType", s.instType)
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*MarketTicker)
	err = sonic.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

// /api/v5/public/position-tiers
type PositionTiersService struct {
	c *Client

	tdMode     string  // 保证金模式
	instType   string  // 产品类型
	instFamily *string // 交易品种，支持多instFamily，半角逗号分隔，最大不超过5个
	instId     *string // 产品ID，支持多instId，半角逗号分隔，最大不超过5个
	ccy        *string // 保证金币种
	tier       *string // 查指定档位
}

func (s *PositionTiersService) TdMode(tdMode string) *PositionTiersService {
	s.tdMode = tdMode
	return s
}

func (s *PositionTiersService) InstType(instType string) *PositionTiersService {
	s.instType = instType
	return s
}

func (s *PositionTiersService) InstFamily(instFamily string) *PositionTiersService {
	s.instFamily = &instFamily
	return s
}

func (s *PositionTiersService) InstId(instId string) *PositionTiersService {
	s.instId = &instId
	return s
}

func (s *PositionTiersService) Ccy(ccy string) *PositionTiersService {
	s.ccy = &ccy
	return s
}

func (s *PositionTiersService) Tier(tier string) *PositionTiersService {
	s.tier = &tier
	return s
}

type PositionTier struct {
	Uly          string `json:"uly"`          // 标的指数，适用于交割/永续/期权
	InstFamily   string `json:"instFamily"`   // 交易品种，适用于交割/永续/期权
	InstId       string `json:"instId"`       // 币对
	Tier         string `json:"tier"`         // 仓位档位
	MinSz        string `json:"minSz"`        // 该档位最少借币量或者持仓数量 杠杆/期权/永续/交割 最小持仓量 默认0；当 ccy 参数生效时，返回 ccy 的最小借币量
	MaxSz        string `json:"maxSz"`        // 该档位最多借币量或者持仓数量 杠杆/期权/永续/交割；当 ccy 参数生效时，返回 ccy 的最大借币量
	Mmr          string `json:"mmr"`          // 仓位维持保证金率
	Imr          string `json:"imr"`          // 最低初始维持保证金率
	MaxLever     string `json:"maxLever"`     // 最高可用杠杆倍数
	OptMgnFactor string `json:"optMgnFactor"` // 期权保证金系数 （仅适用于期权）
	QuoteMaxLoan string `json:"quoteMaxLoan"` // 计价货币 最大借币量（仅适用于杠杆，且instId参数生效时），如 BTC-USDT 里的 USDT最大借币量
	BaseMaxLoan  string `json:"baseMaxLoan"`  // 交易货币 最大借币量（仅适用于杠杆，且instId参数生效时），如 BTC-USDT 里的 BTC最大借币量
}

func (s *PositionTiersService) Do(ctx context.Context, opts ...RequestOption) ([]*PositionTier, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/position-tiers",
		secType:  secTypeNone,
	}
	if s.tdMode != "" {
		r.setParam("tdMode", s.tdMode)
	}
	if s.instType != "" {
		r.setParam("instType", s.instType)
	}
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	if s.ccy != nil {
		r.setParam("ccy", *s.ccy)
	}
	if s.tier != nil {
		r.setParam("tier", *s.tier)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*PositionTier)
	err = sonic.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

// /api/v5/public/mark-price
type MarkPriceService struct {
	c *Client

	instType   string
	instFamily *string
	instId     *string
}

func (s *MarkPriceService) InstType(instType string) *MarkPriceService {
	s.instType = instType
	return s
}

func (s *MarkPriceService) InstFamily(instFamily string) *MarkPriceService {
	s.instFamily = &instFamily
	return s
}

func (s *MarkPriceService) InstId(instId string) *MarkPriceService {
	s.instId = &instId
	return s
}

type MarkPrice struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	MarkPrice string `json:"markPrice"`
	Ts        string `json:"ts"`
}

func (s *MarkPriceService) Do(ctx context.Context, opts ...RequestOption) ([]*MarkPrice, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/mark-price",
		secType:  secTypeNone,
	}
	if s.instType != "" {
		r.setParam("instType", s.instType)
	}
	if s.instFamily != nil {
		r.setParam("instFamily", *s.instFamily)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*MarkPrice)
	err = sonic.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}
