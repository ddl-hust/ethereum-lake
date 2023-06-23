package one_inch_oracle

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./one_inch_oracle.abi --pkg one_inch_oracle --type OneInchOracle --out ./one_inch_oracle.go
var (
	oneInchPriceOralceAddressEthereum = common.HexToAddress("0x07D91f5fb9Bf7798734C3f606dB065549F6893bb")
	oneInchPriceOralceAddressPolygon  = common.HexToAddress("0x7F069df72b7A39bCE9806e3AfaF579E54D8CF2b9")
	oneInchPriceOracleAddressArbitrum = common.HexToAddress("0x735247fb0a604c0adC6cab38ACE16D0DbA31295F")
	oneInchPriceOracleAddressOptimism = common.HexToAddress("0x11DEE30E710B8d4a8630392781Cc3c0046365d4c")
	oneInchPriceOracleAddressBSC      = common.HexToAddress("0xfbD61B037C325b959c0F6A7e69D8f37770C2c550")
)

func OneInchPriceOracle(network constant.Network) (common.Address, error) {
	switch network {
	case constant.NetworkEthereum:
		return oneInchPriceOralceAddressEthereum, nil
	case constant.NetworkPolygon:
		return oneInchPriceOralceAddressPolygon, nil
	case constant.NetworkArbitrumOne:
		return oneInchPriceOracleAddressArbitrum, nil
	case constant.NetworkBSC:
		return oneInchPriceOracleAddressBSC, nil
	case constant.NetworkOptimism:
		return oneInchPriceOracleAddressOptimism, nil
	}

	return common.HexToAddress(""), fmt.Errorf("unsupported network")
}