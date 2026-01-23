package okx_connector

import (
	"context"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
)

type FundingAssetCurrenciesService struct {
	c *Client

	ccy []string
}

func (s *FundingAssetCurrenciesService) Ccy(ccy []string) *FundingAssetCurrenciesService {
	s.ccy = ccy
	return s
}

type FundingAssetCurrency struct {
	Ccy                  string `json:"ccy"`                  // 币种名称，如 BTC
	Name                 string `json:"name"`                 // 币种名称，不显示则无对应名称
	LogoLink             string `json:"logoLink"`             // 币种Logo链接
	Chain                string `json:"chain"`                // 币种链信息，有的币种下有多个链，必须要做区分，如USDT下有USDT-ERC20，USDT-TRC20多个链
	CtAddr               string `json:"ctAddr"`               // 合约地址
	CanDep               bool   `json:"canDep"`               // 当前是否可充值 false：不可链上充值 true：可以链上充值
	CanWd                bool   `json:"canWd"`                // 当前是否可提币 false：不可链上提币 true：可以链上提币
	CanInternal          bool   `json:"canInternal"`          // 当前是否可内部转账 false：不可内部转账 true：可以内部转账
	DepEstOpenTime       string `json:"depEstOpenTime"`       // 充值预期开放时间，Unix时间戳的毫秒数格式，如 1597026383085 如果 canDep 为 true，则返回 ""
	WdEstOpenTime        string `json:"wdEstOpenTime"`        // 提币预期开放时间，Unix时间戳的毫秒数格式，如 1597026383085 如果 canWd 为 true，则返回 ""
	MinDep               string `json:"minDep"`               // 币种单笔最小充值量
	MinWd                string `json:"minWd"`                // 币种单笔最小链上提币量
	MinInternal          string `json:"minInternal"`          // 币种单笔最小内部转账量 无单笔最大内部转账量限制，受24小时内提币额度(wdQuota)限制
	MaxWd                string `json:"maxWd"`                // 币种单笔最大链上提币量
	WdTickSz             string `json:"wdTickSz"`             // 提币精度,表示小数点后的位数。提币手续费精度与提币精度保持一致。内部转账提币精度为小数点后8位。
	WdQuota              string `json:"wdQuota"`              // 过去24小时内提币额度（包含链上提币和内部转账），单位为USD
	UsedWdQuota          string `json:"usedWdQuota"`          // 过去24小时内已用提币额度，单位为USD
	Fee                  string `json:"fee"`                  // 固定的提币手续费数量 适用于链上提币
	BurningFeeRate       string `json:"burningFeeRate"`       // 燃烧费率，如 0.05 代表 5%。 部分币种会收取燃烧费用。燃烧费用按照提币数量（不含gas fee） 乘以 燃烧费率，在提币数量基础上扣除。 适用于链上提币
	MainNet              bool   `json:"mainNet"`              // 当前链是否为主链
	NeedTag              bool   `json:"needTag"`              // 当前链提币是否需要标签（tag/memo）信息，如 EOS该字段为true
	MinDepArrivalConfirm string `json:"minDepArrivalConfirm"` // 充值到账最小网络确认数。币已到账但不可提。
	MinWdUnlockConfirm   string `json:"minWdUnlockConfirm"`   // 提现解锁最小网络确认数
	DepQuotaFixed        string `json:"depQuotaFixed"`        // 充币固定限额，单位为USD 没有充币限制则返回""
	UsedDepQuotaFixed    string `json:"usedDepQuotaFixed"`    // 已用充币固定额度，单位为USD 没有充币限制则返回""
	DepQuoteDailyLayer2  string `json:"depQuoteDailyLayer2"`  // Layer2网络每日充值上限
}

func (s *FundingAssetCurrenciesService) Do(ctx context.Context, opts ...RequestOption) ([]*FundingAssetCurrency, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/asset/currencies",
		secType:  secTypeSigned,
	}

	if len(s.ccy) > 0 {
		r.setParam("ccy", strings.Join(s.ccy, ","))
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*FundingAssetCurrency)
	if err := sonic.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}

type FundingAssetBalancesService struct {
	c *Client

	ccy []string
}

func (s *FundingAssetBalancesService) Ccy(ccy []string) *FundingAssetBalancesService {
	s.ccy = ccy
	return s
}

type FundingAssetBalance struct {
	Ccy              string `json:"ccy"`       // 币种名称，如 BTC
	Balance          string `json:"bal"`       // 币种余额
	FrozenBalance    string `json:"frozenBal"` // 币种冻结余额
	AvailableBalance string `json:"availBal"`  // 币种可用余额
}

func (s *FundingAssetBalancesService) Do(ctx context.Context, opts ...RequestOption) ([]*FundingAssetBalance, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/asset/balances",
		secType:  secTypeSigned,
	}

	if len(s.ccy) > 0 {
		r.setParam("ccy", strings.Join(s.ccy, ","))
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*FundingAssetBalance)
	if err := sonic.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return *result, nil
}

// /api/v5/asset/asset-valuation

type FundingAssetValuationService struct {
	c *Client

	ccy string
}

func (s *FundingAssetValuationService) Ccy(ccy string) *FundingAssetValuationService {
	s.ccy = ccy
	return s
}

type FundingAssetValuation struct {
	TotalBal string `json:"totalBal"` // 账户总资产估值
	Ts       string `json:"ts"`       // 数据更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	Details  struct {
		Funding string `json:"funding"` // 资金账户
		Trading string `json:"trading"` // 交易账户
		Earn    string `json:"earn"`    // 金融账户
	} `json:"details"` // 各个账户的资产估值
}

func (s *FundingAssetValuationService) Do(ctx context.Context, opts ...RequestOption) (*FundingAssetValuation, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/asset/asset-valuation",
		secType:  secTypeSigned,
	}

	if len(s.ccy) > 0 {
		r.setParam("ccy", s.ccy)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*FundingAssetValuation)
	if err := sonic.Unmarshal(data, result); err != nil {
		return nil, err
	}
	if len(*result) > 0 {
		return (*result)[0], nil
	}
	return nil, nil
}
