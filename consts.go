package okx_connector

const Name = "okx-connector-go"

const Version = "0.6.0"

type Channel string

type DepthChannel Channel

const (
	DepthChannelBooks        = "books"
	DepthChannelbooks5       = "books5"
	DepthChannelBboTbt       = "bbo-tbt"
	DepthChannelBooksL2Tbt   = "books-l2-tbt"
	DepthChannelBooks50L2Tbt = "books50-l2-tbt"
)

func (c DepthChannel) Valid() bool {
	return c == DepthChannelBooks || c == DepthChannelbooks5 || c == DepthChannelBboTbt || c == DepthChannelBooksL2Tbt || c == DepthChannelBooks50L2Tbt
}

type KlineChannel Channel

const (
	KlineChannelCandle1s  = "candle1s"
	KlineChannelCandle1m  = "candle1m"
	KlineChannelCandle3m  = "candle3m"
	KlineChannelCandle5m  = "candle5m"
	KlineChannelCandle15m = "candle15m"
	KlineChannelCandle30m = "candle30m"
	KlineChannelCandle1H  = "candle1H"
	KlineChannelCandle2H  = "candle2H"
	KlineChannelCandle4H  = "candle4H"
	KlineChannelCandle6H  = "candle6H"
	KlineChannelCandle12H = "candle12H"
	KlineChannelCandle1D  = "candle1D"
	KlineChannelCandle2D  = "candle2D"
	KlineChannelCandle3D  = "candle3D"
	KlineChannelCandle5D  = "candle5D"
	KlineChannelCandle1W  = "candle1W"
	KlineChannelCandle1M  = "candle1M"
	KlineChannelCandle3M  = "candle3M"

	KlineChannelCandle3Mutc  = "candle3Mutc"
	KlineChannelCandle1Mutc  = "candle1Mutc"
	KlineChannelCandle1Wutc  = "candle1Wutc"
	KlineChannelCandle1Dutc  = "candle1Dutc"
	KlineChannelCandle2Dutc  = "candle2Dutc"
	KlineChannelCandle3Dutc  = "candle3Dutc"
	KlineChannelCandle5Dutc  = "candle5Dutc"
	KlineChannelCandle12Hutc = "candle12Hutc"
	KlineChannelCandle6Hutc  = "candle6Hutc"
)

func (c KlineChannel) Valid() bool {
	return c == KlineChannelCandle1s || c == KlineChannelCandle1m || c == KlineChannelCandle3m || c == KlineChannelCandle5m || c == KlineChannelCandle15m || c == KlineChannelCandle30m || c == KlineChannelCandle1H || c == KlineChannelCandle2H || c == KlineChannelCandle4H || c == KlineChannelCandle6H || c == KlineChannelCandle12H || c == KlineChannelCandle1D || c == KlineChannelCandle2D || c == KlineChannelCandle3D || c == KlineChannelCandle5D || c == KlineChannelCandle1W || c == KlineChannelCandle1M || c == KlineChannelCandle3M || c == KlineChannelCandle3Mutc || c == KlineChannelCandle1Mutc || c == KlineChannelCandle1Wutc || c == KlineChannelCandle1Dutc || c == KlineChannelCandle2Dutc || c == KlineChannelCandle3Dutc || c == KlineChannelCandle5Dutc || c == KlineChannelCandle12Hutc || c == KlineChannelCandle6Hutc
}
