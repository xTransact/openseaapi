package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseamodels"
)

func TestGetNft(t *testing.T) {
	ctx := context.Background()
	cli := NewClient()
	const testCollection = "0xb31d6b5516eed64a874e9f7ab605e359e20b645f"
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := cli.GetNft(ctx, chain.Goerli, &openseamodels.GetNftPayload{
			Address:    common.HexToAddress(testCollection),
			Identifier: "1",
		})
		assert.NoError(t, err)

		data, err := json.Marshal(resp)
		assert.NoError(t, err)
		fmt.Println(string(data))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := cli.GetNft(ctx, chain.Goerli, &openseamodels.GetNftPayload{
			Address:    common.HexToAddress(testCollection),
			Identifier: "2",
		})
		assert.NoError(t, err)

		data, err := json.Marshal(resp)
		assert.NoError(t, err)
		fmt.Println(string(data))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := cli.GetNft(ctx, chain.Goerli, &openseamodels.GetNftPayload{
			Address:    common.HexToAddress(testCollection),
			Identifier: "3",
		})
		assert.NoError(t, err)

		data, err := json.Marshal(resp)
		assert.NoError(t, err)
		fmt.Println(string(data))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := cli.GetNft(ctx, chain.Goerli, &openseamodels.GetNftPayload{
			Address:    common.HexToAddress(testCollection),
			Identifier: "4",
		})
		assert.NoError(t, err)

		data, err := json.Marshal(resp)
		assert.NoError(t, err)
		fmt.Println(string(data))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := cli.GetNft(ctx, chain.Goerli, &openseamodels.GetNftPayload{
			Address:    common.HexToAddress(testCollection),
			Identifier: "5",
		})
		assert.NoError(t, err)

		data, err := json.Marshal(resp)
		assert.NoError(t, err)
		fmt.Println(string(data))
	}()

	wg.Wait()
}
