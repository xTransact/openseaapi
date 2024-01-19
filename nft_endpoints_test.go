package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseamodels"
)

func TestGetAccount(t *testing.T) {
	ctx := context.Background()
	testAddr := common.HexToAddress(os.Getenv("WALLET_ADDRESS"))
	testApiKey := os.Getenv("OPENSEA_API_KEY")

	cli := NewClient(WithApiKey(testApiKey))
	resp, err := cli.GetAccount(ctx, testAddr)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, testAddr.String(), resp.Address.String())
}

func TestGetGoerliAccount(t *testing.T) {
	ctx := context.Background()
	testAddr := common.HexToAddress(os.Getenv("WALLET_ADDRESS"))

	cli := NewClient()
	resp, err := cli.GetAccount(ctx, testAddr, UseTestnets())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, testAddr.String(), resp.Address.String())
}

func TestGetNftsByAccount(t *testing.T) {
	ctx := context.Background()
	testAddr := common.HexToAddress(os.Getenv("WALLET_ADDRESS_ALPHA"))
	testApiKey := os.Getenv("OPENSEA_API_KEY")

	cli := NewClient(WithApiKey(testApiKey))
	resp, err := cli.ListNftsByAccount(ctx, chain.Ethereum, &openseamodels.GetNftsByAccountPayload{
		GetNftsBasePayload: &openseamodels.GetNftsBasePayload{
			BaseQueryParams: &openseamodels.BaseQueryParams{
				Limit: 100,
				Next:  "",
			},
			Address: testAddr,
		},
		Collection: "",
	})

	require.NoError(t, err)
	require.NotNil(t, resp)

	data, err := json.MarshalIndent(resp, "", "  ")
	require.NoError(t, err)
	fmt.Println(string(data))
}

func TestGetNft(t *testing.T) {
	ctx := context.Background()
	testApiKey := os.Getenv("OPENSEA_API_KEY")
	// azuki
	contract := common.HexToAddress("0xed5af388653567af2f388e6224dc7c4b3241c544")
	const tokenID = "1234"

	cli := NewClient(WithApiKey(testApiKey))
	resp, err := cli.GetNft(ctx, chain.Ethereum, &openseamodels.GetNftPayload{
		Address:    contract,
		Identifier: tokenID,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Nft)

	assert.Equal(t, tokenID, resp.Nft.Identifier)
	assert.Equal(t, contract, common.HexToAddress(resp.Nft.Contract))
}

func TestGetGoerliNft(t *testing.T) {
	ctx := context.Background()
	testCollectionContract := common.HexToAddress(os.Getenv("TEST_COLLECTION_CONTRACT"))
	const tokenID = "1"

	cli := NewClient()
	resp, err := cli.GetNft(ctx, chain.Sepolia, &openseamodels.GetNftPayload{
		Address:    testCollectionContract,
		Identifier: tokenID,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Nft)

	assert.Equal(t, tokenID, resp.Nft.Identifier)
	assert.Equal(t, testCollectionContract, common.HexToAddress(resp.Nft.Contract))
}
