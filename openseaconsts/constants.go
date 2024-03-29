package openseaconsts

import "github.com/ethereum/go-ethereum/common"

const (
	PlatformName = "opensea"
	ProtocolName = "seaport"
)

const (
	BaseUrlProd = "https://api.opensea.io"
	BaseUrlTest = "https://testnets-api.opensea.io"
)

const (
	OrderSideAsk = "ask"
	OrderSideBid = "bid"
)

var (
	// Deprecated use v16 instead
	SeaportV15Address = common.HexToAddress("0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC")
	SeaportV16Address = common.HexToAddress("0x0000000000000068F116a894984e2DB1123eB395")
)
