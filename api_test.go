package openseaapi

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

func TestGetListings(t *testing.T) {
	cli := NewClient(WithApiKey(getTestApiKey()))
	listings, err := cli.GetListings(context.Background(), chain.Ethereum, &openseamodels.GetListingsPayload{
		AssetContractAddress: openseaapiutils.StringPtr("0xed5af388653567af2f388e6224dc7c4b3241c544"), // azuki
		TokenIDs: []int{
			4332, 7626,
		},
	})
	require.NoError(t, err)
	fmt.Println(string(listings.ToJson()))
}

func getTestApiKey() string {
	return os.Getenv("OPENSEA_API_KEY")
}
