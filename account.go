package okx_connector

import (
	"context"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
)

type AccountConfigService struct {
	c *Client
}

type AccountConfig struct {
	Uid                 string   `json:"uid"`
	MainUid             string   `json:"mainUid"`             // 当前请求的母账户ID，如果 uid = mainUid，代表当前账号为母账户；如果 uid != mainUid，代表当前账户为子账户。
	AcctLv              string   `json:"acctLv"`              // 账户模式 1：现货模式，2：合约模式，3：跨币种保证金模式，4：组合保证金模式
	AcctStpMode         string   `json:"acctStpMode"`         // 账户自成交保护模式，cancel_maker：撤销挂单，cancel_taker：撤销吃单，cancel_both：撤销挂单和吃单，默认为cancel_maker，用户可通过母账户登录网页修改该配置
	PosMode             string   `json:"posMode"`             // 持仓方式，long_short_mode：开平仓模式，net_mode：买卖模式，仅适用交割/永续
	AutoLoan            bool     `json:"autoLoan"`            // 是否自动借币，true：自动借币，false：非自动借币
	GreeksType          string   `json:"greeksType"`          // 当前希腊字母展示方式，PA：币本位，BS：美元本位
	FeeType             string   `json:"feeType"`             // 手续费类型，0：手续费以获取币种收取，1：手续费以计价币种收取
	Level               string   `json:"level"`               // 当前在平台上真实交易量的用户等级，如 Lv1，代表普通用户等级。
	LevelTmp            string   `json:"levelTmp"`            // 特约用户的临时体验用户等级，如 Lv1
	CtIsoMode           string   `json:"ctIsoMode"`           // 衍生品的逐仓保证金划转模式，automatic：开仓划转，autonomy：自主划转
	StgyType            string   `json:"stgyType"`            // 策略类型，0：普通策略模式，1：delta 中性策略模式
	RoleType            string   `json:"roleType"`            // 用户角色，0：普通用户，1：带单者，2：跟单者
	TraderInsts         []string `json:"traderInsts"`         // 当前账号已经设置的带单合约，仅适用于带单者
	SpotRoleType        string   `json:"spotRoleType"`        // 现货跟单角色。0：普通用户；1：带单者；2：跟单者
	SpotTraderInsts     []string `json:"spotTraderInsts"`     // 当前账号已经设置的带单币对，仅适用于带单者
	OpAuth              string   `json:"opAuth"`              // 是否开通期权交易，0：未开通 1：已经开通
	KycLv               string   `json:"kycLv"`               // 母账户KYC等级 0: 未认证 1: 已完成 level 1 认证 2: 已完成 level 2 认证 3: 已完成 level 3认证 如果请求来自子账户, kycLv 为其母账户的等级 如果请求来自母账户, kycLv 为当前请求的母账户等级
	Label               string   `json:"label"`               // 当前请求API key的备注名，不超过50位字母（区分大小写）或数字，可以是纯字母或纯数字。
	Ip                  string   `json:"ip"`                  // 当前请求API key绑定的ip地址，多个ip用半角逗号隔开，如：117.37.203.58,117.37.203.57。如果没有绑定ip，会返回空字符串""
	Perm                string   `json:"perm"`                // 当前请求的 API key 或 Access token 的权限，read_only：读取 trade：交易 withdraw：提币
	LiquidationGear     string   `json:"liquidationGear"`     // 强平提醒的维持保证金率水平 3 和 -1 代表维持保证金率达到 300% 时，每隔 1 小时 app 和 ”爆仓风险预警推送频道“会推送通知。-1 是初始值，与-3有着同样效果 0 代表不提醒
	EnableSpotBorrow    bool     `json:"enableSpotBorrow"`    // 现货模式下是否支持借币，true：支持，false：不支持
	SpotBorrowAutoRepay bool     `json:"spotBorrowAutoRepay"` // 现货模式下是否支持自动还币，true：支持，false：不支持
	Type                string   `json:"type"`                // 账户类型 0：母账户 1：普通子账户 2：资管子账户 5：托管交易子账户 - Copper 9：资管交易子账户 - Copper 12：托管交易子账户 - Komainu
	SettleCcy           string   `json:"settleCcy"`           // 当前账户的 USD 本位合约结算币种
	SettleCcyList       []string `json:"settleCcyList"`       // 当前账户的 USD 本位合约结算币种列表，如 ["USD", "USDC", "USDG"]。
}

func (s *AccountConfigService) Do(ctx context.Context, opts ...RequestOption) (*AccountConfig, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/config",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AccountConfig)
	if err := sonic.Unmarshal(data, result); err != nil {
		return nil, err
	}
	if len(*result) > 0 {
		return (*result)[0], nil
	}
	return nil, nil
}

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
	Details               []*AssetBalance `json:"details"`
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
func (s *AccountBalanceService) Do(ctx context.Context, opts ...RequestOption) (*AccountBalance, error) {
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
	if err := sonic.Unmarshal(data, result); err != nil {
		return nil, err
	}
	if len(*result) > 0 {
		return (*result)[0], nil
	}
	return nil, nil
}

// /api/v5/account/leverage-info
type AccountLeverageInfoService struct {
	c *Client

	instId  string
	ccy     string
	mgnMode string
}

func (s *AccountLeverageInfoService) InstId(instId string) *AccountLeverageInfoService {
	s.instId = instId
	return s
}

func (s *AccountLeverageInfoService) Ccy(ccy string) *AccountLeverageInfoService {
	s.ccy = ccy
	return s
}

func (s *AccountLeverageInfoService) MgnMode(mgnMode string) *AccountLeverageInfoService {
	s.mgnMode = mgnMode
	return s
}

type AccountLeverageInfo struct {
	InstId  string `json:"instId"`
	Ccy     string `json:"ccy"`
	MgnMode string `json:"mgnMode"`
	PosSide string `json:"posSide"`
	Lever   string `json:"lever"`
}

func (s *AccountLeverageInfoService) Do(ctx context.Context, opts ...RequestOption) ([]*AccountLeverageInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/leverage-info",
		secType:  secTypeSigned,
	}
	if s.instId != "" {
		r.setParam("instId", s.instId)
	}
	if s.ccy != "" {
		r.setParam("ccy", s.ccy)
	}
	if s.mgnMode != "" {
		r.setParam("mgnMode", s.mgnMode)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AccountLeverageInfo)
	if err := sonic.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}

// /api/v5/account/trade-fee
type AccountTradeFeeService struct {
	c *Client

	instType   string
	instId     string
	instFamily string
	ruleType   string
}

func (s *AccountTradeFeeService) InstId(instId string) *AccountTradeFeeService {
	s.instId = instId
	return s
}

func (s *AccountTradeFeeService) InstType(instType string) *AccountTradeFeeService {
	s.instType = instType
	return s
}

func (s *AccountTradeFeeService) InstFamily(instFamily string) *AccountTradeFeeService {
	s.instFamily = instFamily
	return s
}

func (s *AccountTradeFeeService) RuleType(ruleType string) *AccountTradeFeeService {
	s.ruleType = ruleType
	return s
}

type AccountTradeFee struct {
	Level     string     `json:"level"`     // 手续费等级
	Taker     string     `json:"taker"`     // 对于币币/杠杆，为 USDT 交易区的吃单手续费率；对于永续，交割和期权合约，为币本位合约费率
	Maker     string     `json:"maker"`     // 对于币币/杠杆，为 USDT 交易区的挂单手续费率；对于永续，交割和期权合约，为币本位合约费率
	TakerU    string     `json:"takerU"`    // USDT 合约吃单手续费率，仅适用于交割/永续
	MakerU    string     `json:"makerU"`    // USDT 合约挂单手续费率，仅适用于交割/永续
	Delivery  string     `json:"delivery"`  // 交割手续费率
	Exercise  string     `json:"exercise"`  // 行权手续费率
	InstType  string     `json:"instType"`  // 产品类型
	TakerUSDC string     `json:"takerUSDC"` // 对于币币/杠杆，为 USDⓈ&Crypto 交易区的吃单手续费率；对于永续和交割合约，为 USDC 合约费率
	MakerUSDC string     `json:"makerUSDC"` // 对于币币/杠杆，为 USDⓈ&Crypto 交易区的挂单手续费率；对于永续和交割合约，为 USDC 合约费率
	RuleType  string     `json:"ruleType"`  // 交易规则类型 normal：普通交易 pre_market：盘前交易
	Ts        string     `json:"ts"`        // 数据返回时间，Unix时间戳的毫秒数格式，如 1597026383085
	Fiat      []*FiatFee `json:"fiat"`      // 法币费率
}

type FiatFee struct {
	Ccy   string `json:"ccy"`   // 币种
	Taker string `json:"taker"` // 吃单手续费率
	Maker string `json:"maker"` // 挂单手续费率
}

func (s *AccountTradeFeeService) Do(ctx context.Context, opts ...RequestOption) ([]*AccountTradeFee, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/trade-fee",
		secType:  secTypeSigned,
	}
	if s.instId != "" {
		r.setParam("instId", s.instId)
	}
	if s.instType != "" {
		r.setParam("instType", s.instType)
	}
	if s.instFamily != "" {
		r.setParam("instFamily", s.instFamily)
	}
	if s.ruleType != "" {
		r.setParam("ruleType", s.ruleType)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AccountTradeFee)
	if err := sonic.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}
