package okx_connector

const Name = "okx-connector-go"

const Version = "0.6.0"

type Channel string

type DepthChannel Channel

const (
	DepthChannelBooks        DepthChannel = "books"
	DepthChannelbooks5       DepthChannel = "books5"
	DepthChannelBboTbt       DepthChannel = "bbo-tbt"
	DepthChannelBooksL2Tbt   DepthChannel = "books-l2-tbt"
	DepthChannelBooks50L2Tbt DepthChannel = "books50-l2-tbt"
)

func (c DepthChannel) Valid() bool {
	return c == DepthChannelBooks || c == DepthChannelbooks5 || c == DepthChannelBboTbt || c == DepthChannelBooksL2Tbt || c == DepthChannelBooks50L2Tbt
}

type KlineChannel Channel

const (
	KlineChannelCandle1s     KlineChannel = "candle1s"
	KlineChannelCandle1m     KlineChannel = "candle1m"
	KlineChannelCandle3m     KlineChannel = "candle3m"
	KlineChannelCandle5m     KlineChannel = "candle5m"
	KlineChannelCandle15m    KlineChannel = "candle15m"
	KlineChannelCandle30m    KlineChannel = "candle30m"
	KlineChannelCandle1H     KlineChannel = "candle1H"
	KlineChannelCandle2H     KlineChannel = "candle2H"
	KlineChannelCandle4H     KlineChannel = "candle4H"
	KlineChannelCandle6H     KlineChannel = "candle6H"
	KlineChannelCandle12H    KlineChannel = "candle12H"
	KlineChannelCandle1D     KlineChannel = "candle1D"
	KlineChannelCandle2D     KlineChannel = "candle2D"
	KlineChannelCandle3D     KlineChannel = "candle3D"
	KlineChannelCandle5D     KlineChannel = "candle5D"
	KlineChannelCandle1W     KlineChannel = "candle1W"
	KlineChannelCandle1M     KlineChannel = "candle1M"
	KlineChannelCandle3M     KlineChannel = "candle3M"
	KlineChannelCandle3Mutc  KlineChannel = "candle3Mutc"
	KlineChannelCandle1Mutc  KlineChannel = "candle1Mutc"
	KlineChannelCandle1Wutc  KlineChannel = "candle1Wutc"
	KlineChannelCandle1Dutc  KlineChannel = "candle1Dutc"
	KlineChannelCandle2Dutc  KlineChannel = "candle2Dutc"
	KlineChannelCandle3Dutc  KlineChannel = "candle3Dutc"
	KlineChannelCandle5Dutc  KlineChannel = "candle5Dutc"
	KlineChannelCandle12Hutc KlineChannel = "candle12Hutc"
	KlineChannelCandle6Hutc  KlineChannel = "candle6Hutc"
)

func (c KlineChannel) Valid() bool {
	return c == KlineChannelCandle1s || c == KlineChannelCandle1m || c == KlineChannelCandle3m || c == KlineChannelCandle5m || c == KlineChannelCandle15m || c == KlineChannelCandle30m || c == KlineChannelCandle1H || c == KlineChannelCandle2H || c == KlineChannelCandle4H || c == KlineChannelCandle6H || c == KlineChannelCandle12H || c == KlineChannelCandle1D || c == KlineChannelCandle2D || c == KlineChannelCandle3D || c == KlineChannelCandle5D || c == KlineChannelCandle1W || c == KlineChannelCandle1M || c == KlineChannelCandle3M || c == KlineChannelCandle3Mutc || c == KlineChannelCandle1Mutc || c == KlineChannelCandle1Wutc || c == KlineChannelCandle1Dutc || c == KlineChannelCandle2Dutc || c == KlineChannelCandle3Dutc || c == KlineChannelCandle5Dutc || c == KlineChannelCandle12Hutc || c == KlineChannelCandle6Hutc
}
