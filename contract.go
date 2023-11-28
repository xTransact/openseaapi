package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/xTransact/errx/v3"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

// GetContract gets a smart contract for a given chain and address.
// @param ch (chain.Chain): required: The blockchain on which to filter the results.
// @param address: required: The unique public blockchain identifier for the contract.
// DOC: https://docs.opensea.io/reference/get_contract
func (c *client) GetContract(ctx context.Context, ch chain.Chain, address common.Address) (
	resp *openseamodels.Contract, err error) {

	if !common.IsHexAddress(address.String()) {
		return nil, errx.New("invalid address")
	}

	// GET /api/v2/chain/{chain}/contract/{address}
	url := fmt.Sprintf("%s/api/v2/chain/%s/contract/%s",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), address.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, ch.IsTestNet())
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.Contract)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}
