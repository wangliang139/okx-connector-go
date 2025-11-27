package okx_connector

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type AccountBalanceService struct {
	c      *Client
	assets []string
}

type AccountBalance struct {
	AdjEq                 string          `json:"adjEq"`
	AvailEq               string          `json:"availEq"`
	BorrowFroz            string          `json:"borrowFroz"`
	Delta                 string          `json:"delta"`
	DeltaLever            string          `json:"deltaLever"`
	DeltaNeutralStatus    string          `json:"deltaNeutralStatus"`
	Imr                   string          `json:"imr"`
	IsoEq                 string          `json:"isoEq"`
	MgnRatio              string          `json:"mgnRatio"`
	Mmr                   string          `json:"mmr"`
	NotionalUsd           string          `json:"notionalUsd"`
	NotionalUsdForBorrow  string          `json:"notionalUsdForBorrow"`
	NotionalUsdForFutures string          `json:"notionalUsdForFutures"`
	NotionalUsdForOption  string          `json:"notionalUsdForOption"`
	NotionalUsdForSwap    string          `json:"notionalUsdForSwap"`
	OrdFroz               string          `json:"ordFroz"`
	TotalEq               string          `json:"totalEq"`
	UTime                 string          `json:"uTime"`
	Upl                   string          `json:"upl"`
	Assets                []*AssetBalance `json:"details"`
}

type AssetBalance struct {
	AutoLendStatus        string `json:"autoLendStatus"`
	AutoLendMtAmt         string `json:"autoLendMtAmt"`
	AvailBal              string `json:"availBal"`
	AvailEq               string `json:"availEq"`
	BorrowFroz            string `json:"borrowFroz"`
	CashBal               string `json:"cashBal"`
	Ccy                   string `json:"ccy"`
	CrossLiab             string `json:"crossLiab"`
	ColRes                string `json:"colRes"`
	CollateralEnabled     bool   `json:"collateralEnabled"`
	CollateralRestrict    bool   `json:"collateralRestrict"`
	ColBorrAutoConversion string `json:"colBorrAutoConversion"`
	DisEq                 string `json:"disEq"`
	Eq                    string `json:"eq"`
	EqUsd                 string `json:"eqUsd"`
	SmtSyncEq             string `json:"smtSyncEq"`
	SpotCopyTradingEq     string `json:"spotCopyTradingEq"`
	FixedBal              string `json:"fixedBal"`
	FrozenBal             string `json:"frozenBal"`
	FrpType               string `json:"frpType"`
	Imr                   string `json:"imr"`
	Interest              string `json:"interest"`
	IsoEq                 string `json:"isoEq"`
	IsoLiab               string `json:"isoLiab"`
	IsoUpl                string `json:"isoUpl"`
	Liab                  string `json:"liab"`
	MaxLoan               string `json:"maxLoan"`
	MgnRatio              string `json:"mgnRatio"`
	Mmr                   string `json:"mmr"`
	NotionalLever         string `json:"notionalLever"`
	NotionalUsd           string `json:"notionalUsd"`
	NotionalUsdForBorrow  string `json:"notionalUsdForBorrow"`
	NotionalUsdForFutures string `json:"notionalUsdForFutures"`
	NotionalUsdForOption  string `json:"notionalUsdForOption"`
	NotionalUsdForSwap    string `json:"notionalUsdForSwap"`
	OrdFroz               string `json:"ordFroz"`
	RewardBal             string `json:"rewardBal"`
	SpotInUseAmt          string `json:"spotInUseAmt"`
	ClSpotInUseAmt        string `json:"clSpotInUseAmt"`
	MaxSpotInUse          string `json:"maxSpotInUse"`
	SpotIsoBal            string `json:"spotIsoBal"`
	StgyEq                string `json:"stgyEq"`
	Twap                  string `json:"twap"`
	UTime                 string `json:"uTime"`
	Upl                   string `json:"upl"`
	UplLiab               string `json:"uplLiab"`
	SpotBal               string `json:"spotBal"`
	OpenAvgPx             string `json:"openAvgPx"`
	AccAvgPx              string `json:"accAvgPx"`
	SpotUpl               string `json:"spotUpl"`
	SpotUplRatio          string `json:"spotUplRatio"`
	TotalPnl              string `json:"totalPnl"`
	TotalPnlRatio         string `json:"totalPnlRatio"`
}

func (s *AccountBalanceService) Assets(assets []string) *AccountBalanceService {
	s.assets = assets
	return s
}

// Send the request
func (s *AccountBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*AccountBalance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/balance",
		secType:  secTypeSigned,
	}
	if len(s.assets) > 0 {
		r.setParam("ccy", strings.Join(s.assets, ","))
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AccountBalance)
	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}
