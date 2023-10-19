package openseaapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

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
		return nil, errors.New("invalid address")
	}

	// GET /api/v2/chain/{chain}/contract/{address}
	url := fmt.Sprintf("%s/api/v2/chain/%s/contract/%s",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), address.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	c.acceptJson(req)
	if !ch.IsTestNet() {
		c.challenge(req)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp = new(openseamodels.Contract)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}
