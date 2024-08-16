package okx_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Test Connectivity endpoint (GET /api/v5/system/status)
type PingService struct {
	c *Client
}

// Send the request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/system/status",
		secType:  secTypeNone,
	}
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
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

type MarketCandlesService struct {
	c      *Client
	instId string
	bar    *int64
	before *int64
	limit  *int
}

func (s *MarketCandlesService) InstId(instId string) *MarketCandlesService {
	s.instId = instId
	return s
}

func (s *MarketCandlesService) Bar(bar int64) *MarketCandlesService {
	s.bar = &bar
	return s
}

func (s *MarketCandlesService) Before(before int64) *MarketCandlesService {
	s.before = &before
	return s
}

func (s *MarketCandlesService) Limit(limit int) *MarketCandlesService {
	s.limit = &limit
	return s
}

type MarketCandle struct {
	ts          string
	open        string
	high        string
	low         string
	close       string
	vol         string
	volCcy      string
	volCcyQuote string
	confirm     bool
}

// Send the request
func (s *MarketCandlesService) Do(ctx context.Context, opts ...RequestOption) (res []*MarketCandle, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/candles",
		secType:  secTypeNone,
	}
	r.setParam("instId", s.instId)
	if s.bar != nil {
		r.setParam("bar", *s.bar)
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
	candles := make([]*MarketCandle, 0)
	for _, item := range array {
		candle := &MarketCandle{
			ts:          item[0],
			open:        item[1],
			high:        item[2],
			low:         item[3],
			close:       item[4],
			vol:         item[5],
			volCcy:      item[6],
			volCcyQuote: item[7],
			confirm:     item[8] == "1",
		}
		candles = append(candles, candle)
	}
	return candles, nil
}
