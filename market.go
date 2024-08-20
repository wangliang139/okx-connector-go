package okx_connector

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

type SystemStatusResponse struct {
}

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
	err = json.Unmarshal(data, res)
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
		secType:  secTypeNone,
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
	err = json.Unmarshal(data, symbols)
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
	err = json.Unmarshal(data, &array)
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
	err = json.Unmarshal(data, &array)
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
	if err := json.Unmarshal(data, &quotations); err != nil {
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
func (s *MarketDepthService) Do(ctx context.Context, opts ...RequestOption) (res []*Depth, err error) {
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
	if err := json.Unmarshal(data, depth); err != nil {
		return nil, err
	}
	return *depth, nil
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
func (s *MarketDepthFullService) Do(ctx context.Context, opts ...RequestOption) (res []*Depth, err error) {
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
	if err := json.Unmarshal(data, depth); err != nil {
		return nil, err
	}
	return *depth, nil
}
